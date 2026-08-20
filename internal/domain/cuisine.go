package domain

import "strings"

type Cuisine struct {
	ID          int64
	UserID      int64
	Name        string
	Description string
}

func (c *Cuisine) Validate() error {
	if c.UserID <= 0 {
		return ErrInvalidUser
	}
	str := strings.ReplaceAll(c.Name, " ", "")
	if str == "" {
		return ErrEmptyName
	}
	return nil
}
