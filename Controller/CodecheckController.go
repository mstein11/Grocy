package Controller
import (
    "local/Grocy/Godeps/_workspace/src/github.com/martini-contrib/render"
    "local/Grocy/Models"
    "local/Grocy/Godeps/_workspace/src/packages/github.com/go-martini/martini"
    "log"
)


type CodecheckController struct {
}

type ResponseUserAgent struct {
    Useragent string `json:"user-agent"`
}

type SessionResponse struct {
    SessionId string `json: "sessionId"`
    Nonce string `json: "nonce"`
    ExpiresIn int `json: "expiresIn"`
}

type ResponseWrapper struct {
    Result SessionResponse `json: "result"`
}

const username = "marius_stein"

const authType = "DigestQuick"

func (t CodecheckController) GetAllInfosByEan(params martini.Params,renderer render.Render) (int,string){
    ean := params["ean"]
    log.Println("ean: " + ean)

    codecheck := Models.GetCodecheckInstance()
    result := codecheck.GetAllInfoByEan(ean)

    return 200,result
}



func (t CodecheckController) GetInfosByEan(params martini.Params,renderer render.Render) (int,string){

    ean := params["ean"]
    levelOfDetail := params["lod"]
    codecheck := Models.GetCodecheckInstance()
    result := codecheck.GetInfoByEan(ean, levelOfDetail)

    return 200, result
}