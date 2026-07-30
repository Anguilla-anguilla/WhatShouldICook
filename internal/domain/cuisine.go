package domain

type Cuisine struct {
	ID int64
	UserID int64
	Name string
	Description string
}

func (c *Cuisine) Validate() error {
	if c.Name == "" {
		return ErrEmptyName
	}
	return nil
}