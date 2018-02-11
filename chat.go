package bot

// Chat is modelled after Telegram chat but with platform-specific detail
// isolated in Meta field
type Chat struct {
	Type        string
	Title       string
	UserName    string
	FirstName   string
	LastName    string
	Description string
	Meta
}
