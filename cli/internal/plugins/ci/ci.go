package ci

import "github.com/gobuffalo/plugins"

func Plugins() []plugins.Plugin {
	return []plugins.Plugin{
		&Generator{},
	}
}