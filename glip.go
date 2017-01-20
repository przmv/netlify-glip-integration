package app

// GlipWebhook represents Glip webhook payload data.
type GlipWebhook struct {
	Icon     string `json:"icon,omitempty"`
	Activity string `json:"activity,omitempty"`
	Title    string `json:"title,omitempty"`
	Body     string `json:"body"`
}
