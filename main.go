package main

import (
	"flag"
	"net/http"

	"github.com/guotie/config"
	"github.com/smtc/gotua/models"
	"github.com/smtc/goutils"
	"github.com/zenazn/goji"

	"github.com/smtc/gotua/admin"
	_ "github.com/smtc/gotua/models"
)

var (
	configFn = flag.String("config", "./config.json", "config file path")
)

func main() {
	config.ReadCfg(*configFn)

	models.InitDB()
	run()
}

func run() {
	virtualPath := config.GetStringDefault("virtual_path", "")
	// route /admin
	goji.Handle(virtualPath+"/admin/*", admin.AdminMux())
	goji.Get(virtualPath+"/admin", http.RedirectHandler("/admin/", 301))

	// static files
	goji.Get(virtualPath+"/assets/*", http.FileServer(http.Dir("./")))

	goji.NotFound(goutils.NotFound)

	goji.Serve()
}
