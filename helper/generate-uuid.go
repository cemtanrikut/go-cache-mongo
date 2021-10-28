package helper

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	uuid := uuid.New().String()
	generatedUUID := strings.Replace(uuid, "-", "", -1)
	fmt.Println("UUID ", generatedUUID)
	return generatedUUID
}
