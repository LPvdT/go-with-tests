package logging

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

func init() {
	w := os.Stderr

	slog.SetDefault(slog.New(tint.NewHandler(colorable.NewColorable(w), &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.DateTime,
		NoColor:    !isatty.IsTerminal(w.Fd()),
	})))
}
