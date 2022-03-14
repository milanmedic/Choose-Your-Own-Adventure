package server

import (
	"fmt"
	"log"
	"net/http"

	. "cyoa.mmedic.com/m/v2/src/config"
	routes "cyoa.mmedic.com/m/v2/src/routes"
)

func CreateServer(config *Config, errLog *log.Logger) http.Server {
	mux := CreateMux(config)

	server := http.Server{
		Addr:     fmt.Sprintf("localhost:%d", config.PORT),
		Handler:  mux,
		ErrorLog: errLog,
	}

	return server
}

func CreateMux(config *Config) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HelloWorld)
	routes.SetupRoutes(mux, config)

	return mux
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/intro", http.StatusFound)
}
