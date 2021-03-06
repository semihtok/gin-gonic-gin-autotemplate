package main

import (
	"strings"
	"github.com/gin-gonic/gin"
	_"html/template"
	"github.com/gin-gonic/contrib/renders/multitemplate"
	"io/ioutil"
)

func main() {

	router := gin.Default()
	setHtmlPages(router)
	
}


func setHtmlPages(router *gin.Engine) {

	templates := multitemplate.New()

	templateDir := "templates/"

	files, _ := ioutil.ReadDir("./templates")

	for _, f := range files {

		if f.Name() != "layout.html" {

			templates.AddFromFiles(strings.Split(f.Name(), ".")[0],
				templateDir + "layout.html",
				templateDir + f.Name())

		}
	}

	router.HTMLRender = templates

}
