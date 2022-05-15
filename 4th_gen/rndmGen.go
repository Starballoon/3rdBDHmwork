package main

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main/dal/model"
	"main/dal/query"
	"math/rand"
	"strconv"
)

func main() {
	dsn := "root:LLQtT$3v@tcp(127.0.0.1:3306)/homework?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 产生100个随机用户，没有使用OnConflict
	// generateUsers(db)
	// 计数
	fmt.Println(GetMaxVersionCount(db))
}

func GetMaxVersionCount(db *gorm.DB) (int, error) {
	var count int
	db.Raw("SELECT COUNT(uuid) " +
		"FROM user, (SELECT MAX(version) as maxVer FROM user) AS A " +
		"WHERE user.version=A.maxVer").Scan(&count)
	return count, nil
}

func generateUsers(db *gorm.DB) error {
	total := 100
	for i := 0; i < total; i += 1 {
		_ = addUser(db, &model.User{
			UUID:    strconv.Itoa(i),
			Name:    "test",
			Age:     int64(i),
			Version: int64(0),
		})
	}
	for j := rand.Intn(total / 2); j > -1; j -= 1 {
		_ = addUser(db, &model.User{
			UUID:    strconv.Itoa(j),
			Name:    "test",
			Age:     int64(j),
			Version: int64(0),
		})
	}
	return nil
}

func addUser(db *gorm.DB, user *model.User) error {
	u := query.Use(db).User
	tmp, err := u.WithContext(context.Background()).Where(u.UUID.Eq(user.UUID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if tmp == nil {
		err = u.WithContext(context.Background()).Create(user)
	} else {
		_, err = u.WithContext(context.Background()).Where(u.UUID.Eq(user.UUID)).Updates(map[string]interface{}{
			"name":    user.Name,
			"age":     user.Age,
			"version": tmp.Version + 1,
		})
	}
	return err
}
