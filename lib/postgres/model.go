package postgres

type User struct {
	Uid      string `db:"uid"`
	Username string `db:"username"`
}

type Video struct {
	Chunkname        string `db:"chunkname"`
	LastViewed       int64  `db:"lastviewed"`
	Uid              string `db:"uid"`
	ViewCount        int64  `db:"viewcount"`
	Title            string `db:"videotitle"`
	OriginalFileName string `db:"originalfilename"`
}

type Comment struct {
	Id        string `db:"id"`
	Chunkname string `db:"chunkname"`
	Comment   string `db:"comment"`
	AuthorUid string `db:"author_uid"`
	Date      int64  `db:"date"`
}

type comment struct {
	Comment  string `db:"comment"`
	Date     int64  `db:"date"`
	Username string `db:"username"`
}
