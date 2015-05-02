package Controller


type TestController struct {
}

func (t TestController) TestAction() (int,string){
    return 200,"hello world!!!"
}