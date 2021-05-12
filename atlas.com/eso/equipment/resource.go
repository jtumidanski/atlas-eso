package equipment

import (
	"atlas-eso/rest/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

func HandleDeleteEquipment(l logrus.FieldLogger, db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := getEquipmentId(l)(r)

		err := DeleteById(l, db)(id)
		if err != nil {
			l.WithError(err).Errorf("Unable to delete equipment %d.", id)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func HandleCreateRandomEquipment(l logrus.FieldLogger, db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		ei := &InputDataContainer{}
		err := json.FromJSON(ei, r.Body)
		if err != nil {
			l.WithError(err).Errorf("Deserializing instruction.")
			rw.WriteHeader(http.StatusBadRequest)
			err = json.ToJSON(&GenericError{Message: err.Error()}, rw)
			if err != nil {
				l.WithError(err).Errorf("Cannot write error response.")
			}
			return
		}

		att := ei.Data.Attributes
		e, err := CreateRandomEquipment(l, db)(att.ItemId)
		if err != nil {
			l.WithError(err).Errorf("Cannot create equipment.")
			rw.WriteHeader(http.StatusInternalServerError)
			err = json.ToJSON(&GenericError{Message: err.Error()}, rw)
			if err != nil {
				l.WithError(err).Errorf("Cannot write error response.")
			}
			return
		}

		rw.WriteHeader(http.StatusCreated)
		result := makeEquipmentResult(e)
		err = json.ToJSON(result, rw)
		if err != nil {
			l.WithError(err).Errorf("Writing create equipment response.")
		}
	}
}

func HandleCreateEquipment(l logrus.FieldLogger, db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		ei := &InputDataContainer{}
		err := json.FromJSON(ei, r.Body)
		if err != nil {
			l.WithError(err).Errorf("Deserializing instruction.")
			rw.WriteHeader(http.StatusBadRequest)
			err = json.ToJSON(&GenericError{Message: err.Error()}, rw)
			if err != nil {
				l.WithError(err).Errorf("Cannot write error response.")
			}
			return
		}

		att := ei.Data.Attributes
		e, err := CreateEquipment(l, db)(att.ItemId, att.Strength, att.Dexterity, att.Intelligence, att.Luck,
			att.HP, att.MP, att.WeaponAttack, att.MagicAttack, att.WeaponDefense, att.MagicDefense, att.Accuracy,
			att.Avoidability, att.Hands, att.Speed, att.Jump, att.Slots)
		if err != nil {
			l.WithError(err).Errorf("Cannot create equipment.")
			rw.WriteHeader(http.StatusInternalServerError)
			err = json.ToJSON(&GenericError{Message: err.Error()}, rw)
			if err != nil {
				l.WithError(err).Errorf("Cannot write error response.")
			}
			return
		}

		rw.WriteHeader(http.StatusCreated)
		result := makeEquipmentResult(e)
		err = json.ToJSON(result, rw)
		if err != nil {
			l.WithError(err).Errorf("Writing create equipment response.")
		}
	}
}

func HandleGetEquipmentById(l logrus.FieldLogger, db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		id := getEquipmentId(l)(r)

		e, err := GetById(db, id)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve equipment %d.", id)
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		rw.WriteHeader(http.StatusOK)
		result := makeEquipmentResult(e)
		err = json.ToJSON(result, rw)
		if err != nil {
			l.WithError(err).Errorf("Unable to write get equipment by id response.")
		}
	}
}

func makeEquipmentResult(e *Model) DataContainer {
	result := DataContainer{
		Data: DataBody{
			Id:   strconv.FormatUint(uint64(e.Id()), 10),
			Type: "com.atlas.eso.attribute.EquipmentAttributes",
			Attributes: Attributes{
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

func getEquipmentId(l logrus.FieldLogger) func(r *http.Request) uint32 {
	return func(r *http.Request) uint32 {
		vars := mux.Vars(r)
		value, err := strconv.Atoi(vars["equipmentId"])
		if err != nil {
			l.WithError(err).Errorf("Error parsing characterId as uint32")
			return 0
		}
		return uint32(value)
	}
}
