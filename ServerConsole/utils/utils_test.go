package utils

import (
	"testing"
)

/*// GetEquipFromTopic returns equipment name from topic
func GetEquipFromTopic(topic string) string {
	topicParts := strings.Split(topic, "/")
	if len(topicParts) == 1 {
		return topic
	}

	equip := strings.Join([]string{topicParts[0], topicParts[1]}, "/")

	return equip
}
*/

// TestGetHddNumberFromEquip tests GetHddNumberFromEquip
func TestGetHddNumberFromEquip(t *testing.T) {
	hddNumber := GetHddNumberFromEquip("123_qwerty")
	if hddNumber != "qwerty" {
		t.Errorf(`GetHddNumberFromEquip= %v`, hddNumber)
	}

	hddNumber = GetHddNumberFromEquip("qwerty")
	if hddNumber != "" {
		t.Errorf(`GetHddNumberFromEquip= %v`, hddNumber)
	}
}

// TestGetEquipFromTopic tests GetEquipFromTopic
func TestGetEquipFromTopic(t *testing.T) {
	equip := GetEquipFromTopic("qwerty")
	if equip != "qwerty" {
		t.Errorf(`GetEquipFromTopic= %v`, equip)
	}

	equip = GetEquipFromTopic("qwerty/123_qwe/abc")
	if equip != "qwerty/123_qwe" {
		t.Errorf(`GetEquipFromTopic= %v`, equip)
	}
}
