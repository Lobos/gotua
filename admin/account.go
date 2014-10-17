package admin

import (
	"net/http"

	"github.com/smtc/gotua/models"
	"github.com/smtc/goutils"
	"github.com/zenazn/goji/web"
)

func AccountList(w http.ResponseWriter, r *http.Request) {
	models, _ := models.AccountList(0, 20, nil)
	list, _ := goutils.ToMapList(models, []string{"email", "name"}, goutils.FilterModeInclude)
	goutils.Render(w).RenderPage(list, 0, r)
}

func AccountEntity(c web.C, w http.ResponseWriter, r *http.Request) {
}
