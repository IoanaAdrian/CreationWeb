package main

import (
	functionalities "./functionalities"
	ut "./utilities"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//var displayPath = "./testdir"

const MAIN_DIR_PATH = "testdir"

func main() {

	http.HandleFunc("/", handler)
	panic(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	//ip, err := ut.GetClientIp(r)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(ip)

	if len(r.URL.String()) == 1 {

		page, err := ioutil.ReadFile("./static/index.html")
		ut.HandleErr(err)
		fmt.Fprintf(w, string(page))
	}

	if strings.HasPrefix(r.URL.String(), "/"+MAIN_DIR_PATH) {

		functionalities.FileSystemHandle(w, r)

	} else if strings.HasPrefix(r.URL.String(), "/static") {

		functionalities.HandleStaticFiles(w, r)
	} else if strings.HasPrefix(r.URL.String(), "/login") {

		functionalities.Login(w, r)

	} else if strings.Contains(r.URL.String(), "/upload") {
		functionalities.Upload(w,r,functionalities.GetWindowsPath())
	}

}
