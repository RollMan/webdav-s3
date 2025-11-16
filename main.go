package main

import (
	"net/http"
	"log"
)

var Cfg *Config

func main() {
	var err error
	Cfg, err = LoadConfig()
	if err != nil {
		Logoutput("Unable to load config", "error")
		return
	}

	Logoutput("Webdav server started", "info")
	Logoutput("Log level: "+Cfg.Loglevel, "info")
	webdav := NewWebDAVClient()
	Logoutput("Starting server on port "+Cfg.BindAddress, "info")
	http.Handle("/", webdav)
	log.Fatalln(http.ListenAndServe(Cfg.BindAddress, nil))
}
