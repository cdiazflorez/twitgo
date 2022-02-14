package main

import (
	"log"

	"github.com/cdiazflorez/twitgo/bd"
	"github.com/cdiazflorez/twitgo/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}

	handlers.Manejadores()

}
