package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Arya")
	if result != "Hello Arya" {
		// error
		panic("Test failed, Result is not 'Hello Arya'")
	}
}

func TestHelloWorldAndaru(t *testing.T) {
	result := HelloWorld("Andaru")
	if result != "Hello Andaru" {
		// error
		panic("Test failed, Result is not 'Hello Andaru'")
	}
}
