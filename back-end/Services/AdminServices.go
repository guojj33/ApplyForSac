package Services

import (
	"errors"
	"time"

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

func CreateNewRoomAsAdmin(roomName string) (Models.Room, error) {
	newRoom := Models.Room{
		RoomName: roomName,
	}
	_, err := Models.AddRoom(&newRoom)
	if err != nil {
		newRoom = Models.Room{}
	}
	return newRoom, err
}

func CreateAppRecordAsAdmin(adminId string, roomName string, description string, startTime string, endTime string) (int, error) {
	curAdmin, err := GetCurAdmin(adminId)
	if err != nil {
		return 0, err
	} else {
		sTime, serr := time.Parse("2006-01-02-15:04:05", startTime)
		eTime, eerr := time.Parse("2006-01-02-15:04:05", endTime)
		if serr != nil || eerr != nil {
			return -1, errors.New(serr.Error() + "\n" + eerr.Error())
		} else {
			applyUsingTime := Models.TimeDuration{
				StartTime: sTime,
				EndTime:   eTime,
			}
			return curAdmin.CreateAppRecord(roomName, description, applyUsingTime)
		}
	}
}
