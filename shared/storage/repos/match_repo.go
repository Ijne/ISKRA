package repos

type MatchRepo interface {
	SetMatch(id1, id2 int64) error
}
