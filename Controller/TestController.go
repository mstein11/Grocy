package Controller
import (
    "local/Grocy/Godeps/_workspace/src/labix.org/v2/mgo"
    "local/Grocy/Models"
    "local/Grocy/Godeps/_workspace/src/github.com/martini-contrib/render"
)


type TestController struct {
}

func (t TestController) TestAction() (int,string){
    return 200,"hello world!!!"
}

func (t TestController) TestDatabase(db *mgo.Database, renderer render.Render) {
    var testModels []Models.TestModel
    db.C("test").Find(nil).All(&testModels)

    renderer.JSON(200, testModels)
}