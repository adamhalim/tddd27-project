package postgres

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	port   string
	host   string
	user   string
	dbName string

	db *sqlx.DB
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	port = os.Getenv("POSTGRES_PORT")
	if port == "" {
		log.Fatal("no POSTGRES_PORT in .env")
	}
	host = os.Getenv("POSTGRES_HOST")
	if host == "" {
		log.Fatal("no POSTGRES_HOST in .env")
	}
	user = os.Getenv("POSTGRES_USER")
	if user == "" {
		log.Fatal("no POSTGRES_USER in .env")
	}
	dbName = os.Getenv("POSTGRES_DBNAME")
	if dbName == "" {
		log.Fatal("no POSTGRES_DBNAME in .env")
	}
	initDB()
}

func initDB() {
	connString := fmt.Sprintf("port=%s host=%s user=%s dbname=%s sslmode=disable", port, host, user, dbName)
	database, err := sqlx.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	db = database
	createTables()
}

func createTables() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			uid VARCHAR(50) NOT NULL,
			username VARCHAR(50) NOT NULL UNIQUE,
			CONSTRAINT pk_uid PRIMARY KEY(uid)
		)
	`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS videos (
			chunkname VARCHAR(36) NOT NULL,
			lastViewed NUMERIC(14,0) NOT NULL,
			uid VARCHAR(50),
			FOREIGN KEY(uid) REFERENCES users(uid),
			viewcount INTEGER NOT NULL,
			videotitle VARCHAR(100) NOT NULL,
			originalfilename VARCHAR(100) NOT NULL,
			CONSTRAINT pk_chunkname PRIMARY KEY(chunkname)
		)
	`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS comments (
			id SERIAL primary key,
			chunkname VARCHAR(36) NOT NULL,
			comment VARCHAR(2000) NOT NULL,
			author_uid VARCHAR(50) NOT NULL,
			date NUMERIC(14,0) NOT NULL,
			FOREIGN KEY(chunkname) REFERENCES videos(chunkname),
			FOREIGN KEY(author_uid) REFERENCES users(uid)
		)
	`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS likes (
			id SERIAL primary key,
			chunkname VARCHAR(36) NOT NULL,
			author_uid VARCHAR(50) NOT NULL,
			FOREIGN KEY(chunkname) REFERENCES videos(chunkname),
			FOREIGN KEY(author_uid) REFERENCES users(uid),
			UNIQUE (chunkname, author_uid)
		)
	`)

	if err != nil {
		log.Fatal(err)
	}
}

func AddUser(user User) error {
	stmt, err := db.Prepare(`
		INSERT INTO users(uid, username) VALUES($1, $2)
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Uid, user.Username)
	if err != nil {
		return err
	}
	return nil
}

func FindUser(uid string) (User, error) {
	var user User
	err := db.Get(&user, `SELECT * FROM users WHERE uid=$1`, uid)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func ChangeUsername(uid string, newUsername string) error {
	if newUsername == "" {
		return errors.New("error: empty username")
	}
	stmt, err := db.Prepare(`
		UPDATE users
			SET username = $1
			WHERE uid = $2
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(newUsername, uid)
	if err != nil {
		return err
	}
	return nil
}

func UserExists(uid string) bool {
	_, err := FindUser(uid)
	return err == nil
}

func AddVideo(video Video) error {
	stmt, err := db.Prepare(`
		INSERT INTO videos(
			chunkname,
			lastviewed,
			uid,
			viewcount,
			videotitle,
			originalfilename
			)
			VALUES($1, $2, $3, $4, $5, $6)
	`)
	if err != nil {
		return err
	}

	if video.Uid == "" {
		_, err = stmt.Exec(video.Chunkname, video.LastViewed, nil, video.ViewCount, video.Title, video.OriginalFileName)
	} else {
		_, err = stmt.Exec(video.Chunkname, video.LastViewed, video.Uid, video.ViewCount, video.Title, video.OriginalFileName)
	}
	if err != nil {
		return err
	}
	return nil
}

func DeleteVideo(uid string, chunkName string) error {
	vid, err := FindVideo(chunkName)
	if err != nil {
		return err
	}
	if vid.Uid != uid {
		return errors.New("unathorized user")
	}

	err = deleteCommentsFromVideo(chunkName)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`
		DELETE
			FROM videos
		WHERE
			uid = $1
		AND
			chunkname = $2
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(uid, chunkName)
	if err != nil {
		return err
	}

	return nil
}

func deleteCommentsFromVideo(chunkName string) error {
	stmt, err := db.Prepare(`
		DELETE
			FROM comments
		WHERE
			chunkname = $1
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(chunkName)
	if err != nil {
		return err
	}
	return nil
}

func IncrementViewCount(chunkName string) error {
	stmt, err := db.Prepare(`
		UPDATE videos
			SET viewcount = viewcount + 1
			WHERE chunkname = $1
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(chunkName)
	if err != nil {
		return err
	}

	return nil
}

func FindVideo(chunkName string) (Video, error) {
	var video Video
	err := db.Get(&video, `SELECT * FROM videos WHERE chunkname=$1`, chunkName)
	if err != nil {
		return Video{}, err
	}
	return video, nil
}

func FindVideosFromUser(uid string) ([]video, error) {
	videos := []video{}
	err := db.Select(&videos, `
		SELECT 
			chunkname, viewcount, videotitle
		FROM
			videos
		WHERE
			uid = $1
	`, uid)

	if err != nil {
		return nil, err
	}

	return videos, nil
}

func UpdateLastViewed(chunkName string) error {
	stmt, err := db.Prepare(`
		UPDATE videos
			SET lastviewed = $1
			WHERE chunkname = $2
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(time.Now().UnixMilli(), chunkName)
	if err != nil {
		return err
	}
	return nil
}

func AddComment(chunkName string, comment string, authorUid string) error {
	_, err := FindVideo(chunkName)
	if err != nil {
		return err
	}

	_, err = FindUser(authorUid)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`
		INSERT INTO comments(
			chunkname,
			comment,
			author_uid,
			date
		) VALUES($1, $2, $3, $4)
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(chunkName, comment, authorUid, time.Now().UnixMilli())
	if err != nil {
		return err
	}

	return nil
}

func GetComments(chunkName string) ([]comment, error) {

	comments := []comment{}

	err := db.Select(&comments, `
		SELECT 
			c.comment,  c.date, u.username
		FROM 
			comments AS c
		INNER JOIN users AS u ON
			c.author_uid = u.uid
		WHERE 
			c.chunkname = $1
	
	`, chunkName)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func LikeVideo(chunkName string, authorUid string) error {
	_, err := FindVideo(chunkName)
	if err != nil {
		return err
	}

	_, err = FindUser(authorUid)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`
		INSERT INTO likes(
			chunkname,
			author_uid
		) VALUES($1, $2)
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(chunkName, authorUid)
	if err != nil {
		return err
	}

	return nil
}
