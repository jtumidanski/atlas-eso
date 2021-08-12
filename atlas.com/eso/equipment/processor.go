package equipment

import (
	"atlas-eso/rest/equipment_info"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math"
	"math/rand"
)

func CreateEquipment(l logrus.FieldLogger, db *gorm.DB) func(itemId uint32, strength uint16, dexterity uint16, intelligence uint16, luck uint16,
	hp uint16, mp uint16, weaponAttack uint16, magicAttack uint16, weaponDefense uint16, magicDefense uint16,
	accuracy uint16, avoidability uint16, hands uint16, speed uint16, jump uint16, slots uint16) (*Model, error) {
	return func(itemId uint32, strength uint16, dexterity uint16, intelligence uint16, luck uint16, hp uint16, mp uint16, weaponAttack uint16, magicAttack uint16, weaponDefense uint16, magicDefense uint16, accuracy uint16, avoidability uint16, hands uint16, speed uint16, jump uint16, slots uint16) (*Model, error) {
		if strength == 0 && dexterity == 0 && intelligence == 0 && luck == 0 && hp == 0 && mp == 0 && weaponAttack == 0 && weaponDefense == 0 &&
			magicAttack == 0 && magicDefense == 0 && accuracy == 0 && avoidability == 0 && hands == 0 && speed == 0 && jump == 0 &&
			slots == 0 {
			ea, err := equipment_info.GetById(l)(itemId)
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

func CreateRandomEquipment(l logrus.FieldLogger, db *gorm.DB) func(itemId uint32) (*Model, error) {
	return func(itemId uint32) (*Model, error) {
		ea, err := equipment_info.GetById(l)(itemId)
		if err != nil {
			l.WithError(err).Errorf("Unable to get equipment information for %d.", itemId)
			return nil, err
		} else {
			strength := getRandomStat(ea.Data().Attributes.Strength, 5)
			dexterity := getRandomStat(ea.Data().Attributes.Dexterity, 5)
			intelligence := getRandomStat(ea.Data().Attributes.Intelligence, 5)
			luck := getRandomStat(ea.Data().Attributes.Luck, 5)
			hp := getRandomStat(ea.Data().Attributes.HP, 10)
			mp := getRandomStat(ea.Data().Attributes.MP, 10)
			weaponAttack := getRandomStat(ea.Data().Attributes.WeaponAttack, 5)
			magicAttack := getRandomStat(ea.Data().Attributes.MagicAttack, 5)
			weaponDefense := getRandomStat(ea.Data().Attributes.WeaponDefense, 10)
			magicDefense := getRandomStat(ea.Data().Attributes.MagicDefense, 10)
			accuracy := getRandomStat(ea.Data().Attributes.Accuracy, 5)
			avoidability := getRandomStat(ea.Data().Attributes.Avoidability, 5)
			hands := getRandomStat(ea.Data().Attributes.Hands, 5)
			speed := getRandomStat(ea.Data().Attributes.Speed, 5)
			jump := getRandomStat(ea.Data().Attributes.Jump, 5)
			slots := ea.Data().Attributes.Slots

			return create(db, itemId, strength, dexterity, intelligence, luck, hp, mp, weaponAttack, magicAttack, weaponDefense, magicDefense, accuracy, avoidability, hands, speed, jump, slots)
		}

	}
}

func getRandomStat(defaultValue uint16, max uint16) uint16 {
	if defaultValue == 0 {
		return 0
	}
	maxRange := math.Min(math.Ceil(float64(defaultValue)*0.1), float64(max))
	return uint16(float64(defaultValue)-maxRange) + uint16(math.Floor(rand.Float64()*(maxRange*2.0+1.0)))
}

func DeleteById(_ logrus.FieldLogger, db *gorm.DB) func(equipmentId uint32) error {
	return func(equipmentId uint32) error {
		return delete(db, equipmentId)
	}
}
