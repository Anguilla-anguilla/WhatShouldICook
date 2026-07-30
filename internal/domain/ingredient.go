package domain

type Ingredient struct {
	ID   int64
	Name string
}

func (i *Ingredient) Validate() error {
	if i.Name == "" {
		return ErrEmptyName
	}
	return nil
}
