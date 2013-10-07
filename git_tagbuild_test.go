package main

import (
	"testing"
)

func TestIncreaseBuildNumber(t *testing.T) {
	tags := []string{}
	date := "2013-10-06"
	buildNumber := increaseBuildNumber(tags, date)

	if buildNumber != 1 {
		t.Error("increaseBuildNumber should return 1")
	}

	tags = []string{"2013-10-05-01"}
	buildNumber = increaseBuildNumber(tags, date)

	if buildNumber != 1 {
		t.Error("increaseBuildNumber should return 1")
	}

	tags = []string{"2013-10-06-01"}
	buildNumber = increaseBuildNumber(tags, date)

	if buildNumber != 2 {
		t.Error("increaseBuildNumber should return 2")
	}

	tags = []string{"2013-10-06-01", "2013-10-06-02"}
	buildNumber = increaseBuildNumber(tags, date)

	if buildNumber != 3 {
		t.Error("increaseBuildNumber should return 3")
	}
}

func TestFormatNumber(t *testing.T) {
	n := formatNumber(1)
	if n != "01" {
		t.Error("formatBuildNumber should retunr 01")
	}

	n = formatNumber(10)
	if n != "10" {
		t.Error("formatBuildNumber should retunr 01")
	}
}
