package controllers

import (
	"net/http"

	"rest-api/models"

	"github.com/gin-gonic/gin"
)

type GetPersonRequest struct {
	ID uint32 `json:"id"`
}
func GetPerson(c *gin.Context) {
	var (
		person models.Person
		result gin.H
		// request GetPersonRequest
	)
	// id := request.ID
	id := c.Param("id")
	person,err := models.GetPersonByID(id)
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// func GetPersons(c *gin.Context) {
	// var (
	// 	persons []models.Person
	// 	result  gin.H
	// 	person  models.Person
	// )

	// persons, err := person.GetAllPerson()
	// if(err != nil) {
	// 	result = gin.H{
	// 		"result": nil,
	// 		"count":  0,
	// 	}
	// 	c.JSON(http.StatusInternalServerError, result)
	// 	c.Abort()
	// 	return
	// }
	// if len(persons) <= 0 {
	// 	result = gin.H{
	// 		"result": nil,
	// 		"count":  0,
	// 	}
	// } else {
	// 	result = gin.H{
	// 		"result": persons,
	// 		"count":  len(persons),
	// 	}
	// }
	// c.JSON(http.StatusOK, result)
// }

type CreatePersonRequest struct {
	First_Name string `json:"username"`
	Last_Name string `json:"password"`
}
func CreatePerson(c *gin.Context) {
	var (
		person models.Person
		input CreatePersonRequest
		result gin.H
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person.First_Name = input.First_Name
	person.Last_Name = input.Last_Name

	_,err := person.SavePerson()

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

type DeletePersonRequest struct {
	ID uint32 `json:"id"`
}
func DeletePerson(c *gin.Context) {
	var (
		person models.Person
		// request DeletePersonRequest
		result gin.H
	)
	
	// if err := c.ShouldBindJSON(&person); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// id := request.ID
	id := c.Param("id")
	person,err := models.GetPersonByID(id)
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	_,err = person.DeletePerson(uint32(person.ID))
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
