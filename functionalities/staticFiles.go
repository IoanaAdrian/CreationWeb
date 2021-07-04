package functionalities

import (
	ut "../utilities"
	"fmt"
	"io/ioutil"
	"net/http"
)
//HandleStaticFiles TODO: When entering '/static' let the user view the entire directory (https://codemyui.com/ascii-art-and-css-styled-directory-structure/)
func HandleStaticFiles(w http.ResponseWriter, r *http.Request){
	loadPage, err := ioutil.ReadFile("." + r.URL.String())
	if err != nil {
		ut.HandleErr(err)
		fmt.Fprintf(w, "<h1 style='color:red'>Page Unvalaible</h1>")
	} else {
		fmt.Fprintf(w, string(loadPage))
	}
}
