package traceid

import (
	"github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go-opentracing"
)
func IdFromSpan(aspan interface{}) uint64 {
	zspan := aspan.(zipkintracer.Span)
	return zspan.Context().TraceID
}
