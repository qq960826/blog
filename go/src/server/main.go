package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
)
var pagemap map[string]string
var catmap map[string]string

func init() {
	jsonFile,err:=os.Open("pagemap.json")
	if err!=nil{
		log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err=json.Unmarshal(byteValue,&pagemap)
	if err!=nil{
		log.Fatal(err)
	}
	jsonFile,err=os.Open("catmap.json")
	if err!=nil{
		log.Fatal(err)
	}
	byteValue, _ = ioutil.ReadAll(jsonFile)
	err=json.Unmarshal(byteValue,&catmap)
	if err!=nil{
		log.Fatal(err)
	}
}
func main()  {

	r:=gin.New()
	r.Use(JumpMiddleware)
	r.Static("/","./")
	r.Run() // listen and serve on 0.0.0.0:8080
}
func JumpMiddleware(c *gin.Context) {
	if c.Request.URL.Path!="/"{
		c.Next()
		return
	}
	pagenumber:=c.Query("p")
	if path,ok:=pagemap[pagenumber];ok{
		c.Redirect(301,path)
		return
	}
	catnumber:=c.Query("cat")
	if path,ok:=catmap[catnumber];ok{
		c.Redirect(301,path)
		return
	}
	c.Next()
}