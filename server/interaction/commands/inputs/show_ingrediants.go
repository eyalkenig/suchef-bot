package inputs

type ShowIngredients struct{
	CourseID int64
}

func NewShowIngredients(courseID int64) *ShowIngredients {
	return &ShowIngredients{CourseID: courseID}
}
