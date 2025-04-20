package zonotools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExampleFunction(t *testing.T) {
	t.Run("normal case", func(t *testing.T) {
		result := ExampleFunction("hello", "world", " ")
		assert.Equal(t, "hello world", result)
	})

	t.Run("with different separator", func(t *testing.T) {
		result := ExampleFunction("hello", "world", "-")
		assert.Equal(t, "hello-world", result)
	})

	t.Run("with empty strings", func(t *testing.T) {
		result := ExampleFunction("", "", ",")
		assert.Equal(t, ",", result)
	})
}
