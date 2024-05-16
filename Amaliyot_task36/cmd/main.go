package main

import (
	funk "amaliyot/functions"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/courses", funk.GetCourses)
	router.GET("/students", funk.GetStudents)

	router.POST("/student", funk.PostStudent)
	router.POST("/course", funk.PostCourse)

	router.PUT("/course/:id", funk.UpdateCourseByID)
	router.PUT("/student/:id", funk.UpdateStudentByID)

	router.DELETE("/student/:id", funk.DeleteStudentByID)
	router.DELETE("/course/:id", funk.DeleteCourseByID)

	router.Run("localhost:8080")
}
