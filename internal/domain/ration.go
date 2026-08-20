package domain

import "time"

type Ration struct {
	ID        int64
	UserID    int64
	Duration  int64
	CreatedAt time.Time
}

func (r *Ration) Validate() error {
	if r.UserID <= 0 {
		return ErrInvalidUser
	}
	if r.Duration <= 0 {
		return ErrInvalidDuration
	}
	return nil
}