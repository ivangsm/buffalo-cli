package i18n

import (
	"github.com/gobuffalo/buffalo-cli/v2/plugins"
)

func Plugins() []plugins.Plugin {
	return []plugins.Plugin{
		&Generator{},
	}
}
