package utils

import (
	"encoding/json"
	"io/ioutil"

	"ServerConsole/interfaces"
)

// topic storage implementation type
type topicStorage struct {
	//logger
	_log interfaces.ILogger

	//topics
	Topics []string
}

// TopicStorageNew creates an instance of topicStorage
func TopicStorageNew(log interfaces.ILogger, filePath string) interfaces.ITopicStorage {
	data, err := ioutil.ReadFile(filePath) // "topics.json")
	var storage topicStorage
	json.Unmarshal(data, &storage)

	if err != nil {
		log.Error("failed reading data from storage file")
	}

	storage._log = log
	return &storage
}

// GetTopics returns all mqtt topics
func (t *topicStorage) GetTopics() []string {
	return t.Topics
}
