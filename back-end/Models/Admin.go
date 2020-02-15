package Models

import (
	"encoding/json"
	"errors"
)

type Admin struct {
	AdminId  string //key
	Password string
	Email    string
}

func (admin *Admin) String() string {
	b, err := json.Marshal(*admin)
	if err == nil {
		return string(b)
	} else {
		return ""
	}
}

//为新管理员申请账号
func (admin *Admin) RegisterNewAdmin(adminId string, password string, email string) (Admin, error) {
	newAdmin := Admin{
		AdminId:  adminId,
		Password: password,
		Email:    email,
	}
	_, err := AddAdmin(&newAdmin)
	if err == nil {
		return newAdmin, nil
	} else {
		return Admin{}, err
	}
}

func (admin *Admin) GetAppRecordsByRoomName(roomName string) ([]AppRecord, error) {
	if !isRoomExist(roomName) {
		return []AppRecord{}, errors.New("Room doesn't exist.")
	} else {
		room := Rooms[roomName]
		return room.getRoomAppRecords(), nil
	}
}

//获取所有申请记录
func (admin *Admin) GetAllAppRecords() []AppRecord {
	ids := GetAppRecordIds()
	return GetAppRecordsByIds(ids)
}

func (admin Admin) UpdateAppRecord(appRecordId int, updateField string, newValue interface{}) (int, error) {
	return updateAppRecordByAdmin(admin.AdminId, appRecordId, updateField, newValue)
}
