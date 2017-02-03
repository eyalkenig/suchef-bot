package providers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/eyalkenig/suchef-bot/server/models"
)

type BotDataProvider struct {
	db *sql.DB
}

type DBConnectionParams struct {
	User     string
	Password string
	Address  string
	DBName   string
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

const FETCH_USER_PREFERENCE = "SELECT diet_id, sensitivity_id FROM users WHERE (id = ?)"

const FETCH_COURSES_METADATA = `SELECT c.id, c.name, c.description, c.main_photo_url, d.name AS diet, s.name AS sensitivity, t.name AS theme
FROM courses_metadata cm
INNER JOIN courses AS c ON c.id = cm.course_id
LEFT JOIN diets AS d ON cm.metadata_type_id = ? AND cm.value_id=d.type_id
LEFT JOIN sensitivities AS s ON cm.metadata_type_id = ? AND cm.value_id=s.type_id
LEFT JOIN themes AS t ON cm.metadata_type_id = ? AND cm.value_id=t.type_id
WHERE account_id = ?
order by c.id`

const FETCH_COURSE_NAME = "SELECT name FROM courses WHERE id = ? LIMIT 1"
const FETCH_INGREDIENTS = "SELECT name FROM ingredients WHERE course_id = ?"

//ADMINS
const FETCH_ACCOUNT_ID_BY_TOKEN = "SELECT id FROM accounts WHERE (page_access_token = ?) LIMIT 1"
const CREATE_COURSE = "INSERT INTO courses VALUES (NULL, ?, ?, ?, ?)"
const ADD_COURSE_METADATA = "INSERT INTO courses_metadata VALUES (NULL, ?, ?, ?)"
const ADD_COURSE_INGREDIENT = "INSERT INTO ingredients VALUES(NULL, ?, ?)"

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

func (dataProvider *BotDataProvider) FetchUserPreference(userID int64) (dietID, sensitivityID int64, err error) {
	row := dataProvider.db.QueryRow(FETCH_USER_PREFERENCE, userID)

	err = row.Scan(&dietID, &sensitivityID)

	return dietID, sensitivityID, err
}

func (dataProvider *BotDataProvider) FetchCourseName(courseID int64) (name string, err error) {
	row := dataProvider.db.QueryRow(FETCH_COURSE_NAME, courseID)

	err = row.Scan(&name)

	return name, err
}

func (dataProvider *BotDataProvider) FetchCourses(accountID int64) ([]*models.Course, error) {
	rows, err := dataProvider.db.Query(FETCH_COURSES_METADATA,
		models.DietMetadataTypeID,
		models.SensitivityMetadataTypeID,
		models.ThemeMetadataTypeID,
		accountID)

	if err != nil {
		return nil, err
	}

	currentID := int64(-1)
	var currentCourse *models.Course
	var courses []*models.Course
	for rows.Next() {
		var id int64
		var courseName string
		var description string
		var mainPhotoURL string
		var diet *string
		var sensitivity *string
		var theme *string
		err = rows.Scan(&id, &courseName, &description, &mainPhotoURL, &diet, &sensitivity, &theme)
		if err != nil {
			return nil, err
		}
		if currentID == int64(-1) {
			currentCourse = models.NewCourse(id, courseName, description, mainPhotoURL)
			currentID = id
		}
		if currentID != id {
			courses = append(courses, currentCourse)
			currentCourse = models.NewCourse(id, courseName, description, mainPhotoURL)
			currentID = id
		}
		applyCourseMetadata(currentCourse, diet, sensitivity, theme)
	}
	courses = append(courses, currentCourse)

	return courses, nil
}

func (dataProvider *BotDataProvider) Close() error {
	fmt.Println("Closing DB..")
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

func (dataProvider *BotDataProvider) FetchIngredients(courseID int64) ([]*models.Ingredient, error) {
	rows, err := dataProvider.db.Query(FETCH_INGREDIENTS, courseID)

	if err != nil {
		return nil, err
	}

	var ingredients []*models.Ingredient
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, &models.Ingredient{Name: name})
	}

	return ingredients, nil
}

func (dataProvider *BotDataProvider) GetAccountID(token string) (int64, error) {
	row := dataProvider.db.QueryRow(FETCH_ACCOUNT_ID_BY_TOKEN, token)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		return int64(-1), err
	}
	return id, nil
}

func (dataProvider *BotDataProvider) AddCourse(accountID int64, name, description, mainImageURL string) (int64, error) {
	result, err := dataProvider.db.Exec(CREATE_COURSE, accountID, name, description, mainImageURL)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func (dataProvider *BotDataProvider) AddCourseMetadata(courseID, metadataTypeID, valueID int64) error {
	_, err := dataProvider.db.Exec(ADD_COURSE_METADATA, courseID, metadataTypeID, valueID)
	return err
}

func (dataProvider *BotDataProvider) AddCourseIngredient(courseID int64, ingredientName string) error {
	_, err := dataProvider.db.Exec(ADD_COURSE_INGREDIENT, courseID, ingredientName)
	return err
}

func applyCourseMetadata(course *models.Course, diet, sensitivity, theme *string) {
	if diet != nil {
		assignToMetadata(course, models.DietMetadataTypeName, diet)
	}
	if sensitivity != nil {
		assignToMetadata(course, models.SensitivityMetadataTypeName, sensitivity)
	}
	if theme != nil {
		assignToMetadata(course, models.ThemeMetadataTypeName, theme)
	}
}

func assignToMetadata(course *models.Course, metadataType string, metadataValue *string) {
	_, ok := course.Tags[metadataType]
	if ok {
		course.Tags[metadataType] = append(course.Tags[metadataType], metadataValue)
	} else {
		course.Tags[metadataType] = []*string{metadataValue}
	}
}
