package data

import (
	"time"
)

type UserData struct {
	Created              time.Time `json:"created"`
	Email                string    `json:"email"`
	EmailPublic          bool      `json:"email_public"`
	FullName             string    `json:"full_name"`
	HumanReadableWebsite string    `json:"human_readable_website"`
	ID                   string    `json:"id"`
	LanguagesUsedPublic  bool      `json:"languages_used_public"`
	LastHeartbeat        time.Time `json:"last_heartbeat"`
	LastPublic           string    `json:"last_plugin"`
	LastProject          string    `json:"last_project"`
	Location             string    `json:"location"`
	LoggedTimePublic     bool      `json:"logged_time_public"`
	Modified             time.Time `json:"modified"`
	PhotoPublic          bool      `json:"photo_public"`
	Plan                 string    `json:"plan"`
	Timezone             string    `json:"timezone"`
	Username             string    `json:"username"`
	Website              string    `json:"website"`
}

// User maps to the data returned by resource `/api/v1/users/current`
type User struct {
	Data UserData `json:"data"`
}
