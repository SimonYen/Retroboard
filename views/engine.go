package views

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/view"
)

var Engine *view.DjangoEngine

func Init(isDebug bool) {
	Engine = iris.Django("./views/templates", ".html")
	Engine.Reload(isDebug)
}
