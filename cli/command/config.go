package command

import (
	"github.com/urfave/cli"
	"github.com/ihezebin/sdk/utils/stringer"
)

type CmdConfig struct {
	// The name of the command
	Name string
	// A list of aliases for the command
	Aliases []string
	// A short description of the usage of this command
	Usage string
	// Custom text to show on USAGE section of help
	UsageText string
	// A longer explanation of how the command works
	Description string
	// A short description of the arguments of this command
	ArgsUsage string
	// The category the command is part of
	Category string
	// Treat all flags as normal arguments if true
	SkipFlagParsing bool
	// Skip argument reordering which attempts to move flags before arguments,
	// but only works if all flags appear after all arguments. This behavior was
	// removed n version 2 since it only works under specific conditions so we
	// backport here by exposing it as an option for compatibility.
	SkipArgReorder bool
	// Boolean to hide built-in help command
	HideHelp bool
	// Boolean to hide this command from help or completion
	Hidden bool
	// Boolean to enable short-option handling so user can combine several
	// single-character bool arguments into one
	// i.e. foobar -o -v -> foobar -ov
	UseShortOptionHandling bool

	// Full name of command for help, defaults to full command name, including parent commands.
	HelpName string

	// CustomHelpTemplate the text template for the command help topic.
	// cli.go uses text/template to render templates. You can
	// render custom help text by setting this variable.
	CustomHelpTemplate string
}

type CmdOption func(config *CmdConfig)

func newCmdConfig(options ...CmdOption) CmdConfig {
	config := CmdConfig{}
	for _, option := range options {
		option(&config)
	}
	return config
}

func WithName(name string) CmdOption {
	return func(config *CmdConfig) {
		config.Name = name
	}
}

func WithAliases(aliases []string) CmdOption {
	return func(config *CmdConfig) {
		config.Aliases = aliases
	}
}

func WithUsage(usage string) CmdOption {
	return func(config *CmdConfig) {
		config.Usage = usage
	}
}

func WithUsageText(usageText string) CmdOption {
	return func(config *CmdConfig) {
		config.UsageText = usageText
	}
}

func WithDescription(description string) CmdOption {
	return func(config *CmdConfig) {
		config.Description = description
	}
}

func WithArgsUsage(argsUsage string) CmdOption {
	return func(config *CmdConfig) {
		config.ArgsUsage = argsUsage
	}
}

func WithCategory(category string) CmdOption {
	return func(config *CmdConfig) {
		config.Category = category
	}
}

func WithSkipFlagParsing(skipFlagParsing bool) CmdOption {
	return func(config *CmdConfig) {
		config.SkipFlagParsing = skipFlagParsing
	}
}

func WithSkipArgReorder(skipArgReorder bool) CmdOption {
	return func(config *CmdConfig) {
		config.SkipArgReorder = skipArgReorder
	}
}

func WithHideHelp(hideHelp bool) CmdOption {
	return func(config *CmdConfig) {
		config.HideHelp = hideHelp
	}
}

func WithHidden(hidden bool) CmdOption {
	return func(config *CmdConfig) {
		config.Hidden = hidden
	}
}

func WithUseShortOptionHandling(useShortOptionHandling bool) CmdOption {
	return func(config *CmdConfig) {
		config.UseShortOptionHandling = useShortOptionHandling
	}
}

func WithHelpName(helpName string) CmdOption {
	return func(config *CmdConfig) {
		config.HelpName = helpName
	}
}

func WithCustomHelpTemplate(customHelpTemplate string) CmdOption {
	return func(config *CmdConfig) {
		config.CustomHelpTemplate = customHelpTemplate
	}
}

func handleCommandConfig(cmd *cli.Command, config CmdConfig) *cli.Command {
	if stringer.NotEmpty(config.Name) {
		cmd.Name = config.Name
	}
	if len(config.Aliases) > 0 {
		cmd.Aliases = config.Aliases
	}
	if stringer.NotEmpty(config.Usage) {
		cmd.Usage = config.Usage
	}
	if stringer.NotEmpty(config.UsageText) {
		cmd.UsageText = config.UsageText
	}
	if stringer.NotEmpty(config.Description) {
		cmd.Description = config.Description
	}
	if stringer.NotEmpty(config.ArgsUsage) {
		cmd.ArgsUsage = config.ArgsUsage
	}
	if stringer.NotEmpty(config.Category) {
		cmd.Category = config.Category
	}
	if config.SkipFlagParsing {
		cmd.SkipFlagParsing = config.SkipFlagParsing
	}
	if config.SkipArgReorder {
		cmd.SkipArgReorder = config.SkipArgReorder
	}
	if config.HideHelp {
		cmd.HideHelp = config.HideHelp
	}
	if config.Hidden {
		cmd.Hidden = config.Hidden
	}
	if config.UseShortOptionHandling {
		cmd.UseShortOptionHandling = config.UseShortOptionHandling
	}
	if stringer.NotEmpty(config.HelpName) {
		cmd.HelpName = config.HelpName
	}
	if stringer.NotEmpty(config.CustomHelpTemplate) {
		cmd.CustomHelpTemplate = config.CustomHelpTemplate
	}
	return cmd
}
