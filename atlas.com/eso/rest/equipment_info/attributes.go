package equipment_info

import "atlas-eso/rest/json"

type EquipmentInfoDataContainer struct {
	data     json.DataSegment
	included json.DataSegment
}

type EquipmentInfoData struct {
	Id         string                  `json:"id"`
	Type       string                  `json:"type"`
	Attributes EquipmentInfoAttributes `json:"attributes"`
}

type EquipmentInfoAttributes struct {
	Strength      uint16 `json:"strength"`
	Dexterity     uint16 `json:"dexterity"`
	Intelligence  uint16 `json:"intelligence"`
	Luck          uint16 `json:"luck"`
	HP            uint16 `json:"hp"`
	MP            uint16 `json:"mp"`
	WeaponAttack  uint16 `json:"weaponAttack"`
	MagicAttack   uint16 `json:"magicAttack"`
	WeaponDefense uint16 `json:"weaponDefense"`
	MagicDefense  uint16 `json:"magicDefense"`
	Accuracy      uint16 `json:"accuracy"`
	Avoidability  uint16 `json:"avoidability"`
	Hands         uint16 `json:"hands"`
	Speed         uint16 `json:"speed"`
	Jump          uint16 `json:"jump"`
	Slots         uint16 `json:"slots"`
}

func (a *EquipmentInfoDataContainer) UnmarshalJSON(data []byte) error {
	d, i, err := json.UnmarshalRoot(data, json.MapperFunc(EmptyEquipmentInfoData))
	if err != nil {
		return err
	}

	a.data = d
	a.included = i
	return nil
}

func (a *EquipmentInfoDataContainer) Data() *EquipmentInfoData {
	if len(a.data) >= 1 {
		return a.data[0].(*EquipmentInfoData)
	}
	return nil
}

func (a *EquipmentInfoDataContainer) DataList() []EquipmentInfoData {
	var r = make([]EquipmentInfoData, 0)
	for _, x := range a.data {
		r = append(r, *x.(*EquipmentInfoData))
	}
	return r
}

func EmptyEquipmentInfoData() interface{} {
	return &EquipmentInfoData{}
}
