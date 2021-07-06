package utilities

import (
	"fmt"
	"net/http"
	"strings"
)

func StoreEPass(w http.ResponseWriter, s string) {
	fmt.Println(s)
	cookie := &http.Cookie{
		Name:   "EPass",
		Value:  s,
		MaxAge: 86400,
	}
	fmt.Println(s + " was stored")
	http.SetCookie(w, cookie)
}

func CheckEPass(w http.ResponseWriter, r *http.Request, pass string) bool {
	cookie, err := r.Cookie("EPass")
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		res := strings.Split(cookie.Value, " ")
		return Decrypt(res[0], res[1]) == pass
	}
	return true
}
