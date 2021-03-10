package Utils

import "strings"

func GetEquipFromTopic(topic string) string {
	topicParts := strings.Split(topic, "/")
	equip := strings.Join([]string{topicParts[0], topicParts[1]}, "/")

	return equip
}
