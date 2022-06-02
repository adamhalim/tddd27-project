package postgres

type User struct {
	Uid string
}

type Video struct {
	Chunkname  string
	LastViewed int64
	Uid        string
	ViewCount  int64
}
