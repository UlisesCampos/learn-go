package main

import (
	"encoding/json"
	"fmt"
)

//Server struct
type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

//Serverslice ..
type Serverslice struct {
	Servers []Server `json:"servers"`
}

// func Marshal(b interface{}) ([]byte, error) -> of package json is for produce json
func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Mexico_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Shangai_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err", err)
	}
	fmt.Println(string(b))
	// --- Unmarshal --
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}
