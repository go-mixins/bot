package bot

// User is modelled after Telegram chat user but without specific fields
type User struct {
	FirstName    string
	LastName     string
	UserName     string
	LanguageCode string
}
