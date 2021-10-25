package models

import (
	"fmt"
	"testing"
)

//
//type Auth struct {
//	ID       int    `gorm:"primary_key" json:"id"`
//	Username string `json:"username"`
//	Password string `json:"password"`
//}

func TestGetTags(t *testing.T) {
	var tag Tag
	id := 1
	db.Where("id = ?", id).First(&tag)
	fmt.Println("tag:", tag)
}
