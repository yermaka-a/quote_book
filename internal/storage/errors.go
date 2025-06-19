package storage

import "errors"

var (
	ErrListIsEmpty     = errors.New("list is empty")
	ErrRecordNotFound  = errors.New("record isn't found")
	ErrMapIsEmpty      = errors.New("map is empty")
	ErrHasNoSubRecords = errors.New("record has no subrecords")
)
