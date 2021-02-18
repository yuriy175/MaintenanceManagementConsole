package BL

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type topicsContent struct {
	Topics []string
}

type ITopicStorage interface {
	getTopics() []string
}

type TopicStorage struct{}

func (t *TopicStorage) getTopics() []string {
	data, err := ioutil.ReadFile("topics.json")
	var content topicsContent
	json.Unmarshal(data, &content)

	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	return content.Topics
}
