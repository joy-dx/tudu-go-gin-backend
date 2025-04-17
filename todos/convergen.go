//go:build convergen

package todos

//go:generate go run github.com/reedom/convergen@v0.7.0
type Convergen interface {
	// :skip BaseModel
	// :skip Content
	TodoToListing(*Todo) *TodoListing
	// :skip BaseModel
	// :skip Id
	// :skip Complete
	TodoPostToTodo(*TodosPostRequest) *Todo
	// :skip BaseModel
	// :skip Id
	TodoPutToTodo(*TodosPutByIdRequest) *Todo
}
