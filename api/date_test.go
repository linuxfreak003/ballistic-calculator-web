package api_test

import (
	"testing"

	handler "github.com/linuxfreak003/ballistic-calculator-web/api"
	"github.com/stretchr/testify/assert"
)

func TestDate(t *testing.T) {
	t.Run("TestMe Func", func(t *testing.T) {
		assert.Equal(t, 4, handler.TestMe(2))
	})
}
