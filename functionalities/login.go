package functionalities

import (
	ut "../utilities"
	"fmt"
	"net/http"
)

const PASSOWRD = ""

func Login(w http.ResponseWriter, r *http.Request){
	/*
	TODO:
		When logging in, generate a cookie, store it in a database -- for multiple cookies it wil take a lot of time
		When logging in, generate a cookie that stores the password, but encrypted (do that by calling the C++ program executable); to try to access the db
		the cookie is checked by the reverse function
	https://golangbyexample.com/cookies-golang/
	*/
	ut.HandleErr(r.ParseForm())
	pass := r.FormValue("password")
	fmt.Println(pass)
	if pass==PASSOWRD{
		ut.StoreEPass(w,ut.Encrypt(pass))
		http.Redirect(w, r, "http://moth5002.go.ro:3000/softHoarders", 302)
	}

}

func GetPass() string {
	return PASSOWRD
}
