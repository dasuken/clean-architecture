package main
import (
	"github.com/noguchidaisuke/clean-architecture/infrastructure"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", infrastructure.Router))
}
