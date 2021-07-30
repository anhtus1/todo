package Models

type Todo struct {
	ID uint            `json:"id"`
	Text string       `json:"text"`
	Day string `json:"day"`
	Reminder bool	 `json:"reminder"`
}

func (b *Todo) TableName() string {
	return "todo"
}
