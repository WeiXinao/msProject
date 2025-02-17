package middleware

import "net/http"

type ProjectAuthMiddleware struct {
	ProjectClient  projectservice.ProjectService
	TaskClient     taskservice.TaskService
}

func NewProjectAuthMiddleware() *ProjectAuthMiddleware {
	return &ProjectAuthMiddleware{}
}

func (m *ProjectAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		

		// Passthrough to next handler if need
		next(w, r)
	}
}
