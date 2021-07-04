package utilities

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

func FileDownload(w http.ResponseWriter, path string, fileName string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	//w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	_, _ = io.Copy(w, bytes.NewReader(content))
}
