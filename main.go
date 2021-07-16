package main

import (
	"./functionalities"
	ut "./utilities"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)


var testing = false

func main() {
	if testing {
		ut.PathVisualisation()
	} else {
		ut.SetConfig()
		http.HandleFunc("/", handler)
		panic(http.ListenAndServe(":8080", nil))
	}
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
	} else if strings.HasPrefix(r.URL.String(), "/login") {

		functionalities.Login(w, r)

	} else if ut.CheckEPass(w, r, ut.GetPassword()) {

		if strings.HasPrefix(r.URL.String(), "/"+ut.GetMainDirPath()) {

			functionalities.FileSystemHandle(w, r)

		} else if strings.HasPrefix(r.URL.String(), "/static") {

			functionalities.HandleStaticFiles(w, r)
		} else if strings.Contains(r.URL.String(), "/upload") {

			functionalities.Upload(w, r, functionalities.GetWindowsPath())

		} else if strings.Contains(r.URL.String(), "/mkdir") {

			functionalities.MakeDir(w, r, functionalities.GetWindowsPath())

		}
	}

}
