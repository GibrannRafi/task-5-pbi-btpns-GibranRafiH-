package models

type User struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaLengkap string `gorm:"varchar(300)" json:"nama_lengkap"`
	Username    string `json:"username" gorm:"varchar(300)"`
	Password    string `gorm:"varchar(300)" json:"password"`
}
