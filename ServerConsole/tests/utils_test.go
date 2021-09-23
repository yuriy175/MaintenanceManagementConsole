package tests

import (
	"testing"

	"ServerConsole/utils"
)


// TestGetHddNumberFromEquip tests GetHddNumberFromEquip
func TestGetHddNumberFromEquip(t *testing.T) {
	hddNumber := utils.GetHddNumberFromEquip("123_qwerty")
	if hddNumber != "qwerty" {
		t.Errorf(`GetHddNumberFromEquip= %v`, hddNumber)
	}

	hddNumber = utils.GetHddNumberFromEquip("qwerty")
	if hddNumber != "" {
		t.Errorf(`GetHddNumberFromEquip= %v`, hddNumber)
	}
}

// TestGetEquipFromTopic tests GetEquipFromTopic
func TestGetEquipFromTopic(t *testing.T) {
	equip := utils.GetEquipFromTopic("qwerty")
	if equip != "qwerty" {
		t.Errorf(`GetEquipFromTopic= %v`, equip)
	}

	equip = utils.GetEquipFromTopic("qwerty/123_qwe/abc")
	if equip != "qwerty/123_qwe" {
		t.Errorf(`GetEquipFromTopic= %v`, equip)
	}
}
