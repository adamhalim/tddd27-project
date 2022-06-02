package postgres

type User struct {
	Uid string
}

type Video struct {
	Chunkname  string `db:"chunkname"`
	LastViewed int64  `db:"lastviewed"`
	Uid        string `db:"uid"`
	ViewCount  int64  `db:"viewcount"`
}
