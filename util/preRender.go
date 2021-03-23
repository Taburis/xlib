package main

import (
	"github.com/Taburis/xlib/manager"
	//"fmt"
	"io/ioutil"
	//"golang.org/x/net/html"
	//html2 "golang.org/x/net/html"
	//"strings"
)

func loadFile(path string) (string,error){
	cont , error := ioutil.ReadFile(path)
	if error != nil {
        return "", error
    }
    return string(cont), error
}

func main(){
	m := manager.New()
	m.ProduceHTML("../../Documents/myLib","../web/dist") // test generating the html lib from MD
	m.ProduceIndex("../web/dist/notes", "../web/dist/category")
}
//func main(sourcepath string, destpath string){
//	p.ProduceHTML(sourcepath,destpath) // test generating the html lib from MD
//	p.ProduceIndex(fmt.Sprintf("%s/notes", destpath), fmt.Sprintf("%s/category", destpath))
//}
