package Models
import (
    "gopkg.in/jmcvetta/napping.v1"
    "time"
    "local/Grocy/Util"
    "encoding/base64"
    "log"
    "os"
    "encoding/json"
    "fmt"
    "encoding/hex"
    "crypto/sha256"
    "crypto/hmac"
    "net/http"
    "strconv"
)

const username = "marius_stein"
const authType = "DigestQuick"
const deviceId = "TEST_123456"

const codecheck_auth_url = "http://www.codecheck.info/WebService/rest/session/auth"
const codecheck_ean_url = "http://www.codecheck.info/WebService/rest/prod/ean2/1/"

var codecheckApiInstance *CodecheckApi

type CodecheckApi struct {
    session *codecheckSession
}

type codecheckSession struct {
    session *napping.Session
    validTill time.Time
}

type responseUserAgent struct {
    Useragent string `json:"user-agent"`
}

type sessionResponse struct {
    SessionId string `json: "sessionId"`
    Nonce string `json: "nonce"`
    ExpiresIn int `json: "expiresIn"`
}

type responseWrapper struct {
    Result sessionResponse `json: "result"`
}

func GetCodecheckInstance() *CodecheckApi {
    if (codecheckApiInstance == nil) {
        log.Println("Getting new CodecheckApi instance!")
        codecheckApiInstance = &CodecheckApi{&codecheckSession{}}
    }

    return codecheckApiInstance
}

func (s *codecheckSession) GetRequestSession() *napping.Session {
    if (s.session == nil || s.validTill.Before(time.Now())) {
        apiKey := os.Getenv("CODECHECK_API_SECRET")

        clientNonce := Util.RandString(16)
        decodedClientNonceBytes, err := base64.StdEncoding.DecodeString(clientNonce)
        if (err != nil) {
            log.Fatal(clientNonce)
        }

        decodedClientNonce := string(decodedClientNonceBytes)

        params := napping.Params{"authType": authType, "username" : username, "clientNonce" : clientNonce, "deviceId": deviceId}
        session := napping.Session{nil, false,nil,nil,nil}

        res := responseUserAgent{}
        resp, err := session.Post(codecheck_auth_url, &params, &res, nil)
        if err != nil {
            log.Fatal(err)
        }


        var sessionResult responseWrapper
        err = json.Unmarshal([]byte(resp.RawText()), &sessionResult)
        //fmt.Println("sessionId: " + sessionResult.Result.SessionId, "nonce: " + sessionResult.Result.Nonce, "expiresIn: " + strconv.Itoa(sessionResult.Result.ExpiresIn))

        serverNonceBytes, err := base64.StdEncoding.DecodeString(sessionResult.Result.Nonce)
        if (err != nil) {
            fmt.Println(err.Error())
        }
        serverNonce := string(serverNonceBytes)


        toHash := username + serverNonce + decodedClientNonce

        decodedKey,err := hex.DecodeString(apiKey)
        if (err != nil) {
            log.Print(err.Error())
        }


        hasher := hmac.New(sha256.New, decodedKey)
        hasher.Write([]byte(toHash))
        hashSum := hasher.Sum(nil)
        hash := base64.StdEncoding.EncodeToString(hashSum)


        headerString := "DigestQuick nonce=\"" + sessionResult.Result.Nonce + "\",mac=\"" + hash + "\""


        header := http.Header{}
        header.Add("Authorization", headerString)
        header.Add("Content-Type", "application/json; charset=utf-8")
        res = responseUserAgent{}
        session = napping.Session{nil,false,nil, &header,nil}
        s.session = &session



        duration, err := time.ParseDuration(strconv.Itoa(sessionResult.Result.ExpiresIn) + "s")
        s.validTill = time.Now().Add(duration)

    }

    return s.session
}

func (api CodecheckApi) TestApiCall() string {
    resp, err := api.session.GetRequestSession().Get("http://www.codecheck.info/WebService/rest/prod/ean2/1/401407/4006939082489", nil, nil, nil)
    if (err != nil) {
        log.Fatal(err)
    }

    return resp.RawText()
}

func (api CodecheckApi) GetAllInfoByEan(ean string) string {
    //401407 means all available details
    levelOfDetail := "401407"
    return api.GetInfoByEan(ean,levelOfDetail)
}

func (api CodecheckApi) GetInfoByEan(ean, levelOfDetail string) string {
    url := codecheck_ean_url + levelOfDetail + "/" + ean
    log.Println(url)
    log.Println("ean: " + ean)
    resp, err := api.session.GetRequestSession().Get(url, nil, nil, nil)
    if (err != nil) {
        log.Fatal(err)
    }

    return resp.RawText()
}