package graph

import "github.com/tainmat/graphql/internal/database"

type Resolver struct {
	CategoryDB *database.Category
	CourseDB   *database.Course
}
