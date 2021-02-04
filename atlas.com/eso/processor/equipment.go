package processor

import (
	"atlas-eso/database/equipment"
	"atlas-eso/domain"
	"atlas-eso/rest/equipment_info"
	"gorm.io/gorm"
)

func CreateEquipment(db *gorm.DB, itemId uint32, strength uint16, dexterity uint16, intelligence uint16, luck uint16,
	hp uint16, mp uint16, weaponAttack uint16, magicAttack uint16, weaponDefense uint16, magicDefense uint16,
	accuracy uint16, avoidability uint16, hands uint16, speed uint16, jump uint16, slots uint16) (*domain.Equipment, error) {

	if strength == 0 && dexterity == 0 && intelligence == 0 && luck == 0 && hp == 0 && mp == 0 && weaponAttack == 0 && weaponDefense == 0 &&
		magicAttack == 0 && magicDefense == 0 && accuracy == 0 && avoidability == 0 && hands == 0 && speed == 0 && jump == 0 &&
		slots == 0 {
		ea, err := equipment_info.EquipmentInfo().GetById(itemId)
		if err != nil {
			return nil, err
		} else {
			attr := ea.Data().Attributes
			return equipment.CreateEquipment(db, itemId, attr.Strength, attr.Dexterity, attr.Intelligence, attr.Luck,
				attr.HP, attr.MP, attr.WeaponAttack, attr.MagicAttack, attr.WeaponDefense, attr.MagicDefense, attr.Accuracy,
				attr.Avoidability, attr.Hands, attr.Speed, attr.Jump, attr.Slots)
		}
	} else {
		return equipment.CreateEquipment(db, itemId, strength, dexterity, intelligence, luck, hp, mp, weaponAttack,
			magicAttack, weaponDefense, magicDefense, accuracy, avoidability, hands, speed, jump, slots)
	}
}
