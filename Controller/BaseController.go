package Controller
import (
    "local/Grocy/Godeps/_workspace/src/packages/github.com/go-martini/martini"
)


var baseControllerInstance *BaseController

type BaseController struct {
    TestController TestController
    ClientInterfaceController ClientInterfaceController
    CodecheckController CodecheckController
}

//Here the routing is configured
func (c BaseController) HandleRouting(router martini.Router) {
    router.Get("/", c.TestController.TestAction)
    router.Get("/test/database", c.TestController.TestDatabase)

    router.Get("/ClientApi/GetInfoForEan", c.ClientInterfaceController.GetInfoForEan)
    router.Get("/codecheck/GetAllInfosByEan/:ean", c.CodecheckController.GetAllInfosByEan)
    router.Get("/codecheck/GetInfosByEan/:ean/:lod", c.CodecheckController.GetInfosByEan)
}

func GetBaseController() *BaseController {
    if (baseControllerInstance == nil) {
        baseControllerInstance = getBaseController()
    }

    return baseControllerInstance
}

func getBaseController() *BaseController {
    return &BaseController{TestController{}, ClientInterfaceController{}, CodecheckController{}}
}

