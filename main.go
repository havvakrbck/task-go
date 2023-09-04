package main

import (
	"fmt"
	"github.com/havvakrbck/gorm"
	"github.com/havvakrbck/gorm/dialects/mysql"
)

var db *gorm.DB

func main() {
	
	 
	var err error
	db, err = gorm.Open("mysql", "kullanici:parola@tcp(localhost:3306)/veritabani_adi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
 
	
	createPlan(1, "Örnek Plan", "Açıklama", "Aktif", "2023-09-04", "10:00")
	listPlans(1)
	updatePlan(1, "Tamamlandı")
	deletePlan(1)
	if checkForConflict(1, "2023-09-04", "10:00") {
		fmt.Println("Çakışma bulundu!")
	} else {
		fmt.Println("Çakışma bulunamadı.")
	}
}
type Plan struct {
    gorm.Model
    StudentID   int
    Title       string
    Description string
    Status      string
    Date        string
    Time        string
}

func createPlan(studentID int, title string, description string, status string, date string, time string) {
	plan := Plan{
		StudentID:   studentID,
		Title:       title,
		Description: description,
		Status:      status,
		Date:        date,
		Time:        time,
	}

	if err := db.Create(&plan).Error; err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Plan created successfully")
}

func listPlans(studentID int) {
	var plans []Plan

	if err := db.Where("student_id = ?", studentID).Find(&plans).Error; err != nil {
		fmt.Println(err)
		return
	}

	for _, plan := range plans {
		fmt.Printf("ID: %d, Title: %s, Date: %s, Time: %s, Status: %s\n", plan.ID, plan.Title, plan.Date, plan.Time, plan.Status)
	}
}

func updatePlan(planID int, newStatus string) {
	var plan Plan

	if err := db.First(&plan, planID).Error; err != nil {
		fmt.Println(err)
		return
	}

	plan.Status = newStatus

	if err := db.Save(&plan).Error; err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Plan updated successfully")
}

func deletePlan(planID int) {
	if err := db.Where("id = ?", planID).Delete(&Plan{}).Error; err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Plan deleted successfully")
}

func checkForConflict(studentID int, date string, time string) bool {
	var count int

	if err := db.Model(&Plan{}).
		Where("student_id = ? AND date = ? AND time = ?", studentID, date, time).
		Count(&count).Error; err != nil {
		fmt.Println(err)
		return false
	}

	return count > 0
}

