package main

import (
	ut "./Utilities"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type tmp struct {
	Name string
	Path string
}

//var displayPath = "./testdir"

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

	if strings.HasPrefix(r.URL.String(), "/testdir") {
		var path = "." + r.URL.String()
		windowsPath := ut.Convert(path)
		fmt.Println(windowsPath)

		instance, err := os.Stat(windowsPath)
		if err != nil {
			panic(err)
		}

		if instance.Mode().IsDir() {


			pageBeginning, err := ioutil.ReadFile("./hpages/directoryStructure.html")

			if err != nil {
				panic(err)
			}
			fmt.Fprintf(w, string(pageBeginning))
			fmt.Fprintf(w,"<p>"+windowsPath+"</p>")
			fmt.Fprintf(w,"<div class='folder-wrap'>")

			file, err := ioutil.ReadDir(windowsPath)
			if err != nil {
				panic(err)
			}
			for _, f := range file {
				t := template.New("ex")
				p := tmp{f.Name(), windowsPath[1:]}

				if f.IsDir() {
					t, _ = t.Parse("<a href='{{.Path}}/{{.Name}}' class='tile folder'><div><i class='mdi mdi-folder'></i><h3>{{.Name}}</h3><p>Description</p></div></a>")
					_ = t.Execute(w, p)
				} else {
					t, _ = t.Parse("<a href='{{.Path}}/{{.Name}}' class='tile form'><div><i class='mdi mdi-file-document'></i><h3>{{.Name}}</h3><p>Description</p></div></a>")
					_ = t.Execute(w, p)
				}
			}
		} else {
			content, err := ioutil.ReadFile(windowsPath)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Disposition", "attachment; filename="+instance.Name())
			w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
			_, _ = io.Copy(w, bytes.NewReader(content))
		}
	} else if strings.HasPrefix(r.URL.String(), "/hpages") {
		loadPage, err := ioutil.ReadFile("." + r.URL.String())
		if err != nil {
			fmt.Fprintf(w, "<h1 style='color:red'>Page Unvalaible</h1>")
		} else {
			fmt.Fprintf(w, string(loadPage))
		}
	}

}
