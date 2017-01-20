package app

// NetlifyWebhook represents Netlify deploy webhook data object.
type NetlifyWebhook struct {
	Error     string `json:"error_message"`
	URL       string `json:"url"`
	Name      string `json:"name"`
	State     string `json:"state"`
	Branch    string `json:"branch"`
	DeployURL string `json:"deploy_ssl_url"`
}
