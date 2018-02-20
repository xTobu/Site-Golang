package Controllers

import (
	"net/http"

	"Site-Golang/Models"

	"github.com/gin-gonic/gin"
)

//Student use mysql
func Student(c *gin.Context) {
	if res, count := Models.DBGetStudents(); count > 0 {
		n := Models.Students{Students: res}
		c.JSON(http.StatusOK, n)
	} else {
		n := Models.Students{}
		c.JSON(http.StatusOK, n)
		// c.JSON(http.StatusNoContent)
	}
	// if res, count := sqldb.DBGetStudents(); res != nil {
	// 	if cou := count; cou >= 0 {
	// 		n := sqldb.StudentS{res}
	// 		c.JSON(http.StatusOK, n)
	// 	}
	// }
}

//Insert use mysql handler 新增一筆學生資料
func Insert(c *gin.Context) {

	if r := Models.DBInsertStudent(c.Request.FormValue("name"), c.Request.FormValue("email")); r == true {
		c.JSON(http.StatusOK, gin.H{
			"result": "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "fail",
		})
		// c.JSON(http.StatusNoContent)
	}
}
