package user

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type UserInfo struct {
	UserID uuid.UUID `json:"UserID"`
	Age    int       `json:"Age"`
}

func GetInfo(db *sql.DB, userID uuid.UUID) (UserInfo, error) {
	var ui UserInfo
	var age int
	err := db.QueryRow("SELECT age FROM users WHERE id = $1", userID).Scan(&age)
	if err != nil {
		if err == sql.ErrNoRows {
			return ui, fmt.Errorf("user %d not found: %s", userID, err)
		}

		return ui, fmt.Errorf("error querying database: %s", err)
	}

	ui = UserInfo{
		UserID: userID,
		Age:    age,
	}

	return ui, nil
}
