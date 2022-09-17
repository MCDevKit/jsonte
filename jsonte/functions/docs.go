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
	stat, err := os.Stat("docs")
	if err != nil {
		if !os.IsNotExist(err) {
			return utils.WrapErrorf(err, "Failed to check for docs directory!")
		}
	} else if !stat.IsDir() {
		return utils.WrappedErrorf("docs is not a directory!")
	}
	err = os.RemoveAll("docs")
	if err != nil {
		if !os.IsNotExist(err) {
			return utils.WrapErrorf(err, "Failed to remove docs directory!")
		}
	}
	err = os.Mkdir("docs", 0755)
	if err != nil {
		return utils.WrapErrorf(err, "Failed to create docs directory")
	}
	for _, group := range groups {
		err = os.Mkdir(fmt.Sprintf("docs/%s-functions", group.Name), 0755)
		if err != nil {
			return utils.WrapErrorf(err, "Failed to create docs/%s directory", group.Name)
		}
		err = ioutil.WriteFile(fmt.Sprintf("docs/%s-functions/index.md", group.Name), []byte(fmt.Sprintf(`---
layout: page
title: %[1]s
parent: JSON Templating Engine
has_children: true
---

# %[1]s

%[2]s
`, group.Title, group.Summary)), 0644)
		if err != nil {
			return utils.WrapErrorf(err, "Failed to write docs/%s/index.md", group.Name)
		}
		for _, fns := range functions {
			for _, fn := range fns {
				if fn.Group != group.Name || fn.Docs.Summary == "" {
					continue
				}
				err = ioutil.WriteFile(fmt.Sprintf("docs/%s-functions/%s.md", group.Name, fn.Name), []byte(fmt.Sprintf(`---
layout: page
grand_parent: JSON Templating Engine
parent: %[2]s
title: %[1]s
---

# %[1]s

%[3]s
%[4]s
%[5]s
`, fn.Name, group.Title, fn.Docs.Summary, generateArgumentDocs(fn.Docs.Arguments), prepareExample(fn.Docs.Example))), 0644)
				if err != nil {
					return utils.WrapErrorf(err, "Failed to write docs/%s/%s.md", group.Name, fn.Name)
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
		s += fmt.Sprintf("- `%s` - %s\n", arg.Name, arg.Summary)
	}
	return s
}
