package Controller
import (
    "local/Grocy/Godeps/_workspace/src/labix.org/v2/mgo"
    "local/Grocy/Godeps/_workspace/src/github.com/martini-contrib/render"
    "local/Grocy/Godeps/_workspace/src/packages/github.com/go-martini/martini"
    "local/Grocy/Models"
)

type ClientInterfaceController struct {
}

func (t ClientInterfaceController) TestAction() (int,string){
    return 200,"hello world!!!"
}

func (t ClientInterfaceController) GetInfoForEan(params martini.Params, db *mgo.Database, renderer render.Render) {
    var ean = params["ean"]

    var reweInfos = Models.GetReweInfosForEan(ean)



    renderer.JSON(200, reweInfos)
}