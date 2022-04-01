package equipment

import (
	"atlas-eso/database"
	"atlas-eso/model"
	"gorm.io/gorm"
)

func byIdEntityProvider(id uint32) database.EntityProvider[entity] {
	return func(db *gorm.DB) model.Provider[entity] {
		var result entity
		err := db.First(&result, id).Error
		if err != nil {
			return model.ErrorProvider[entity](err)
		}
		return model.FixedProvider(result)
	}
}
