package main

import (
	"github.com/Taburis/xlib/manager"
	//"fmt"
	"io/ioutil"
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
	m.ProduceHTML("../../myLib","../web/dist") // test generating the html lib from MD
	m.ProduceIndex("../web/dist/notes", "../web/dist/category")
	err:= m.CopyDirectory("../../myLib/images", "../web/public/images")
    if err!=nil { panic(err)}
}
