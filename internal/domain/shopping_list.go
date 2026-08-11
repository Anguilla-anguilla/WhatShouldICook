package domain

import "time"

type ShoppingList struct {
	ID        int64
	RationID  int64
	CreatedAt time.Time
}

func (r *ShoppingList) Validate() error {
	if r.RationID == 0 {
		return ErrEmptyRation
	}
	return nil
}
