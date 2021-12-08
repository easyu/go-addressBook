package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db *gorm.DB

type UserInfo struct {
	UserId    int `gorm:"primaryKey"`
	UserName  string
	LoginName string
	Password  string
}
type AddressBook struct {
	AddressBookId int `gorm:"primaryKey"`
	Name          string
	PhoneNumber   string
	IsValid       int `gorm:"default:1"`
	OwnerUserId   int
}

func main() {
	db, _ = InitDB()
	r := gin.Default()
	//用户操作
	userGroup := r.Group("/user")
	{
		//查询单个用户
		userGroup.GET("/findOne", findOneUserHandle)
		//新增用户
		userGroup.POST("/save", saveUserHandle)
		//修改用户
		userGroup.POST("/update", updateUserHandle)
		//删除用户
		//userGroup.GET("/delete", deleteUserHandle)
	}
	addressBookGroup := r.Group("/addressBook")
	{
		//新增通讯录
		addressBookGroup.POST("/save", saveOne)
		//修改通讯录
		addressBookGroup.POST("/update", updateOne)
		//删除通讯录
		addressBookGroup.GET("/delete", deleteOne)
		//分页查询通讯录
		//查询单个通讯录详情
		addressBookGroup.GET("/findOne", findOne)
	}
	r.Run(":8001")

}

func deleteOne(c *gin.Context) {
	id := c.Query("id")
	//db.Delete(&AddressBook{}, id)
	var deleteFlag string
	var addressBook AddressBook
	result := db.First(&addressBook, id)
	if result.RowsAffected > 0 {
		result.Row().Scan(&addressBook)
		addressBook.IsValid = 0
		db.Save(&addressBook)
		deleteFlag = "Success"
	} else {
		deleteFlag = "Failed"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "delete " + deleteFlag,
	})
}

func findOne(c *gin.Context) {
	id := c.Query("id")
	var addressBook AddressBook
	result := db.First(&addressBook, id)
	result.Row().Scan(&addressBook)
	c.JSON(http.StatusOK, gin.H{
		"result": addressBook,
	})
}

func updateOne(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		if err != nil {
			log.Fatalf("新增通讯录异常:%v\n", err)
		}
	}
	var addressBook AddressBook
	json.Unmarshal(data, &addressBook)
	tx := db.Save(&addressBook)
	tx.Row().Scan(addressBook)
	fmt.Printf("修改成功%#v\n", &addressBook)
}

func saveOne(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		log.Fatalf("新增通讯录异常:%v\n", err)
	}
	var addressBook AddressBook
	json.Unmarshal(data, &addressBook)
	tx := db.Create(&addressBook)
	tx.Row().Scan(&addressBook)
	log.Printf("新增通讯录成功:%v\n", addressBook)
	c.JSON(http.StatusOK, gin.H{
		"result": addressBook,
	})
}

func updateUserHandle(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		log.Fatalf("修改异常：%v\n", err)
		return
	}
	var userInfo UserInfo
	json.Unmarshal(data, &userInfo)
	tx := db.Save(&userInfo)
	tx.Row().Scan(userInfo)
	fmt.Printf("修改成功%#v\n", &userInfo)
}

func saveUserHandle(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		log.Fatalf("新增异常：%v\n", err)
		return
	}
	var userInfo UserInfo
	json.Unmarshal(data, &userInfo)
	tx := db.Create(&userInfo)
	tx.Row().Scan(userInfo)
	fmt.Printf("保存成功：%#v\n", &userInfo)
}

func findOneUserHandle(c *gin.Context) {
	userId := c.Query("userId")
	var userInfo UserInfo
	result := db.First(&userInfo, userId)
	result.Row().Scan(&userInfo)
	c.JSON(http.StatusOK, gin.H{
		"result": userInfo,
	})
}
