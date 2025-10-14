package models

import (
	"fmt"
	"workshop-registration/config"

	"github.com/go-sql-driver/mysql"
)

func RegisterAttendee(attendee_id, workshop_id int) {
	tx, err := config.DB.Begin()
	if err != nil {
		panic(err)
	}

	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM registrations WHERE workshop_id=? AND status='registered'", workshop_id).Scan(&count)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	var capacity int
	err = tx.QueryRow("SELECT capacity from workshops where id=?", workshop_id).Scan(&capacity)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	if capacity >= count {
		tx.Rollback()
		fmt.Println("Capacity is full")
		return
	}
	_, err = tx.Exec("INSERT INTO registrations(attendee_id, workshop_id) values(?,?)", attendee_id, workshop_id)
	if err != nil {
		if sqlerr, ok := err.(*mysql.MySQLError); ok && sqlerr.Number == 1062 {
			fmt.Println("Attendee already registered. ")
		} else {
			tx.Rollback()
			panic(err)
		}
		return
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	fmt.Println("Attendee Registered")

}

func GetWorkshopAttendees(workshop_id int) {
	query := "SELECT a.id, a.name , a.email from attendees a inner join registrations r on a.id = r.attendee_id where workshop_id=?"
	rows, err := config.DB.Query(query, workshop_id)

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
