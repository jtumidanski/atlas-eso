package equipment_info

import (
	"atlas-eso/rest/requests"
	"fmt"
)

const (
	itemInformationServicePrefix string = "/ms/iis/"
	itemInformationService              = requests.BaseRequest + itemInformationServicePrefix
	itemInformationResource             = itemInformationService + "equipment/"
	itemInformationById                 = itemInformationResource + "%d"
)

var EquipmentInfo = func() *equipmentInfo {
	return &equipmentInfo{}
}

type equipmentInfo struct {
}

func (a *equipmentInfo) GetById(id uint32) (*EquipmentInfoDataContainer, error) {
	ar := &EquipmentInfoDataContainer{}
	err := requests.Get(fmt.Sprintf(itemInformationById, id), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
