package Services

import (
	"../Models"
)

func GetCurAdmin(adminId string) (*Models.Admin, error) {
	curAdmin, err := Models.GetAdmin(adminId)
	if err != nil {
		return nil, err
	} else {
		return curAdmin, nil
	}
}

func RegisterNewAdminAsAdmin(curAdminId string, newAdminId string, password string, email string) (Models.Admin, error) {
	curAdmin, err := GetCurAdmin(curAdminId)
	if err != nil {
		return Models.Admin{}, err
	} else {
		return curAdmin.RegisterNewAdmin(newAdminId, password, email)
	}
}

func GetAppRecordsByRoomNameAsAdmin(adminId string, roomName string) ([]Models.AppRecord, error) {
	curAdmin, err := GetCurAdmin(adminId)
	if err != nil {
		return []Models.AppRecord{}, err
	} else {
		return curAdmin.GetAppRecordsByRoomName(roomName)
	}
}

func GetAllAppRecordsAsAdmin(adminId string) ([]Models.AppRecord, error) {
	curAdmin, err := GetCurAdmin(adminId)
	if err != nil {
		return []Models.AppRecord{}, err
	} else {
		return curAdmin.GetAllAppRecords(), nil
	}
}

func UpdateAppRecordAsAdmin(adminId string, appRecordId int, updateField string, newValue interface{}) (int, error) {
	curAdmin, err := GetCurAdmin(adminId)
	if err != nil {
		return -1, err
	} else {
		return curAdmin.UpdateAppRecord(appRecordId, updateField, newValue)
	}
}
