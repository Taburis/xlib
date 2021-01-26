package main

import (
	"io/ioutil"
	"html/template"
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/contrib/static"
	"net/http"
	"fmt"
	//"os"
)


func loadFile(path string) (string,error){


	cont , error := ioutil.ReadFile(path)
	if error != nil {
        return "", error
    }
    return string(cont), error
}

func main(){
	//gin.SetMode(gin.ReleaseMode)
	distRoot:= string("./dist")
	r:=gin.Default()
	//r.Delims("{{","}}")
	r.SetFuncMap(template.FuncMap{
		"loadFile" : loadFile,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/css", "./public/css")
	r.Static("/js", "./public/js")

   	cont , _ := loadFile(fmt.Sprintf("%s/%s", distRoot, "golang.html"))
	r.GET("/golang", func (c *gin.Context) {
		c.HTML(http.StatusOK, "bookcase.html", gin.H{
			"content" : template.HTML(cont),
		})
  	})

  
	r.Run(":3000")
}