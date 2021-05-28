package main
//
//import (
//	"fmt"
//	"log"
//	"net/http"
//	"io/ioutil"
//)
//
//var PORT = ":8080"
//
//func checkError(e error) {
//	if e != nil {
//		panic(e)
//	}
//}
//
//func handler(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path[1:] == "" {
//		_, err := fmt.Fprintf(w, "Hi there, I love %s !", r.URL.Path[1:])
//		checkError(err)
//	}
//}
//
//func viewHandler(w http.ResponseWriter, r *http.Request) {
//	title := r.URL.Path[len("/view/"):]
//	p, _ := loadPage(title)
//	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
//}
//
//func main() {
//	http.HandleFunc("/", handler)
//	log.Fatal(http.ListenAndServe(PORT, nil))
//}
