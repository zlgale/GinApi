package user

import (
	"GinApi/api/db"
	"fmt"
)

type User struct {
	UserId       *int `form:"user_id" json:"user_id"`
	UserRealname *string `form:"user_realname" json:"user_realname"`
	UserNickname *string `form:"user_nickname" json:"user_nickname"`
	UserPassword *string `form:"user_password" json:"user_password"`
	UserAge      *int    `form:"user_age" json:"user_age"`
	UserSex      *int    `form:"user_sex" json:"user_sex"`
	UserAdress   *string `form:"user_adress" json:"user_adress"`
	UserPhone    *string `form:"user_phone" json:"user_phone"`
	UserQQ       *int    `form:"user_qq" json:"user_qq"`
	UserWeChat   *string `form:"user_wechat" json:"user_wechat"`
}

// 用户注册
func UserInsert(user_nickname, user_password, user_age, user_sex, user_phone string) int64 {
	//db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	//checkErr(err)
	stmt, err := db.DB.Prepare(`INSERT into user_info (user_nickname,user_password,user_age,user_sex,user_phone) values (?,?,?,?,?)`)
	checkErr(err)
	res, err := stmt.Exec(user_nickname, user_password, user_age, user_sex, user_phone)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("UserInsert",id)
	return id
}

// 根据ID-查询用戶信息
func UserQueryByNickId(user_id int) User {
	row := (db.DB).QueryRow(`SELECT
		user_id,
		user_realname,
		user_nickname,
		user_age,
		user_sex,
		user_adress,
		user_phone,
		user_qq,
		user_wechat
		FROM user_info where user_id = ?`, user_id)
	var u User
	err := row.Scan(&u.UserId, &u.UserRealname, &u.UserNickname,&u.UserAdress, &u.UserPhone, &u.UserWeChat, &u.UserAge, &u.UserSex, &u.UserQQ)
	checkErr(err)
	fmt.Println("UserQueryByNickId",*u.UserNickname)
	return u
}

// 查询所有用户
func UserListQuery() (users []User) {
	rows, err := (db.DB).Query(`SELECT * FROM user_info`)
	checkErr(err)
	defer rows.Close()
	var u User
	var arr []User
	for rows.Next() {
		//var user_id *int
		//var user_nickname *string
		//var user_age *int
		//var user_sex *int

		rows.Columns()
		//err = rows.Scan(&user_id, &user_nickname, &user_age, &user_sex)
		err := rows.Scan(&u.UserId, &u.UserPassword,&u.UserRealname, &u.UserNickname,&u.UserAdress, &u.UserPhone, &u.UserWeChat, &u.UserAge, &u.UserSex, &u.UserQQ)
		checkErr(err)
		//u = User{
		//	UserId:       user_id,
		//	UserNickname: user_nickname,
		//	UserAge:      user_age,
		//	UserSex:      user_sex,
		//}
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
