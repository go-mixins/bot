package bot

// User is modelled after Telegram chat user but with platform-specific detail
// isolated in Meta field
type User struct {
	FirstName    string
	LastName     string
	UserName     string
	LanguageCode string
	Meta
}
