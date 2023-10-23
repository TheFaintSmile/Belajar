package dto

type (
	CourseListResponse struct {
		ID       uint   `json:"id" validate:"required"`
		Name     string `json:"name" validate:"required"`
		Lecturer string `json:"lecturer" validate:"required"`
	}
)

type AddCourseInput struct {
	Name     string `json:"name" binding:"required"`
	Lecturer string `json:"lecturer" binding:"required"`
	LevelID  uint   `json:"level_id" gorm:"foreign_key" binding:"required"`
}

type UpdateCourseInformationInput struct {
	Name     string `json:"name" binding:"required"`
	Lecturer string `json:"lecturer" binding:"required"`
	LevelID  uint   `json:"level_id" gorm:"foreign_key" binding:"required"`
}

type UpdateWeekInCourseInput struct {
	Name     string `json:"name" binding:"required"`
}