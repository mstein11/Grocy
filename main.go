package main

import (
	"local/Grocy/Controller"
	"local/Grocy/Godeps/_workspace/src/github.com/martini-contrib/render"
	"local/Grocy/Godeps/_workspace/src/labix.org/v2/mgo"
	"local/Grocy/Godeps/_workspace/src/packages/github.com/go-martini/martini"
	"os"
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
	martiniApp.Use(render.Renderer())
	martiniApp.Use(db())

	martiniApp.Action(martiniRouter.Handle)
	martiniApp.Run()
}

/*
   the function returns a martini.Handler which is called on each request. We simply clone
   the session for each request and close it when the request is complete. The call to c.Map
   maps an instance of *mgo.Database to the request context. Then *mgo.Database
   is injected into each handler function.
*/
func db() martini.Handler {
	session, err := mgo.Dial(os.Getenv("MGO_DB_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}

	return func(c martini.Context) {
		s := session.Clone()
		c.Map(s.DB(os.Getenv("MONGO_DB"))) // local
		defer s.Close()
		c.Next()
	}
}
