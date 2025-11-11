package repos

type LikeRepo interface {
	SetLike(id1, id2 int64) error
	IsLike(id1, id2 int64) bool
	GetAllLikesToUser(id int64) ([]int64, error)
	GetAllLikesFromUser(id int64) ([]int64, error)
}
