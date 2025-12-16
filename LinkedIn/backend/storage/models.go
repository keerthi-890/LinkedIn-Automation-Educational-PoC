package storage

import "time"

type Config struct {
	ID    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Profile struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Location  string    `json:"location"`
	ProfileID string    `json:"profile_id"` // External ID from mock site
	Status    string    `json:"status"`     // New, Connected, Messaged
	CreatedAt time.Time `json:"created_at"`
}

type ActivityLog struct {
	ID        int       `json:"id"`
	Action    string    `json:"action"`
	Details   string    `json:"details"`
	Timestamp time.Time `json:"timestamp"`
}

type Stats struct {
	ProfilesFound       int `json:"profiles_found"`
	RequestsSent        int `json:"requests_sent"`
	AcceptedConnections int `json:"accepted_connections"`
	MessagesSent        int `json:"messages_sent"`
}
