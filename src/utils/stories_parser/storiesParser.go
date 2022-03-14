package storiesparser

import (
	"encoding/json"

	. "cyoa.mmedic.com/m/v2/src/models/story"
)

type StoriesParser struct{}

func CreateStoriesParser() *StoriesParser {
	return &StoriesParser{}
}

func (sp *StoriesParser) ParseIntoStories(contents []byte) (map[string]Story, error) {
	var stories map[string]Story = make(map[string]Story)

	err := json.Unmarshal(contents, &stories)

	if err != nil {
		return nil, err
	}

	return stories, nil
}
