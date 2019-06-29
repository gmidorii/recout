package app

import (
	"fmt"
	"strings"
	"time"
)

const (
	pixelaURL         = "https://pixe.la/v1/users"
	pixelaHeaderToken = "X-USER-TOKEN"
	pixelaTokenLen    = 32
	pixelaGraphIDLen  = 8
	prefixAccountID   = "rec"

	durationDay = 24 * time.Hour
	hoursPerDay = 24
	idLen       = 16
)

type Container struct {
	Env       string
	Now       time.Time
	Location  *time.Location
	Generator RandomGenerator
}

func NewContainer(env string, now time.Time, location time.Location, generator RandomGenerator) Container {
	return Container{
		Env:       env,
		Now:       now,
		Location:  &location,
		Generator: generator,
	}
}

type RandomGenerator interface {
	Do(length int) string
}

func subDate(before, after time.Time) int {
	trBefore := before.Truncate(durationDay)
	trAfter := after.Truncate(durationDay)
	hours24Divisible := trAfter.Sub(trBefore).Hours()
	return int(hours24Divisible) / hoursPerDay
}

func generateGraphName(origin string) string {
	return strings.Replace(strings.ToLower(origin), " ", "", -1)
}

func encodeAccountID(origin string) string {
	return fmt.Sprintf("%v%v", prefixAccountID, strings.ToLower(origin))
}

func decodeAccountID(encoded string) string {
	return strings.Split(encoded, prefixAccountID)[1]
}
