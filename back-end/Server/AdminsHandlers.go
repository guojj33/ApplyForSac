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

func AdminsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		admins := Models.GetAllAdmins()
		info := struct {
			Count  int            `json:"count"`
			Admins []Models.Admin `json:"admins"`
		}{
			Count:  len(admins),
			Admins: admins,
		}
		formatter.JSON(w, http.StatusOK, info)
	}
}

func GetAdminByAdminIdHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		adminId := vars["adminId"]
		admin, err := Models.GetAdmin(adminId)
		if err == nil {
			formatter.JSON(w, http.StatusOK, *admin)
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}

func GetAllAppRecordsAsAdminHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		adminId := vars["adminId"]
		appRecords, err := Services.GetAllAppRecordsAsAdmin(adminId)
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

func UpdateAppRecordAsAdminHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		adminId := vars["adminId"]
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
			_, err = Services.UpdateAppRecordAsAdmin(adminId, appRecordId, updateAppRecordRequest.UpdateField, updateAppRecordRequest.NewValue)
		case "int":
			var newValueInt int
			newValueInt, err = strconv.Atoi(updateAppRecordRequest.NewValue)
			if err != nil {
				formatter.Text(w, http.StatusBadRequest, "Failed to parse request.\n"+err.Error()+"\n")
				return
			}
			_, err = Services.UpdateAppRecordAsAdmin(adminId, appRecordId, updateAppRecordRequest.UpdateField, newValueInt)
		}
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to update appRecord.\n"+err.Error()+"\n")
			return
		} else {
			formatter.Text(w, http.StatusOK, "Update succeed.\n")
			appRecord, _ := Models.GetAppRecord(appRecordId)
			formatter.JSON(w, http.StatusOK, *appRecord)
		}
	}
}

func GetAppRecordByRoomNameAsAdminHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		roomName := vars["roomName"]
		adminId := vars["adminId"]
		appRecords, err := Services.GetAppRecordsByRoomNameAsAdmin(adminId, roomName)
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

func RegisterNewAdminAsAdminHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		adminId := vars["adminId"]
		payload, _ := ioutil.ReadAll(req.Body)
		newRegisterReq := struct {
			AdminId  string
			Password string
			Email    string
		}{}
		err := json.Unmarshal(payload, &newRegisterReq)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse request.\n"+err.Error()+"\n")
			return
		}
		_, err = Services.RegisterNewAdminAsAdmin(adminId, newRegisterReq.AdminId, newRegisterReq.Password, newRegisterReq.Email)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to create admin.\n"+err.Error()+"\n")
			return
		} else {
			formatter.Text(w, http.StatusOK, "Register succeed.\n")
		}
	}
}

func CreateNewRoomAsAdminHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		newRoomReq := struct {
			RoomName string
		}{}
		err := json.Unmarshal(payload, &newRoomReq)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse request.\n"+err.Error()+"\n")
			return
		}
		_, err = Services.CreateNewRoomAsAdmin(newRoomReq.RoomName)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to create room.\n"+err.Error()+"\n")
			return
		} else {
			formatter.Text(w, http.StatusOK, "Creation succeed.\n")
		}
	}
}

func CreateAppRecordAsAdminHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		adminId := vars["adminId"]
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
		newAppRecordId, err := Services.CreateAppRecordAsAdmin(
			adminId, newAppRecordRequest.RoomName,
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
