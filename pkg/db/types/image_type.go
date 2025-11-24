package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type ImageType string

const (
	Backend  ImageType = "backend"
	SQL      ImageType = "sql"
	Web      ImageType = "web"
	NOSQL    ImageType = "nosql"
	FrontEnd ImageType = "frontend"
)

var validImageTypes = map[ImageType]bool{
	Backend:  true,
	SQL:      true,
	Web:      true,
	NOSQL:    true,
	FrontEnd: true,
}

func (t *ImageType) Value() (driver.Value, error) {
	return string(*t), nil
}

func (t *ImageType) Scan(value interface{}) error {
	strVal, ok := value.(string)
	if !ok {
		return errors.New("invalid type for ImageType")
	}
	imageType := ImageType(strVal)
	if !validImageTypes[imageType] {
		return fmt.Errorf("invalid ImageType: %s", strVal)
	}
	*t = imageType
	return nil
}
