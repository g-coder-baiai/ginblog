package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string		`gorm:"type:varchar(20)" json:"username"`
	Password string		`gorm:"type:varchar(20)" json:"password"`
	Role int            `gorm:"type:int" json:"role"`     //1为管理者，2为用户
}

//查询用户是否存在
func CheckUser(name string)(code int){
	var user User
	db.Select("id").Where("username = ?",name).First(&user)
	if user.ID>0{
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}


// 创建用户
func CreateUser(data *User)int{
	// data.Password=ScryptPw(data.Password)

	err:= db.Create(&data).Error
	if err!=nil{
		return errmsg.ERROR   // 500
	}
	return errmsg.SUCCSE
}

// GetUser 查询用户
func GetUser(id int) (User, int) {
	var user User
	err := db.Limit(1).Where("ID = ?", id).Find(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCSE
}

// 查询用户列表
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64

	if username != "" {
		db.Select("id,username,role,created_at").Where(
			"username LIKE ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		db.Model(&users).Where(
			"username LIKE ?", username+"%",
		).Count(&total)
		return users, total
	}
	db.Select("id,username,role,created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	db.Model(&users).Count(&total)

	if err != nil {
		return users, 0
	}
	return users, total
}

// CheckUpUser 更新查询
func CheckUpUser(id int, name string) (code int) {
	var user User
	db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCSE
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE
}


// EditUser 编辑用户信息
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}


//删除用户
func DeleteUser(id int)int{
	var user User
	err:=db.Where("id=?",id).Delete(&user).Error
	if err!=nil{
		return errmsg.ERROR
	}

	return errmsg.SUCCSE

}

// BeforeCreate 密码加密&权限控制
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	//log.Println(u.Password)

	u.Role = 2
	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return nil
}


//密码加密
func ScryptPw(password string)string{
	const KeyLen=10
	salt := make([]byte, 8)
	salt=[]byte{13,24,56,34,44,55,11,31}

	HashPw,err:=scrypt.Key([]byte(password),salt,2048,8,1,KeyLen)
	if err!=nil{
		log.Fatal(err)
	}
	fpw:=base64.StdEncoding.EncodeToString(HashPw)
	return fpw

}