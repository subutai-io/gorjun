package main

import (
	"net/http"
	"github.com/subutai-io/agent/log"
	"github.com/subutai-io/cdn/config"
	"github.com/subutai-io/cdn/db"
	"github.com/subutai-io/cdn/server/app"
)

var (
	srv *http.Server
)

func init() {
	app.InitFilters()
	app.InitValidators()
}

// main starts CDN server
func main() {
	defer db.Close()
	log.Info("Server has started. " + "Listening at " + "127.0.0.1:8080")
	http.HandleFunc("/kurjun/rest/apt/info", app.FileSearch)
	http.HandleFunc("/kurjun/rest/apt/list", app.FileSearch)
	http.HandleFunc("/kurjun/rest/raw/info", app.FileSearch)
	http.HandleFunc("/kurjun/rest/raw/list", app.FileSearch)
	http.HandleFunc("/kurjun/rest/template/info", app.FileSearch)
	http.HandleFunc("/kurjun/rest/template/list", app.FileSearch)
	srv = &http.Server{
		Addr:    ":" + config.Network.Port,
		Handler: nil,
	}
	srv.ListenAndServe()
}
