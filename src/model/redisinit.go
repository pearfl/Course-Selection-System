package model

import (
	"course_select/src/database"
)

func CourseToCount() error {

	// 清空redis并将mysql中的表注入redis
	get := database.RedisClient.Get()
	get.Flush()

	var courses []Course

	result := db.Model(&Course{}).Find(&courses)
	err := result.Error

	for _, v := range courses {
		get.Do("HSET", "CourseToCount", v.CourseID, v.Capacity-v.CapSelected)
	}

	return err
}

func CourseToTeacher() error {

	// 清空redis并将mysql中的表注入redis
	get := database.RedisClient.Get()
	get.Flush()

	var binds []Bind

	result := db.Model(&Bind{}).Find(&binds)
	err := result.Error

	for _, v := range binds {
		get.Do("HSET", "CourseToTeacher", v.CourseID, v.TeacherID)
	}

	return err
}

func CourseToName() error {

	// 清空redis并将mysql中的表注入redis
	get := database.RedisClient.Get()
	get.Flush()

	var courses []Course

	result := db.Model(&Course{}).Find(&courses)
	err := result.Error

	for _, v := range courses {
		get.Do("HSET", "CourseToName", v.CourseID, v.Name)
	}

	return err
}