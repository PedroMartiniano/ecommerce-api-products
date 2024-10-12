package vo

import "github.com/google/uuid"

type UUID struct {
	value string
}

func NewUUID(id string) UUID {
	if id == "" {
		id, _ := uuid.NewV7()

		return UUID{
			value: id.String(),
		}
	}

	return UUID{
		value: id,
	}
}

func (u UUID) GetValue() string {
	return u.value
}
