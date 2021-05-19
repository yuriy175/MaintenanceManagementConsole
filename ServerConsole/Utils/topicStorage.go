package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"../interfaces"
)

// topic storage implementation type
type topicStorage struct {
	Topics []string
}

// TopicStorageNew creates an instance of topicStorage
func TopicStorageNew() interfaces.ITopicStorage {
	data, err := ioutil.ReadFile("topics.json")
	var storage topicStorage
	json.Unmarshal(data, &storage)

	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	return &storage
}

// GetTopics returns all mqtt topics
func (t *topicStorage) GetTopics() []string {
	return t.Topics
}
