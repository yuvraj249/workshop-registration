package models

import (
	"fmt"
	"workshop-registration/config"
)

func InsertAttendee(name, email string) int64 {
	query := "INSERT INTO attendees(name,string) values(?,?)"
	result, err := config.DB.Exec(query, name, email)
	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()
	fmt.Printf("Id inserted : %d", id)
	return id
}

func GetAttendee() {
	rows, err := config.DB.Query("SELECT * FROM attendees")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name, email string
		rows.Scan(&id, &name, &email)
		fmt.Printf("Id: %d, Name: %s, Email: %s\n", id, name, email)

	}
}
