package helper

import (
	"fmt"
	"testing"
)

/*
	fail
	failNow
	error
	fatal
*/

func TestHelloWorldArya(t *testing.T) {
	result := HelloWorld("Arya")
	if result != "Hello Arya" {
		// error
		t.Error("Result must be 'Hello Arya'")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldAndaru(t *testing.T) {
	result := HelloWorld("Andaru")
	if result != "Hello Andaru" {
		// error
		t.Fatal("Result must be 'Hello Andaru'")
	}

	fmt.Println("TestHelloWorldAndaru Done")
}
