package score

type ScoreRequest struct {
	Initials string
}

func NewScoreRequest() ScoreRequest {
	return ScoreRequest{}
}
