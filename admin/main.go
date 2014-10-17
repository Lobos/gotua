package admin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/guotie/config"
	"github.com/smtc/goutils"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

func AdminMux() *web.Mux {
	virtualPath := config.GetStringDefault("virtual_path", "")

	mux := web.New()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Get(virtualPath+"/admin/", indexHandler)
	mux.Get(virtualPath+"/admin/menu", menuHandler)

	mux.Get(virtualPath+"/admin/account/", AccountList)
	mux.Get(virtualPath+"/admin/account/:id", AccountEntity)

	mux.Get(virtualPath+"/admin/table/", TableList)
	mux.Get(virtualPath+"/admin/table/:id", TableEntity)
	mux.Post(virtualPath+"/admin/table/", TableSave)
	mux.Delete(virtualPath+"/admin/table/", TableDelete)

	mux.Get(virtualPath+"/admin/column/type", ColumnType)
	mux.Get(virtualPath+"/admin/column/filter", ColumnFilter)
	mux.Get(regexp.MustCompile(`/admin/column/(?P<table_id>.+)/$`), ColumnList)
	mux.Get(regexp.MustCompile(`/admin/column/(?P<table_id>.+)/(?P<id>.+)$`), ColumnEntity)
	mux.Post(regexp.MustCompile(`/admin/column/(?P<table_id>.+)/$`), ColumnSave)
	mux.Delete(regexp.MustCompile(`/admin/column/(?P<table_id>.+)/$`), ColumnDelete)

	mux.Get(regexp.MustCompile(`/admin/dynamic/struct/(?P<table>.+)/(?P<method>.+)$`), DynamicStruct)
	mux.Get(regexp.MustCompile(`/admin/dynamic/api/(?P<table>.+)/$`), DynamicList)
	mux.Post(regexp.MustCompile(`/admin/dynamic/api/(?P<table>.+)/$`), DynamicSave)
	mux.Delete(regexp.MustCompile(`/admin/dynamic/api/(?P<table>.+)/$`), DynamicDelete)
	mux.Get(regexp.MustCompile(`/admin/dynamic/api/(?P<table>.+)/(?P<id>.+)$`), DynamicEntity)

	mux.Get(regexp.MustCompile(`/admin/(?P<model>.+)\.(?P<fn>.+):(?P<param>.+)$`), templateHandler)
	mux.Get(regexp.MustCompile(`/admin/(?P<model>.+)\.(?P<fn>.+)$`), templateHandler)

	mux.NotFound(goutils.NotFound)
	return mux
}

func indexHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	h := goutils.HttpHandler(c, w, r)
	h.RenderHtml("/admin/main.html")
}

/*
模板页暂时以 model.fn:param 分级
*/
func templateHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	temp := fmt.Sprintf("/admin/%s_%s.html", c.URLParams["model"], c.URLParams["fn"])
	goutils.Render(w).RenderHtml(temp)
}

func menuHandler(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadFile("./admin/menu.json")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}
