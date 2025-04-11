package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
	fail
	failNow
	error
	fatal
*/

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Not Supported on Windows")
	}

	result := HelloWorld("Kumalala")
	require.Equal(t, "Hello Kumalala", result, "Result must be 'Hello Kumalala'")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Acumalaka")
	assert.Equal(t, "Hello Acumalaka", result, "Result must be 'Hello Acumalaka'")
	fmt.Println("TestHelloWorld wth Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Acumalaka")
	require.Equal(t, "Hello Acumalaka", result, "Result must be 'Hello Acumalaka'")
	fmt.Println("TestHelloWorld wth Require Done")
}

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
