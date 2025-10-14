package main

import (
	"workshop-registration/config"
	"workshop-registration/models"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	att1 := models.InsertAttendee("Yuvraj", "yuvraj@example.com")
	att2 := models.InsertAttendee("Akash", "akash@example.com")

	ws1 := models.InsertWorkshop("Go Concurrency", 2)
	ws2 := models.InsertWorkshop("Building APIs with Go", 3)

	models.RegisterAttendee(int(att1), int(ws1))
	models.RegisterAttendee(int(att2), int(ws1))
	models.RegisterAttendee(int(att2), int(ws2))

	println("\nAttendees for Workshop 1:")
	models.GetWorkshopAttendees(int(ws1))
}
