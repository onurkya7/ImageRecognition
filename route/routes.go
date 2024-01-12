package route

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	router *mux.Router
)

func CreateRouter() {
	router = mux.NewRouter()
}

func ServerStarter() {
	fmt.Println("Server is running on port 3000")

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{
			"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization",
		}),
		handlers.AllowedMethods([]string{
			"POST", "GET", "PUT", "DELETE",
		}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	err := http.ListenAndServe(":3000", cors(router))
	if err != nil {
		log.Fatal(err)
	}
}
