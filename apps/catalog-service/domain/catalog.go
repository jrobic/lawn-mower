package domain

import (
	"fmt"
	"strconv"
	"time"
)

type Mower struct {
	ID        string     `json:"id"`
	CreatedAt *Timestamp `json:"createdAt,omitempty"`
	UpdatedAt *Timestamp `json:"updatedAt,omitempty"`
	DeletedAt *Timestamp `json:"deletedAt,omitempty"`

	Name string `json:"name"`
}

type CreateMowerDTO struct {
	Name string `json:"name"`
}

type UpdateMowerDTO struct {
	Name string `json:"name,omitempty"`
}

func (m Mower) String() string {
	return fmt.Sprintf("mower %s (#%v)", m.Name, m.ID)
}

type Timestamp time.Time

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	ts := time.Time(*t).Unix()
	stamp := fmt.Sprint(ts)
	return []byte(stamp), nil
}
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}
	*t = Timestamp(time.Unix(int64(ts), 0))
	return nil
}

func (t *Timestamp) String() string {
	return time.Time(*t).String()
}
