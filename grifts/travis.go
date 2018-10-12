package grifts

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/markbates/grift/grift"
	"github.com/pkg/errors"
)

var _ = grift.Add("travis", func(c *grift.Context) error {
	goos := runtime.GOOS

	cmd := exec.CommandContext(c, "make", "ci-test")
	switch goos {
	case "windows":
		co := exec.CommandContext(c, "choco", "install", "make", "-y")
		fmt.Println(co.Args)
		co.Stderr = os.Stderr
		co.Stdout = os.Stdout
		if err := co.Run(); err != nil {
			return errors.WithStack(err)
		}
	case "linux":
		// cmd
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	fmt.Println(cmd.Args)
	return cmd.Run()
})
