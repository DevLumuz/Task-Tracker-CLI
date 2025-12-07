package core

type Task struct {
	Id        string
	Message   string
	Status    string // 'not-done', 'done', 'in-progress', 'deleted'
	CreatedAt string
	UpdatedAt string
}
