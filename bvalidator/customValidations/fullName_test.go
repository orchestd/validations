package customValidations

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func Test_ValidateFullName(t *testing.T) {
	allowedNames := []string{
		"John Smith",
		"John Michael Smith",
		"Jonathan Smith",
		"José García",
		"François Dupont",
		"Jean Pierre",
		"John  Smith",
		"John Michael David Smith",
		"Анна Иванова",
		"محمد علي",
		"שם מלא",
	}

	convey.Convey("given allowed full names", t, func() {
		for _, name := range allowedNames {
			v := isValidFullName(name)
			convey.So(v, convey.ShouldBeTrue)
		}
	})

	blockedNames := []string{
		"John",
		"J Smith",
		"John S",
		"J S",
		"John1 Smith",
		"John Smith2",
		"123 Smith",
		"John 123",
		"John-Smith",
		"John_Smith",
		"John.Smith",
		"John@Smith",
		"John",
		" Smith",
		"John Smith ",
		"John  ",
		" John Smith",
		"",
		"   ",
		"שם",
	}

	convey.Convey("given blocked full names", t, func() {
		for _, name := range blockedNames {
			v := isValidFullName(name)
			convey.So(v, convey.ShouldBeFalse)
		}
	})
}
