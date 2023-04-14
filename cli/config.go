package cli

import (
	"github.com/urfave/cli"
	"github.com/ihezebin/sdk/utils/stringer"
	"io"
	"reflect"
	"time"
)

type Config struct {
	// The name of the program. Defaults to path.Base(os.Args[0])
	Name string
	// Full name of command for help, defaults to Name
	HelpName string
	// Description of the program.
	Usage string
	// Text to override the USAGE section of help
	UsageText string
	// Description of the program argument format.
	ArgsUsage string
	// Version of the program
	Version string
	// Description of the program
	Description string
	// Boolean to enable bash completion commands
	EnableBashCompletion bool
	// Boolean to hide built-in help command
	HideHelp bool
	// Boolean to hide built-in version flag and the VERSION section of help
	HideVersion bool
	// Compilation date
	Compiled time.Time
	// List of all authors who contributed
	Authors []Author
	// Copyright of the binary if any
	Copyright string
	// Name of Author (Note: Use App.Authors, this is deprecated)
	Author string
	// Email of Author (Note: Use App.Authors, this is deprecated)
	Email string
	// Writer writer to write output to
	Writer io.Writer
	// ErrWriter writes error output
	ErrWriter io.Writer
	// CustomAppHelpTemplate the text template for app help topic.
	// cli.go uses text/template to render templates. You can
	// render custom help text by setting this variable.
	CustomAppHelpTemplate string
	// Boolean to enable short-option handling so user can combine several
	// single-character bool arguements into one
	// i.e. foobar -o -v -> foobar -ov
	UseShortOptionHandling bool
}

type Author struct {
	Name  string // The Author name
	Email string // The Author email
}

type Option func(config *Config)

func newConfig(options ...Option) Config {
	config := Config{}
	for _, option := range options {
		option(&config)
	}
	return config
}

func WithName(name string) Option {
	return func(config *Config) {
		config.Name = name
	}
}

func WithHelpName(helpName string) Option {
	return func(config *Config) {
		config.HelpName = helpName
	}
}

func WithUsage(usage string) Option {
	return func(config *Config) {
		config.Usage = usage
	}
}

func WithUsageText(usageText string) Option {
	return func(config *Config) {
		config.UsageText = usageText
	}
}

func WithArgsUsage(argsUsage string) Option {
	return func(config *Config) {
		config.ArgsUsage = argsUsage
	}
}

func WithVersion(version string) Option {
	return func(config *Config) {
		config.Version = version
	}
}

func WithDescription(description string) Option {
	return func(config *Config) {
		config.Description = description
	}
}

func WithEnableBashCompletion(enableBashCompletion bool) Option {
	return func(config *Config) {
		config.EnableBashCompletion = enableBashCompletion
	}
}

func WithHideHelp(hideHelp bool) Option {
	return func(config *Config) {
		config.HideHelp = hideHelp
	}
}

func WithHideVersion(hideVersion bool) Option {
	return func(config *Config) {
		config.HideVersion = hideVersion
	}
}

func WithCompiled(compiled time.Time) Option {
	return func(config *Config) {
		config.Compiled = compiled
	}
}

func WithAuthors(authors []Author) Option {
	return func(config *Config) {
		config.Authors = authors
	}
}

func WithCopyright(copyright string) Option {
	return func(config *Config) {
		config.Copyright = copyright
	}
}

func WithAuthor(author string) Option {
	return func(config *Config) {
		config.Author = author
	}
}

func WithEmail(email string) Option {
	return func(config *Config) {
		config.Email = email
	}
}

func WithWriter(writer io.Writer) Option {
	return func(config *Config) {
		config.Writer = writer
	}
}

func WithErrWriter(errWriter io.Writer) Option {
	return func(config *Config) {
		config.ErrWriter = errWriter
	}
}

func WithCustomAppHelpTemplate(customAppHelpTemplate string) Option {
	return func(config *Config) {
		config.CustomAppHelpTemplate = customAppHelpTemplate
	}
}

func WithUseShortOptionHandling(useShortOptionHandling bool) Option {
	return func(config *Config) {
		config.UseShortOptionHandling = useShortOptionHandling
	}
}

func handleAppConfig(app *cli.App, config Config) *cli.App {
	if stringer.NotEmpty(config.Name) {
		app.Name = config.Name
	}
	if stringer.NotEmpty(config.HelpName) {
		app.HelpName = config.HelpName
	}
	if stringer.NotEmpty(config.Usage) {
		app.Usage = config.Usage
	}
	if stringer.NotEmpty(config.UsageText) {
		app.UsageText = config.UsageText
	}
	if stringer.NotEmpty(config.ArgsUsage) {
		app.ArgsUsage = config.ArgsUsage
	}
	if stringer.NotEmpty(config.Version) {
		app.Version = config.Version
	}
	if stringer.NotEmpty(config.Description) {
		app.Description = config.Description
	}
	if config.EnableBashCompletion {
		app.EnableBashCompletion = config.EnableBashCompletion
	}
	if config.HideHelp {
		app.HideHelp = config.HideHelp
	}
	if config.HideVersion {
		app.HideVersion = config.HideVersion
	}
	if !reflect.ValueOf(config.Compiled).IsZero() {
		app.Compiled = config.Compiled
	}
	if stringer.NotEmpty(config.Copyright) {
		app.Copyright = config.Copyright
	}
	if stringer.NotEmpty(config.Author) {
		app.Author = config.Author
	}
	if stringer.NotEmpty(config.Email) {
		app.Email = config.Email
	}
	if len(config.Authors) > 0 {
		auths := make([]cli.Author, 0, len(config.Authors))
		for _, author := range config.Authors {
			auths = append(auths, cli.Author{Name: author.Name, Email: author.Email})
		}
		app.Authors = auths
	}
	if config.Writer != nil {
		app.Writer = config.Writer
	}
	if config.ErrWriter != nil {
		app.ErrWriter = config.ErrWriter
	}
	if stringer.NotEmpty(config.CustomAppHelpTemplate) {
		app.CustomAppHelpTemplate = config.CustomAppHelpTemplate
	}
	if config.UseShortOptionHandling {
		app.UseShortOptionHandling = config.UseShortOptionHandling
	}
	return app
}
