package database

import (
	"fmt"

	databaseSym "github.com/taubyte/go-sdk-smartops/symbols/resource/database"
)

func (d *Database) Name() (string, error) {
	var nameSize uint32
	err := databaseSym.GetNameSize(d.Uint32(), &nameSize)
	if err != 0 {
		return "", fmt.Errorf("Getting database name size failed with: %s", err)
	}
	if nameSize == 0 {
		return "", fmt.Errorf("Cannot have a database with no name")
	}

	name := make([]byte, nameSize)
	err = databaseSym.GetName(d.Uint32(), &name[0])
	if err != 0 {
		return "", fmt.Errorf("Getting database name failed with: %s", err)
	}

	return string(name), nil
}
