package model

import (
	"github.com/smartystreets/assertions"
	"testing"
)

func TestUtils(t *testing.T) {
	fromString, err := GetSexTypeFromString("man")
	assertions.ShouldEqual(fromString, Male)
	assertions.ShouldBeNil(err)
	male, err2 := GetSexTypeFromString("female")
	assertions.ShouldEqual(male, FEMALE)
	assertions.ShouldBeNil(err2)
	other, err3 := GetSexTypeFromString("other")
	assertions.ShouldBeNil(err3)
	assertions.ShouldEqual(other, Other)
}
