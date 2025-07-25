// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"context"

	conventions "go.opentelemetry.io/otel/semconv/v1.9.0"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/filter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/receiver"
)

type EventAttributeOption interface {
	apply(plog.LogRecord)
}

type eventAttributeOptionFunc func(plog.LogRecord)

func (eaof eventAttributeOptionFunc) apply(lr plog.LogRecord) {
	eaof(lr)
}

func WithOptionalIntAttrEventAttribute(optionalIntAttrAttributeValue int64) EventAttributeOption {
	return eventAttributeOptionFunc(func(dp plog.LogRecord) {
		dp.Attributes().PutInt("optional_int_attr", optionalIntAttrAttributeValue)
	})
}

func WithOptionalStringAttrEventAttribute(optionalStringAttrAttributeValue string) EventAttributeOption {
	return eventAttributeOptionFunc(func(dp plog.LogRecord) {
		dp.Attributes().PutStr("optional_string_attr", optionalStringAttrAttributeValue)
	})
}

type eventDefaultEvent struct {
	data   plog.LogRecordSlice // data buffer for generated log records.
	config EventConfig         // event config provided by user.
}

func (e *eventDefaultEvent) recordEvent(ctx context.Context, timestamp pcommon.Timestamp, stringAttrAttributeValue string, overriddenIntAttrAttributeValue int64, enumAttrAttributeValue string, sliceAttrAttributeValue []any, mapAttrAttributeValue map[string]any, options ...EventAttributeOption) {
	if !e.config.Enabled {
		return
	}
	dp := e.data.AppendEmpty()
	dp.SetEventName("default.event")
	dp.SetTimestamp(timestamp)

	if span := trace.SpanContextFromContext(ctx); span.IsValid() {
		dp.SetTraceID(pcommon.TraceID(span.TraceID()))
		dp.SetSpanID(pcommon.SpanID(span.SpanID()))
	}
	dp.Attributes().PutStr("string_attr", stringAttrAttributeValue)
	dp.Attributes().PutInt("state", overriddenIntAttrAttributeValue)
	dp.Attributes().PutStr("enum_attr", enumAttrAttributeValue)
	dp.Attributes().PutEmptySlice("slice_attr").FromRaw(sliceAttrAttributeValue)
	dp.Attributes().PutEmptyMap("map_attr").FromRaw(mapAttrAttributeValue)

	for _, op := range options {
		op.apply(dp)
	}
}

// emit appends recorded event data to a events slice and prepares it for recording another set of log records.
func (e *eventDefaultEvent) emit(lrs plog.LogRecordSlice) {
	if e.config.Enabled && e.data.Len() > 0 {
		e.data.MoveAndAppendTo(lrs)
	}
}

func newEventDefaultEvent(cfg EventConfig) eventDefaultEvent {
	e := eventDefaultEvent{config: cfg}
	if cfg.Enabled {
		e.data = plog.NewLogRecordSlice()
	}
	return e
}

type eventDefaultEventToBeRemoved struct {
	data   plog.LogRecordSlice // data buffer for generated log records.
	config EventConfig         // event config provided by user.
}

func (e *eventDefaultEventToBeRemoved) recordEvent(ctx context.Context, timestamp pcommon.Timestamp, stringAttrAttributeValue string, overriddenIntAttrAttributeValue int64, enumAttrAttributeValue string, sliceAttrAttributeValue []any, mapAttrAttributeValue map[string]any) {
	if !e.config.Enabled {
		return
	}
	dp := e.data.AppendEmpty()
	dp.SetEventName("default.event.to_be_removed")
	dp.SetTimestamp(timestamp)

	if span := trace.SpanContextFromContext(ctx); span.IsValid() {
		dp.SetTraceID(pcommon.TraceID(span.TraceID()))
		dp.SetSpanID(pcommon.SpanID(span.SpanID()))
	}
	dp.Attributes().PutStr("string_attr", stringAttrAttributeValue)
	dp.Attributes().PutInt("state", overriddenIntAttrAttributeValue)
	dp.Attributes().PutStr("enum_attr", enumAttrAttributeValue)
	dp.Attributes().PutEmptySlice("slice_attr").FromRaw(sliceAttrAttributeValue)
	dp.Attributes().PutEmptyMap("map_attr").FromRaw(mapAttrAttributeValue)

}

