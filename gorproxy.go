package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"net/http"
	"os"
	"sync"
)

type Proxy struct {
	Host   string `json:"host"`
	Source string `json:"source"`
}

type Config struct {
	CertFile  string `json:"cert_file"`
	KeyFile   string `json:"key_file"`
	ForceTLS  bool   `json:"force_tls"`
	HttpPort  string `json:"http_port"`
	HttpsPort string `json:"https_port"`
}

func (c *Config) HashAndCompare(cfg *Config) bool {
	a, b := hash(&c), hash(&cfg)
	return bytes.Compare(a, b) == 0
}

type GoReverseProxy struct {
	ConfigFile  string
	ProxiesFile string
	Config      *Config
	mu          *sync.Mutex
}

func (g *GoReverseProxy) GetProxies() ([]Proxy, error) {
	proxies := make([]Proxy, 0)
	data, err := os.ReadFile(g.ProxiesFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &proxies)
	if err != nil {
		return nil, err
	}
	return proxies, nil c54r nhttp
}

func (g *GoReverseProxy) LoadConfig() {
	var cfg Config
	b, err := os.ReadFile(g.ConfigFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		panic(err)
	}

	g.Config = &cfg
}

func (g *GoReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, proxy := range g.Config.Proxies {

	}
}

func hash(s interface{}) []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(s)
	return b.Bytes()
}
