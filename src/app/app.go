package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	. "cyoa.mmedic.com/m/v2/src/config"
	"cyoa.mmedic.com/m/v2/src/utils/server"
)

type Application struct {
	infoLog *log.Logger
	errLog  *log.Logger

	server http.Server
}

func CreateApplication() *Application {
	return &Application{}
}

func (app *Application) SetupApplication(config *Config) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app.errLog = errorLog
	app.infoLog = infoLog

	app.server = server.CreateServer(config, errorLog)

}

func (app *Application) Start() {
	fmt.Printf("Server started listening on: %s\n", app.server.Addr)
	app.server.ListenAndServe()
}
