package main

import (
	"local/Grocy/Controller"
	"local/Grocy/Godeps/_workspace/src/packages/github.com/go-martini/martini"
)

var martiniApp *martini.Martini
var martiniRouter martini.Router

func main() {
	martiniApp = martini.New()
	martiniRouter = martini.NewRouter()

	var baseController = Controller.GetBaseController()
	baseController.HandleRouting(martiniRouter)

	//Not quite sure why they are needed
	martiniApp.Use(martini.Recovery())
	martiniApp.Use(martini.Logger())

	martiniApp.Action(martiniRouter.Handle)
	martiniApp.Run()
}
