package helpers

import (
	"backend/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gorm.io/gorm"
)

var dbConnection *gorm.DB

//MakeRequest ...
func MakeRequest(url string) ([]byte, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return []byte(body), err
}

//ParseJSON ...
func ParseJSON(body []byte) (*models.CatResponse, error) {
	var s = new(models.CatResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

// GetDB ...
type GetDB struct {
	db *gorm.DB
}

//GetCategory ...
func GetCategory(url string) {
	// var conn *GetDB
	body, _ := MakeRequest(url)
	s, _ := ParseJSON(body)
	log.Print(s)

	// conn.db.Create(&s)
}

func Save(category models.Category) error {
	if err := dbConnection.Create(&category).Error; err != nil {
		return err
	}
	return nil
}
