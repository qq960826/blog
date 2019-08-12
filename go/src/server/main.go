package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
)
var jumpmap map[string]string
func main()  {
	jsonFile,err:=os.Open("map.json")
	if err!=nil{
		log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err=json.Unmarshal(byteValue,&jumpmap)
	if err!=nil{
		log.Fatal(err)
	}
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
	querynumber:=c.Query("p")
	if querynumber==""{
		c.Next()
		return
	}
	if path,ok:=jumpmap[querynumber];ok{
		c.Redirect(301,path)
		return
	}
	c.Next()
}