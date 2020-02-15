package Models

import "encoding/json"

type Room struct {
	RoomName string //key

	AppRecordIds []int
}

func (room *Room) String() string {
	b, err := json.Marshal(*room)
	if err == nil {
		return string(b)
	} else {
		return ""
	}
}

func (room Room) getRoomAppRecords() []AppRecord {
	return GetAppRecordsByIds(room.AppRecordIds)
}
