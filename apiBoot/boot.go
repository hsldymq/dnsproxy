package apiBoot

import (
	"context"
	"github.com/AdguardTeam/dnsproxy/internal/cmd"
	"github.com/AdguardTeam/golibs/logutil/slogutil"
	"io"
	"log/slog"
)

func StartProxy(logWriter io.Writer, conf *cmd.Configuration) error {
	l := slogutil.New(&slogutil.Config{
		Output: logWriter,
		Format: slogutil.FormatDefault,
		Level:  slog.LevelInfo,
		// TODO(d.kolyshev): Consider making configurable.
		AddTimestamp: true,
	})

	return StartProxy2(l, conf)
}

func StartProxy2(l *slog.Logger, conf *cmd.Configuration) error {
	return cmd.RunProxy(context.Background(), l, conf)
}
