package grift

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gobuffalo/makr"
	"github.com/pkg/errors"
)

func init() {
	fmt.Println("github.com/gobuffalo/buffalo/generators/refresh has been deprecated in v0.13.0, and will be removed in v0.14.0. Use github.com/gobuffalo/buffalo/genny/refresh directly.")
}

//Run allows to create a new grift task generator
func (gg Generator) Run(root string, data makr.Data) error {
	g := makr.New()
	defer g.Fmt(root)

	header := tmplHeader
	path := filepath.Join("grifts", gg.Name.File()+".go")

	if _, err := os.Stat(path); err == nil {
		template, err := ioutil.ReadFile(path)
		if err != nil {
			return errors.WithStack(err)
		}
		header = string(template)
	}

	g.Add(makr.NewFile(path, header+tmplBody))

	data["opts"] = gg
	return g.Run(root, data)
}