// emit appends recorded event data to a events slice and prepares it for recording another set of log records.
func (e *eventDefaultEventToBeRemoved) emit(lrs plog.LogRecordSlice) {
	if e.config.Enabled && e.data.Len() > 0 {
		e.data.MoveAndAppendTo(lrs)
	}
}

func newEventDefaultEventToBeRemoved(cfg EventConfig) eventDefaultEventToBeRemoved {
	e := eventDefaultEventToBeRemoved{config: cfg}
	if cfg.Enabled {
		e.data = plog.NewLogRecordSlice()
	}
	return e
}

type eventDefaultEventToBeRenamed struct {
	data   plog.LogRecordSlice // data buffer for generated log records.
	config EventConfig         // event config provided by user.
}

func (e *eventDefaultEventToBeRenamed) recordEvent(ctx context.Context, timestamp pcommon.Timestamp, stringAttrAttributeValue string, booleanAttrAttributeValue bool, booleanAttr2AttributeValue bool, options ...EventAttributeOption) {
	if !e.config.Enabled {
		return
	}
	dp := e.data.AppendEmpty()
	dp.SetEventName("default.event.to_be_renamed")
	dp.SetTimestamp(timestamp)

	if span := trace.SpanContextFromContext(ctx); span.IsValid() {
		dp.SetTraceID(pcommon.TraceID(span.TraceID()))
		dp.SetSpanID(pcommon.SpanID(span.SpanID()))
	}
	dp.Attributes().PutStr("string_attr", stringAttrAttributeValue)
	dp.Attributes().PutBool("boolean_attr", booleanAttrAttributeValue)
	dp.Attributes().PutBool("boolean_attr2", booleanAttr2AttributeValue)

	for _, op := range options {
		op.apply(dp)
	}
}

// emit appends recorded event data to a events slice and prepares it for recording another set of log records.
func (e *eventDefaultEventToBeRenamed) emit(lrs plog.LogRecordSlice) {
	if e.config.Enabled && e.data.Len() > 0 {
		e.data.MoveAndAppendTo(lrs)
	}
}

func newEventDefaultEventToBeRenamed(cfg EventConfig) eventDefaultEventToBeRenamed {
	e := eventDefaultEventToBeRenamed{config: cfg}
	if cfg.Enabled {
		e.data = plog.NewLogRecordSlice()
	}
	return e
}

// LogsBuilder provides an interface for scrapers to report logs while taking care of all the transformations
// required to produce log representation defined in metadata and user config.
type LogsBuilder struct {
	config                         LogsBuilderConfig // config of the logs builder.
	logsBuffer                     plog.Logs
	logRecordsBuffer               plog.LogRecordSlice
	buildInfo                      component.BuildInfo // contains version information.
	resourceAttributeIncludeFilter map[string]filter.Filter
	resourceAttributeExcludeFilter map[string]filter.Filter
	eventDefaultEvent              eventDefaultEvent
	eventDefaultEventToBeRemoved   eventDefaultEventToBeRemoved
	eventDefaultEventToBeRenamed   eventDefaultEventToBeRenamed
}

// LogBuilderOption applies changes to default logs builder.
type LogBuilderOption interface {
	apply(*LogsBuilder)
}

