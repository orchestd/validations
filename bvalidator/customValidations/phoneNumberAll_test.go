package customValidations

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_ValidatePhoneNumberAll(t *testing.T) {
	allowedNumbers := []string{
		"0500000000",
		"050-0000000",
		"050 0000000",
		"050.0000000",
		"(050)-0000000",
		"030000000",
		"03-0000000",
		"(03)0000000",
		"03 0000000",
		"050-000-0000",
	}

	convey.Convey("given allowed phone numbers", t, func() {
		for _, n := range allowedNumbers {
			v := isValidPhoneNumberAll(n)
			convey.So(v, convey.ShouldBeTrue)
		}
	})

	blockedNumbers := []string{
		"xxxxxxxxxx",
		"0500000000000000",
		"050+0000000",
		"050--0000000",
		"(050)(0000000)",
	}

	convey.Convey("given blocked phone numbers", t, func() {
		for _, n := range blockedNumbers {
			v := isValidPhoneNumberAll(n)
			convey.So(v, convey.ShouldBeFalse)
		}
	})
}
