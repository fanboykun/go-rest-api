package models

import (
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	First_Name string
	Last_Name  string
}

func (p *Person) SavePerson() (*Person, error) {
	var err error
	err = DB.Create(&p).Error
	if err != nil {
		return &Person{}, err
	}
	return p, nil
}

// func (p *Person) GetAllPerson() ([]Person, error) {
// 	var err error
// 	err = DB.Find(&p).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return p, nil
// }

func GetPersonByID(id string) (Person, error) {
	var p Person
	err := DB.Where("id = ?", id).First(&p).Error
	if(err != nil){
		return p, err
	}
	return p, nil
}

func (p *Person) UpdatePerson(uid uint32) (*Person, error) {
	return nil, nil
}

func (p *Person) DeletePerson(uid uint32) (int64, error) {
	var err error
	err = DB.Delete(&p).Error
	if(err != nil){
		return 0, err
	}
	return 0, nil
}