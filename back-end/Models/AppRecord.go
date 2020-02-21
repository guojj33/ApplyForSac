package Models

import (
	"encoding/json"
	"errors"
)

type AppRecord struct {
	AppRecordId int //key

	RoomName    string
	ApplyUserId string
	Description string

	ReviewAdminId string //审核管理员
	CheckAdminId  string //签到管理员

	ApplyStatus  ApplyStatusType  //当前申请状态，由用户更改
	ReviewStatus ReviewStatusType //当前审核状态，由审核管理员更改

	CheckStatus CheckStatusType //当前签到状态，由签到管理员更改

	ApplyUsingTime TimeDuration //申请使用时间
}

func (appRecord *AppRecord) String() string {
	b, err := json.Marshal(appRecord)
	if err == nil {
		return string(b)
	} else {
		return ""
	}
}

//检查申请是否无冲突
func (appRecord AppRecord) isAppRecordValid() error {
	//判断房间是否存在
	if !isRoomExist(appRecord.RoomName) {
		return errors.New("Room doesn't exists.")
	}
	//判断时间是否有效
	if !isTimeDurationValid(appRecord.ApplyUsingTime) {
		return errors.New("Invalid duration.")
	}
	//判断时间是否重叠
	//正在申请且没有被拒绝的申请
	existAppRecords := Rooms[appRecord.RoomName].getRoomAppRecords()
	for _, tmpAppRecord := range existAppRecords {
		if tmpAppRecord.ApplyStatus == ApplyStatus_Applying &&
			tmpAppRecord.ReviewStatus != ReviewStatus_Rejected &&
			appRecord.ApplyUsingTime.compareWith(tmpAppRecord.ApplyUsingTime) == TimeDurComp_Overlapping {
			return errors.New("Application records clash.")
		}
	}
	return nil
}

func createAppRecordByUser(roomName string, userId string, description string, applyUsingTime TimeDuration) (int, error) {
	var appRecord *AppRecord = new(AppRecord)

	//修改申请数据
	appRecord.AppRecordId = NextAppRecordId

	appRecord.RoomName = roomName
	appRecord.ApplyUserId = userId
	appRecord.Description = description
	appRecord.ApplyUsingTime = applyUsingTime

	appRecord.ApplyStatus = ApplyStatus_Applying
	appRecord.ReviewStatus = ReviewStatus_Waiting

	err := AppRecord.isAppRecordValid(*appRecord)
	if err == nil {
		println("NextAppRecordId:", NextAppRecordId)
		NextAppRecordId++
		//将申请记录添加到各个对象中
		AddAppRecord(appRecord) //包含数据库操作
		return appRecord.AppRecordId, nil
	} else {
		return -1, err
	}
}

//管理员创建的申请，申请人为实际不存在的 SAC ，申请会加入到 appRecords 和 room 中
//申请的目的是暂停普通用户在该时间对该房间进行申请
//如果已存在该事件的用户申请，自动拒绝其申请
func createAppRecordByAdmin(roomName string, description string, applyUsingTime TimeDuration) (int, error) {
	var appRecord *AppRecord = new(AppRecord)
	appRecord.AppRecordId = NextAppRecordId
	appRecord.RoomName = roomName
	appRecord.ApplyUserId = "SAC"
	appRecord.ApplyUsingTime = applyUsingTime
	appRecord.Description = description

	appRecord.ApplyStatus = ApplyStatus_Applying
	appRecord.ReviewStatus = ReviewStatus_Accepted
	appRecord.CheckStatus = CheckStatus_Normal

	//拒绝所有时间冲突的其他用户的申请
	//判断房间是否存在
	if !isRoomExist(appRecord.RoomName) {
		return -1, errors.New("Room doesn't exists.")
	}
	//判断时间是否有效
	if !isTimeDurationValid(appRecord.ApplyUsingTime) {
		return -1, errors.New("Invalid duration.")
	}
	//判断时间是否重叠
	//正在申请且没有被拒绝的申请
	existAppRecords := Rooms[appRecord.RoomName].getRoomAppRecords()
	for _, tmpAppRecord := range existAppRecords {
		if tmpAppRecord.ApplyStatus == ApplyStatus_Applying &&
			tmpAppRecord.ReviewStatus != ReviewStatus_Rejected &&
			appRecord.ApplyUsingTime.compareWith(tmpAppRecord.ApplyUsingTime) == TimeDurComp_Overlapping {
			if tmpAppRecord.ApplyUserId == "SAC" {
				return -1, errors.New("Application records clash.")
			}
			_, err := updateAppRecordByAdmin("SAC", tmpAppRecord.AppRecordId, "ReviewStatus", int(ReviewStatus_Rejected))
			if err != nil {
				return -1, errors.New("Update appRecord failed.")
			}
		}
	}
	//拒绝成功
	NextAppRecordId++
	AddAppRecordToAppRecords(appRecord)
	AddAppRecordToRoom(appRecord.AppRecordId, appRecord.RoomName)
	return appRecord.AppRecordId, nil
}

