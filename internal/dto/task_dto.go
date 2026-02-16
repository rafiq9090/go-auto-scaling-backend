package dto

type CreateTaskRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTaskRequest struct {
	Name      string `json:"name"`
	Completed *bool  `json:"completed"`
}
