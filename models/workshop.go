package models

import (
	"fmt"
	"workshop-registration/config"
)

func InsertWorkshop(title string, capacity int) int64 {
	query := "INSERT INTO workshops(title , capacity) VALUES(?,?)"
	result, err := config.DB.Exec(query, title, capacity)
	if err != nil {
		panic(err)
	}

	id, _ := result.LastInsertId()
	fmt.Println("Worskhop id inserted: ", id)
	return id
}

func GetWorkshop() {
	rows, err := config.DB.Query("SELECT * FROM workshops")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var title, capacity string
		rows.Scan(&id, &title, &capacity)
		fmt.Printf("Id: %d, Title: %s, Capacity: %s\n", id, title, capacity)
	}
	fmt.Println()

}

func DeleteWorkshop(id int64) {
	query := "DELETE FROM attendees WHERE id=?"
	result, err := config.DB.Exec(query, id)
	if err != nil {
		panic(err)
	}
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Deleted workshop id : %d (Rows Affected : %d)\n", id, rowsAffected)
}
