package Controller
import (
    "local/Grocy/Godeps/_workspace/src/packages/github.com/go-martini/martini"
)


var baseControllerInstance *BaseController

type BaseController struct {
    TestController TestController
}

//Here the routing is configured
func (c BaseController) HandleRouting(router martini.Router) {
    router.Get("/", c.TestController.TestAction)
}

func GetBaseController() *BaseController {
    if (baseControllerInstance == nil) {
        baseControllerInstance = getBaseController()
    }

    return baseControllerInstance
}

func getBaseController() *BaseController {
    return &BaseController{TestController{}}
}

