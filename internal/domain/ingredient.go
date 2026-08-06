package domain

import "strings"

type Ingredient struct {
	ID   int64
	Name string
}

func (i *Ingredient) Validate() error {
	str := strings.ReplaceAll(i.Name, " ", "")
	if str == "" {
		return ErrEmptyName
	}
	return nil
}
