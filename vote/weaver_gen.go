// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package vote

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/eminetto/microservices-serviceweaver/vote/Writer",
		Iface: reflect.TypeOf((*Writer)(nil)).Elem(),
		Impl:  reflect.TypeOf(writer{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return writer_local_stub{impl: impl.(Writer), tracer: tracer, writeMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/eminetto/microservices-serviceweaver/vote/Writer", Method: "Write", Remote: false, Generated: true})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return writer_client_stub{stub: stub, writeMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/eminetto/microservices-serviceweaver/vote/Writer", Method: "Write", Remote: true, Generated: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return writer_server_stub{impl: impl.(Writer), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return writer_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[Writer] = (*writer)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*writer)(nil)

// Local stub implementations.

type writer_local_stub struct {
	impl         Writer
	tracer       trace.Tracer
	writeMetrics *codegen.MethodMetrics
}

// Check that writer_local_stub implements the Writer interface.
var _ Writer = (*writer_local_stub)(nil)

func (s writer_local_stub) Write(ctx context.Context, a0 *Vote) (r0 uuid.UUID, err error) {
	// Update metrics.
	begin := s.writeMetrics.Begin()
	defer func() { s.writeMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "vote.Writer.Write", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Write(ctx, a0)
}

// Client stub implementations.

type writer_client_stub struct {
	stub         codegen.Stub
	writeMetrics *codegen.MethodMetrics
}

// Check that writer_client_stub implements the Writer interface.
var _ Writer = (*writer_client_stub)(nil)

func (s writer_client_stub) Write(ctx context.Context, a0 *Vote) (r0 uuid.UUID, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.writeMetrics.Begin()
	defer func() { s.writeMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "vote.Writer.Write", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Encode arguments.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_Vote_dfbd89aa(enc, a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	dec.DecodeBinaryUnmarshaler(&r0)
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][24]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.24.2 (codegen
version v0.24.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

type writer_server_stub struct {
	impl    Writer
	addLoad func(key uint64, load float64)
}

// Check that writer_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*writer_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s writer_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Write":
		return s.write
	default:
		return nil
	}
}

func (s writer_server_stub) write(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 *Vote
	a0 = serviceweaver_dec_ptr_Vote_dfbd89aa(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Write(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.EncodeBinaryMarshaler(&r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type writer_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that writer_reflect_stub implements the Writer interface.
var _ Writer = (*writer_reflect_stub)(nil)

func (s writer_reflect_stub) Write(ctx context.Context, a0 *Vote) (r0 uuid.UUID, err error) {
	err = s.caller("Write", ctx, []any{a0}, []any{&r0})
	return
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*Vote)(nil)

type __is_Vote[T ~struct {
	ID       uuid.UUID
	Email    string
	TalkName string "json:\"talk_name\""
	Score    int    "json:\"score,string\""
	weaver.AutoMarshal
}] struct{}

var _ __is_Vote[Vote]

func (x *Vote) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Vote.WeaverMarshal: nil receiver"))
	}
	enc.EncodeBinaryMarshaler(&x.ID)
	enc.String(x.Email)
	enc.String(x.TalkName)
	enc.Int(x.Score)
}

func (x *Vote) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Vote.WeaverUnmarshal: nil receiver"))
	}
	dec.DecodeBinaryUnmarshaler(&x.ID)
	x.Email = dec.String()
	x.TalkName = dec.String()
	x.Score = dec.Int()
}

// Encoding/decoding implementations.

func serviceweaver_enc_ptr_Vote_dfbd89aa(enc *codegen.Encoder, arg *Vote) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		(*arg).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_ptr_Vote_dfbd89aa(dec *codegen.Decoder) *Vote {
	if !dec.Bool() {
		return nil
	}
	var res Vote
	(&res).WeaverUnmarshal(dec)
	return &res
}
