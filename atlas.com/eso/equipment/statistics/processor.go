package statistics

import (
	"atlas-eso/model"
	"atlas-eso/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func ByIdModelProvider(l logrus.FieldLogger, span opentracing.Span) func(id uint32) model.Provider[Model] {
	return func(id uint32) model.Provider[Model] {
		return requests.Provider[attributes, Model](l, span)(requestById(id), makeModel)
	}
}

func GetById(l logrus.FieldLogger, span opentracing.Span) func(id uint32) (Model, error) {
	return func(id uint32) (Model, error) {
		return ByIdModelProvider(l, span)(id)()
	}
}

func makeModel(body requests.DataBody[attributes]) (Model, error) {
	attr := body.Attributes
	return Model{
		strength:      attr.Strength,
		dexterity:     attr.Dexterity,
		intelligence:  attr.Intelligence,
		luck:          attr.Luck,
		hp:            attr.HP,
		mp:            attr.MP,
		weaponAttack:  attr.WeaponAttack,
		magicAttack:   attr.MagicAttack,
		weaponDefense: attr.WeaponDefense,
		magicDefense:  attr.MagicDefense,
		accuracy:      attr.Accuracy,
		avoidability:  attr.Avoidability,
		hands:         attr.Hands,
		speed:         attr.Speed,
		jump:          attr.Jump,
		slots:         attr.Slots,
	}, nil
}
