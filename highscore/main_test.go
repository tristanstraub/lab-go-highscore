package main

import (
	jsn "allthethings/highscore/json"
	"allthethings/highscore/score"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonDecode(t *testing.T) {
	r := []byte("{}")
	scoreRequest := jsn.Decode(r, func() interface{} { return score.NewScoreRequest() }).(score.ScoreRequest)

	assert.Equal(t, scoreRequest, score.ScoreRequest{})
}
