package functionalities

import (
	ut "../utilities"
	"net/http"
	"os"
)

func MakeDir(w http.ResponseWriter, r *http.Request, path string) {
	r.ParseForm()
	ut.HandleErr(os.Mkdir(path+"/"+r.FormValue("folderName"), 0200))
	http.Redirect(w, r, "http://moth5002.go.ro:3000"+path[1:], 302)
}
