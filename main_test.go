package main

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func Test_store(t *testing.T) {
	t.Run("正常にinsertされること", func(t *testing.T) {
		args := Post{
			Title: "title",
			Body:  "body",
			Media: `{"mime_type":"image/png","file_name":"test.png"}`,
		}

		store(args)

		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/demo?parseTime=true&loc=Asia%2FTokyo")
		if err != nil {
			log.Println(err)
			return
		}

		t.Cleanup(func() {
			_, err := db.Exec("DELETE FROM posts WHERE title = ?", args.Title)
			if err != nil {
				t.Errorf("failed to clean up test data: %v", err)
			}
			db.Close()
		})

		var got Post

		query := "SELECT title, body, media FROM posts WHERE title = ?"
		err = db.QueryRow(query, args.Title).Scan(&got.Title, &got.Body, &got.Media)

		assert.Equal(t, args.Title, got.Title)
		assert.Equal(t, args.Body, got.Body)
		assert.JSONEq(t, args.Media, got.Media)
	})
}
