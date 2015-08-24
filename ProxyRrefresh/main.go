package main

import (
	"./service"
	log "github.com/cihub/seelog"
)

func main(){
	log.Info("Initialize!")

	var proxy  = service.Proxy{}
	var refresh = service.Refresh{}

	log.Info("Run!")

	refresh.Run( &proxy )
}