func NewLogsBuilder(lbc LogsBuilderConfig, settings receiver.Settings) *LogsBuilder {
	if !lbc.Events.DefaultEvent.enabledSetByUser {
		settings.Logger.Warn("[WARNING] Please set `enabled` field explicitly for `default.event`: This event will be disabled by default soon.")
	}
	if lbc.Events.DefaultEventToBeRemoved.Enabled {
		settings.Logger.Warn("[WARNING] `default.event.to_be_removed` should not be enabled: This event is deprecated and will be removed soon.")
	}
	if lbc.Events.DefaultEventToBeRenamed.enabledSetByUser {
		settings.Logger.Warn("[WARNING] `default.event.to_be_renamed` should not be configured: This event is deprecated and will be renamed soon.")
	}
	if !lbc.ResourceAttributes.StringResourceAttrDisableWarning.enabledSetByUser {
		settings.Logger.Warn("[WARNING] Please set `enabled` field explicitly for `string.resource.attr_disable_warning`: This resource_attribute will be disabled by default soon.")
	}
	if lbc.ResourceAttributes.StringResourceAttrRemoveWarning.enabledSetByUser || lbc.ResourceAttributes.StringResourceAttrRemoveWarning.EventsInclude != nil || lbc.ResourceAttributes.StringResourceAttrRemoveWarning.EventsExclude != nil {
		settings.Logger.Warn("[WARNING] `string.resource.attr_remove_warning` should not be configured: This resource_attribute is deprecated and will be removed soon.")
	}
	if lbc.ResourceAttributes.StringResourceAttrToBeRemoved.Enabled {
		settings.Logger.Warn("[WARNING] `string.resource.attr_to_be_removed` should not be enabled: This resource_attribute is deprecated and will be removed soon.")
	}
	lb := &LogsBuilder{
		config:                         lbc,
		logsBuffer:                     plog.NewLogs(),
		logRecordsBuffer:               plog.NewLogRecordSlice(),
		buildInfo:                      settings.BuildInfo,
		eventDefaultEvent:              newEventDefaultEvent(lbc.Events.DefaultEvent),
		eventDefaultEventToBeRemoved:   newEventDefaultEventToBeRemoved(lbc.Events.DefaultEventToBeRemoved),
		eventDefaultEventToBeRenamed:   newEventDefaultEventToBeRenamed(lbc.Events.DefaultEventToBeRenamed),
		resourceAttributeIncludeFilter: make(map[string]filter.Filter),
		resourceAttributeExcludeFilter: make(map[string]filter.Filter),
	}
	if lbc.ResourceAttributes.MapResourceAttr.EventsInclude != nil {
		lb.resourceAttributeIncludeFilter["map.resource.attr"] = filter.CreateFilter(lbc.ResourceAttributes.MapResourceAttr.EventsInclude)
	}
	if lbc.ResourceAttributes.MapResourceAttr.EventsExclude != nil {
		lb.resourceAttributeExcludeFilter["map.resource.attr"] = filter.CreateFilter(lbc.ResourceAttributes.MapResourceAttr.EventsExclude)
	}
	if lbc.ResourceAttributes.OptionalResourceAttr.EventsInclude != nil {
		lb.resourceAttributeIncludeFilter["optional.resource.attr"] = filter.CreateFilter(lbc.ResourceAttributes.OptionalResourceAttr.EventsInclude)
	}
	if lbc.ResourceAttributes.OptionalResourceAttr.EventsExclude != nil {
		lb.resourceAttributeExcludeFilter["optional.resource.attr"] = filter.CreateFilter(lbc.ResourceAttributes.OptionalResourceAttr.EventsExclude)
	}
	if lbc.ResourceAttributes.SliceResourceAttr.EventsInclude != nil {
		lb.resourceAttributeIncludeFilter["slice.resource.attr"] = filter.CreateFilter(lbc.ResourceAttributes.SliceResourceAttr.EventsInclude)
	}
	if lbc.ResourceAttributes.SliceResourceAttr.EventsExclude != nil {
		lb.resourceAttributeExcludeFilter["slice.resource.attr"] = filter.CreateFilter(lbc.ResourceAttributes.SliceResourceAttr.EventsExclude)
	}
	if lbc.ResourceAttributes.StringEnumResourceAttr.EventsInclude != nil {
		lb.resourceAttributeIncludeFilter["string.enum.resource.attr"] = filter.CreateFilter(lbc.ResourceAttributes.StringEnumResourceAttr.EventsInclude)
	}
	if lbc.ResourceAttributes.StringEnumResourceAttr.EventsExclude != nil {
		lb.resourceAttributeExcludeFilter["string.enum.resource.attr"] = filter.CreateFilter(lbc.ResourceAttributes.StringEnumResourceAttr.EventsExclude)
	}
	if lbc.ResourceAttributes.StringResourceAttr.EventsInclude != nil {
		lb.resourceAttributeIncludeFilter["string.resource.attr"] = filter.CreateFilter(lbc.ResourceAttributes.StringResourceAttr.EventsInclude)
	}
	if lbc.ResourceAttributes.StringResourceAttr.EventsExclude != nil {
		lb.resourceAttributeExcludeFilter["string.resource.attr"] = filter.CreateFilter(lbc.ResourceAttributes.StringResourceAttr.EventsExclude)
	}
	if lbc.ResourceAttributes.StringResourceAttrDisableWarning.EventsInclude != nil {
		lb.resourceAttributeIncludeFilter["string.resource.attr_disable_warning"] = filter.CreateFilter(lbc.ResourceAttributes.StringResourceAttrDisableWarning.EventsInclude)
	}
	if lbc.ResourceAttributes.StringResourceAttrDisableWarning.EventsExclude != nil {
		lb.resourceAttributeExcludeFilter["string.resource.attr_disable_warning"] = filter.CreateFilter(lbc.ResourceAttributes.StringResourceAttrDisableWarning.EventsExclude)
	}
	if lbc.ResourceAttributes.StringResourceAttrRemoveWarning.EventsInclude != nil {
		lb.resourceAttributeIncludeFilter["string.resource.attr_remove_warning"] = filter.CreateFilter(lbc.ResourceAttributes.StringResourceAttrRemoveWarning.EventsInclude)
	}
	if lbc.ResourceAttributes.StringResourceAttrRemoveWarning.EventsExclude != nil {
		lb.resourceAttributeExcludeFilter["string.resource.attr_remove_warning"] = filter.CreateFilter(lbc.ResourceAttributes.StringResourceAttrRemoveWarning.EventsExclude)
	}
	if lbc.ResourceAttributes.StringResourceAttrToBeRemoved.EventsInclude != nil {
		lb.resourceAttributeIncludeFilter["string.resource.attr_to_be_removed"] = filter.CreateFilter(lbc.ResourceAttributes.StringResourceAttrToBeRemoved.EventsInclude)
	}
	if lbc.ResourceAttributes.StringResourceAttrToBeRemoved.EventsExclude != nil {
		lb.resourceAttributeExcludeFilter["string.resource.attr_to_be_removed"] = filter.CreateFilter(lbc.ResourceAttributes.StringResourceAttrToBeRemoved.EventsExclude)
	}

	return lb
}

