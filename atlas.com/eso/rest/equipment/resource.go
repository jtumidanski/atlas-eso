package equipment

import (
	"atlas-eso/database/equipment"
	"atlas-eso/domain"
	"atlas-eso/processor"
	"atlas-eso/rest/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

func CreateEquipment(l *log.Logger, db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		ei := &EquipmentInputDataContainer{}
		err := json.FromJSON(ei, r.Body)
		if err != nil {
			l.Println("[ERROR] deserializing instruction", err)
			rw.WriteHeader(http.StatusBadRequest)
			json.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		att := ei.Data.Attributes
		e, err := processor.CreateEquipment(db, att.ItemId, att.Strength, att.Dexterity, att.Intelligence, att.Luck,
			att.HP, att.MP, att.WeaponAttack, att.MagicAttack, att.WeaponDefense, att.MagicDefense, att.Accuracy,
			att.Avoidability, att.Hands, att.Speed, att.Jump, att.Slots)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			json.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		result := makeEquipmentResult(e)
		json.ToJSON(result, rw)
	}
}

func GetEquipmentById(_ *log.Logger, db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		id := getEquipmentId(r)

		e, err := equipment.GetById(db, id)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		rw.WriteHeader(http.StatusOK)
		result := makeEquipmentResult(e)
		json.ToJSON(result, rw)
	}
}

func makeEquipmentResult(e *domain.Equipment) EquipmentDataContainer {
	result := EquipmentDataContainer{
		Data: EquipmentData{
			Id:   strconv.FormatUint(uint64(e.Id()), 10),
			Type: "com.atlas.eso.attribute.EquipmentAttributes",
			Attributes: EquipmentAttributes{
				ItemId:        e.ItemId(),
				Strength:      e.Strength(),
				Dexterity:     e.Dexterity(),
				Intelligence:  e.Intelligence(),
				Luck:          e.Luck(),
				HP:            e.HP(),
				MP:            e.MP(),
				WeaponAttack:  e.WeaponAttack(),
				MagicAttack:   e.MagicAttack(),
				WeaponDefense: e.WeaponDefense(),
				MagicDefense:  e.MagicDefense(),
				Accuracy:      e.Accuracy(),
				Avoidability:  e.Avoidability(),
				Hands:         e.Hands(),
				Speed:         e.Speed(),
				Jump:          e.Jump(),
				Slots:         e.Slots(),
			},
		},
	}
	return result
}

func getEquipmentId(r *http.Request) uint32 {
	vars := mux.Vars(r)
	value, err := strconv.Atoi(vars["equipmentId"])
	if err != nil {
		log.Println("Error parsing characterId as uint32")
		return 0
	}
	return uint32(value)
}
