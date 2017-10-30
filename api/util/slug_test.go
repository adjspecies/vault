// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package util_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/adjspecies/vault/api/util"
)

func TestSlug(t *testing.T) {
	Convey("It should be able to create a slug", t, func() {
		So(util.NewSlug("rose"), ShouldEqual, "rose")
		So(util.NewSlug("the doctor"), ShouldEqual, "the-doctor")
		So(util.NewSlug("bad$$$wolf"), ShouldEqual, "bad-wolf")
	})
}
