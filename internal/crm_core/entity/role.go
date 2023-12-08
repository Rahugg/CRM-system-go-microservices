package entity

type Role struct {
	ID   uint   `gorm:"primary_key, AUTO_INCREMENT"`
	Name string `gorm:"unique"`
}
