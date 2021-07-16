package helpers

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// MakeDirectoryIfNotExists if not exist make dir
func MakeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}

func DisableLogging(enable bool) {

	if enable {
		log.Print("INIT APP: LOGGING IS DISABLED")
		log.SetOutput(ioutil.Discard)
	}

}



// MakeUUID UUID newUUID generates a random UUID
func MakeUUID(smallUUID ...bool) (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x%x%x", uuid[0:4], uuid[4:6], uuid[6:8]), nil
	//return fmt.Sprintf("%x%x%x%x%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil

}


// S2b converts a string true/T/"1" to a true/false
func S2b(value string)  bool {
	r, _ := strconv.ParseBool(value)
	return r
}

func I2b(b int) bool {
	if b == 1 {
		return true
	}
	return false
}

func B2i(b bool) int8 {
	if b {
		return 1
	}
	return 0
}



func arrayContains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

