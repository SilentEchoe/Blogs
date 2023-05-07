package main

/*
1.用户访问app demo 服务,app 验证用户是否登录,如果没有登录,跳转到dex登录页面
2.用户在dex登录页面填写用户名和密码，dex验证用户名和密码是否正确，如果正确，dex生成一个code，重定向到app demo服务
3.app demo服务拿到code，再次请求dex，dex验证code是否正确，如果正确，dex生成一个id_token和access_token，重定向到app demo服务
4.app demo服务拿到id_token和access_token，再次请求dex，dex验证id_token和access_token是否正确，如果正确，dex生成一个refresh_token，重定向到app demo服务
*/

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/login", DexLogin)

	// Start the server on port
	http.ListenAndServe(":5555", nil)
}

func DexLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
