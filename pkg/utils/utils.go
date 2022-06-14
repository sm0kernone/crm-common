package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var (
	DayDuration   = int(24 * time.Hour.Seconds())
	WeekDuration  = 7 * DayDuration
	MonthDuration = 30 * DayDuration
)

func InListString(entry string, list []string) bool {
	for _, listEntry := range list {
		if entry == listEntry {
			return true
		}
	}
	return false
}

// AnyOfDstListInSrcList checks whether any string in src list is the same as in dst list
func AnyOfDstListInSrcList(src, dst []string) bool {
	for _, val := range src {
		if InListString(val, dst) {
			return true
		}
	}
	return false
}

func InListInt(entry int, list []int) bool {
	for _, listEntry := range list {
		if entry == listEntry {
			return true
		}
	}
	return false
}

func InCountryMap(entry string, reqMap map[string]string) (bool, string) {
	for _, v := range reqMap {
		if NormalizeCountry(entry) == NormalizeCountry(v) {
			return true, v
		}
	}

	return false, ""
}

func ReplaceRegex(patternRegex, src string) string {
	var re = regexp.MustCompile(`(?m)` + patternRegex)
	return re.ReplaceAllString(src, "")
}

// RemoveSpaces removes redundant spaces in string - starts a loop while string contains more than 2 spaces in a row and replaces them to 1 space
func RemoveSpaces(val string) string {
	for {
		if strings.Contains(val, "  ") {
			val = strings.ReplaceAll(val, "  ", " ")
		} else {
			return val
		}
	}
}

func NormalizeField(val string) string {
	pattern := "[^a-zA-Z0-9--- ]+"
	return strings.TrimSpace(strings.Title(strings.ToLower(RemoveSpaces(ReplaceRegex(pattern, val)))))
}

func NormalizeCountry(val string) string {
	pattern := "[^a-zA-Z0-9]+"
	return strings.ToLower(ReplaceRegex(pattern, val))
}

func NormalizeCountryWithSpace(val string) string {
	pattern := "[^a-zA-Z0-9 ]+"

	return strings.TrimSpace(strings.Title(strings.ToLower(RemoveSpaces(ReplaceRegex(pattern, val)))))
}

// GetKeyForCloner returns key to store in redis for cloner, accepts entity (e.g. leads) and returns redis key
func GetKeyForCloner(entity string) string {
	return fmt.Sprintf("%v:%v", entity, randStringRunes(32))
}

func randStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func CheckErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func EncodeMD5(req string) string {

	hash := md5.Sum([]byte(req))
	return hex.EncodeToString(hash[:])
}

func getIndexInIntSlice(s []int, element int) int {
	for idx, v := range s {
		if v == element {
			return idx
		}
	}

	return 0
}

func RemoveElementFromIntSlice(s []int, element int) []int {
	idx := getIndexInIntSlice(s, element)

	return append(s[:idx], s[idx+1:]...)
}
