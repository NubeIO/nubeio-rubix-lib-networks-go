package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"nubeio-rubix-lib-rest-go/helpers"
	"gorm.io/gorm"
	"log"
	mathRand "math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

)

var (
	minSecureKeyLength = 8
	errShortSecureKey  = errors.New("length of secure key does not meet with minimum requirements")
)



type NotFoundError struct {
	item string
	id string
	err error
	msg string
}

func (e *NotFoundError) Error() string {
	if e.err == nil {
		m:= fmt.Sprintf("%s with id#%s %s", e.item, e.id, e.msg)
		log.Println(m)
		return m
	}
	m:= fmt.Sprintf("%s with id#%s %s", e.item, e.id, e.err)
	log.Println(m)
	return m
}






// ArgsIsNil check if arg is nil
func ArgsIsNil(value string) bool {
	if value == "" {
		return true
	} else {
		return false
	}
}


func WithChildren(value string) (bool, bool)  {
	if value == "" {
		return false, false
	} else if value == "all" { // will return children's children
		return false, true
	} else  {
		return helpers.S2b(value), false
	}
}

func ArgsCheck(value string) bool {
	if value == "" {
		return false
	} else  {
		return helpers.S2b(value)
	}
}


// Offset returns the starting number of result for pagination
func Offset(offset string) int {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 0
	}
	return offsetInt
}

// Limit returns the number of result for pagination
func Limit(limit string) int {
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 25
	}
	return limitInt
}

// SortOrder returns the string for sorting and orderin data
func SortOrder(table, sort, order string) string {
	return table + "." + ToSnakeCase(sort) + " " + ToSnakeCase(order)
}

// Search adds where to search keywords
func Search(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			db = db.Where("name LIKE ?", "%"+search+"%")
			db = db.Or("description LIKE ?", "%"+search+"%")
		}
		return db
	}
}

// ToSnakeCase changes string to database table
func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}

// FallbackInsecureKey fallback method for sercure key
func FallbackInsecureKey(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789" +
		"~!@#$%^&*()_+{}|<>?,./:"

	if err := checkSecureKeyLen(length); err != nil {
		return "", err
	}

	var seededRand *mathRand.Rand = mathRand.New(
		mathRand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b), nil
}

// GenerateSecureKey generates a secure key width a given length
func GenerateSecureKey(length int) (string, error) {
	key := make([]byte, length)

	if err := checkSecureKeyLen(length); err != nil {
		return "", err
	}
	_, err := rand.Read(key)
	if err != nil {
		return FallbackInsecureKey(length)
	}
	// encrypted key length > provided key length
	keyEnc := base64.StdEncoding.EncodeToString(key)
	return keyEnc, nil
}

func checkSecureKeyLen(length int) error {
	if length < minSecureKeyLength {
		return errShortSecureKey
	}
	return nil
}
