package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTableHello(b *testing.B) {
	benchmark := []struct {
		name     string
		status   string
		expected string
	}{
		{
			name:   "John",
			status: "Active",
		},
		{
			name:   "Denny",
			status: "Active",
		},
	}
	for _, benchmark := range benchmark {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHello(benchmark.name, benchmark.status)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("John", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("John", "Active")
		}
	})
	b.Run("Denny", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Denny", "Active")
		}
	})

}

func BenchmarkSayHello(b *testing.B) {

	for i := 0; i < b.N; i++ {
		SayHello("John", "Active")
	}
}

func TestTableHello(t *testing.T) {
	test := []struct {
		name     string
		status   string
		expected string
	}{
		{
			name:     "John",
			status:   "Active",
			expected: "Hi Hello John Active",
		},
		{
			name:     "Denny",
			status:   "Active",
			expected: "Hi Hello Denny Active",
		},
	}
	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			result := SayHello(test.name, test.status)
			require.Equal(t, test.expected, result)
		})
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
