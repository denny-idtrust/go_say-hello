package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestTableHellorlf(t *testing.T) {
	test := []struct {
		name     string
		status   string
		request  string
		expected string
	}{
		{
			name:     "John",
			status:   "Active",
			request:  "Hi Hello John Active",
			expected: "Hi Hello John Active",
		},
	}
}

func TestSubTest(t *testing.T) {
	t.Run("John", func(t *testing.T) {
		result := SayHello("John", "Active")
		assert.Equal(t, "Hi Hello John Active", result)
	})

	t.Run("Denny", func(t *testing.T) {
		result := SayHello("Denny", "Active")
		assert.Equal(t, "Hi Hello Denny Active", result)
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")
	m.Run()
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}
	result := SayHello("John", "Active")
	assert.Equal(t, "Hi Hello John Active", result, "Result must be Hi Hello John Active")

}

func TestHelloJohnAssert(t *testing.T) {
	result := SayHello("John", "Active")
	assert.Equal(t, "Hi Hello John Active", result, "Result must be Hi Hello John Active")
	fmt.Println("TestHelloJohnAssert is done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := SayHello("John", "Active")
	require.Equal(t, "Hi Hello John Active", result, "Result must be Hi Hello John Active")
	fmt.Println("TestHelloWorldRequire is done")
}

func TestSayHello(t *testing.T) {
	result := SayHello("John", "Active")
	assert.Equal(t, "Hi Hello John Active", result)
	fmt.Println("TestHelloJohn is done")
}

func TestSayHelloDenny(t *testing.T) {
	result := SayHello("John", "Active")
	if result != "Hi Hello John Active" {
		t.Fatalf("Result must be Hi Hello John Active, but got %s", result)
	}
	fmt.Println("TestHelloDenny is done")
}
