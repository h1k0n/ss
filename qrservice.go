package main

import (
        "encoding/json"
"strconv"
"fmt"
        "log"
        "net/http"
"image/png"
"io/ioutil"
        "time"

        "github.com/go-chi/chi"
        "github.com/boombuler/barcode/qr"
        "github.com/boombuler/barcode"
        guuid "github.com/google/uuid"
)

const format = `[Interface]
PrivateKey= %s
Address = %s
DNS = 8.8.8.8

[Peer]
PublicKey = %s
Endpoint = %s
AllowedIPs = 0.0.0.0/0
`

type RespS struct {
        ServerPort      int
        ServerPublicKey  string
        ClientPrivateKey string

        PeerIP string
}

func main() {
        tmpclient := http.Client{
                Timeout: 10 * time.Second,
        }
        r := chi.NewRouter()
        r.Get("/", func(w http.ResponseWriter, r *http.Request) {
                req, err := http.NewRequest("GET", "http://18.181.47.5:8801/v1/peers/config", nil)
                uuid := guuid.New()
                req.Header.Add("UUID", uuid.String())
                req.Header.Add("Authentication", "zhimakaimen")
                resp, err := tmpclient.Do(req)
                if err != nil {
                        log.Fatal(err)
                }
                defer resp.Body.Close()

                responseData, err := ioutil.ReadAll(resp.Body)
                var respS RespS
                json.Unmarshal(responseData, &respS)
                if err != nil {
                        log.Fatal(err)
                }
                dataString := fmt.Sprintf(format, respS.ClientPrivateKey, respS.PeerIP, respS.ServerPublicKey, "18.181.47.5:"+strconv.Itoa(respS.ServerPort))
                log.Println("config",dataString)
                qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
                qrCode, _ = barcode.Scale(qrCode, 512, 512)

                png.Encode(w, qrCode)
        })
        http.ListenAndServe(":3000", r)

}
