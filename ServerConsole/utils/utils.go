package utils

import "strings"

// GetEquipFromTopic returns equipment name from topic
func GetEquipFromTopic(topic string) string {
	topicParts := strings.Split(topic, "/")
	if len(topicParts) == 1{
		return topic
	}

	equip := strings.Join([]string{topicParts[0], topicParts[1]}, "/")

	return equip
}
