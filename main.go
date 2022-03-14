package main

import (
	"flag"

	. "cyoa.mmedic.com/m/v2/src/app"
	"cyoa.mmedic.com/m/v2/src/config"
)

func main() {
	port := flag.Int("PORT", 3000, "PORT that the server will listen on.")
	filename := flag.String("Filename", "story.json", "File contatining the stories for the application")

	flag.Parse()

	app := CreateApplication()

	config := config.CreateConfig()
	config.SetFilename(*filename)
	config.SetPORT(*port)

	app.SetupApplication(config)
	app.Start()
}
