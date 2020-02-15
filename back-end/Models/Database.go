package Models

import (
	"log"
	"os"

	"../Code"
	"github.com/boltdb/bolt"
)

var db *bolt.DB

func OpenDB(dbName string) {
	isExisted := false
	var err error
	if _, err := os.Open(dbName); err == nil {
		isExisted = true
	}
	db, err = bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if !isExisted {
		err = db.Update(func(tx *bolt.Tx) error {
			//创建桶
			tx.CreateBucketIfNotExists([]byte("Users"))
			tx.CreateBucketIfNotExists([]byte("Admins"))
			tx.CreateBucketIfNotExists([]byte("Rooms"))
			tx.CreateBucketIfNotExists([]byte("AppRecords"))
			return nil
		})
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func CloseDB() {
	db.Close()
}

func LoadDataFromDB() (map[string]*User, map[string]*Admin, map[string]*Room, map[int]*AppRecord, int, error) {
	var err error
	var users map[string]*User
	var admins map[string]*Admin
	var rooms map[string]*Room
	var appRecords map[int]*AppRecord
	var nextAppRecord int = -1
	if users, err = GetAllUsersFromDB(); err != nil {
		return users, admins, rooms, appRecords, nextAppRecord, err
	}
	if admins, err = GetAllAdminsFromDB(); err != nil {
		return users, admins, rooms, appRecords, nextAppRecord, err
	}
	if rooms, err = GetAllRoomsFromDB(); err != nil {
		return users, admins, rooms, appRecords, nextAppRecord, err
	}
	if appRecords, err = GetAllAppRecordsFromDB(); err != nil {
		return users, admins, rooms, appRecords, nextAppRecord, err
	}
	nextAppRecord = GetCount("AppRecords")
	return users, admins, rooms, appRecords, nextAppRecord, err
}

func GetAllUsersFromDB() (map[string]*User, error) {
	users := make(map[string]*User, GetCount("Users"))
	e := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Users"))
		er := b.ForEach(func(k, v []byte) error {
			curUserId := string(k[:])
			var curUser User
			err := Code.Decode(v, &curUser)
			if err == nil {
				users[curUserId] = &curUser
			}
			return err
		})
		return er
	})
	return users, e
}

func GetAllAdminsFromDB() (map[string]*Admin, error) {
	admins := make(map[string]*Admin, GetCount("Admins"))
	e := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Admins"))
		er := b.ForEach(func(k, v []byte) error {
			curAdminId := string(k[:])
			var curAdmin Admin
			err := Code.Decode(v, &curAdmin)
			if err == nil {
				admins[curAdminId] = &curAdmin
			}
			return err
		})
		return er
	})
	return admins, e
}

func GetAllRoomsFromDB() (map[string]*Room, error) {
	rooms := make(map[string]*Room, GetCount("Rooms"))
	e := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Rooms"))
		er := b.ForEach(func(k, v []byte) error {
			curRoomName := string(k[:])
			var curRoom Room
			err := Code.Decode(v, &curRoom)
			if err == nil {
				rooms[curRoomName] = &curRoom
			}
			return err
		})
		return er
	})
	return rooms, e
}

func GetAllAppRecordsFromDB() (map[int]*AppRecord, error) {
	appRecords := make(map[int]*AppRecord, GetCount("AppRecords"))
	e := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("AppRecords"))
		er := b.ForEach(func(k, v []byte) error {
			curAppRecordId := Code.BytesToInt(k[:])
			var curAppRecord AppRecord
			err := Code.Decode(v, &curAppRecord)
			if err == nil {
				appRecords[curAppRecordId] = &curAppRecord
			}
			return err
		})
		return er
	})
	return appRecords, e
}

func AddUserToDB(user *User) error {
	userIdByte := []byte(user.UserId)
	userByte, err := Code.Encode(user)
	if err != nil {
		return err
	} else {
		err = AddValue([]byte("Users"), userIdByte, userByte)
		return err
	}
}

func AddAdminToDB(admin *Admin) error {
	adminIdByte := []byte(admin.AdminId)
	adminByte, err := Code.Encode(admin)
	if err != nil {
		return err
	} else {
		err = AddValue([]byte("Admins"), adminIdByte, adminByte)
		return err
	}
}

func AddRoomToDB(room *Room) error {
	roomNameByte := []byte(room.RoomName)
	roomByte, err := Code.Encode(room)
	if err != nil {
		return err
	} else {
		err = AddValue([]byte("Rooms"), roomNameByte, roomByte)
		return err
	}
}

func AddAppRecordToDB(appRecord *AppRecord) error {
	appRecordIdByte := Code.IntToBytes(appRecord.AppRecordId)
	appRecordByte, err := Code.Encode(appRecord)
	if err != nil {
		return err
	} else {
		err = AddValue([]byte("AppRecords"), appRecordIdByte, appRecordByte)
		return err
	}
}

func UpdateUserToDB(user *User) error {
	return AddUserToDB(user)
}

func UpdateAdminToDB(admin *Admin) error {
	return AddAdminToDB(admin)
}

func UpdateRoomToDB(room *Room) error {
	return AddRoomToDB(room)
}

func UpdateAppRecordToDB(appRecord *AppRecord) error {
	return AddAppRecordToDB(appRecord)
}

func GetValue(bucketName []byte, key []byte) ([]byte, error) {
	var value []byte
	err := db.View(func(tx *bolt.Tx) error {
		byteLen := len(tx.Bucket([]byte(bucketName)).Get(key))
		value = make([]byte, byteLen)
		copy(value[:], tx.Bucket([]byte(bucketName)).Get(key)[:])
		return nil
	})
	return value, err
}

func AddValue(bucketName []byte, key []byte, value []byte) error {
	err := db.Update(func(tx *bolt.Tx) error {
		if err := tx.Bucket(bucketName).Put(key, value); err != nil {
			return err
		}
		return nil
	})
	return err
}

func ChangeValue(bucketName []byte, key []byte, value []byte) error {
	return AddValue(bucketName, key, value)
}

func GetCount(bucketName string) int {
	count := 0
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucketName))
		b.ForEach(func(k, v []byte) error {
			count++
			return nil
		})
		return nil
	})
	return count
}
