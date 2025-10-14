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
	fmt.Println("Id inserted: ", id)
	return id
}

func GetAttendee() {
	rows, err := config.DB.Query("SELECT * FROM attendees")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, email string
		rows.Scan(&id, &name, &email)
		fmt.Printf("Id: %d, Name: %s, Email: %s\n", id, name, email)

	}
	fmt.Println()
}

func DeleteAttendee(id int64) {
	query := "DELETE FROM attendees WHERE id=?"
	result, err := config.DB.Exec(query, id)
	if err != nil {
		panic(err)
	}
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Deleted attendee id : %d (Rows Affected : %d)\n", id, rowsAffected)
}
