package providers

import (
	"database/sql"
	"fmt"
	"github.com/eyalkenig/suchef-bot/server/models"
	_ "github.com/go-sql-driver/mysql"
)

type BotDataProvider struct {
	db *sql.DB
}

const FETCH_USER_QUERY = "SELECT id, account_id, external_user_id, first_name, last_name, profile_pic, locale, timezone, gender, diet_id, sensitivity_id FROM users WHERE (account_id = ? AND external_user_id = ?)"
const CREATE_USER_QUERY = "INSERT INTO users (id, account_id, external_user_id, first_name, last_name, profile_pic, locale, timezone, gender, created_at, updated_at) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())"

const FETCH_LAST_INTERACTION_QUERY = "SELECT id FROM interactions WHERE (user_id = ?) ORDER BY id DESC LIMIT 1"
const FETCH_INTERACTION_STATE_ID_QUERY = "select state_id FROM interaction_states WHERE (interaction_id = ?)"
const ADD_INTERACTION_QUERY = "INSERT INTO interactions (id, user_id, created_at, updated_at) VALUES (NULL, ?, NOW(), NOW())"
const ADD_INTERACTION_STATE_QUERY = "INSERT INTO interaction_states (id, interaction_id, state_id, created_at, updated_at) VALUES (NULL, ?, ?, NOW(), NOW())"
const UPDATE_INTERACTION_STATE_QUERY = "UPDATE interaction_states SET state_id = ? WHERE (interaction_id = ?)"

const UPDATE_USER_DIET = "UPDATE users SET diet_id = ? WHERE (id = ?)"
const UPDATE_USER_SENSITIVITY = "UPDATE users SET sensitivity_id = ? WHERE (id = ?)"

func NewBotDataProvider(connParams DBConnectionParams) (dataProvider *BotDataProvider, err error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", connParams.User, connParams.Password, connParams.Address, connParams.DBName)
	fmt.Println("connecting to mysql: " + connectionString)
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

func (dataProvider *BotDataProvider) InitState(userID, initialStateID int64) (err error) {
	result, err := dataProvider.db.Exec(ADD_INTERACTION_QUERY, userID)
	if err != nil {
		return err
	}

	interactionID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	_, err = dataProvider.db.Exec(ADD_INTERACTION_STATE_QUERY, interactionID, initialStateID)
	if err != nil {
		return err
	}
	return nil
}

func (dataProvider *BotDataProvider) FetchCurrentState(userID int64) (int64, error) {
	interactionID, err := dataProvider.getLastUserInteraction(userID)

	if err != nil {
		return -1, err
	}
	row := dataProvider.db.QueryRow(FETCH_INTERACTION_STATE_ID_QUERY, interactionID)

	var stateID int64
	err = row.Scan(&stateID)

	if err != nil {
		return -1, err
	}

	return stateID, nil
}

func (dataProvider *BotDataProvider) SetCurrentState(userID, stateID int64) (err error) {
	interactionID, err := dataProvider.getLastUserInteraction(userID)

	if err != nil {
		return err
	}

	_, err = dataProvider.db.Exec(UPDATE_INTERACTION_STATE_QUERY, stateID, interactionID)

	return err
}

func (dataProvider *BotDataProvider) SetUserDiet(userID, dietTypeID int64) (err error) {
	_, err = dataProvider.db.Exec(UPDATE_USER_DIET, dietTypeID, userID)
	return err
}

func (dataProvider *BotDataProvider) SetSensitivity(userID, sensitivityTypeID int64) (err error) {
	_, err = dataProvider.db.Exec(UPDATE_USER_SENSITIVITY, sensitivityTypeID, userID)
	return err
}

func (dataProvider *BotDataProvider) Close() error {
	return dataProvider.db.Close()
}

func (dataProvider *BotDataProvider) getLastUserInteraction(userID int64) (int64, error) {
	row := dataProvider.db.QueryRow(FETCH_LAST_INTERACTION_QUERY, userID)

	var interactionID int64
	err := row.Scan(&interactionID)

	if err != nil {
		return -1, err
	}

	return interactionID, nil
}
