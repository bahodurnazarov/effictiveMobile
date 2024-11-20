package main

import (
	"github.com/bahodurnazarov/effictiveMobile/internal/app"
	"github.com/bahodurnazarov/effictiveMobile/pkg/db"
	"github.com/bahodurnazarov/effictiveMobile/pkg/utils"
	"log"
)

func main() {
	cfg, err := utils.Init()
	if err != nil {
		log.Fatalln("Failed init cfg", err.Error())
	}
	a := app.New(cfg)
	defer db.CloseDB()
	log.Fatalln(a.Run(cfg))
}
