package postgres

import (
	"fmt"
	"log"
	"os"

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
			CONSTRAINT pk_uid PRIMARY KEY(uid)
		)
	`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS videos (
			chunkname VARCHAR(36) NOT NULL,
			lastViewed NUMERIC(11,0) NOT NULL,
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
}

func AddUser(user User) error {
	stmt, err := db.Prepare(`
		INSERT INTO users(uid) VALUES($1)
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Uid)
	if err != nil {
		return err
	}
	return nil
}

func FindUser(uid string) User {
	var user User
	row := db.QueryRow(`SELECT uid FROM users WHERE uid=$1`, uid)
	err := row.Scan(&user.Uid)
	if err != nil {
		return User{}
	}
	return user
}

func UserExists(uid string) bool {
	user := FindUser(uid)
	return user != User{}
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
