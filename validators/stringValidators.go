package validators

import (
	log "github.com/sirupsen/logrus"
	"regexp"
)


var isAlphaNumeric = regexp.MustCompile(`^[A-Za-z\d]+$`).MatchString

func ContainsOnlyAlphanumeric(list []string) bool {
	for _, i := range list {
		if !isAlphaNumeric(i) {
			log.Warnf("Medallion Id was not alphanumeric: %s", i)
			return false
		}
	}

	return true
}