// NewResourceBuilder returns a new resource builder that should be used to build a resource associated with for the emitted logs.
func (lb *LogsBuilder) NewResourceBuilder() *ResourceBuilder {
	return NewResourceBuilder(lb.config.ResourceAttributes)
}

// ResourceLogsOption applies changes to provided resource logs.
type ResourceLogsOption interface {
	apply(plog.ResourceLogs)
}

type resourceLogsOptionFunc func(plog.ResourceLogs)

func (rlof resourceLogsOptionFunc) apply(rl plog.ResourceLogs) {
	rlof(rl)
}

// WithLogsResource sets the provided resource on the emitted ResourceLogs.
// It's recommended to use ResourceBuilder to create the resource.
func WithLogsResource(res pcommon.Resource) ResourceLogsOption {
	return resourceLogsOptionFunc(func(rl plog.ResourceLogs) {
		res.CopyTo(rl.Resource())
	})
}

// AppendLogRecord adds a log record to the logs builder.
func (lb *LogsBuilder) AppendLogRecord(lr plog.LogRecord) {
	lr.MoveTo(lb.logRecordsBuffer.AppendEmpty())
}

// EmitForResource saves all the generated logs under a new resource and updates the internal state to be ready for
// recording another set of log records as part of another resource. This function can be helpful when one scraper
// needs to emit logs from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceLogsOption arguments.
func (lb *LogsBuilder) EmitForResource(options ...ResourceLogsOption) {
	rl := plog.NewResourceLogs()
	rl.SetSchemaUrl(conventions.SchemaURL)
	ils := rl.ScopeLogs().AppendEmpty()
	ils.Scope().SetName(ScopeName)
	ils.Scope().SetVersion(lb.buildInfo.Version)
	lb.eventDefaultEvent.emit(ils.LogRecords())
	lb.eventDefaultEventToBeRemoved.emit(ils.LogRecords())
	lb.eventDefaultEventToBeRenamed.emit(ils.LogRecords())

	for _, op := range options {
		op.apply(rl)
	}

	if lb.logRecordsBuffer.Len() > 0 {
		lb.logRecordsBuffer.MoveAndAppendTo(ils.LogRecords())
		lb.logRecordsBuffer = plog.NewLogRecordSlice()
	}

	for attr, filter := range lb.resourceAttributeIncludeFilter {
		if val, ok := rl.Resource().Attributes().Get(attr); ok && !filter.Matches(val.AsString()) {
			return
		}
	}
	for attr, filter := range lb.resourceAttributeExcludeFilter {
		if val, ok := rl.Resource().Attributes().Get(attr); ok && filter.Matches(val.AsString()) {
			return
		}
	}

	if ils.LogRecords().Len() > 0 {
		rl.MoveTo(lb.logsBuffer.ResourceLogs().AppendEmpty())
	}
}

