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

func BenchmarkSub(b *testing.B) {
	b.Run("Arya", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Arya")
		}
	})
	b.Run("Rizki", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Arya")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Arya")
	}
}

func BenchmarkHelloWorldRizki(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Rizki")
	}
}

func TestTableHelloWorld(t *testing.T) {
	test := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Arya",
			request:  "Arya",
			expected: "Hello Arya",
		},
		{
			name:     "Rizki",
			request:  "Rizki",
			expected: "Hello Rizki",
		},
		{
			name:     "Andaru",
			request:  "Andaru",
			expected: "Hello Andaru",
		},
		{
			name:     "Acumalaka",
			request:  "Acumalaka",
			expected: "Hello Acumalaka",
		},
		{
			name:     "Kumalala",
			request:  "Kumalala",
			expected: "Hello Kumalala",
		},
	}

	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Arya", func(t *testing.T) {
		result := HelloWorld("Arya")
		require.Equal(t, "Hello Arya", result, "Result must be 'Hello Arya'")
	})
	t.Run("Rizki", func(t *testing.T) {
		result := HelloWorld("Rizki")
		require.Equal(t, "Hello Rizki", result, "Result must be 'Hello Rizki'")
	})
}

func TestMain(m *testing.M) {
	// before test
	fmt.Println("Before Unit Test")

	m.Run()

	// after test
	fmt.Println("After Unit Test")
}

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
