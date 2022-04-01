package statistics

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

func requestById(id uint32) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(itemInformationById, id))
}
