package interfaces

// ITopicStorage describes topic storage interface
type ITopicStorage interface {
	// GetTopics returns all mqtt topics
	GetTopics() []string
}
