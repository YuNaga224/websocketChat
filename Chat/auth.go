package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/stretchr/gomniauth"
	gomniauthcommon "github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
)

type ChatUser interface {
	UniqueID() string
	AvatarURL() string
}

type chatUser struct {
	gomniauthcommon.User
	uniqueID string
}

func (u chatUser) UniqueID() string {
	return u.uniqueID
}

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("auth"); err == http.ErrNoCookie || cookie.Value == "" {
		//認証されていないのでloginページにリダイレクトする
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		//何かしらのエラーが発生
		panic(err.Error())
	} else {
		//認証済みなのでラップされたハンドラを呼び出す
		h.next.ServeHTTP(w, r)
	}
}

// 任意のhttp.Handlerをラップした*authHandlerを生成するためのヘルパー関数
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

// loginHandlerはサードパーティへのログイン処理を受け持つ
func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		// URLに指定されたプロバイダー名を引数として認証プロバイダーのオブジェクトを取得
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln("認証プロバイダーの取得に失敗しました:", provider, "-", err)
		}
		// 認証プロセスのためのURLを取得
		loginUrl, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			log.Fatalln("GetBeginAuthURLの呼び出し中にエラーが発生しました", provider, "-", err)

		}
		//認証処理のためのページにリダイレクト
		w.Header().Set("Location", loginUrl)
		w.WriteHeader(http.StatusTemporaryRedirect)
		log.Println("TODO: ログイン処理", provider)
	case "callback":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln("認証プロバイダーの取得に失敗しました", provider, "-", err)

		}
		//URIからGETメソッドで提供された認可コードを取り出す
		//その情報をもとに認証プロバイダーにアクセスしてアクセストークンを取得しcredsにセット
		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			log.Fatalln("認証を成功できませんでした", provider, "-", err)
		}
		//アクセストークンをもとにユーザー情報をjson形式で取得
		user, err := provider.GetUser(creds)
		if err != nil {
			log.Fatalln("ユーザーの取得に失敗しました", provider, "-", err)
		}
		chatUser := &chatUser{User: user}
		m := md5.New()
		io.WriteString(m, strings.ToLower(user.Email()))
		chatUser.uniqueID = fmt.Sprintf("%x", m.Sum(nil))
		// グローバル変数avatarsを使用してGetAvatarURLメソッドを呼び出す
		avatarURL, err := avatars.GetAvatarURL(chatUser)
		if err != nil {
			log.Fatalln("GetAvatarURLに失敗しました", "-", err)
		}

		//userのNameフィールドの値をBase64形式でエンコード
		authCookieValue := objx.New(map[string]interface{}{
			"userid":     chatUser.uniqueID,
			"name":       user.Name(),
			"avatar_url": avatarURL,
		}).MustBase64()
		//authという名前でBase64形式の値をクッキーにセット
		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: authCookieValue,
			Path:  "/",
		})
		w.Header()["Location"] = []string{"/chat"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "アクション%sには非対応です", action)
	}

}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "auth",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	w.Header()["Location"] = []string{"/chat"}
	w.WriteHeader(http.StatusTemporaryRedirect)
}
