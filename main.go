package main

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gorm_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	user := User{Name: "cl", Age: 18, Birthday: time.Now()}
	db.NewRecord(user)
	db.Create(&user)
	db.NewRecord(user)
	defer db.Close()

}

/**
默认表名是结构体的复数形式
*/

func (User) TableName() string {
	return "profiles" //设置User的表名为profiles
}

//列名是字段名的蛇形小写   Age --> age    CreatedAt --> created_at
type User struct { //默认表名是users
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"`       //string 默认长度为255，使用这种tag重设
	Num      int    `gorm:"AUTO_INCREMENT"` //自增

	CreditCard       CreditCard //拥有一个  userid为外键
	Emails           []Email
	BillingAddress   Address // One-To-One (属于 - 本表的BillingAddressID作外键)
	BillingAddressID sql.NullInt64

	ShippingAddress   Address // One-To-One (属于 - 本表的ShippingAddressID作外键)
	ShippingAddressID int

	IgnoreMe  int        `gorm:"-"`                         // 忽略这个字段
	Languages []Language `gorm:"many2many:user_languages;"` // Many-To-Many , 'user_languages'是连接表
}

type CreditCard struct {
	//模型
	gorm.Model
	UserID uint
	Number string
}

type Email struct {
	ID         int
	UserID     int    `gorm:"index"`                          // 外键 (属于), tag `index`是为该列创建索引
	Email      string `gorm:"type:varchar(100);unique_index"` // `type`设置sql类型, `unique_index` 为该列设置唯一索引
	Subscribed bool
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"` // 创建索引并命名，如果找到其他相同名称的索引则创建组合索引
	Code string `gorm:"index:idx_name_code"` // `unique_index` also works
}

type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"` // 设置字段为非空并唯一
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}

//重设列名
type Animal struct {
	AnimalId int64 `gorm:"column:beast_id;primary_key"` // 设置列名为`beast_id` 并且为主键
}
