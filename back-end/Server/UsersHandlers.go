package Server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"../Models"
	"../Services"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func UsersHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		users := Models.GetAllUsers()
		info := struct {
			Count int           `json:"count"`
			Users []Models.User `json:"users"`
		}{
			Count: len(users),
			Users: users,
		}
		formatter.JSON(w, http.StatusOK, info)
	}
}

func GetUserByUserIdHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		userId := vars["userId"]
		user, err := Models.GetUser(userId)
		if err == nil {
			formatter.JSON(w, http.StatusOK, *user)
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}

func GetUserAppRecordsAsUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		userId := vars["userId"]
		appRecords, err := Services.GetUserAppRecordsAsUser(userId)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, err.Error()+"\n")
			return
		} else {
			info := struct {
				Count      int `json:"Count"`
				AppRecords []Models.AppRecord
			}{
				Count:      len(appRecords),
				AppRecords: appRecords,
			}
			formatter.JSON(w, http.StatusOK, info)
		}
	}
}

func CreateAppRecordAsUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		userId := vars["userId"]
		println(userId)
		payload, _ := ioutil.ReadAll(req.Body)
		newAppRecordRequest := struct {
			RoomName    string
			Description string
			StartTime   string
			EndTime     string
		}{}
		err := json.Unmarshal(payload, &newAppRecordRequest)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse request.\n"+err.Error()+"\n")
			return
		}
		newAppRecordId, err := Services.CreateAppRecordAsUser(
			userId, newAppRecordRequest.RoomName,
			newAppRecordRequest.Description,
			newAppRecordRequest.StartTime, newAppRecordRequest.EndTime)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Creation failed.\n"+err.Error()+"\n")
			return
		} else {
			println("Creation succeed.")
			//formatter.Text(w, http.StatusOK, "Creation succeed.\n")
			newAppRecord, _ := Models.GetAppRecord(newAppRecordId)
			formatter.JSON(w, http.StatusOK, *newAppRecord)
		}
	}
}

func UpdateAppRecordAsUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		userId := vars["userId"]
		appRecordIdStr := vars["appRecordId"]
		appRecordId, err := strconv.Atoi(appRecordIdStr)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse request.\n")
			return
		}
		updateAppRecordRequest := struct {
			UpdateField string
			NewValue    string
			ValueType   string
		}{}
		payload, _ := ioutil.ReadAll(req.Body)
		err = json.Unmarshal(payload, &updateAppRecordRequest)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse request.\n"+err.Error()+"\n")
			return
		}
		switch updateAppRecordRequest.ValueType {
		case "string":
			_, err = Services.UpdateAppRecordAsUser(userId, appRecordId, updateAppRecordRequest.UpdateField, updateAppRecordRequest.NewValue)
		case "int":
			var newValueInt int
			newValueInt, err = strconv.Atoi(updateAppRecordRequest.NewValue)
			if err != nil {
				formatter.Text(w, http.StatusBadRequest, "Failed to parse request.\n"+err.Error()+"\n")
				return
			}
			_, err = Services.UpdateAppRecordAsUser(userId, appRecordId, updateAppRecordRequest.UpdateField, newValueInt)
		}
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to update appRecord.\n"+err.Error()+"\n")
			return
		} else {
			println("Update succeed.")
			//formatter.Text(w, http.StatusOK, "Update succeed.\n")
			appRecord, _ := Models.GetAppRecord(appRecordId)
			formatter.JSON(w, http.StatusOK, *appRecord)
		}
	}
}

func GetAppRecordByRoomNameAsUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		roomName := vars["roomName"]
		userId := vars["userId"]
		appRecords, err := Services.GetAppRecordsByRoomNameAsUser(userId, roomName)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to get appRecords.\n"+err.Error()+"\n")
			return
		} else {
			info := struct {
				Count      int `json:"Count"`
				AppRecords []Models.AppRecord
			}{
				Count:      len(appRecords),
				AppRecords: appRecords,
			}
			formatter.JSON(w, http.StatusOK, info)
		}
	}
}
