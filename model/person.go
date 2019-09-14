package model

type Person struct {
	//gorm.Model
	Id    int64  `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
	Age   int64  `gorm:"column:age"`
}

func (Person) TableName() string {
	return "person"
}
