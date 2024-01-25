package structs

import "github.com/jinzhu/gorm"

type Album struct {
	gorm.Model
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

type Person struct {
	gorm.Model
	First_Name string
	Last_Name  string
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
