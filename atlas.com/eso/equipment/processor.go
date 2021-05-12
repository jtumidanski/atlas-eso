package equipment

import (
	"atlas-eso/rest/equipment_info"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CreateEquipment(l logrus.FieldLogger, db *gorm.DB) func(itemId uint32, strength uint16, dexterity uint16, intelligence uint16, luck uint16,
	hp uint16, mp uint16, weaponAttack uint16, magicAttack uint16, weaponDefense uint16, magicDefense uint16,
	accuracy uint16, avoidability uint16, hands uint16, speed uint16, jump uint16, slots uint16) (*Model, error) {
	return func(itemId uint32, strength uint16, dexterity uint16, intelligence uint16, luck uint16, hp uint16, mp uint16, weaponAttack uint16, magicAttack uint16, weaponDefense uint16, magicDefense uint16, accuracy uint16, avoidability uint16, hands uint16, speed uint16, jump uint16, slots uint16) (*Model, error) {
		if strength == 0 && dexterity == 0 && intelligence == 0 && luck == 0 && hp == 0 && mp == 0 && weaponAttack == 0 && weaponDefense == 0 &&
			magicAttack == 0 && magicDefense == 0 && accuracy == 0 && avoidability == 0 && hands == 0 && speed == 0 && jump == 0 &&
			slots == 0 {
			ea, err := equipment_info.EquipmentInfo().GetById(itemId)
			if err != nil {
				l.WithError(err).Errorf("Unable to get equipment information for %d.", itemId)
				return nil, err
			} else {
				attr := ea.Data().Attributes
				return create(db, itemId, attr.Strength, attr.Dexterity, attr.Intelligence, attr.Luck,
					attr.HP, attr.MP, attr.WeaponAttack, attr.MagicAttack, attr.WeaponDefense, attr.MagicDefense, attr.Accuracy,
					attr.Avoidability, attr.Hands, attr.Speed, attr.Jump, attr.Slots)
			}
		} else {
			return create(db, itemId, strength, dexterity, intelligence, luck, hp, mp, weaponAttack,
				magicAttack, weaponDefense, magicDefense, accuracy, avoidability, hands, speed, jump, slots)
		}
	}
}

func DeleteById(_ logrus.FieldLogger, db *gorm.DB) func(equipmentId uint32) error {
	return func(equipmentId uint32) error {
		return delete(db, equipmentId)
	}
}
