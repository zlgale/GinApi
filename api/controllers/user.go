package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "github.com/astaxie/beego"
	_ "io/ioutil"
	_ "path"
	_ "runtime"
	"GinApi/api/models/user"
	"fmt"
	"strconv"
	//"github.com/gin-gonic/gin/json"
)

func UserLoginHandler(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  "failed",
		"message": "login failed",
	})
}
func UserLogoutHandler(c *gin.Context) {
	name := c.Query("name")
	message := name + " is logout"
	//sessions.AuthLogout(c)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
	})
}
func CreateUserHandler(c *gin.Context) {

	//uid := c.Query("name")
	//username := c.Query("phone")
	//departname := c.Query("pwd")
	//created := c.Query("gender")
	user_nickname, user_password, user_age, user_sex, user_phone := "大风", "123456", "18", "1", "15555555555"

	var userInfo user.User
	userId := user.UserInsert(user_nickname, user_password, user_age, user_sex, user_phone)
	userInfo = user.UserQueryByNickId(int(userId))
	fmt.Println("CreateUserHandler", userInfo)
	//userJson, err := json.Marshal(&userInfo)
	//checkErr(err)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "用户注册成功！",
		"data": gin.H{
			"user_info": userInfo,
		},
	})
}

func UserListHandler(c *gin.Context) {
	users := user.UserListQuery()
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   users,
		"msg":    "ok",
	})
}
func UserByNameHandler(c *gin.Context) {
	// 获取传入的参数

	// 1、参数放在请求的Body中传递,使用gin.Context中的PostForm方法解析,参数是form中获得,即从Body中获得,忽略URL中的参数
	name0 := c.Request.FormValue("name")
	fmt.Println("name0:", name0)
	name1 := c.PostForm("name")
	fmt.Println("name1:", name1)
	// 2、正常的URL中的参数传递(http://localhost:8888/api/v1?name=大风)
	name2 := c.Query("name")
	fmt.Println("name2:", name2)
	// 3、(http://localhost:8888/api/v1/大风)
	name3 := c.Param("name")
	fmt.Println("name3:", name3)
	name4 := c.Params.ByName("name")
	// String转Int
	fmt.Println("name4:", name4)

	var u user.User
	if len(name0) > 0 {
		name00, err := strconv.Atoi(name0)
		checkErr(err)
		u = user.UserQueryByNickId(name00)
	}
	if len(name2) > 0 {
		name22, err := strconv.Atoi(name0)
		checkErr(err)
		u = user.UserQueryByNickId(name22)
	}
	if len(name3) > 0 {
		name33, err := strconv.Atoi(name0)
		checkErr(err)
		u = user.UserQueryByNickId(name33)
	}
	if len(name4) > 0 {
		name44, err := strconv.Atoi(name0)
		checkErr(err)
		u = user.UserQueryByNickId(name44)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"user":   u,
	})
}

// bind JSON数据
func BindJSONHandler(c *gin.Context) {
	//c.Header("Content-Type", "application/json; charset=utf-8")
	name0 := c.Request.FormValue("name")
	fmt.Println("name0:", name0)
	name := c.PostForm("name")
	if len(name) > 0 {
		name00, err := strconv.Atoi(name0)
		checkErr(err)
		u := user.UserQueryByNickId(name00)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"user":   u,
		})
	}
	//var json user.User
	//// binding JSON,本质是将request中的Body中的数据按照JSON格式解析到json变量中
	////if c.BindWith(&json, binding.JSON) == nil {
	//if c.BindJSON(&json) == nil {
	//	fmt.Println("json.Username--->", json.Username)
	//
	//	if len(json.Username) > 0 {
	//		u := user.UserQueryByUserName(json.Username)
	//		c.JSON(http.StatusOK, gin.H{
	//			"status": http.StatusOK,
	//			"user":   u,
	//		})
	//	}
	//	if json.Username == "大风" {
	//		c.JSON(http.StatusOK, gin.H{"JSON=== status": "you are logged in"})
	//	} else {
	//		c.JSON(http.StatusUnauthorized, gin.H{"JSON=== status": "unauthorized"})
	//	}
	//} else {
	//	c.JSON(404, gin.H{"JSON=== status": "binding JSON error!"})
	//}
}

// 下面测试bind FORM数据
func BindFormHandler(c *gin.Context) {
	//c.Header("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	name0 := c.Request.FormValue("name")
	fmt.Println("name0:", name0)
	name := c.PostForm("name")
	fmt.Println("name1:", name)
	if len(name) > 0 {
		name00, err := strconv.Atoi(name0)
		checkErr(err)
		u := user.UserQueryByNickId(name00)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"user":   u,
		})
	}
	//var form user.User
	//// 本质是将c中的request中的BODY数据解析到form中
	//// 方法一: 对于FORM数据直接使用Bind函数, 默认使用使用form格式解析,if c.Bind(&form) == nil
	//// 方法二: 使用BindWith函数,如果你明确知道数据的类型
	//if c.Bind(&form) == nil {
	//	//if c.BindWith(&form, binding.Form) == nil {
	//	fmt.Println("form.Username--->", form.Username)
	//	if len(form.Username) > 0 {
	//		u := user.UserQueryByUserName(form.Username)
	//		c.JSON(http.StatusOK, gin.H{
	//			"status": http.StatusOK,
	//			"user":   u,
	//		})
	//	}
	//	if form.Username == "大风" {
	//		c.JSON(http.StatusOK, gin.H{"JSON=== status": "you are logged in"})
	//	} else {
	//		c.JSON(http.StatusUnauthorized, gin.H{"JSON=== status": "unauthorized"})
	//	}
	//} else {
	//	c.JSON(404, gin.H{"FORM=== status": "binding FORM error!"})
	//}
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
