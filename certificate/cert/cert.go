package cert

import (
	"fmt"
	"strings"
	"time"
)

var MaxLenCourse = 20
var MaxLenName = 10

type Cert struct {
	Course,
	Name string
	Date time.Time

	LabelTitle,
	LabelPresented,
	LabelParticipation,
	LabelDate,
	LabelCompletion string
}

type Saver interface {
	Save(c Cert) error
}

func New(course, name, date string) (*Cert, error) {
	c, err := validateCourse(course)
	if err != nil {
		return nil, err
	}
	n, err := validateName(name)
	if err != nil {
		return nil, err
	}
	d, e := parseDate(date)

	cert := &Cert{
		Course:             c,
		Name:               n,
		Date:               d,
		LabelTitle:         fmt.Sprintf("%v Certificate - %v", c, n),
		LabelCompletion:    "Certificate of completion",
		LabelPresented:     "This Certificate is Presented To",
		LabelParticipation: fmt.Sprintf("For participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date %v", d.Format("02/01/2006")),
	}

	return cert, nil
}

func parseDate(d string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", d)
	if err != nil {
		return t, err
	}
	return t, nil
}

func validateCourse(c string) (string, error) {
	c, err := validateStr(c, MaxLenCourse)
	if err != nil {
		return "", err
	}
	suffix := " course"
	if !strings.HasSuffix(c, suffix) {
		c += suffix
	}
	return strings.ToTitle(c), nil
}

func validateStr(str string, maxLen int) (string, error) {
	c := strings.TrimSpace(str)
	if len(c) <= 0 || len(c) > maxLen {
		return c, fmt.Errorf("Invalid string, got='%s', len=%d", c, len(c))
	}
	return c, nil
}

func validateName(n string) (string, error) {
	n, err := validateStr(n, MaxLenName)
	if err != nil {
		return "", err
	}
	return strings.ToTitle(n), nil
}
