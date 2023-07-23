package customzap

import "github.com/aliworkshop/logger/logger"

func (zl *zapLogger) withMeta(fields logger.Field) *zapLogger {
	l := zl.clone()
	l.Meta = append(zl.Meta, fields)
	return l
}
