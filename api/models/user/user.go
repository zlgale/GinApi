package user

import (
	"GinApi/api/db"
	"fmt"
)

type User struct {
	Uid        string `form:"uid" json:"uid"`
	Username   string `form:"username json:"username" binding:"required`
	Departname string `form:"departname json:"departname"`
	Created    string `form:"created json:"created"`
}

//check user password
//func (user *User) CheckPassword(rawPassword string) bool {
//	password := genPassword(
//		rawPassword, user.Salt)
//
//	return password == user.Password
//}
//
//func UserQueryById(uid int) (user User) {
//
//	o := orm.NewOrm()
//	u := User{Id: uid}
//
//	err := o.Read(&u)
//
//	if err == orm.ErrNoRows {
//		fmt.Println("查询不到")
//	} else if err == orm.ErrMissPK {
//		fmt.Println("找不到主键")
//	} else {
//		fmt.Println(u.Id, u.Name)
//	}
//
//	return u
//}
// 插入数据
func UserInsert(uid, username, departname, created string) bool {
	//db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	//checkErr(err)
	stmt, err := db.DB.Prepare(`INSERT into userinfo (uid,username,departname,created) values (?,?,?,?)`)
	checkErr(err)
	res, err := stmt.Exec(uid, username, departname, created)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	return true
}

// 根据姓名查询
func UserQueryByUserName(name string) User {
	//rows, err := (models.DB).Query(`SELECT * FROM userinfo where username = ?`,"大风")
	rows := (db.DB).QueryRow(`SELECT * FROM userinfo where username = ?`, name)
	////普通demo
	var u User
	var uid, username, departname, created string
	err := rows.Scan(&uid, &username, &departname, &created)
	checkErr(err)
	u = User{
		Uid:        uid,
		Username:   username,
		Departname: departname,
		Created:    created,
	}
	return u
}

// 查询所有用户
func UserListQuery() (users []User) {
	rows, err := (db.DB).Query(`SELECT * FROM userinfo`)
	checkErr(err)
	////普通demo
	var u User
	var arr []User
	for rows.Next() {
		var uid string
		var username string
		var departname string
		var created string

		rows.Columns()
		err = rows.Scan(&uid, &username, &departname, &created)
		checkErr(err)
		u = User{
			Uid:        uid,
			Username:   username,
			Departname: departname,
			Created:    created,
		}
		arr = append(arr, u)
	}
	return arr
}

// 更新数据
func UserUpdate() {
	//db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	//checkErr(err)

	stmt, err := db.DB.Prepare(`UPDATE user SET user_age=?,user_sex=? WHERE user_id=?`)
	checkErr(err)
	res, err := stmt.Exec(21, 2, 1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

// 删除数据
func UserRemove() {
	//db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	//checkErr(err)

	stmt, err := db.DB.Prepare(`DELETE FROM user WHERE user_id=?`)
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
