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
	distRoot:= string("./dist/mpEnv")
	r:=gin.Default()
	//r.Delims("{{","}}")
	r.SetFuncMap(template.FuncMap{
		"loadFile" : loadFile,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/css", "./public/css")
	r.Static("/js", "./public/js")

   	//cont , _ := loadFile(fmt.Sprintf("%s/%s", distRoot, "golang.html"))
//
	

  	cont , _ := loadFile(fmt.Sprintf("%s/%s", distRoot, "DRL.html"))
  	toc , _ := loadFile(fmt.Sprintf("%s/%s", distRoot, "DRL.html_toc"))
  	fmt.Println(toc)
  	r.GET("/DRL", func (c *gin.Context) {
		c.HTML(http.StatusOK, "bookcase.html", gin.H{
			"toc" : template.HTML(toc),
			"content" : template.HTML(cont),
		})
  	})

  	//r.GET("/ServerSetup", func (c *gin.Context) {
	//	c.HTML(http.StatusOK, "bookcase.html", gin.H{
	//		"toc" : template.HTML(toc),
	//		"content" : template.HTML(cont),
	//	})
  	//})
  
	r.Run(":3000")
}
