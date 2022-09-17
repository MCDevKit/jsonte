package functions

import (
	"fmt"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"io/ioutil"
	"os"
	"strings"
)

type Docs struct {
	Summary   string
	Arguments []Argument
	Example   string
}

type Argument struct {
	Name     string
	Summary  string
	Optional bool
}

type Group struct {
	Name    string
	Title   string
	Summary string
}

func GenerateDocs() error {
	Init()
	dir, err := os.ReadDir(".")
	if err != nil {
		return utils.WrapErrorf(err, "Failed to read current directory")
	}
	for _, file := range dir {
		if file.IsDir() && strings.HasSuffix(file.Name(), "-functions") {
			err := os.RemoveAll(file.Name())
			if err != nil {
				return utils.WrapErrorf(err, "Failed to remove directory %s", file.Name())
			}
		}
	}
	for _, group := range groups {
		err = os.Mkdir(fmt.Sprintf("%s-functions", group.Name), 0755)
		if err != nil {
			return utils.WrapErrorf(err, "Failed to create %s directory", group.Name)
		}
		err = ioutil.WriteFile(fmt.Sprintf("%s-functions/index.md", group.Name), []byte(fmt.Sprintf(`---
layout: page
title: %[1]s
parent: JSON Templating Engine
has_children: true
---

# %[1]s

%[2]s
`, group.Title, group.Summary)), 0644)
		if err != nil {
			return utils.WrapErrorf(err, "Failed to write %s/index.md", group.Name)
		}
		for _, fns := range functions {
			for _, fn := range fns {
				if fn.Group != group.Name || fn.Docs.Summary == "" {
					continue
				}
				summary := fn.Docs.Summary
				if fn.IsUnsafe {
					summary += "\n\n**This method is marked as unsafe. It can be disabled in certain environments.**"
				}
				err = ioutil.WriteFile(fmt.Sprintf("%s-functions/%s.md", group.Name, fn.Name), []byte(fmt.Sprintf(`---
layout: page
grand_parent: JSON Templating Engine
parent: %[2]s
title: %[1]s
---

# %[1]s

%[3]s
%[4]s
%[5]s
`, fn.Name, group.Title, summary, generateArgumentDocs(fn.Docs.Arguments), prepareExample(fn.Docs.Example))), 0644)
				if err != nil {
					return utils.WrapErrorf(err, "Failed to write %s/%s.md", group.Name, fn.Name)
				}
			}
		}
	}
	return nil
}

func prepareExample(example string) string {
	if example == "" {
		return ""
	}
	s := "## Example\n"
	example = strings.ReplaceAll(example, "<code>", "```json")
	example = strings.ReplaceAll(example, "</code>", "```")
	example = strings.ReplaceAll(example, "<pre>", "```json")
	example = strings.ReplaceAll(example, "</pre>", "```")
	example = strings.ReplaceAll(example, "{{", "{{\"{{")
	example = strings.ReplaceAll(example, "}}", "\"}}}}")
	s += example
	return s
}

func generateArgumentDocs(arguments []Argument) string {
	if len(arguments) == 0 {
		return ""
	}
	s := "## Arguments\n\n"
	for _, arg := range arguments {
		op := ""
		if arg.Optional {
			op = "(optional) "
		}
		s += fmt.Sprintf("- `%s` - %s%s\n", arg.Name, op, arg.Summary)
	}
	return s
}
