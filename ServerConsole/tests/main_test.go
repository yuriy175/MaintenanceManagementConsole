package tests

import (
	"fmt"

	"testing"
	"ServerConsole/interfaces"
	"ServerConsole/tests/mocks"
)

var mockIoCProvider interfaces.IIoCProvider

func init() {
    mockIoCProvider = mocks.InitMockIoc()

	mockIoCProvider.GetEquipsService().Start()

	fmt.Println("Test Go inited")
}

func TestMain(t *testing.T) {
	
}
