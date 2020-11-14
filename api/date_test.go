package api_test

import (
	"testing"

	"github.com/linuxfreak003/ballistic-calculator-web/api"
	"github.com/stretchr/testify/assert"
)

func TestDate(t *testing.T) {
	t.Run("TestMe Func", func(t *testing.T) {
		assert.Equal(t, 4, api.TestMe(2))
	})
}
