package tests

import (
	"github.com/cosmic-chichu/ab-go/requests"
	"github.com/cosmic-chichu/ab-go/service"
	"testing"
)

func TestReadUrls(t *testing.T) {
	args := &service.Flags{
		UrlFile: "urls.txt"}
	d := &requests.Dispatcher{
		Args: args}

	val1 := d.ReadUrl()
	val2 := d.ReadUrl()

	if val1 == val2 {
		t.Errorf("Url values equals  %s == %s", val1, val2)
	}

	val3 := d.ReadUrl()
	if val1 != val3 {
		t.Errorf("Url values must be equals  %s == %s", val1, val3)
	}
}
