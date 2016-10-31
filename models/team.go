package models

// Team maps directly to the teams database table.
type Team struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	URLSlug string `json:"url_slug" db:"url_slug"`
	IconURL string `json:"icon_url" db:"icon_url"`
	LeadID  int64  `json:"lead_id" db:"lead_id"`
}

// TeamJSON has additional fields we will be sending to the client.
type TeamJSON struct {
	Team

	Lead User `json:"lead"`
}