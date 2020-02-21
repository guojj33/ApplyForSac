package Models

import (
	"errors"
)

var Users map[string]*User = make(map[string]*User)          //UserId -> User
var Admins map[string]*Admin = make(map[string]*Admin)       //AdminId -> Admin
var Rooms map[string]*Room = make(map[string]*Room)          //RoomName -> Room
var AppRecords map[int]*AppRecord = make(map[int]*AppRecord) //AppRecordId -> AppRecord
var NextAppRecordId int                                      //下一个 AppRecords 的 Id，递增	可能得加个锁

//资源访问

func GetAllUsers() []User {
	users := make([]User, len(Users))
	index := 0
	for k := range Users {
		users[index] = *Users[k]
		index++
	}
	return users
}

func GetAllAdmins() []Admin {
	admins := make([]Admin, len(Admins))
	index := 0
	for k := range Admins {
		admins[index] = *Admins[k]
		index++
	}
	return admins
}

func GetAllRooms() []Room {
	rooms := make([]Room, len(Rooms))
	index := 0
	for k := range Rooms {
		rooms[index] = *Rooms[k]
		index++
	}
	return rooms
}

func GetAllAppRecords() []AppRecord {
	appRecords := make([]AppRecord, len(AppRecords))
	index := 0
	for k := range AppRecords {
		appRecords[index] = *AppRecords[k]
		index++
	}
	return appRecords
}

func GetUser(userId string) (*User, error) {
	if Users[userId] == nil {
		return nil, errors.New("User doesn't exist.")
	} else {
		return Users[userId], nil
	}
}

func GetAdmin(adminId string) (*Admin, error) {
	if Admins[adminId] == nil {
		println(adminId)
		return nil, errors.New("Admin doesn't exist.")
	} else {
		return Admins[adminId], nil
	}
}

func GetRoom(roomName string) (*Room, error) {
	if Rooms[roomName] == nil {
		return nil, errors.New("Room doesn't exist.")
	} else {
		return Rooms[roomName], nil
	}
}

func GetAppRecord(appRecordId int) (*AppRecord, error) {
	if AppRecords[appRecordId] == nil {
		return nil, errors.New("AppRecord doesn't exist.")
	} else {
		return AppRecords[appRecordId], nil
	}
}

//服务实现

func GetAppRecordIds() []int {
	index := 0
	appRecordIds := make([]int, len(AppRecords))
	for id := range AppRecords {
		appRecordIds[index] = id
		index++
	}
	return appRecordIds
}

func GetAppRecordsByIds(ids []int) []AppRecord {
	appRecords := []AppRecord{}
	for _, i := range ids { //少写了 _, 导致获取错误
		appRecords = append(appRecords, *AppRecords[i])
	}
	return appRecords
}

//将申请添加到总记录，返回添加后记录总数
func AddAppRecordToAppRecords(appRecord *AppRecord) int {
	AppRecords[appRecord.AppRecordId] = appRecord
	ar := AppRecords[appRecord.AppRecordId]
	AddAppRecordToDB(ar)
	return len(AppRecords)
}

//将申请添加到对应房间记录，返回添加后记录总数
func AddAppRecordToRoom(appRecordId int, roomName string) int {
	oldAppRecords := Rooms[roomName].AppRecordIds
	Rooms[roomName].AppRecordIds = append(oldAppRecords, appRecordId)
	UpdateRoomToDB(Rooms[roomName])
	return len(Rooms[roomName].AppRecordIds)
}

//将申请添加到对应用户的记录，返回添加后记录总数
func AddAppRecordToUser(appRecordId int, userId string) int {
	oldAppRecords := Users[userId].AppRecordIds
	Users[userId].AppRecordIds = append(oldAppRecords, appRecordId)
	UpdateUserToDB(Users[userId])
	return len(Users[userId].AppRecordIds)
}

func AddAppRecord(appRecord *AppRecord) (int, int, int) {
	return AddAppRecordToAppRecords(appRecord),
		AddAppRecordToRoom(appRecord.AppRecordId, appRecord.RoomName),
		AddAppRecordToUser(appRecord.AppRecordId, appRecord.ApplyUserId)
}

func AddUser(user *User) (int, error) {
	if Users[user.UserId] != nil || user.UserId == "SAC" {
		return len(Users), errors.New("User id has already been used.")
	} else {
		Users[user.UserId] = user
		return len(Users), AddUserToDB(user)
	}
}

func AddAdmin(admin *Admin) (int, error) {
	if Admins[admin.AdminId] != nil || admin.AdminId == "SAC" {
		return len(Admins), errors.New("Admin id has already been used.")
	} else {
		Admins[admin.AdminId] = admin
		return len(Admins), AddAdminToDB(admin)
	}
}

func AddRoom(room *Room) (int, error) {
	if Rooms[room.RoomName] != nil {
		return len(Rooms), errors.New("Room name has already been used.")
	} else {
		Rooms[room.RoomName] = room
		return len(Rooms), AddRoomToDB(room)
	}
}

func isAccountExist(id string) bool {
	return Admins[id] != nil || Users[id] != nil
}

func isRoomExist(roomName string) bool {
	return Rooms[roomName] != nil
}

func isAppRecordExist(appRecordId int) bool {
	return AppRecords[appRecordId] != nil
}

func GetAccountType(id string) AccountType {
	if Users[id] != nil {
		return AccountType_User
	} else if Admins[id] != nil {
		return AccountType_Admin
	} else {
		return AccountType_NotExist
	}
}

func CheckAccountInfo(id string, password string) (AccountType, error) {
	accountType := GetAccountType(id)
	switch accountType {
	case AccountType_NotExist:
		return accountType, errors.New("Account doesn't exist.")
	case AccountType_User:
		if Users[id].Password == password {
			return accountType, nil
		} else {
			return accountType, errors.New("Password is incorrect.")
		}
	case AccountType_Admin:
		if Admins[id].Password == password {
			return accountType, nil
		} else {
			return accountType, errors.New("Password is incorrect.")
		}
	}
	return accountType, errors.New("Unknown error.")
}

//添加用户、管理员、房间
func InitDB() {
	OpenDB("./myDatabase.db")
	var err error
	Users, Admins, Rooms, AppRecords, NextAppRecordId, err = LoadDataFromDB()
	if err == nil {
		var PianoRoom1 Room
		PianoRoom1.RoomName = "PianoRoom1"
		AddRoom(&PianoRoom1)

		var admin Admin
		admin.AdminId = "Admin"
		admin.Password = "123"
		admin.Email = "Admin@admin.com"
		AddAdmin(&admin)

		var JJ User
		JJ.UserId = "JJ"
		JJ.Password = "123"
		JJ.Email = "JJ@JJ.com"
		AddUser(&JJ)
	}
}
