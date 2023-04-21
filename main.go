// package treeJiazu
package main

import (
	"flag"
	"fmt"
	"github.com/devfeel/dotweb"
	"log"
	"os"
	"treeJiazu/treebuild"
)

type App struct {
	Web *dotweb.DotWeb
}

type ResBody struct {
	Status string `json:"status"`
	Result string `json:"result"`
}

func NewApp() *App {
	var a = &App{}
	a.Web = dotweb.New()
	return a
}

var app = NewApp()

func main() {
	treebuild.BuildWenShiJiaZu()
	treebuild.UpdateInfo()
	db, root := treebuild.GetRootByName("温友贵")
	defer db.Close()
	root.PrintDetailByName("温有福")
	root.PrintTree(0)
	var err error

	var (
		version = flag.Bool("version", false, "version v1.0")
		port    = flag.Int(
			"port", 8081, "listen port.")
	)

	flag.Parse()

	if *version {
		fmt.Println("v1.0")
		os.Exit(0)
	}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	InitRoute(app.Web.HttpServer)
	log.Println("Start China Phone Query Server on ", *port)
	app.Web.StartServer(*port)
}
func InitRoute(server *dotweb.HttpServer) {
	server.GET("/", indexHandler)
}

func indexHandler(ctx dotweb.Context) error {
	name := ctx.QueryString("name")

	if name == "" {
		log.Println("ERROR: 查无此人")
		return ctx.WriteString("欢迎查询温室宗族相关人员")
	}

	var message = ResBody{
		Status: "failed",
		Result: "",
	}
	log.Println(name)
	db, root := treebuild.GetRootByName("温友贵")
	defer db.Close()
	message.Status = "success"
	message.Result = root.PrintDetailByNameStr(name)
	log.Printf("indexHandler  %+\n", message)
	return ctx.WriteJson(message)
}



