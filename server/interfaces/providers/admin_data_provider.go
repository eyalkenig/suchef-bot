package providers

type AdminDataProvider interface {
	AddCourse(accountID int64, name, description, mainImageURL string) (int64, error)
	AddCourseMetadata(courseID, metadataTypeID, valueID int64) error
}
