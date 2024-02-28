package scheduler

type Scheduled interface {
	SelectCandidateNodes()
	Score()
	Pick()
}
