package models

import (
	"DataCertPlatform/db_mysql"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Id          int     `form:"id"`
	Phone       string  `form:"phone"`
	Password    string  `form:"password"`
}


/*
*将用户的信息保存到数据库中
 */
func (u User) AddUser()  (int64, error){
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	pwdBytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(pwdBytes)//把脱敏的密码的MD5值重新赋值
	rs, err := db_mysql.Db.Exec("insert into user(phone,password) value(?,?) ",
		u.Phone, u.Password)
	//错误早发现，早解决
	if err != nil{//保存数据遇到的错误
		return -1, err
	}
	id,err := rs.RowsAffected()
	if err != nil {
		return -1, err

	}
	return id,nil
}
/*
*查询用户信息
 */
func (u User) QueryUser() (*User,error) {
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	pwdBytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(pwdBytes)//把脱敏的密码的md5值重新赋值为密码进行储存
	row := db_mysql.Db.QueryRow("select phone from user where phone = ? and password = ?",
		u.Phone, u.Password)
	err := row.Scan(&u.Phone)
	if err != nil{
		return nil, err
	}
	return &u,nil
}