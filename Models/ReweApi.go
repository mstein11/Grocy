package Models
import (
    "net/http"
    "fmt"
    "io/ioutil"
)

type ReweApiModel struct {

}

func GetReweInfosForEan(ean string) (string){
    resp, err := http.Get("https://shop.rewe.de/services/search?q=5000112548341&format=json&start=0&rows=10&indent=true")
    if (err != nil) {
        fmt.Print(err)
    }

    byteString, err := ioutil.ReadAll(resp.Body)
    if (err != nil) {
        fmt.Print(err)
    }

    return string(byteString)
}
