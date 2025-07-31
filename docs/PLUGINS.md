# Plugins

This document describes how to create and register new plugins for AIHUNTER-X.

## Creating a Plugin

A plugin is a Go package that implements the `dispatcher.Plugin` interface. This interface has three methods:

- `Commands() []*cobra.Command`: This method should return a list of `cobra.Command` objects that the plugin provides.
- `PreInit() error`: This method is called before any command is executed. It can be used to perform any initialization that the plugin requires.
- `OnExit() error`: This method is called when the application is about to exit. It can be used to perform any cleanup that the plugin requires.

Here is an example of a simple plugin:

```go
package example

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Plugin is the example plugin.
type Plugin struct{}

// Commands returns the commands for the example plugin.
func (p *Plugin) Commands() []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "example",
		Short: "An example plugin.",
		Long:  `An example plugin.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from the example plugin!")
		},
	}
	return []*cobra.Command{cmd}
}

// PreInit is called before the command is executed.
func (p *Plugin) PreInit() error {
	fmt.Println("PreInit called from the example plugin!")
	return nil
}

// OnExit is called when the application is about to exit.
func (p *Plugin) OnExit() error {
	fmt.Println("OnExit called from the example plugin!")
	return nil
}
```

## Registering a Plugin

To register a plugin, you need to call the `dispatcher.Register` function in the `init` function of your plugin's main package.

Here is an example of how to register the example plugin:

```go
package main

import (
	"github.com/user/aihunter-x/aihunter-core-cli/cmd/aihunter-x/example"
	"github.com/user/aihunter-x/aihunter-core-cli/core/dispatcher"
)

func init() {
	dispatcher.Register(&example.Plugin{})
}
```