func updateAppRecordByUser(userId string, appRecordId int, updateField string, newValue interface{}) (int, error) {
	failed := false
	var err error
	switch updateField {
	case "ApplyStatus":
		_, err = cancelAppRecordByUser(appRecordId, userId)
		if err != nil {
			failed = true
		}
	default:
		failed = true
	}
	if failed {
		return -1, errors.New("User: Update failed.\n" + err.Error())
	} else {
		return appRecordId, nil
	}
}

func updateAppRecordByAdmin(adminId string, appRecordId int, updateField string, newValue interface{}) (int, error) {
	failed := false
	var err error
	switch updateField {
	case "ReviewStatus":
		newValueRst := ReviewStatusType(newValue.(int))
		err = reviewAppRecordByAdmin(appRecordId, adminId, newValueRst)
		if err != nil {
			failed = true
		}
	case "CheckStatus":
		newValueCst := CheckStatusType(newValue.(int))
		err = checkAppRecordByAdmin(appRecordId, adminId, newValueCst)
		if err != nil {
			failed = true
		}
	default:
		failed = true
	}
	if failed {
		return -1, errors.New("Admin: Update failed.\n" + err.Error())
	} else {
		return appRecordId, nil
	}
}

func cancelAppRecordByUser(applyRecordId int, userId string) (int, error) {
	appRecord := AppRecords[applyRecordId]
	if appRecord.ApplyUserId == userId { //申请发起者是本人
		if appRecord.ApplyStatus == ApplyStatus_Canceled { //已经被取消
			return appRecord.AppRecordId, errors.New("The application record has already been canceled.")
		} else {
			appRecord.ApplyStatus = ApplyStatus_Canceled
			UpdateAppRecordToDB(appRecord) //数据库操作
			return appRecord.AppRecordId, nil
		}
	} else {
		return -1, errors.New("Cancellation denied.")
	}
}

func reviewAppRecordByAdmin(applyRecordId int, adminId string, reviewStatus ReviewStatusType) error {
	appRecord := AppRecords[applyRecordId]
	if appRecord.ApplyStatus == ApplyStatus_Applying { //申请未被取消
		if appRecord.ReviewStatus == reviewStatus { //不允许做状态相同的更改
			switch appRecord.ReviewStatus {
			case ReviewStatus_Accepted:
				return errors.New("The application record reviewing has already been accepted.")
			case ReviewStatus_Rejected:
				return errors.New("The application record reviewing has already been rejected.")
			}
		} else {
			appRecord.ReviewStatus = reviewStatus
			appRecord.ReviewAdminId = adminId
			UpdateAppRecordToDB(appRecord) //数据库操作
			return nil
		}
	} else {
		return errors.New("The application record reviewing has been canceled by the user.")
	}
	return errors.New("Unknown error.")
}

func checkAppRecordByAdmin(applyRecordId int, adminId string, checkStatus CheckStatusType) error {
	appRecord := AppRecords[applyRecordId]
	if appRecord.ApplyStatus == ApplyStatus_Applying { //申请未被取消
		if appRecord.ReviewStatus == ReviewStatus_Accepted { //申请已通过
			if appRecord.CheckStatus == checkStatus { //不允许做状态相同的更改
				switch appRecord.CheckStatus {
				case CheckStatus_Normal:
					return errors.New("The application record checking has already been set to NORMAL")
				case CheckStatus_Late:
					return errors.New("The application record checking has already been set to LATE.")
				}
			} else {
				appRecord.CheckStatus = checkStatus
				appRecord.CheckAdminId = adminId
				UpdateAppRecordToDB(appRecord) //数据库操作
				return nil
			}
		} else {
			return errors.New("The application record checking has not been accepted.")
		}
	} else {
		return errors.New("The application record checking has been canceled by the user.")
	}
	return errors.New("Unknown error.")
}
