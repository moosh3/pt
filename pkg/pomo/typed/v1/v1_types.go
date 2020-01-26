package pomo

// Pomo ...
type Pomo struct {
	UUID           string `json:"uuid,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAd      string `json:"updated_at,omitempty"`
	Description    string `json:"description,omitempty"`
	StartedAt      string `json:"started_at,omitempty"`
	EndedAt        string `json:"ended_at,omitempty"`
	LocalStartedAt string `json:"local_started_at,omitempty"`
	LocalEndedAt   string `json:"local_ended_at,omitempty"`
	Length         int    `json:"length,omitempty"`
	Abandonded     bool   `json:"abandoned,omitempty"`
	Manual         bool   `json:"manual,omitempty"`
}

type PomoList struct {
	pomos []Pomo
}

// Todo ...
type Todo struct {
	UUID               string `json:"uuid,omitempty"`
	CreatedAt          string `json:"created_at,omitempty"`
	UpdatedAd          string `json:"updated_at,omitempty"`
	Description        string `json:"description,omitempty"`
	Notice             string `json"notice,omitempty"`
	Pin                bool   `json:"pin,omitempty"`
	Completed          bool   `json:"completed,omitmpy"`
	CompletedAt        string `json:"completed_at,omitempty"`
	RepeatType         string `json:"repeat_type,omitempty"`
	RepeatTime         string `json:"remind_time,omitempty"`
	estimatedPomoCount string `json:"estimated_pomo_count,omitempty"`
	CostedPomoCount    string `json:"costed_pomo_count,omitempty"`
	SubTodos           string `json:"sub_todos,omitempty"`
}

type TodoList struct {
	todos []Todo
}

// SubTodo ...
type SubTodo struct {
	UUID        string `json:"uuid,omitempty"`
	ParentUUID  string `json:"parent_uuid,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAd   string `json:"updated_at,omitempty"`
	Description string `json:"description,omitempty"`
	Completed   bool   `json:"completed,omitmpy"`
	CompletedAt string `json:"completed_at,omitempty"`
}

type SubTodoList struct {
	subtodos []SubTodo
}
