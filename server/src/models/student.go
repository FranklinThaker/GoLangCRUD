package models

type Student struct {
	ID     int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Fname  string `gorm:"type:varchar(100)"`
	Lname  string `gorm:"type:varchar(100)"`
	Age    int    `gorm:"size:10"`
	Email  string `gorm:"type:varchar(100)"`
	Mobile int    `gorm:"size:10"`
}
