package Services

import (
	"errors"
	"time"

	"../Models"
)

func GetCurUser(userId string) (*Models.User, error) {
	curUser, err := Models.GetUser(userId)
	if err != nil {
		return nil, err
	} else {
		return curUser, nil
	}
}

func GetAppRecordsByRoomNameAsUser(userId string, roomName string) ([]Models.AppRecord, error) {
	curUser, err := GetCurUser(userId)
	if err != nil {
		return []Models.AppRecord{}, err
	} else {
		return curUser.GetAppRecordsByRoomName(roomName)
	}
}

func GetUserAppRecordsAsUser(userId string) ([]Models.AppRecord, error) {
	curUser, err := GetCurUser(userId)
	if err != nil {
		return []Models.AppRecord{}, err
	} else {
		return curUser.GetUserAppRecords(), nil
	}
}

func CreateAppRecordAsUser(userId string, roomName string, description string, startTime string, endTime string) (int, error) {
	curUser, err := GetCurUser(userId)
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
			return curUser.CreateAppRecord(roomName, description, applyUsingTime)
		}
	}
}

func UpdateAppRecordAsUser(userId string, appRecordId int, updateField string, newValue interface{}) (int, error) {
	curUser, err := GetCurUser(userId)
	if err != nil {
		return 0, err
	} else {
		return curUser.UpdateAppRecord(appRecordId, updateField, newValue)
	}
}

// func CancelAppRecordAsUser(userId string, applyRecordId int) (int, error) {
// 	curUser, err := GetCurUser(userId)
// 	if err != nil {
// 		return 0, err
// 	} else {
// 		return curUser.CancelAppRecord(applyRecordId)
// 	}
// }
