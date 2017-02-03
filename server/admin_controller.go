package server

import (
	"github.com/eyalkenig/suchef-bot/server/models"
	"github.com/eyalkenig/suchef-bot/server/interfaces/providers"
	"errors"
)

type AdminController struct {
	dataProvider providers.AdminDataProvider
}

func NewAdminController(adminDataProvider providers.AdminDataProvider) *AdminController{
	return &AdminController{dataProvider: adminDataProvider}
}

func (admin *AdminController) AddCourse(accountID int64, course *models.Course) error {
	err := ensureTags(course.Tags)
	if err != nil {
		return err
	}
	newCourseID, err := admin.dataProvider.AddCourse(accountID, course.Name, course.Description, course.ImageURL)
	if err != nil {
		return err
	}
	for _, dietName := range course.Tags[models.DietMetadataTypeName] {
		diet, _ := models.GetDietByName(dietName)
		err = admin.dataProvider.AddCourseMetadata(newCourseID, models.DietMetadataTypeID, diet.ID)
		if err != nil {
			return err
		}
	}
	for _, sensitivityName := range course.Tags[models.SensitivityMetadataTypeName] {
		sensitivity, _ := models.GetSensitivityByName(sensitivityName)
		err = admin.dataProvider.AddCourseMetadata(newCourseID, models.SensitivityMetadataTypeID, sensitivity.ID)
		if err != nil {
			return err
		}
	}
	for _, themeName := range course.Tags[models.ThemeMetadataTypeName] {
		theme, _ := models.GetThemeByName(themeName)
		err = admin.dataProvider.AddCourseMetadata(newCourseID, models.ThemeMetadataTypeID, theme.ID)
		if err != nil {
			return err
		}
	}

	for _, ingredient := range course.Ingredients {
		err = admin.dataProvider.AddCourseIngredient(newCourseID, ingredient.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func ensureTags(tags map[string][]*string) error {
	for _, dietName := range tags[models.DietMetadataTypeName] {
		_, notFound := models.GetDietByName(dietName)
		if notFound != nil {
			return errors.New("diet not found: " + *dietName)
		}
	}
	for _, sensitivityName := range tags[models.SensitivityMetadataTypeName] {
		_, notFound := models.GetSensitivityByName(sensitivityName)
		if notFound != nil {
			return errors.New("sensitivity not found: " + *sensitivityName)
		}
	}
	for _, themeName := range tags[models.ThemeMetadataTypeName] {
		_, notFound := models.GetThemeByName(themeName)
		if notFound != nil {
			return errors.New("theme not found: " + *themeName)
		}
	}
	return nil
}
