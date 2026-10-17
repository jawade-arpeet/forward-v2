package dao

type PlatformInfo struct {
	ID      string `db:"id"`
	Name    string `db:"name"`
	LogoURL string `db:"logo_url"`
}
