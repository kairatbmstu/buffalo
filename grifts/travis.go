package grifts

import (
	"os/exec"
	"runtime"

	"github.com/markbates/grift/grift"
)

var _ = grift.Add("travis", func(c *grift.Context) error {
	cmd := exec.CommandContext(c, "make", "ci-test")
	os := runtime.GOOS
	switch os {
	case "linux":
		cmd
	}
	return nil
})
