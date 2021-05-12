package equipment

import (
	"gorm.io/gorm"
)

func GetById(db *gorm.DB, id uint32) (*Model, error) {
	var result equipment
	err := db.First(&result, id).Error
	if err != nil {
		return nil, err
	}
	return makeEquipment(&result), nil
}