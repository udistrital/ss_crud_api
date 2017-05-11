package main

import (
	_ "github.com/udistrital/ss_crud_api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
)

func init() {
	//orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres2016@127.0.0.1:5432/udistrital?sslmode=disable&search_path=seguridad_social")

	orm.RegisterDataBase("default", "postgres", "postgres://"+beego.AppConfig.String("PGuser")+":"+
		beego.AppConfig.String("PGpass")+"@"+
		beego.AppConfig.String("PGurls")+"/"+
		beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+
		beego.AppConfig.String("PGschemas")+"")
}

func main() {

	orm.Debug = true
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()

	/*
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
	*/
	/*
		orm.Debug = true
		beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
			AllowHeaders: []string{"Origin", "x-requested-with",
				"content-type",
				"accept",
				"origin",
				"authorization",
				"x-csrftoken"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
		}))
		if beego.BConfig.RunMode == "dev" {
			beego.BConfig.WebConfig.DirectoryIndex = true
			beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		}

		beego.Run() */
}
