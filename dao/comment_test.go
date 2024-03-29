package dao

import (
	"gorm.io/gorm/utils/tests"
	"testing"
)

func TestInsertComment(t *testing.T) {
	InitDb()
	id, _, err := InsertComment(1, 1, "test ljp")
	tests.AssertEqual(t, err, nil)
	tests.AssertEqual(t, id, 1)
}

func TestDeleteComment(t *testing.T) {
	InitDb()
	err := DeleteComment(1)
	tests.AssertEqual(t, err, nil)
}
