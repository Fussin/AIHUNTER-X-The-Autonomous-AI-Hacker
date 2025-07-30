package plugin

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

type testPlugin struct{}

func (p *testPlugin) Commands() []*cobra.Command {
	return []*cobra.Command{
		{
			Use: "test",
		},
	}
}

func TestRegister(t *testing.T) {
	// Reset plugins
	plugins = []Plugin{}

	Register(&testPlugin{})
	assert.Len(t, plugins, 1)
}

func TestAddCommands(t *testing.T) {
	// Reset plugins
	plugins = []Plugin{}

	Register(&testPlugin{})
	rootCmd := &cobra.Command{}
	AddCommands(rootCmd)
	assert.True(t, rootCmd.HasSubCommands())
	cmd, _, err := rootCmd.Find([]string{"test"})
	assert.NoError(t, err)
	assert.NotNil(t, cmd)
}
