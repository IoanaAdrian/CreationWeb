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
	ip,err := ut.GetClientIp(r)
	if err!=nil{
		panic(err)
	}
	fmt.Println(ip)
	if strings.HasPrefix(r.URL.String(), "/testdir") {
		var path = "." + r.URL.String()
		windowsPath := ut.Convert(path)
		fmt.Println(windowsPath)
		//var windowsPath = strings.Replace(path, "%20", "` ", -1)
		instance, err := os.Stat(windowsPath)
		if err != nil {
			panic(err)
		}
		if instance.Mode().IsDir() {
			file, err := ioutil.ReadDir(windowsPath)
			if err != nil {
				panic(err)
			}
			for _, f := range file {
				t := template.New("ex")
				p := tmp{f.Name(), windowsPath[1:]}
				t, _ = t.Parse("<a href='{{.Path}}/{{.Name}}'>{{.Name}}</a>")
				_ = t.Execute(w, p)
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
	}

}
