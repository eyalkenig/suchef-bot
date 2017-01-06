package providers

import (
	"github.com/eyalkenig/suchef-bot/server/models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type BotDataProvider struct {
	db *sql.DB
}

const FETCH_USER_QUERY = "SELECT id, account_id, external_user_id, first_name, last_name, profile_pic, locale, timezone, gender, diet_id, sensitivity_id FROM users WHERE (account_id = ? AND external_user_id = ?)"
const CREATE_USER_QUERY = "INSERT INTO users (id, account_id, external_user_id, first_name, last_name, profile_pic, locale, timezone, gender, created_at, updated_at) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())"

func NewBotDataProvider(connParams DBConnectionParams) (dataProvider *BotDataProvider, err error){
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", connParams.User, connParams.Password, connParams.Address, connParams.DBName)
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return nil, err
	}

	return &BotDataProvider{db: db}, nil
}

func (dataProvider *BotDataProvider) FetchUser(accountID int64, externalUserID string) (*models.User, error) {
	row := dataProvider.db.QueryRow(FETCH_USER_QUERY, accountID, externalUserID)

	var user models.User
	err := row.Scan(&user.ID,
			&user.AccountID,
			&user.ExternalUserID,
			&user.FirstName,
			&user.LastName,
			&user.ProfilePicture,
			&user.Locale,
			&user.Timezone,
			&user.Gender,
			&user.DietID,
			&user.SensitivityID)

	if err != nil {
		return nil, nil
	}

	return &user, nil
}

func (dataProvider *BotDataProvider) CreateUser(accountID int64, externalUserID, firstName, lastName, gender, profilePicURL, locale string, timezone int) (userID int64, err error) {
	result, err := dataProvider.db.Exec(CREATE_USER_QUERY, accountID, externalUserID, firstName, lastName, profilePicURL, locale, timezone, gender)
	if err != nil {
		return -1, err
	}

	userID, err = result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return userID, nil
}
