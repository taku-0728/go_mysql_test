package main

import (
	"bufio"
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Title string
	Body  string
	Media string
}

func main() {
	var argLimit int = 3
	args := make([]string, 0, 3)
	s := bufio.NewScanner(os.Stdin)

	for i := 0; i < argLimit; i++ {
		s.Scan()
		args = append(args, s.Text())
	}

	post := Post{
		Title: args[0],
		Body:  args[1],
		Media: args[2],
	}

	store(post)
}

func store(p Post) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/demo?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	query := "INSERT INTO posts (title, body, media) VALUE(?, ?, ?)"
	_, err = db.Exec(query, p.Title, p.Body, p.Media)
	if err != nil {
		log.Println(err)
	}
	return
}
