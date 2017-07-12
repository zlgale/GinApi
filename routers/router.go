package routers


import (
    _ "github.com/astaxie/beego"
    "net/http"
    "fmt"
    "GinAPI/controllers"
    "github.com/gin-gonic/gin"
)


func InitRouter() http.Handler {
    // 获得路由实例
    router := gin.New()
    //添加中间件
    router.Use(gin.Logger())
    router.Use(gin.Recovery())
    // TODO:注意':'必须要匹配,'*'选择匹配,即存在就匹配,否则可以不考虑
    v1 := router.Group("api/v1")
    {
        v1.POST("/user/add", controllers.CreateUserHandler)
        // 1、参数放在请求的Body中传递
        v1.POST("/user/byname1", controllers.UserByNameHandler)
        // 2、正常的URL中的参数传递(http://localhost:8888/api/v1/user?name=大风)
        v1.POST("/user/byname2", controllers.UserByNameHandler)
        // 3、(http://localhost:8888/api/v1/user/name=大风)
        v1.POST("/user/byname3/:name", controllers.UserByNameHandler)
        // 4、binding数据，gin内置了几种数据的绑定例如JSON, XML。简单来说, 即根据Body数据类型, 将数据赋值到指定的结构体变量中. (类似于序列化和反序列化)
        // 4.1、bind JSON 数据。对应客户端：application/json
        v1.POST("/user/login_json", controllers.BindJSONHandler)
        // 4.2、bind Form 数据。对应客户端：application/x-www-form-urlencoded
        v1.POST("/user/login_form", controllers.BindFormHandler)

        v1.GET("/user/list", controllers.UserListHandler)
        v1.POST("/user/list2", controllers.UserListHandler)
    }
    fmt.Println("[Plugin Router Profile]...")
    return router
}