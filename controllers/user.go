package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "github.com/astaxie/beego"
	_ "io/ioutil"
	_ "path"
	_ "runtime"
	"GinAPI/models/user"
	"fmt"
)

// @api post /login 用户登录
//
// @apiRequest json
// @apiParam username string 用户名
// @apiParam password string 密码
// @apiExample json
// {
//     "username": "admin",
//     "password": "123456"
// }
//
// @apiSuccess 201 成功登录
// @apiParam expries int 过期的时间，单位秒
// @apiParam token string 存储 token
//
// @apiError 401 账号密码验证错误
func UserLoginHandler(c *gin.Context) {
	//user login
	//name := c.Query("name")
	//pwd := c.Query("pwd")
	//
	//u, _:= user.UserQueryByName(name)
	//
	//valid := u.CheckPassword(pwd)
	//
	//if valid{
	//        sessions.AuthLogin(c, strconv.Itoa(u.Id))
	//        c.JSON(http.StatusOK, gin.H{
	//            "status":  "success",
	//        })
	//    }else{
	//        c.JSON(http.StatusUnauthorized, gin.H{
	//            "status":  "failed",
	//            "message": "login failed",
	//        })
	//    }
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  "failed",
		"message": "login failed",
	})
}

// @Title User Logout
// @API_GROUP User
// @Description 用户登出接口
// @Success 200 {object} 
// @Param   name     query   string false       "user name"
// @Failure 400 no enough input
// @Failure 500 get  common error
// @router /user/logout [get]
func UserLogoutHandler(c *gin.Context) {
	name := c.Query("name")
	message := name + " is logout"
	//sessions.AuthLogout(c)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
	})
}

//@Title User Create
//@API_GROUP User
//@Description 创建用户接口
//@Success 200 {object}
//@Param   name     query   string false       "user name"
//@Param   phone     query   string false       "user name"
//@Param   pwd     query   string false       "user name"
//@Failure 400 no enough input
//@Failure 500 get  common error
//@router /user/create [get]
func CreateUserHandler(c *gin.Context) {

	//uid := c.Query("name")
	//username := c.Query("phone")
	//departname := c.Query("pwd")
	//created := c.Query("gender")
	uid, username, departname, created := "001", "大风", "系统开发部", "2017-07-05 11:50:38"
	success := user.UserInsert(uid, username, departname, created)
	c.JSON(http.StatusOK, gin.H{
		"status":     http.StatusOK,
		"is_created": success,
	})
}

//"userinfo":{
//	"uid":uid,
//	"username":username,
//	"departname":departname,
//	"created":created,
//},
// @Title User Query By ID
// @API_GROUP User
// @Description 查询用户接口通过用户ID
// @Success 200 {object} 
// @Param   uid     query   string false       "user id"
// @Failure 400 no enough input
// @Failure 500 get  common error
// @router /user/query [get]
//func UserQueryByIdHandler(c *gin.Context) {
//
//    suid := c.Query("uid")
//    uid , error := strconv.Atoi(suid)
//    if error != nil {
//        c.JSON(400, gin.H{
//            "status":  "fail",
//            "msg": "字符串转换成整数失败",
//        })
//        return
//    }
//
//    u := user.UserQueryById(uid)
//
//
//    c.JSON(http.StatusOK, gin.H {
//        "status":  "success",
//        "user": u,
//    })
//
//}

//@Title User List
//@API_GROUP User
//@Description 查询用户列表接口
func UserListHandler(c *gin.Context) {
	users := user.UserListQuery()

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"users":  users,
	})
}

// @Title User List
// @API_GROUP User
// @Description 查询个人信息接口
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
	fmt.Println("name4:", name4)

	var u user.User
	if len(name1) > 0 {
		u = user.UserQueryByUserName(name1)
	}
	if len(name2) > 0 {
		u = user.UserQueryByUserName(name2)
	}
	if len(name3) > 0 {
		u = user.UserQueryByUserName(name3)
	}
	if len(name4) > 0 {
		u = user.UserQueryByUserName(name4)
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
		u := user.UserQueryByUserName(name)
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
		u := user.UserQueryByUserName(name)
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
