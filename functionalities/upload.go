package functionalities

import (
	ut "../utilities"
	"io/ioutil"
	"net/http"
)

func Upload(w http.ResponseWriter, r *http.Request, path string) {

	//TODO: track process of uploading (https://stackoverflow.com/questions/13452914/how-do-i-track-the-upload-progress-in-golang/13454533)

	file, handler, err := r.FormFile("uploadFile") // Use the second argument to get things such as fileSize or fileName
	ut.HandleErr(err)
	fileBytes, err := ioutil.ReadAll(file)
	ut.HandleErr(err)
	ut.HandleErr(ioutil.WriteFile(path+"/"+handler.Filename, fileBytes, 0200))

	http.Redirect(w, r, "http://moth5002.go.ro:3000"+path[1:], 302)
}
