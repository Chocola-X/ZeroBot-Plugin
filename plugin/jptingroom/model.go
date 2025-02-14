package jptingroom

import (
	"time"

	sql "github.com/FloatTech/sqlite"
)

type item struct {
	ID       int64     `db:"id"`
	Title    string    `db:"title"`
	PageURL  string    `db:"page_url"`
	Category string    `db:"category"`
	Intro    string    `db:"intro"`
	AudioURL string    `db:"audio_url"`
	Content  string    `db:"content"`
	Datetime time.Time `db:"datetime"`
}

var db sql.Sqlite

func getRandomAudioByCategory(category string) (t item) {
	_ = db.Find("item", &t, "WHERE category = ? ORDER BY RANDOM() limit 1", category)
	return
}

func getRandomAudioByCategoryAndKeyword(category string, keyword string) (t item) {
	_ = db.Find("item", &t,
		"WHERE category = ? and (title LIKE ? OR content LIKE ?) ORDER BY RANDOM() limit 1",
		category, "%"+keyword+"%", "%"+keyword+"%")
	return
}
