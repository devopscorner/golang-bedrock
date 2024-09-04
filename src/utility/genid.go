// utility/genid.go
package utility

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

func GenerateID() string {
	t := time.Now().UnixNano()
	id_time := strconv.FormatInt(t, 10)
	return id_time
}

func GenerateFileID() string {
	return uuid.New().String()
}
