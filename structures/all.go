package structures
import ()

type GetPointsResult struct {
	UserID       uint
	TotalPoints int
}
type Body struct{
	UserID uint
	CourseID uint	
	Key string
	Value string
}
type Correct struct{
	IsCorrect bool
}