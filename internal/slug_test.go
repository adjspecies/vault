// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package internal_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/adjspecies/vault/internal"
)

func TestSlug(t *testing.T) {
	Convey("It should be able to create a slug", t, func() {
		So(internal.NewSlug("rose"), ShouldEqual, "rose")
		So(internal.NewSlug("the doctor"), ShouldEqual, "the-doctor")
		So(internal.NewSlug("bad$$$wolf"), ShouldEqual, "bad-wolf")
	})
}
