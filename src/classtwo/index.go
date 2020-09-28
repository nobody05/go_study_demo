package classtwo


import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"net/http"
)

/**
 * @Author: gaoz
 * @Date: 2020/9/28
 */

// 这里用的gorm包 可以查看官方文档
var connectDB *gorm.DB
var err error

func HttpStart() {
	connectDB, err = InitDB()
	if err != nil {
		panic(err.Error())
		return
	}
	defer connectDB.Close()


	// 处理登录
	http.HandleFunc("/login", Login)
	// 处理注册
	http.HandleFunc("/register", Register)

	errs := http.ListenAndServe("127.0.0.1:9210", nil)
	if errs != nil {
		fmt.Println("server listen error"+ errs.Error())
		return
	}



	fmt.Println("hello go")
}

func Login(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("request login")

}

func Register(writer http.ResponseWriter, request *http.Request) {
	// 逻辑  请求的json数据转为model 、insert到表中
	fmt.Println("request register")
	// 从请求body解析数据->转为字节数组
	bytes,_ := ioutil.ReadAll(request.Body)
	// & 取地址符，指向内存地址
	registerUser := &RegisterRequest{}
	// 字符串转为对象 类似于 json_decode， unmarshal会直接更新到对象
	jsonErr := json.Unmarshal(bytes, registerUser)
	if jsonErr != nil {
		fmt.Println("json err "+ jsonErr.Error())
		return
	}

	// 获取到用户名
	fmt.Println(registerUser.Username)
	fmt.Println(registerUser.Password)

	// 构建插入对象
	user := GoUser{
		Username:registerUser.Username,
		Password:MD5Password(registerUser.Password),
		Status:1,
	}

	// 插入成功
	insertId := user.CreateUser(connectDB, user)
	if insertId == 0 {
		fmt.Println("insert error")
		return
	}

	writer.Header().Set("content-type", "application/json")
	responseText, err := json.Marshal(SuccessResponse(user, "success"))
	if err != nil {
		fmt.Println("jsonmarshal err"+ err.Error())
		responseText, _ = json.Marshal(FailResponse("jsonmarshal err"+ err.Error()))
	}
	writer.Write(responseText)
}



/**
连接数据库
*/
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		// 多返回值
		return nil, err
	}

	return db, nil
}
