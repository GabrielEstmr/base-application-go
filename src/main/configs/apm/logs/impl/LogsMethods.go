package main_configs_apm_logs_impl

import "go.opentelemetry.io/otel/trace"

type LogsMethods interface {
	DEBUG(
		span trace.Span,
		msg string,
		args ...any,
	)

	WARN(
		span trace.Span,
		msg string,
		args ...any,
	)

	INFO(
		span trace.Span,
		msg string,
		args ...any,
	)

	ERROR(
		span trace.Span,
		msg string,
		args ...any,
	)
}
