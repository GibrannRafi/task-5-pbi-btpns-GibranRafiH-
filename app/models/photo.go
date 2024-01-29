package models

type Photo struct {
	// User     User `gorm:"foreignkey:UserID"`
	Id       int64  `gorm:"primaryKey" json:"id"`
	Title    string `gorm:"varchar(300)" json:"title"`
	Caption  string `gorm:"varchar(300)" json:"caption"`
	PhotoUrl string `gorm:"varchar(300)" json:"photo_url"`
}
