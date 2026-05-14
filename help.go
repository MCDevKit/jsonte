//go:build !wasm && !js

package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/MCDevKit/jsonte/jsonte/functions"
)

var htmlTagRe = regexp.MustCompile(`<[^>]+>`)

func runHelp(args []string) error {
	if len(args) == 0 {
		printGeneralHelp()
		return nil
	}
	topic := args[0]
	if strings.EqualFold(topic, "functions") {
		printFunctionList()
		return nil
	}
	// Check if topic matches a group name
	groups := functions.GetGroups()
	for name, group := range groups {
		if strings.EqualFold(topic, name) {
			printGroupHelp(group)
			return nil
		}
	}
	// Check if topic matches a function name
	allFunctions := functions.GetFunctions()
	key := strings.ToLower(topic)
	if fns, ok := allFunctions[key]; ok {
		printFunctionHelp(fns)
		return nil
	}
	fmt.Printf("No help found for %q.\n\n", topic)
	fmt.Println("Use `jsonte help functions` to list all available functions.")
	return nil
}

func printGeneralHelp() {
	fmt.Println("jsonte - JSON Templating Engine")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  jsonte help                  Show this help")
	fmt.Println("  jsonte help functions         List all available functions")
	fmt.Println("  jsonte help <functionName>    Show help for a specific function")
	fmt.Println("  jsonte help <groupName>       List all functions in a group")
	fmt.Println()
	fmt.Println("Function groups:")
	groups := functions.GetGroups()
	names := make([]string, 0, len(groups))
	for name := range groups {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		g := groups[name]
		fmt.Printf("  %-20s %s\n", name, g.Summary)
	}
}

func printFunctionList() {
	groups := functions.GetGroups()
	allFunctions := functions.GetFunctions()

	// Build group -> functions mapping (deduplicated by name within group)
	type fnEntry struct {
		name    string
		summary string
	}
	groupFunctions := make(map[string][]fnEntry)
	seen := make(map[string]bool)
	for _, fns := range allFunctions {
		for _, fn := range fns {
			key := fn.Group + ":" + fn.Name
			if seen[key] || fn.Docs.Summary == "" {
				continue
			}
			seen[key] = true
			groupFunctions[fn.Group] = append(groupFunctions[fn.Group], fnEntry{fn.Name, fn.Docs.Summary})
		}
	}

	groupNames := make([]string, 0, len(groups))
	for name := range groups {
		groupNames = append(groupNames, name)
	}
	sort.Strings(groupNames)

	fmt.Println("Available functions:")
	fmt.Println()
	for _, groupName := range groupNames {
		g := groups[groupName]
		entries := groupFunctions[groupName]
		if len(entries) == 0 {
			continue
		}
		sort.Slice(entries, func(i, j int) bool { return entries[i].name < entries[j].name })
		fmt.Printf("[%s] %s\n", g.Name, g.Title)
		for _, e := range entries {
			fmt.Printf("  %-30s %s\n", e.name, e.summary)
		}
		fmt.Println()
	}
	fmt.Println("Use `jsonte help <functionName>` for detailed help on a specific function.")
}

func printGroupHelp(group functions.Group) {
	allFunctions := functions.GetFunctions()

	type fnEntry struct {
		name    string
		summary string
	}
	var entries []fnEntry
	seen := make(map[string]bool)
	for _, fns := range allFunctions {
		for _, fn := range fns {
			if fn.Group != group.Name || seen[fn.Name] {
				continue
			}
			seen[fn.Name] = true
			entries = append(entries, fnEntry{fn.Name, fn.Docs.Summary})
		}
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].name < entries[j].name })

	fmt.Printf("%s\n", group.Title)
	fmt.Printf("%s\n\n", group.Summary)
	fmt.Println("Functions:")
	for _, e := range entries {
		if e.summary == "" {
			fmt.Printf("  %s\n", e.name)
		} else {
			fmt.Printf("  %-30s %s\n", e.name, e.summary)
		}
	}
	fmt.Printf("\nUse `jsonte help <functionName>` for detailed help on a specific function.\n")
}

func printFunctionHelp(fns []functions.JsonFunction) {
	fn := fns[0]
	fmt.Printf("Function: %s\n", fn.Name)
	if fn.Docs.Summary != "" {
		fmt.Printf("\n%s\n", fn.Docs.Summary)
	}
	if fn.IsUnsafe {
		fmt.Println("\n[!] This function is marked unsafe and can be disabled in safe mode.")
	}
	if fn.Deprecated {
		fmt.Println("\n[!] This function is deprecated.")
	}
	// Show all overloads
	for _, f := range fns {
		if len(f.Docs.Arguments) > 0 {
			fmt.Println("\nArguments:")
			for _, arg := range f.Docs.Arguments {
				optional := ""
				if arg.Optional {
					optional = " (optional)"
				}
				varargs := ""
				if arg.VarArgs {
					varargs = "..."
				}
				fmt.Printf("  %s%s%s - %s\n", arg.Name, varargs, optional, arg.Summary)
			}
		}
	}
	if fn.Docs.Example != "" {
		fmt.Println("\nExample:")
		example := htmlTagRe.ReplaceAllString(fn.Docs.Example, "")
		example = strings.TrimSpace(example)
		for _, line := range strings.Split(example, "\n") {
			fmt.Printf("  %s\n", line)
		}
	}
}
