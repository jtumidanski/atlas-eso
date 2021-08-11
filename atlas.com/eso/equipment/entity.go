package equipment

import "gorm.io/gorm"

func Migration(db *gorm.DB) error {
	return db.AutoMigrate(&equipment{})
}

type equipment struct {
	ID            uint32 `gorm:"primaryKey;autoIncrement;not null"`
	ItemId        uint32 `gorm:"not null;default=0"`
	Strength      uint16 `gorm:"not null;default=0"`
	Dexterity     uint16 `gorm:"not null;default=0"`
	Intelligence  uint16 `gorm:"not null;default=0"`
	Luck          uint16 `gorm:"not null;default=0"`
	Hp            uint16 `gorm:"not null;default=0"`
	Mp            uint16 `gorm:"not null;default=0"`
	WeaponAttack  uint16 `gorm:"not null;default=0"`
	MagicAttack   uint16 `gorm:"not null;default=0"`
	WeaponDefense uint16 `gorm:"not null;default=0"`
	MagicDefense  uint16 `gorm:"not null;default=0"`
	Accuracy      uint16 `gorm:"not null;default=0"`
	Avoidability  uint16 `gorm:"not null;default=0"`
	Hands         uint16 `gorm:"not null;default=0"`
	Speed         uint16 `gorm:"not null;default=0"`
	Jump          uint16 `gorm:"not null;default=0"`
	Slots         uint16 `gorm:"not null;default=0"`
}
