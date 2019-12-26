package buildcmd

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/gobuffalo/buffalo-cli/plugins"
	"github.com/gobuffalo/here"
	"github.com/stretchr/testify/require"
)

func newRefCtx(t *testing.T, root string) (BuilderContext, here.Info) {
	ctx := plugins.WithIO(context.Background(), plugins.DiscardIO())
	bctx := WithBuilderContext(ctx, nil)
	return bctx, newRef(t, root)
}

func newRef(t *testing.T, root string) here.Info {
	t.Helper()

	info := here.Info{
		Dir:        root,
		ImportPath: "github.com/markbates/coke",
		Name:       "main",
		Root:       root,
		Module: here.Module{
			Path:  "github.com/markbates/coke",
			Main:  true,
			Dir:   root,
			GoMod: filepath.Join(root, "go.mod"),
		},
	}

	return info
}

func Test_BuildCmd_Subcommands(t *testing.T) {
	r := require.New(t)

	b := &builder{}
	all := plugins.Plugins{
		plugins.Background("foo"),
		&beforeBuilder{},
		b,
		&afterBuilder{},
		plugins.Background("bar"),
		&buildVersioner{},
		&templatesValidator{},
		&packager{},
	}

	bc := &BuildCmd{
		pluginsFn: all.ScopedPlugins,
	}

	plugs := bc.SubCommands()
	r.Len(plugs, 1)
	r.Equal(b, plugs[0])
}

func Test_BuildCmd_ScopedPlugins(t *testing.T) {
	r := require.New(t)

	all := plugins.Plugins{
		plugins.Background("foo"),
		&builder{},
		&beforeBuilder{},
		&afterBuilder{},
		plugins.Background("bar"),
		&buildVersioner{},
		&buildImporter{},
		&templatesValidator{},
		&packager{},
	}

	bc := &BuildCmd{
		pluginsFn: all.ScopedPlugins,
	}

	plugs := bc.ScopedPlugins()
	r.NotEqual(all, plugs)

	ep := plugins.Plugins(plugs).ExposedPlugins()

	tot := len(all) - 2
	r.Equal(tot, len(ep))

}
