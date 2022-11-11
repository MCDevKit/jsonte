package main

import (
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"strings"
)

type flag struct {
	Flag
	Type                   string
	BoolDestination        *bool
	StringDestination      *string
	StringSliceDestination *[]string
	Found                  bool
}

type Flag struct {
	Name        string
	Usage       string
	Description string
	Required    bool
	OnSet       func()
}

type Action struct {
	Name        string
	Hidden      bool
	Usage       string
	Description string
	Function    func(args []string) error
}

type App struct {
	flags   []flag
	actions map[string]Action
}

func (a *App) BoolFlag(f Flag, destination *bool) {
	a.flags = append(a.flags, flag{
		Flag:            f,
		Type:            "bool",
		BoolDestination: destination,
	})
}

func (a *App) StringFlag(f Flag, destination *string) {
	a.flags = append(a.flags, flag{
		Flag:              f,
		Type:              "string",
		StringDestination: destination,
	})
}

func (a *App) StringSliceFlag(f Flag, destination *[]string) {
	a.flags = append(a.flags, flag{
		Flag:                   f,
		Type:                   "stringSlice",
		StringSliceDestination: destination,
	})
}

func (a *App) Action(action Action) {
	if a.actions == nil {
		a.actions = make(map[string]Action)
	}
	a.actions[action.Name] = action
}

func (a *App) Run(args []string, onParse func()) error {
	cleanArgs := make([]string, 0)
	for i := 1; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "--") {
			found := false
			for _, flag := range a.flags {
				if arg == "--"+flag.Name {
					found = true
					switch flag.Type {
					case "bool":
						if flag.Found {
							return burrito.WrappedErrorf("flag --%s is already set", flag.Name)
						}
						*flag.BoolDestination = true
						flag.Found = true
					case "string":
						if flag.Found {
							return burrito.WrappedErrorf("flag --%s is already set", flag.Name)
						}
						*flag.StringDestination = args[i+1]
						flag.Found = true
						i++
						break
					case "stringSlice":
						*flag.StringSliceDestination = append(*flag.StringSliceDestination, args[i+1])
						i++
						flag.Found = true
						break
					default:
						return burrito.WrappedErrorf("unknown flag type: %s", flag.Type)
					}
				}
			}
			if !found {
				a.PrintHelp()
				return burrito.WrappedErrorf("unknown flag: %s", arg)
			}
		} else {
			cleanArgs = append(cleanArgs, arg)
		}
	}
	if onParse != nil {
		onParse()
	}
	for i, f := range a.flags {
		if f.Required && !f.Found {
			return fmt.Errorf("flag --%s is required", f.Name)
		}
		if f.Found && f.OnSet != nil {
			f.OnSet()
		}
		a.flags[i] = f
	}
	if len(cleanArgs) == 0 {
		a.PrintHelp()
		return nil
	}
	action, ok := a.actions[cleanArgs[0]]
	if !ok {
		a.PrintHelp()
		return burrito.WrappedErrorf("unknown command: %s", cleanArgs[0])
	}

	return action.Function(cleanArgs[1:])
}

func (a *App) PrintHelp() {
	fmt.Println("jsonte - JSON Templating Engine")
	fmt.Println("")
	fmt.Println("Usage: jsonte <command> [arguments...] [options]")
	fmt.Println("")
	for _, flag := range a.flags {
		fmt.Printf("  --%s\t%s\n", flag.Name, flag.Usage)
	}
	fmt.Println("")
	fmt.Println("Commands:")
	for _, action := range a.actions {
		if action.Hidden {
			continue
		}
		fmt.Printf("  %s\t%s\n", action.Name, action.Usage)
	}
}

func NewApp() *App {
	return &App{}
}
