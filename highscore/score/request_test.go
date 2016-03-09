package score

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScoreRequest(t *testing.T) {
	assert.Equal(t, NewScoreRequest(), ScoreRequest{})
}
