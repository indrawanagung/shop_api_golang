package util

import "gorm.io/gorm"

func TruncateDB(db *gorm.DB) error {
	return db.Exec("DELETE FROM public.users").Error
	// DELETE FROM users
}
