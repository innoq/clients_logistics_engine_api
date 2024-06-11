// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/v1/logistics.proto

/*
Package logistics_v1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package logistics_v1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

var (
	filter_LogisticsEngineAPI_MoveUnit_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_LogisticsEngineAPI_MoveUnit_0(ctx context.Context, marshaler runtime.Marshaler, client LogisticsEngineAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq MoveUnitRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_LogisticsEngineAPI_MoveUnit_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.MoveUnit(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_LogisticsEngineAPI_MoveUnit_0(ctx context.Context, marshaler runtime.Marshaler, server LogisticsEngineAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq MoveUnitRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_LogisticsEngineAPI_MoveUnit_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.MoveUnit(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_LogisticsEngineAPI_UnitReachedWarehouse_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_LogisticsEngineAPI_UnitReachedWarehouse_0(ctx context.Context, marshaler runtime.Marshaler, client LogisticsEngineAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UnitReachedWarehouseRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_LogisticsEngineAPI_UnitReachedWarehouse_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.UnitReachedWarehouse(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_LogisticsEngineAPI_UnitReachedWarehouse_0(ctx context.Context, marshaler runtime.Marshaler, server LogisticsEngineAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UnitReachedWarehouseRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_LogisticsEngineAPI_UnitReachedWarehouse_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.UnitReachedWarehouse(ctx, &protoReq)
	return msg, metadata, err

}

func request_LogisticsEngineAPI_MetricsReport_0(ctx context.Context, marshaler runtime.Marshaler, client LogisticsEngineAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DefaultRequest
	var metadata runtime.ServerMetadata

	msg, err := client.MetricsReport(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_LogisticsEngineAPI_MetricsReport_0(ctx context.Context, marshaler runtime.Marshaler, server LogisticsEngineAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DefaultRequest
	var metadata runtime.ServerMetadata

	msg, err := server.MetricsReport(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterLogisticsEngineAPIHandlerServer registers the http handlers for service LogisticsEngineAPI to "mux".
// UnaryRPC     :call LogisticsEngineAPIServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterLogisticsEngineAPIHandlerFromEndpoint instead.
func RegisterLogisticsEngineAPIHandlerServer(ctx context.Context, mux *runtime.ServeMux, server LogisticsEngineAPIServer) error {

	mux.Handle("POST", pattern_LogisticsEngineAPI_MoveUnit_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/logistics.api.v1.LogisticsEngineAPI/MoveUnit", runtime.WithHTTPPathPattern("/v1/cargo_unit/move"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_LogisticsEngineAPI_MoveUnit_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LogisticsEngineAPI_MoveUnit_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_LogisticsEngineAPI_UnitReachedWarehouse_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/logistics.api.v1.LogisticsEngineAPI/UnitReachedWarehouse", runtime.WithHTTPPathPattern("/v1/warehouse/cargo_unit/reached"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_LogisticsEngineAPI_UnitReachedWarehouse_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LogisticsEngineAPI_UnitReachedWarehouse_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_LogisticsEngineAPI_MetricsReport_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/logistics.api.v1.LogisticsEngineAPI/MetricsReport", runtime.WithHTTPPathPattern("/v1/report"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_LogisticsEngineAPI_MetricsReport_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LogisticsEngineAPI_MetricsReport_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterLogisticsEngineAPIHandlerFromEndpoint is same as RegisterLogisticsEngineAPIHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterLogisticsEngineAPIHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterLogisticsEngineAPIHandler(ctx, mux, conn)
}

// RegisterLogisticsEngineAPIHandler registers the http handlers for service LogisticsEngineAPI to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterLogisticsEngineAPIHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterLogisticsEngineAPIHandlerClient(ctx, mux, NewLogisticsEngineAPIClient(conn))
}

// RegisterLogisticsEngineAPIHandlerClient registers the http handlers for service LogisticsEngineAPI
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "LogisticsEngineAPIClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "LogisticsEngineAPIClient"
// doesn't go through the normal gRPC flow (creating a gRPC grpc_client etc.) then it will be up to the passed in
// "LogisticsEngineAPIClient" to call the correct interceptors.
func RegisterLogisticsEngineAPIHandlerClient(ctx context.Context, mux *runtime.ServeMux, client LogisticsEngineAPIClient) error {

	mux.Handle("POST", pattern_LogisticsEngineAPI_MoveUnit_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/logistics.api.v1.LogisticsEngineAPI/MoveUnit", runtime.WithHTTPPathPattern("/v1/cargo_unit/move"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_LogisticsEngineAPI_MoveUnit_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LogisticsEngineAPI_MoveUnit_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_LogisticsEngineAPI_UnitReachedWarehouse_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/logistics.api.v1.LogisticsEngineAPI/UnitReachedWarehouse", runtime.WithHTTPPathPattern("/v1/warehouse/cargo_unit/reached"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_LogisticsEngineAPI_UnitReachedWarehouse_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LogisticsEngineAPI_UnitReachedWarehouse_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_LogisticsEngineAPI_MetricsReport_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/logistics.api.v1.LogisticsEngineAPI/MetricsReport", runtime.WithHTTPPathPattern("/v1/report"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_LogisticsEngineAPI_MetricsReport_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_LogisticsEngineAPI_MetricsReport_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_LogisticsEngineAPI_MoveUnit_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "cargo_unit", "move"}, ""))

	pattern_LogisticsEngineAPI_UnitReachedWarehouse_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v1", "warehouse", "cargo_unit", "reached"}, ""))

	pattern_LogisticsEngineAPI_MetricsReport_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "report"}, ""))
)

var (
	forward_LogisticsEngineAPI_MoveUnit_0 = runtime.ForwardResponseMessage

	forward_LogisticsEngineAPI_UnitReachedWarehouse_0 = runtime.ForwardResponseMessage

	forward_LogisticsEngineAPI_MetricsReport_0 = runtime.ForwardResponseMessage
)
