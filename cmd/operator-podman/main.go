package main

import (
	"context"
	"os"

	"github.com/earthboundkid/versioninfo/v2"
	"github.com/urfave/cli/v3"

	_ "github.com/go-orb/plugins/codecs/json"
	_ "github.com/go-orb/plugins/codecs/yaml"
	_ "github.com/go-orb/plugins/log/slog"
)

// Version is the version of the operator-docker-compose application.
//
//nolint:gochecknoglobals
var Version = versioninfo.Short()

func main() {
	cmd := &cli.Command{
		Name:    "octoctl",
		Version: Version,
		Usage:   "podman-compose Operator",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				Usage:    "Set the config file",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "log-level",
				Aliases: []string{"l"},
				Value:   "info",
				Usage:   "Set the log level (debug, info, warn, error)",
			},
		},
		Commands: []*cli.Command{
			startCmd,
			stopCmd,
			restartCmd,
			execCmd,
			logsCmd,
			composeCmd,
			statusCmd,
			showCmd,
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		os.Exit(1)
	}
}
