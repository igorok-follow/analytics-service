package models

type RegisterEvent struct {
	ApiKey string   `json:"api_key"`
	Events []*Event `json:"events"`
}

type Event struct {
	UserId    string `json:"user_id"` // string with len !< 5 chars
	EventType string `json:"event_type"`
	Unix      int64  `json:"time"`
}