// Emit returns all the logs accumulated by the logs builder and updates the internal state to be ready for
// recording another set of logs. This function will be responsible for applying all the transformations required to
// produce logs representation defined in metadata and user config.
func (lb *LogsBuilder) Emit(options ...ResourceLogsOption) plog.Logs {
	lb.EmitForResource(options...)
	logs := lb.logsBuffer
	lb.logsBuffer = plog.NewLogs()
	return logs
}

// RecordDefaultEventEvent adds a log record of default.event event.
func (lb *LogsBuilder) RecordDefaultEventEvent(ctx context.Context, timestamp pcommon.Timestamp, stringAttrAttributeValue string, overriddenIntAttrAttributeValue int64, enumAttrAttributeValue AttributeEnumAttr, sliceAttrAttributeValue []any, mapAttrAttributeValue map[string]any, options ...EventAttributeOption) {
	lb.eventDefaultEvent.recordEvent(ctx, timestamp, stringAttrAttributeValue, overriddenIntAttrAttributeValue, enumAttrAttributeValue.String(), sliceAttrAttributeValue, mapAttrAttributeValue, options...)
}

// RecordDefaultEventToBeRemovedEvent adds a log record of default.event.to_be_removed event.
func (lb *LogsBuilder) RecordDefaultEventToBeRemovedEvent(ctx context.Context, timestamp pcommon.Timestamp, stringAttrAttributeValue string, overriddenIntAttrAttributeValue int64, enumAttrAttributeValue AttributeEnumAttr, sliceAttrAttributeValue []any, mapAttrAttributeValue map[string]any) {
	lb.eventDefaultEventToBeRemoved.recordEvent(ctx, timestamp, stringAttrAttributeValue, overriddenIntAttrAttributeValue, enumAttrAttributeValue.String(), sliceAttrAttributeValue, mapAttrAttributeValue)
}

// RecordDefaultEventToBeRenamedEvent adds a log record of default.event.to_be_renamed event.
func (lb *LogsBuilder) RecordDefaultEventToBeRenamedEvent(ctx context.Context, timestamp pcommon.Timestamp, stringAttrAttributeValue string, booleanAttrAttributeValue bool, booleanAttr2AttributeValue bool, options ...EventAttributeOption) {
	lb.eventDefaultEventToBeRenamed.recordEvent(ctx, timestamp, stringAttrAttributeValue, booleanAttrAttributeValue, booleanAttr2AttributeValue, options...)
}
