// +build integration_test

package core

import (
	"testing"

	"github.com/gobuffalo/buffalo-cli/v2/internal/v1/genny/docker"
	"github.com/gobuffalo/genny/gentest"
	"github.com/gobuffalo/meta"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	r := require.New(t)
	app := meta.Named("coke", ".")
	(&app).PackageRoot("coke")

	gg, err := New(&Options{
		App: app,
	})
	r.NoError(err)

	run := gentest.NewRunner()
	run.WithGroup(gg)

	r.NoError(run.Run())

	res := run.Results()

	cmds := []string{
		"go mod init coke",
	}
	r.NoError(gentest.CompareCommands(cmds, res.Commands))

	expected := commonExpected
	for _, e := range expected {
		_, err = res.Find(e)
		r.NoError(err)
	}

	unexpected := []string{
		"Dockerfile",
		"database.yml",
		"models/models.go",
		".buffalo.dev.yml",
		"assets/css/application.scss.css",
		"public/assets/application.js",
	}
	for _, u := range unexpected {
		_, err = res.Find(u)
		r.Error(err)
	}

}

func Test_New_Docker(t *testing.T) {
	r := require.New(t)

	app := meta.New(".")
	app.WithModules = false

	gg, err := New(&Options{
		Docker: &docker.Options{},
	})
	r.NoError(err)

	run := gentest.NewRunner()
	run.WithGroup(gg)

	r.NoError(run.Run())

	res := run.Results()

	expected := append(commonExpected, "Dockerfile")
	for _, e := range expected {
		_, err := res.Find(e)
		r.NoError(err)
	}
}

var commonExpected = []string{
	"main.go",
	"actions/app.go",
	"actions/actions_test.go",
	"actions/render.go",
	"actions/home.go",
	"actions/home_test.go",
	"fixtures/sample.toml",
	"grifts/init.go",
	".codeclimate.yml",
	".env",
	"inflections.json",
	"README.md",
}
