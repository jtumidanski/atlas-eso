package equipment

import (
	"atlas-eso/equipment/statistics"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math"
	"math/rand"
)

func Create(l logrus.FieldLogger, db *gorm.DB, span opentracing.Span) func(itemId uint32, strength uint16, dexterity uint16, intelligence uint16, luck uint16,
	hp uint16, mp uint16, weaponAttack uint16, magicAttack uint16, weaponDefense uint16, magicDefense uint16,
	accuracy uint16, avoidability uint16, hands uint16, speed uint16, jump uint16, slots uint16) (*Model, error) {
	return func(itemId uint32, strength uint16, dexterity uint16, intelligence uint16, luck uint16, hp uint16, mp uint16, weaponAttack uint16, magicAttack uint16, weaponDefense uint16, magicDefense uint16, accuracy uint16, avoidability uint16, hands uint16, speed uint16, jump uint16, slots uint16) (*Model, error) {
		if strength == 0 && dexterity == 0 && intelligence == 0 && luck == 0 && hp == 0 && mp == 0 && weaponAttack == 0 && weaponDefense == 0 &&
			magicAttack == 0 && magicDefense == 0 && accuracy == 0 && avoidability == 0 && hands == 0 && speed == 0 && jump == 0 &&
			slots == 0 {
			ea, err := statistics.GetById(itemId)(l, span)
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

func CreateRandom(l logrus.FieldLogger, db *gorm.DB, span opentracing.Span) func(itemId uint32) (*Model, error) {
	return func(itemId uint32) (*Model, error) {
		ea, err := statistics.GetById(itemId)(l, span)
		if err != nil {
			l.WithError(err).Errorf("Unable to get equipment information for %d.", itemId)
			return nil, err
		} else {
			attr := ea.Data().Attributes
			strength := getRandomStat(attr.Strength, 5)
			dexterity := getRandomStat(attr.Dexterity, 5)
			intelligence := getRandomStat(attr.Intelligence, 5)
			luck := getRandomStat(attr.Luck, 5)
			hp := getRandomStat(attr.HP, 10)
			mp := getRandomStat(attr.MP, 10)
			weaponAttack := getRandomStat(attr.WeaponAttack, 5)
			magicAttack := getRandomStat(attr.MagicAttack, 5)
			weaponDefense := getRandomStat(attr.WeaponDefense, 10)
			magicDefense := getRandomStat(attr.MagicDefense, 10)
			accuracy := getRandomStat(attr.Accuracy, 5)
			avoidability := getRandomStat(attr.Avoidability, 5)
			hands := getRandomStat(attr.Hands, 5)
			speed := getRandomStat(attr.Speed, 5)
			jump := getRandomStat(attr.Jump, 5)
			slots := attr.Slots

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
