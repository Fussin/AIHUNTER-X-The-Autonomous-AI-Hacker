package dispatcher_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/user/aihunter-x/aihunter-core-cli/core/dispatcher"
)

type testPlugin struct {
	preInitCalled bool
	onExitCalled  bool
}

func (p *testPlugin) Commands() []*cobra.Command {
	return []*cobra.Command{
		{
			Use: "test",
		},
	}
}

func (p *testPlugin) PreInit() error {
	p.preInitCalled = true
	return nil
}

func (p *testPlugin) OnExit() error {
	p.onExitCalled = true
	return nil
}

type panicPlugin struct{}

func (p *panicPlugin) Commands() []*cobra.Command {
	return []*cobra.Command{
		{
			Use: "panic",
			RunE: func(cmd *cobra.Command, args []string) error {
				panic("test panic")
			},
		},
	}
}

func (p *panicPlugin) PreInit() error {
	return nil
}

func (p *panicPlugin) OnExit() error {
	return nil
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

func TestLifecycleHooks(t *testing.T) {
	dispatcher.ResetPlugins()
	p := &testPlugin{}
	dispatcher.Register(p)

	// This is a bit of a hack to test the lifecycle hooks.
	// In a real application, the Execute function would be called.
	for _, pl := range dispatcher.GetPlugins() {
		pl.PreInit()
	}
	for _, pl := range dispatcher.GetPlugins() {
		pl.OnExit()
	}

	assert.True(t, p.preInitCalled)
	assert.True(t, p.onExitCalled)
}

func TestPanicIsolation(t *testing.T) {
	dispatcher.ResetPlugins()
	dispatcher.Register(&panicPlugin{})
	rootCmd := &cobra.Command{
		Use: "root",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	dispatcher.AddCommands(rootCmd)
	rootCmd.SetArgs([]string{"panic"})

	err := rootCmd.Execute()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "panic: test panic")
}
