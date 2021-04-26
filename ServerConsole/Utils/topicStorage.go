package Utils

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"../Interfaces"
)

type topicStorage struct {
	Topics []string
}

func TopicStorageNew() Interfaces.ITopicStorage {
	data, err := ioutil.ReadFile("topics.json")
	var storage topicStorage
	json.Unmarshal(data, &storage)

	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	return &storage
}

func (t *topicStorage) GetTopics() []string {
	return t.Topics
}
