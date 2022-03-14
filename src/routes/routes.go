package routes

import (
	"net/http"

	. "cyoa.mmedic.com/m/v2/src/config"
	filereader "cyoa.mmedic.com/m/v2/src/utils/file_reader"
	storiesparser "cyoa.mmedic.com/m/v2/src/utils/stories_parser"
)

func SetupRoutes(mux *http.ServeMux, config *Config) *http.ServeMux {
	fr := filereader.CreateFileReader()
	sp := storiesparser.CreateStoriesParser()

	fileContents, err := fr.ReadFile(config.GetFilename())

	if err != nil {
		panic(err)
	}

	stories, err := sp.ParseIntoStories(fileContents)

	if err != nil {
		panic(err)
	}

	for title, story := range stories {
		storyCopy := story
		mux.Handle("/"+title, &storyCopy)
	}

	return mux
}
