package Models

import (
	"encoding/json"
	"errors"
)

type User struct {
	UserId   string //key
	Password string
	Email    string

	AppRecordIds []int //用户创建的申请的 Id
}

func (user *User) String() string {
	b, err := json.Marshal(*user)
	if err == nil {
		return string(b)
	} else {
		return ""
	}
}

//根据房间名查找申请记录
func (user User) GetAppRecordsByRoomName(roomName string) ([]AppRecord, error) {
	if !isRoomExist(roomName) {
		return []AppRecord{}, errors.New("Room doesn't exist.")
	} else {
		room := Rooms[roomName]
		return room.getRoomAppRecords(), nil
	}
}

//获取用户创建的申请记录
func (user User) GetUserAppRecords() []AppRecord {
	return GetAppRecordsByIds(user.AppRecordIds)
}

//创建申请
func (user User) CreateAppRecord(roomName string, description string, applyUsingTime TimeDuration) (int, error) {
	return createAppRecordByUser(roomName, user.UserId, description, applyUsingTime)
}

//修改申请
func (user User) UpdateAppRecord(appRecordId int, updateField string, newValue interface{}) (int, error) {
	//检查用户是否拥有此申请
	existed := false
	for tmpId := range user.AppRecordIds {
		if tmpId == appRecordId {
			existed = true
			break
		}
	}
	if existed {
		return updateAppRecordByUser(user.UserId, appRecordId, updateField, newValue)
	} else {
		return -1, errors.New("User doesn't have this appRecord.")
	}
}
