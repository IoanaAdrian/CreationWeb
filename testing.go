package main

import "io/ioutil"

type Page struct {
	Title string
	Body  []uint8
}

func (p *Page) save() error {
	fileName := p.Title + ".txt"
	return ioutil.WriteFile(fileName, p.Body, 0600)
}

func main() {

}
