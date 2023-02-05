package utils

import (
	"testing"
	"tiktok/models"
	"time"
)

func TestMappingMap2Struct(t *testing.T) {
	mp := map[string]string{
		"ID":        "1",
		"CreatedAt": "2006-01-02T15:04:05Z",
	}
	vd := Scan[models.Message](mp)
	t.Log(vd)

	CreatedAt, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	if err != nil {
		t.Error(err)
	}
	if vd.ID == 1 && vd.CreatedAt == CreatedAt {
		t.Log("MappingMap2Struct[models.Video] pass")
	} else {
		t.Error("MappingMap2Struct[models.Video] fail")
	}
}
