package user

import (
	"GinApi/api/db"
	"fmt"
)

type User struct {
	user_id       string `form:"user_id" json:"user_id"`
	user_realname string `form:"user_realname" json:"user_realname"`
	user_nickname string `form:"user_nickname" json:"user_nickname"`
	user_password string `form:"user_password" json:"user_password"`
	user_age      int    `form:"user_age" json:"user_age"`
	user_sex      int    `form:"user_sex" json:"user_sex"`
	user_adress   string `form:"user_adress" json:"user_adress"`
	user_phone    string `form:"user_phone" json:"user_phone"`
	user_qq       int    `form:"user_qq" json:"user_qq"`
	user_wechat   string `form:"user_wechat" json:"user_wechat"`
}


// 用户注册
func UserInsert(user_nickname, user_password, user_age, user_sex, user_phone string) bool {
	//db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	//checkErr(err)
	stmt, err := db.DB.Prepare(`INSERT into user_info (user_nickname,user_password,user_age,user_sex,user_phone) values (?,?,?,?,?)`)
	checkErr(err)
	res, err := stmt.Exec(user_nickname, user_password, user_age, user_sex, user_phone)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	return true
}

// 根据昵称查询
func UserQueryByNickName(nickname string) User {
	rows := (db.DB).QueryRow(`SELECT * FROM user_info where user_nickname = ?`, nickname)
	////普通demo
	var u User
	var user_id, user_nickname string
	var user_age, user_sex int
	err := rows.Scan(&user_id, &user_nickname, &user_age, &user_sex)
	checkErr(err)
	u = User{
		user_id:        user_id,
		user_nickname:   user_nickname,
		user_age: user_age,
		user_sex:    user_sex,
	}
	return u
}

// 查询所有用户
func UserListQuery() (users []User) {
	rows, err := (db.DB).Query(`SELECT * FROM user_info`)
	checkErr(err)
	var u User
	var arr []User
	for rows.Next() {
		var user_id string
		var user_nickname string
		var user_age int
		var user_sex int

		rows.Columns()
		err = rows.Scan(&user_id, &user_nickname, &user_age, &user_sex)
		checkErr(err)
		u = User{
			user_id:        user_id,
			user_nickname:   user_nickname,
			user_age: user_age,
			user_sex:    user_sex,
		}
		arr = append(arr, u)
	}
	return arr
}

// 昵称修改
func UpdateUserNickName() {
	//stmt, err := db.DB.Prepare(`UPDATE user_info SET user_nickname=?,user_sex=? WHERE user_id=?`)
	stmt, err := db.DB.Prepare(`UPDATE user_info SET user_nickname=?`)
	checkErr(err)
	res, err := stmt.Exec(21, 2, 1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

// 删除数据
func UserRemove() {
	stmt, err := db.DB.Prepare(`DELETE FROM user_info WHERE user_phone=?`)
	checkErr(err)
	res, err := stmt.Exec(1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
