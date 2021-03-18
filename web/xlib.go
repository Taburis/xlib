package main

import (
	"io/ioutil"
	"strings"
	"os"
	"html/template"
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/contrib/static"
	"net/http"
	"fmt"
	"path/filepath"
	//"os"
)

func getList(fileType string, path string) []string{
	// iterating over the folders (path) and return the list of 
	// paths of files that satisfies the given type (fileType)
	var list strings.Builder
	filepath.Walk(path, func (path string, info os.FileInfo, err error) error{
			if err != nil {
				return err
			}
			s := strings.Split(path, ".")
			if fileType == s[len(s)-1]{
				list.WriteString(path)
				list.WriteString("\n")
				//fmt.Printf("found: %s\n",path)
			}
			return nil
		})
	return strings.Split(list.String(), "\n")
}

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

   	//cont , _ := loadFile(fmt.Sprintf("%s/%s", distRoot, "golang.html"))
//
	//cont , _ := loadFile(fmt.Sprintf("%s/%s", distRoot, "golang.html"))
  	//toc , _ := loadFile(fmt.Sprintf("%s/%s", distRoot, "golang.html_toc"))
  	//fmt.Println(toc)
  	//r.GET("/golang", func (c *gin.Context) {
	//	c.HTML(http.StatusOK, "bookcase.html", gin.H{
	//		"toc" : template.HTML(toc),
	//		"content" : template.HTML(cont),
	//	})
  	//})
	cont , _ := loadFile(fmt.Sprintf("%s/category/notesIndex.html",distRoot))
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content" : template.HTML(cont),
		})
	})

	listHtml := getList("html", fmt.Sprintf("%s/notes",distRoot))
	listTOC := getList("html_toc", fmt.Sprintf("%s/notes",distRoot))
	for i:=0 ; i<len(listHtml); i++{
		content , _ := loadFile(listHtml[i])
  		toc , _ := loadFile(listTOC[i])
  		if len(listHtml[i])==0 {continue }
  		tmp := listHtml[i][len(distRoot)-1:len(listHtml[i])-5]
  		r.GET(fmt.Sprintf("/%s",tmp), func (c *gin.Context) {
			c.HTML(http.StatusOK, "bookcase.html", gin.H{
				"toc" : template.HTML(toc),
				"content" : template.HTML(content),
			})
  		})
	}

  	//cont1 , _ := loadFile(fmt.Sprintf("%s/%s",
  	//	fmt.Sprintf("%s/notes/Machine_Learning",distRoot), "DRL.html"))
  	//toc , _ := loadFile(fmt.Sprintf("%s/%s",
  	//	fmt.Sprintf("%s/notes/Machine_Learning",distRoot), "DRL.html_toc"))
  	//r.GET("/notes/DRL", func (c *gin.Context) {
	//	c.HTML(http.StatusOK, "bookcase.html", gin.H{
	//		"toc" : template.HTML(toc),
	//		"content" : template.HTML(cont1),
	//	})
  	//})

	r.Run(":80")
}
