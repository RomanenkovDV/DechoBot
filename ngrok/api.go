package ngrok

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type TunnelData struct {
	Tunnels []Tunnel `json:"tunnels"`
}

type Tunnel struct {
	Proto     string `json:"proto"`
	PublicUrl string `json:"public_url"`
}

func parseTunnelsResponse(resp *http.Response) (string, bool) {
	status := false
	tunnerUrl := ""
	if resp.StatusCode == http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		var tunnelData TunnelData
		json.Unmarshal(data, &tunnelData)
		for _, tunnel := range tunnelData.Tunnels {
			if tunnel.Proto == "https" {
				status = tunnel.PublicUrl != ""
				tunnerUrl = tunnel.PublicUrl
			}
		}
	}

	return tunnerUrl, status
}

func FetchTunnelAddress(webInterfaceUrl string) (string, bool) {
	status := false
	tunnelUrl := ""
	resp, err := http.Get(webInterfaceUrl + "/api/tunnels")
	if resp != nil {
		defer resp.Body.Close()
		tunnelUrl, status = parseTunnelsResponse(resp)
	} else {
		log.Println("Can not fetch tunnel address")
		if err != nil {
			log.Println(err)
		}
	}

	return tunnelUrl, status
}
