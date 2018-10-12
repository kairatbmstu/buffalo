package grifts

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/markbates/grift/grift"
)

var _ = grift.Add("travis", func(c *grift.Context) error {
	cmd := exec.CommandContext(c, "make", "ci-test")
	goos := runtime.GOOS
	switch goos {
	case "linux":
		// cmd
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
})
