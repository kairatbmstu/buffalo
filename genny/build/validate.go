package build

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/plush"
	"github.com/markbates/safe"
	"github.com/pkg/errors"
)

type TemplateValidator func(f genny.File) error

func ValidateTemplates(box packr.Box, tvs []TemplateValidator) genny.RunFn {
	if len(tvs) == 0 {
		return func(r *genny.Runner) error {
			return nil
		}
	}
	return func(r *genny.Runner) error {
		var errs []string
		box.Walk(func(path string, file packr.File) error {
			info, err := file.FileInfo()
			if err != nil {
				return errors.WithStack(err)
			}
			if info.IsDir() {
				return nil
			}

			f := genny.NewFile(path, file)
			for _, tv := range tvs {
				err := safe.Run(func() {
					if err := tv(f); err != nil {
						errs = append(errs, fmt.Sprintf("template error in file %s: %s", path, err.Error()))
					}
				})
				if err != nil {
					return errors.WithStack(err)
				}
			}

			return nil
		})
		if len(errs) == 0 {
			return nil
		}
		return errors.New(strings.Join(errs, "\n"))
	}
}

func PlushValidator(f genny.File) error {
	if !genny.HasExt(f, ".md") && !genny.HasExt(f, ".html") {
		return nil
	}
	_, err := plush.Parse(f.String())
	return err
}

func GoTemplateValidator(f genny.File) error {
	if !genny.HasExt(f, ".tmpl") {
		return nil
	}
	t := template.New(f.Name())
	_, err := t.Parse(f.String())
	return err
}
