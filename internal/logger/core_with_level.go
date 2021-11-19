package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type coreWithLevel struct {
	zapcore.Core
	level zapcore.Level
}

//Enabled comment for linter
func (c *coreWithLevel) Enabled(l zapcore.Level) bool {
	return c.level.Enabled(l)
}

//Check comment for linter
func (c *coreWithLevel) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}

//With comment for linter
func (c *coreWithLevel) With(fields []zapcore.Field) zapcore.Core {
	return &coreWithLevel{
		c.Core.With(fields),
		c.level,
	}
}

//WithLevel comment for linter
func WithLevel(lvl zapcore.Level) zap.Option {
	return zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return &coreWithLevel{core, lvl}
	})
}