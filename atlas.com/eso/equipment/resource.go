package equipment

import (
	"atlas-eso/json"
	"atlas-eso/rest"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

const (
	createRandomEquipment = "create_random_equipment"
	createEquipment       = "create_equipment"
	getEquipment          = "get_equipment"
	deleteEquipment       = "delete_equipment"
)

func InitResource(router *mux.Router, l logrus.FieldLogger, db *gorm.DB) {
	eRouter := router.PathPrefix("/equipment").Subrouter()
	eRouter.HandleFunc("", registerCreateRandomEquipment(l, db)).Queries("random", "{random}").Methods(http.MethodPost)
	eRouter.HandleFunc("", registerCreateEquipment(l, db)).Methods(http.MethodPost)
	eRouter.HandleFunc("/{equipmentId}", registerGetEquipment(l, db)).Methods(http.MethodGet)
	eRouter.HandleFunc("/{equipmentId}", registerDeleteEquipment(l, db)).Methods(http.MethodDelete)
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

type equipmentIdHandler func(equipmentId uint32) http.HandlerFunc

func parseEquipmentId(l logrus.FieldLogger, next equipmentIdHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		value, err := strconv.Atoi(vars["equipmentId"])
		if err != nil {
			l.WithError(err).Errorf("Error parsing characterId as uint32")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next(uint32(value))(w, r)
	}
}

func registerDeleteEquipment(l logrus.FieldLogger, db *gorm.DB) http.HandlerFunc {
	return rest.RetrieveSpan(deleteEquipment, func(span opentracing.Span) http.HandlerFunc {
		return parseEquipmentId(l, func(equipmentId uint32) http.HandlerFunc {
			return handleDeleteEquipment(l, db)(span)(equipmentId)
		})
	})
}

func handleDeleteEquipment(l logrus.FieldLogger, db *gorm.DB) func(span opentracing.Span) func(equipmentId uint32) http.HandlerFunc {
	return func(span opentracing.Span) func(equipmentId uint32) http.HandlerFunc {
		return func(equipmentId uint32) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				err := DeleteById(l, db)(equipmentId)
				if err != nil {
					l.WithError(err).Errorf("Unable to delete equipment %d.", equipmentId)
					w.WriteHeader(http.StatusNotFound)
					return
				}

				w.WriteHeader(http.StatusNoContent)
			}
		}
	}
}

func registerCreateRandomEquipment(l logrus.FieldLogger, db *gorm.DB) http.HandlerFunc {
	return rest.RetrieveSpan(createRandomEquipment, func(span opentracing.Span) http.HandlerFunc {
		return handleCreateRandomEquipment(l, db)(span)
	})
}

func handleCreateRandomEquipment(l logrus.FieldLogger, db *gorm.DB) rest.SpanHandler {
	return func(span opentracing.Span) http.HandlerFunc {
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
			e, err := CreateRandom(l, db, span)(att.ItemId)
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
}

func registerCreateEquipment(l logrus.FieldLogger, db *gorm.DB) http.HandlerFunc {
	return rest.RetrieveSpan(createEquipment, func(span opentracing.Span) http.HandlerFunc {
		return handleCreateEquipment(l, db)(span)
	})
}

func handleCreateEquipment(l logrus.FieldLogger, db *gorm.DB) rest.SpanHandler {
	return func(span opentracing.Span) http.HandlerFunc {
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
			e, err := Create(l, db, span)(att.ItemId, att.Strength, att.Dexterity, att.Intelligence, att.Luck,
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
}

func registerGetEquipment(l logrus.FieldLogger, db *gorm.DB) http.HandlerFunc {
	return rest.RetrieveSpan(getEquipment, func(span opentracing.Span) http.HandlerFunc {
		return parseEquipmentId(l, func(equipmentId uint32) http.HandlerFunc {
			return handleGetEquipment(l, db)(span)(equipmentId)
		})
	})
}

func handleGetEquipment(l logrus.FieldLogger, db *gorm.DB) func(span opentracing.Span) func(equipmentId uint32) http.HandlerFunc {
	return func(span opentracing.Span) func(equipmentId uint32) http.HandlerFunc {
		return func(equipmentId uint32) http.HandlerFunc {
			return func(rw http.ResponseWriter, r *http.Request) {
				e, err := GetById(db, equipmentId)
				if err != nil {
					l.WithError(err).Errorf("Unable to retrieve equipment %d.", equipmentId)
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
