package Models

import (
	"errors"
	"sort"
	"strings"
)

var Users map[string]*User = make(map[string]*User)          //UserId -> User
var Admins map[string]*Admin = make(map[string]*Admin)       //AdminId -> Admin
var Rooms map[string]*Room = make(map[string]*Room)          //RoomName -> Room
var AppRecords map[int]*AppRecord = make(map[int]*AppRecord) //AppRecordId -> AppRecord
var NextAppRecordId int                                      //下一个 AppRecords 的 Id，递增	可能得加个锁
var Comments map[int]*Comment = make(map[int]*Comment)       //CommentId -> Comment
var NextCommentId int                                        //下一个评论的 Id
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

//自定义房间排序
type RoomSlice []Room

func (r RoomSlice) Len() int {
	return len(r)
}

func (r RoomSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r RoomSlice) Less(i, j int) bool {
	if strings.Compare(r[i].RoomName, r[j].RoomName) == -1 {
		return true
	}
	return false
}

func GetAllRooms() []Room {
	rooms := make([]Room, len(Rooms))
	index := 0
	for k := range Rooms {
		rooms[index] = *Rooms[k]
		index++
	}
	//按照名字排序
	sort.Sort(RoomSlice(rooms))
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

//自定义评论排序
type CommentsSlice []Comment

func (c CommentsSlice) Len() int {
	return len(c)
}

func (c CommentsSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c CommentsSlice) Less(i, j int) bool {
	return c[i].CommentId < c[j].CommentId
}

func GetAllComments() []Comment {
	comments := make([]Comment, len(Comments))
	index := 0
	for k := range Comments {
		comments[index] = *Comments[k]
		index++
	}
	//按照 id 排序
	sort.Sort(CommentsSlice(comments))
	return comments
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
func AddAppRecordToAppRecords(appRecord *AppRecord) (int, error) {
	err := AddAppRecordToDB(appRecord)
	if err != nil {
		return -1, err
	}
	AppRecords[appRecord.AppRecordId] = appRecord
	return len(AppRecords), nil
}

//将申请添加到对应房间记录，返回添加后记录总数
func AddAppRecordToRoom(appRecordId int, roomName string) (int, error) {
	oldAppRecords := Rooms[roomName].AppRecordIds
	Rooms[roomName].AppRecordIds = append(oldAppRecords, appRecordId)
	return len(Rooms[roomName].AppRecordIds), UpdateRoomToDB(Rooms[roomName])
}

//将申请添加到对应用户的记录，返回添加后记录总数
func AddAppRecordToUser(appRecordId int, userId string) (int, error) {
	oldAppRecords := Users[userId].AppRecordIds
	Users[userId].AppRecordIds = append(oldAppRecords, appRecordId)
	return len(Users[userId].AppRecordIds), UpdateUserToDB(Users[userId])
}

func AddAppRecord(appRecord *AppRecord) (int, error, int, error, int, error) {
	i1, e1 := AddAppRecordToAppRecords(appRecord)
	i2, e2 := AddAppRecordToRoom(appRecord.AppRecordId, appRecord.RoomName)
	i3, e3 := AddAppRecordToUser(appRecord.AppRecordId, appRecord.ApplyUserId)
	return i1, e1, i2, e2, i3, e3
}

//添加评论，成功则返回评论的 id ，否则返回 -1
func AddComment(name string, content string) (int, error) {
	var comment Comment
	comment.CommentId = NextCommentId
	comment.Name = name
	comment.Content = content
	err := AddCommentToDB(&comment)
	if err != nil {
		return -1, err
	}
	NextCommentId++
	Comments[comment.CommentId] = &comment
	return comment.CommentId, nil
}

//应该先加到数据库无误，再修改内存数据
func AddUser(user *User) (int, error) {
	if Users[user.UserId] != nil || user.UserId == "SAC" || Admins[user.UserId] != nil {
		return len(Users), errors.New("User id has already been used.")
	} else {
		err := AddUserToDB(user)
		if err != nil {
			return -1, err
		}
		Users[user.UserId] = user
		return len(Users), nil
	}
}

func AddAdmin(admin *Admin) (int, error) {
	if Admins[admin.AdminId] != nil || admin.AdminId == "SAC" || Users[admin.AdminId] != nil {
		return len(Admins), errors.New("Admin id has already been used.")
	} else {
		err := AddAdminToDB(admin)
		if err != nil {
			return -1, err
		}
		Admins[admin.AdminId] = admin
		return len(Admins), nil
	}
}

func AddRoom(room *Room) (int, error) {
	if Rooms[room.RoomName] != nil {
		return len(Rooms), errors.New("Room name has already been used.")
	} else {
		err := AddRoomToDB(room)
		if err != nil {
			return -1, err
		}
		Rooms[room.RoomName] = room
		return len(Rooms), nil
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
	Users, Admins, Rooms, AppRecords, NextAppRecordId, Comments, NextCommentId, err = LoadDataFromDB()
	if err == nil {
		roomNames := []string{"钢琴房1", "钢琴房2", "钢琴房3", "钢琴房4", "钢琴房5", "钢琴房6", "钢琴房7", "钢琴房8",
			"音乐沙龙", "文化沙龙", "会议室", "舞蹈练习室", "音乐练习室", "排练创作室", "乐团排练室", "乐队练习室1",
			"乐队练习室2", "绘画室", "讨论室1", "讨论室2", "讨论室3", "讨论室5", "双排键工作室", "棚拍室", "录音室"}

		for _, roomName := range roomNames {
			var r Room
			r.RoomName = roomName
			AddRoom(&r)
		}

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

		AddComment("Admin", "自带管理员\n用户名: Admin\n密码: 123")
	}
}
