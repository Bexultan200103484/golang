package main
import (
	"log"
	"github.com/Bexultan200103484/golang/goproject"
	"github.com/Bexultan200103484/golang/goproject/pkg/handler"
)
func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
	
}