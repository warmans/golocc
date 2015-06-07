package fixture_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	Convey("test something", t, func() {
		So(1, ShouldEqual, 1)
		assert.Equal(t, 1, 1)
	})

}
