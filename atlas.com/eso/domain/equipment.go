package domain

type Equipment struct {
	id            uint32
	itemId        uint32
	strength      uint16
	dexterity     uint16
	intelligence  uint16
	luck          uint16
	hp            uint16
	mp            uint16
	weaponAttack  uint16
	magicAttack   uint16
	weaponDefense uint16
	magicDefense  uint16
	accuracy      uint16
	avoidability  uint16
	hands         uint16
	speed         uint16
	jump          uint16
	slots         uint16
}

func (e Equipment) Id() uint32 {
	return e.id
}

func (e Equipment) ItemId() uint32 {
	return e.itemId
}

func (e Equipment) Strength() uint16 {
	return e.strength
}

func (e Equipment) Dexterity() uint16 {
	return e.dexterity
}

func (e Equipment) Intelligence() uint16 {
	return e.intelligence
}

func (e Equipment) Luck() uint16 {
	return e.luck
}

func (e Equipment) HP() uint16 {
	return e.hp
}

func (e Equipment) MP() uint16 {
	return e.mp
}

func (e Equipment) WeaponAttack() uint16 {
	return e.weaponAttack
}

func (e Equipment) MagicAttack() uint16 {
	return e.magicAttack
}

func (e Equipment) WeaponDefense() uint16 {
	return e.weaponDefense
}

func (e Equipment) MagicDefense() uint16 {
	return e.magicDefense
}

func (e Equipment) Accuracy() uint16 {
	return e.accuracy
}

func (e Equipment) Avoidability() uint16 {
	return e.avoidability
}

func (e Equipment) Hands() uint16 {
	return e.hands
}

func (e Equipment) Speed() uint16 {
	return e.speed
}

func (e Equipment) Jump() uint16 {
	return e.jump
}

func (e Equipment) Slots() uint16 {
	return e.slots
}

type equipmentBuilder struct {
	id            uint32
	itemId        uint32
	strength      uint16
	dexterity     uint16
	intelligence  uint16
	luck          uint16
	hp            uint16
	mp            uint16
	weaponAttack  uint16
	magicAttack   uint16
	weaponDefense uint16
	magicDefense  uint16
	accuracy      uint16
	avoidability  uint16
	hands         uint16
	speed         uint16
	jump          uint16
	slots         uint16
}

func NewEquipmentBuilder(id uint32) *equipmentBuilder {
	return &equipmentBuilder{id: id}
}

func (e *equipmentBuilder) SetItemId(itemId uint32) *equipmentBuilder {
	e.itemId = itemId
	return e
}

func (e *equipmentBuilder) SetStrength(strength uint16) *equipmentBuilder {
	e.strength = strength
	return e
}

func (e *equipmentBuilder) SetDexterity(dexterity uint16) *equipmentBuilder {
	e.dexterity = dexterity
	return e
}

func (e *equipmentBuilder) SetIntelligence(intelligence uint16) *equipmentBuilder {
	e.intelligence = intelligence
	return e
}

func (e *equipmentBuilder) SetLuck(luck uint16) *equipmentBuilder {
	e.luck = luck
	return e
}

func (e *equipmentBuilder) SetHp(hp uint16) *equipmentBuilder {
	e.hp = hp
	return e
}

func (e *equipmentBuilder) SetMp(mp uint16) *equipmentBuilder {
	e.mp = mp
	return e
}

func (e *equipmentBuilder) SetWeaponAttack(weaponAttack uint16) *equipmentBuilder {
	e.weaponAttack = weaponAttack
	return e
}

func (e *equipmentBuilder) SetMagicAttack(magicAttack uint16) *equipmentBuilder {
	e.magicAttack = magicAttack
	return e
}

func (e *equipmentBuilder) SetWeaponDefense(weaponDefense uint16) *equipmentBuilder {
	e.weaponDefense = weaponDefense
	return e
}

func (e *equipmentBuilder) SetMagicDefense(magicDefense uint16) *equipmentBuilder {
	e.magicDefense = magicDefense
	return e
}

func (e *equipmentBuilder) SetAccuracy(accuracy uint16) *equipmentBuilder {
	e.accuracy = accuracy
	return e
}

func (e *equipmentBuilder) SetAvoidability(avoidability uint16) *equipmentBuilder {
	e.avoidability = avoidability
	return e
}

func (e *equipmentBuilder) SetHands(hands uint16) *equipmentBuilder {
	e.hands = hands
	return e
}

func (e *equipmentBuilder) SetSpeed(speed uint16) *equipmentBuilder {
	e.speed = speed
	return e
}

func (e *equipmentBuilder) SetJump(jump uint16) *equipmentBuilder {
	e.jump = jump
	return e
}

func (e *equipmentBuilder) SetSlots(slots uint16) *equipmentBuilder {
	e.slots = slots
	return e
}

func (e *equipmentBuilder) Build() Equipment {
	return Equipment{
		id:            e.id,
		itemId:        e.itemId,
		strength:      e.strength,
		dexterity:     e.dexterity,
		intelligence:  e.intelligence,
		luck:          e.luck,
		hp:            e.hp,
		mp:            e.mp,
		weaponAttack:  e.weaponAttack,
		magicAttack:   e.magicAttack,
		weaponDefense: e.weaponDefense,
		magicDefense:  e.magicDefense,
		accuracy:      e.accuracy,
		avoidability:  e.avoidability,
		hands:         e.hands,
		speed:         e.speed,
		jump:          e.jump,
		slots:         e.slots,
	}
}
