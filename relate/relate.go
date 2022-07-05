package relate

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Profile   Profile `gorm:"ForeignKey:ProfileRefer;AssociationForeignKey:Refer"` //使用ProfileRefer 做为外键
	ProfileId int
}

type Profile struct {
	gorm.Model
	Name string
}

func main() {
	// db, err := gorm.Open("mysql", "root:123456@/gorm_test?charset=utf8&parseTime=True&loc=Local")
	// db.Model(&user).Related(&profile)
}
