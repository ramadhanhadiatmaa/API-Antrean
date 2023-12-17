package models

type Antrean struct {
	Id  string `gorm:"type:varchar(100); primaryKey" json:"id"`
	Num string `gorm:"type:varchar(25)" json:"num"`
}