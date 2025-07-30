package dispatcher_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/user/aihunter-x/aihunter-core-cli/core/dispatcher"
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
	dispatcher.ResetPlugins()
	dispatcher.Register(&testPlugin{})
	// assert.Len(t, dispatcher.plugins, 1) // This is not possible as plugins is not exported
}

func TestAddCommands(t *testing.T) {
	dispatcher.ResetPlugins()
	dispatcher.Register(&testPlugin{})
	rootCmd := &cobra.Command{}
	dispatcher.AddCommands(rootCmd)
	assert.True(t, rootCmd.HasSubCommands())
	cmd, _, err := rootCmd.Find([]string{"test"})
	assert.NoError(t, err)
	assert.NotNil(t, cmd)
}
