package api

import (
	"encoding/json"
	"lecturer/event"
	"lecturer/helpers"
	"lecturer/models"
	"lecturer/services"
	"net/http"

	"github.com/go-chi/chi"
	amqp "github.com/rabbitmq/amqp091-go"
)

type API struct {
	Rabbit *amqp.Connection
	Helpers helpers.Helpers
	LecturerService services.LecturerService
	ClassService services.ClassRoomService
}

/* 

CHECK FOR ERRORS DURING SETUP, IF POSSIBLE

*/

func NewAPI(rabbit *amqp.Connection, helpers *helpers.Helpers, lecturerService *services.LecturerService, classRoomService *services.ClassRoomService) (API/*, error*/) {
	api := API{
		Rabbit: rabbit,
		Helpers: *helpers,
		LecturerService: *lecturerService,
		ClassService: *classRoomService,
	}

	/*err :=*/ //api.start()

	/*if err != nil {
		return API{}, err
	}*/

	return api//, nil
}

func (api *API) Start() /*error*/ {
	api.registerRoutes()
}

func (api *API) registerRoutes() {
	mux := chi.NewRouter()

	mux.HandleFunc("/createLecturer", api.CreateLecturer)
	mux.HandleFunc("/getAllLecturers", api.GetAllLecturers)
	mux.HandleFunc("/addLecturerToClass", api.AddLecturerToClass)
	mux.HandleFunc("/sendMessage", api.SendMessage)
	mux.HandleFunc("/getAllClasses", api.GetAllClasses)
	mux.HandleFunc("/getLecturerByID", api.GetLecturerByID)

	http.ListenAndServe(":8080", mux)
}

func (api *API ) SendMessage(w http.ResponseWriter, r *http.Request) {
	var requestPayload models.RequestPayload
	
	err := api.Helpers.ReadJSON(w, r, &requestPayload)
	
	if err != nil {
		api.Helpers.ErrorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "message":
		api.putMessageOnQueue(w, requestPayload.Message)
	default:
		api.Helpers.ErrorJSON(w, err)
	}
}

func (api *API ) putMessageOnQueue(w http.ResponseWriter, msg models.MessagePayload) {

	err := api.pushToQueue(msg.From, msg.To, msg.Message)
	if err != nil {
		api.Helpers.ErrorJSON(w, err)
		return
	}

	var payload models.JsonResponse
	payload.Error = false
	payload.Message = "Send message via RabbitMQ"

	api.Helpers.WriteJSON(w, http.StatusAccepted, payload)
}

func (api *API)  pushToQueue(from string, to []string, message string) error {
	emitter, err := event.NewEventEmitter(api.Rabbit)
	if err != nil {
		return err
	}

	payload := models.MessagePayload{
		From:    from,
		To:      to,
		Message: message,
	}

	j, _ := json.MarshalIndent(&payload, "", "\t")
	err = emitter.Push(string(j), "Messages")
	
	if err != nil {
		return err
	}
	return nil
}

func (api *API ) CreateLecturer(w http.ResponseWriter, lect *http.Request) {
	var lecturerPayload models.LecturerPayload
	error := api.Helpers.ReadJSON(w, lect, &lecturerPayload)

	if error!= nil {
		api.Helpers.ErrorJSON(w, error)
		return 
	}

	lect3, error := api.LecturerService.CreateLecturer(lecturerPayload)

	if(error != nil){
		api.Helpers.ErrorJSON(w, error)
		return
	}

	var payload models.JsonResponse

	payload.Error = false
	payload.Message = "Created lecturer"
	payload.Data = lect3

	api.Helpers.WriteJSON(w, http.StatusAccepted, payload)
}

func (api *API) GetAllClasses(w http.ResponseWriter, r *http.Request) {
	classes := api.ClassService.GetAllClasses()

	var payload models.JsonResponse

	payload.Error = false
	payload.Message = "All classes"
	payload.Data = classes

	api.Helpers.WriteJSON(w, http.StatusAccepted, payload)
}

func (api *API ) AddLecturerToClass(w http.ResponseWriter, lect *http.Request) {
	var classLecturerPayload models.ClassLecturerPayload
	error := api.Helpers.ReadJSON(w, lect, &classLecturerPayload)

	if(error != nil){
		api.Helpers.ErrorJSON(w, error)
		return
	}

	message, error := api.LecturerService.AddLecturerToClass(classLecturerPayload.ClassId, classLecturerPayload.LecturerId)

	if(error != nil){
		api.Helpers.ErrorJSON(w, error)
		return
	}

	var payload models.JsonResponse
	payload.Error = false
	payload.Message = message
	payload.Data = nil

	api.Helpers.WriteJSON(w, http.StatusAccepted, payload)
}

func (api *API ) GetAllLecturers(w http.ResponseWriter, request *http.Request)  {
	lecturers, err := api.LecturerService.GetAllLecturers()

	if err != nil {
		api.Helpers.ErrorJSON(w, err)
		return
	}

	var payload models.JsonResponse
	payload.Error = false
	payload.Message = "Retrieved lecturers"
	payload.Data = lecturers

	api.Helpers.WriteJSON(w, http.StatusOK, payload)
}

func (api *API) GetLecturerByID (w http.ResponseWriter, request *http.Request) {
	var idPayload models.IDPayload
	err := api.Helpers.ReadJSON(w, request, &idPayload)

	if err != nil {
		api.Helpers.ErrorJSON(w, err)
		return
	}

	lecturer, err := api.LecturerService.GetLecturerByID(idPayload.ID)

	if err != nil {
		api.Helpers.ErrorJSON(w, err)
		return
	}

	var payload models.JsonResponse
	payload.Error = false
	payload.Message = "Retrieved lecturer"
	payload.Data = lecturer

	api.Helpers.WriteJSON(w, http.StatusOK, payload)
}