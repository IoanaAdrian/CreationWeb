package functionalities

import (
	ut "../utilities"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type element struct {
	Name string
	Path string
}

type pathS struct {
	DirName string
	URL string
}
var windowsPath string

func GetWindowsPath() string{
	return windowsPath
}

func FileSystemHandle(w http.ResponseWriter, r *http.Request) {
	var path = "." + r.URL.String()
	windowsPath = ut.HtmlPathToWindows(path)
	fmt.Println(windowsPath)

	instance, err := os.Stat(windowsPath)
	if err != nil {
		panic(err)
	}

	if instance.Mode().IsDir() {

		page := template.Must(template.ParseFiles("./static/directoryStructure/main.html"))
		tmp := pathS{windowsPath,ut.GetURL()}
		ut.HandleErr(page.Execute(w, tmp))

		//fmt.Fprintf(w, "<p>"+windowsPath+"</p>")
		//fmt.Fprintf(w, "<div class='folder-wrap'>")

		file, err := ioutil.ReadDir(windowsPath)
		if err != nil {
			panic(err)
		}
		for _, f := range file {
			t := template.New("ex")
			p := element{f.Name(), windowsPath[1:]}

			if f.IsDir() {
				t, _ = t.Parse("<a href='{{.Path}}/{{.Name}}' class='tile folder'><div><i class='mdi mdi-folder'></i><h3>{{.Name}}</h3><p>Description</p></div></a>")
				_ = t.Execute(w, p)
			} else {
				t, _ = t.Parse("<a href='{{.Path}}/{{.Name}}' class='tile form'><div><i class='mdi mdi-file-document'></i><h3>{{.Name}}</h3><p>Description</p></div></a>")
				_ = t.Execute(w, p)
			}
		}
	} else {
		ut.FileDownload(w, windowsPath, instance.Name())
	}
}
