package main

import (
	"github.com/NatalNW7/gopportunities/config"
	"github.com/NatalNW7/gopportunities/router"
)

var (
	logger *config.Logger
)

func main(){
	logger = config.GetLogger("root")
	err := config.Init()
	if err != nil {
		logger.Errorf("Erro ao initializar as configurações: %v", err)
		return
	}
	router.Initialize()
}