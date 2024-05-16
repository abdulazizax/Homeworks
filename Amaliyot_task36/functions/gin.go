package functions

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostStudent(c *gin.Context) {
	db := DBConnect()

	man := Manager{
		DB: db,
	}
	var newStudent Student

	if err := c.BindJSON(&newStudent); err != nil {
		log.Fatal(err)
	}

	result, err := man.InsertIntoStudents(newStudent)
	if err != nil {
		c.IndentedJSON(http.StatusNotImplemented, "There is any mistake!")
	}

	c.IndentedJSON(http.StatusCreated, result)
}

func PostCourse(c *gin.Context) {
	db := DBConnect()

	man := Manager{
		DB: db,
	}
	var newCourse Course

	if err := c.BindJSON(&newCourse); err != nil {
		log.Fatal(err)
	}

	result, err := man.InsertIntoCourse(newCourse)
	if err != nil {
		c.IndentedJSON(http.StatusNotImplemented, "There is any mistake!")
	}

	c.IndentedJSON(http.StatusCreated, result)
}

func UpdateCourseByID(c *gin.Context) {
	db := DBConnect()

	man := Manager{
		DB: db,
	}
	id := c.Param("id")

	var newCourse Course

	if err := c.BindJSON(&newCourse); err != nil {
		log.Fatal(err)
	}

	result, err := man.AlterCourseTableById(newCourse, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotImplemented, "There is any mistake!")
	}

	c.IndentedJSON(http.StatusCreated, result)
}

func UpdateStudentByID(c *gin.Context) {
	db := DBConnect()

	man := Manager{
		DB: db,
	}
	id := c.Param("id")

	var newStudent Student

	if err := c.BindJSON(&newStudent); err != nil {
		log.Fatal(err)
	}

	result, err := man.AlterStudentTableById(newStudent, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotImplemented, "There is any mistake!")
		return
	}

	c.IndentedJSON(http.StatusCreated, result)
}

func GetCourses(c *gin.Context) {
	db := DBConnect()

	man := Manager{
		DB: db,
	}

	result, err := man.GetAllCourses()
	if err != nil {
		c.IndentedJSON(http.StatusNotImplemented, "No information!")
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func GetStudents(c *gin.Context) {
	db := DBConnect()

	man := Manager{
		DB: db,
	}

	result, err := man.GetAllStudents()
	if err != nil {
		c.IndentedJSON(http.StatusNotImplemented, "No information!")
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func DeleteCourseByID(c *gin.Context) {
	db := DBConnect()

	man := Manager{
		DB: db,
	}
	id := c.Param("id")

	result, err := man.DeleteFromCourseById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotImplemented, "There is any mistake!")
	}

	c.IndentedJSON(http.StatusCreated, result)
}

func DeleteStudentByID(c *gin.Context) {
	db := DBConnect()

	man := Manager{
		DB: db,
	}
	id := c.Param("id")

	result, err := man.DeleteFromStudentsById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotImplemented, "There is any mistake!")
	}

	c.IndentedJSON(http.StatusCreated, result)
}
