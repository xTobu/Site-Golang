package Controllers

import (
	"net/http"

	"Site-Golang/Models"

	"github.com/gin-gonic/gin"
)

//APIStudent use mysql
func APIStudent(c *gin.Context) {
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

//StudentReq struct
type StudentReq struct {
	Name  string `form:"name" json:"name"`
	Email string `form:"email" json:"email"`
}

//APIInsert use mysql handler 新增一筆學生資料
func APIInsert(c *gin.Context) {
	var studentData StudentReq

	// err := json.NewDecoder(c.Request.Body).Decode(&studentData)
	err := c.Bind(&studentData)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"result": "Error in request",
		})
		return
	}
	// c.Request.FormValue("email")
	if r := Models.DBInsertStudent(studentData.Name, studentData.Email); r == true {
		c.JSON(http.StatusOK, gin.H{
			"result": "success",
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"result": "fail",
		})
		// c.JSON(http.StatusNoContent)
	}
}
