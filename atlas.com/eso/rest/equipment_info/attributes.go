package equipment_info

import (
	"atlas-eso/rest/response"
	"encoding/json"
)

type EquipmentInfoDataContainer struct {
	data     response.DataSegment
	included response.DataSegment
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

func (c *EquipmentInfoDataContainer) MarshalJSON() ([]byte, error) {
	t := struct {
		Data     interface{} `json:"data"`
		Included interface{} `json:"included"`
	}{}
	if len(c.data) == 1 {
		t.Data = c.data[0]
	} else {
		t.Data = c.data
	}
	return json.Marshal(t)
}

func (c *EquipmentInfoDataContainer) UnmarshalJSON(data []byte) error {
	d, i, err := response.UnmarshalRoot(data, response.MapperFunc(EmptyEquipmentInfoData))
	if err != nil {
		return err
	}

	c.data = d
	c.included = i
	return nil
}

func (c *EquipmentInfoDataContainer) Data() *EquipmentInfoData {
	if len(c.data) >= 1 {
		return c.data[0].(*EquipmentInfoData)
	}
	return nil
}

func (c *EquipmentInfoDataContainer) DataList() []EquipmentInfoData {
	var r = make([]EquipmentInfoData, 0)
	for _, x := range c.data {
		r = append(r, *x.(*EquipmentInfoData))
	}
	return r
}

func EmptyEquipmentInfoData() interface{} {
	return &EquipmentInfoData{}
}
