package equipment_info

import (
	"atlas-eso/rest/requests"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	itemInformationServicePrefix string = "/ms/iis/"
	itemInformationService              = requests.BaseRequest + itemInformationServicePrefix
	itemInformationResource             = itemInformationService + "equipment/"
	itemInformationById                 = itemInformationResource + "%d"
)

func GetById(l logrus.FieldLogger, span opentracing.Span) func(id uint32) (*EquipmentInfoDataContainer, error) {
	return func(id uint32) (*EquipmentInfoDataContainer, error) {
		ar := &EquipmentInfoDataContainer{}
		err := requests.Get(l, span)(fmt.Sprintf(itemInformationById, id), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}
