// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: service.proto

package mbotpbconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	mbotpb "github.com/rkuprov/mbot/pkg/gen/mbotpb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// MBotServerServiceName is the fully-qualified name of the MBotServerService service.
	MBotServerServiceName = "mbot.MBotServerService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MBotServerServiceCreateCustomerProcedure is the fully-qualified name of the MBotServerService's
	// CreateCustomer RPC.
	MBotServerServiceCreateCustomerProcedure = "/mbot.MBotServerService/CreateCustomer"
	// MBotServerServiceGetCustomerProcedure is the fully-qualified name of the MBotServerService's
	// GetCustomer RPC.
	MBotServerServiceGetCustomerProcedure = "/mbot.MBotServerService/GetCustomer"
	// MBotServerServiceGetCustomersAllProcedure is the fully-qualified name of the MBotServerService's
	// GetCustomersAll RPC.
	MBotServerServiceGetCustomersAllProcedure = "/mbot.MBotServerService/GetCustomersAll"
	// MBotServerServiceUpdateCustomerProcedure is the fully-qualified name of the MBotServerService's
	// UpdateCustomer RPC.
	MBotServerServiceUpdateCustomerProcedure = "/mbot.MBotServerService/UpdateCustomer"
	// MBotServerServiceDeleteCustomerProcedure is the fully-qualified name of the MBotServerService's
	// DeleteCustomer RPC.
	MBotServerServiceDeleteCustomerProcedure = "/mbot.MBotServerService/DeleteCustomer"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	mBotServerServiceServiceDescriptor               = mbotpb.File_service_proto.Services().ByName("MBotServerService")
	mBotServerServiceCreateCustomerMethodDescriptor  = mBotServerServiceServiceDescriptor.Methods().ByName("CreateCustomer")
	mBotServerServiceGetCustomerMethodDescriptor     = mBotServerServiceServiceDescriptor.Methods().ByName("GetCustomer")
	mBotServerServiceGetCustomersAllMethodDescriptor = mBotServerServiceServiceDescriptor.Methods().ByName("GetCustomersAll")
	mBotServerServiceUpdateCustomerMethodDescriptor  = mBotServerServiceServiceDescriptor.Methods().ByName("UpdateCustomer")
	mBotServerServiceDeleteCustomerMethodDescriptor  = mBotServerServiceServiceDescriptor.Methods().ByName("DeleteCustomer")
)

// MBotServerServiceClient is a client for the mbot.MBotServerService service.
type MBotServerServiceClient interface {
	CreateCustomer(context.Context, *connect.Request[mbotpb.CreateCustomerRequest]) (*connect.Response[mbotpb.CreateCustomerResponse], error)
	GetCustomer(context.Context, *connect.Request[mbotpb.GetCustomerRequest]) (*connect.Response[mbotpb.GetCustomerResponse], error)
	GetCustomersAll(context.Context, *connect.Request[mbotpb.GetCustomersAllRequest]) (*connect.Response[mbotpb.GetCustomersAllResponse], error)
	UpdateCustomer(context.Context, *connect.Request[mbotpb.UpdateCustomerRequest]) (*connect.Response[mbotpb.UpdateCustomerResponse], error)
	DeleteCustomer(context.Context, *connect.Request[mbotpb.DeleteCustomerRequest]) (*connect.Response[mbotpb.DeleteCustomerResponse], error)
}

// NewMBotServerServiceClient constructs a client for the mbot.MBotServerService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMBotServerServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) MBotServerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &mBotServerServiceClient{
		createCustomer: connect.NewClient[mbotpb.CreateCustomerRequest, mbotpb.CreateCustomerResponse](
			httpClient,
			baseURL+MBotServerServiceCreateCustomerProcedure,
			connect.WithSchema(mBotServerServiceCreateCustomerMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getCustomer: connect.NewClient[mbotpb.GetCustomerRequest, mbotpb.GetCustomerResponse](
			httpClient,
			baseURL+MBotServerServiceGetCustomerProcedure,
			connect.WithSchema(mBotServerServiceGetCustomerMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getCustomersAll: connect.NewClient[mbotpb.GetCustomersAllRequest, mbotpb.GetCustomersAllResponse](
			httpClient,
			baseURL+MBotServerServiceGetCustomersAllProcedure,
			connect.WithSchema(mBotServerServiceGetCustomersAllMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateCustomer: connect.NewClient[mbotpb.UpdateCustomerRequest, mbotpb.UpdateCustomerResponse](
			httpClient,
			baseURL+MBotServerServiceUpdateCustomerProcedure,
			connect.WithSchema(mBotServerServiceUpdateCustomerMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteCustomer: connect.NewClient[mbotpb.DeleteCustomerRequest, mbotpb.DeleteCustomerResponse](
			httpClient,
			baseURL+MBotServerServiceDeleteCustomerProcedure,
			connect.WithSchema(mBotServerServiceDeleteCustomerMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// mBotServerServiceClient implements MBotServerServiceClient.
type mBotServerServiceClient struct {
	createCustomer  *connect.Client[mbotpb.CreateCustomerRequest, mbotpb.CreateCustomerResponse]
	getCustomer     *connect.Client[mbotpb.GetCustomerRequest, mbotpb.GetCustomerResponse]
	getCustomersAll *connect.Client[mbotpb.GetCustomersAllRequest, mbotpb.GetCustomersAllResponse]
	updateCustomer  *connect.Client[mbotpb.UpdateCustomerRequest, mbotpb.UpdateCustomerResponse]
	deleteCustomer  *connect.Client[mbotpb.DeleteCustomerRequest, mbotpb.DeleteCustomerResponse]
}

// CreateCustomer calls mbot.MBotServerService.CreateCustomer.
func (c *mBotServerServiceClient) CreateCustomer(ctx context.Context, req *connect.Request[mbotpb.CreateCustomerRequest]) (*connect.Response[mbotpb.CreateCustomerResponse], error) {
	return c.createCustomer.CallUnary(ctx, req)
}

// GetCustomer calls mbot.MBotServerService.GetCustomer.
func (c *mBotServerServiceClient) GetCustomer(ctx context.Context, req *connect.Request[mbotpb.GetCustomerRequest]) (*connect.Response[mbotpb.GetCustomerResponse], error) {
	return c.getCustomer.CallUnary(ctx, req)
}

// GetCustomersAll calls mbot.MBotServerService.GetCustomersAll.
func (c *mBotServerServiceClient) GetCustomersAll(ctx context.Context, req *connect.Request[mbotpb.GetCustomersAllRequest]) (*connect.Response[mbotpb.GetCustomersAllResponse], error) {
	return c.getCustomersAll.CallUnary(ctx, req)
}

// UpdateCustomer calls mbot.MBotServerService.UpdateCustomer.
func (c *mBotServerServiceClient) UpdateCustomer(ctx context.Context, req *connect.Request[mbotpb.UpdateCustomerRequest]) (*connect.Response[mbotpb.UpdateCustomerResponse], error) {
	return c.updateCustomer.CallUnary(ctx, req)
}

// DeleteCustomer calls mbot.MBotServerService.DeleteCustomer.
func (c *mBotServerServiceClient) DeleteCustomer(ctx context.Context, req *connect.Request[mbotpb.DeleteCustomerRequest]) (*connect.Response[mbotpb.DeleteCustomerResponse], error) {
	return c.deleteCustomer.CallUnary(ctx, req)
}

// MBotServerServiceHandler is an implementation of the mbot.MBotServerService service.
type MBotServerServiceHandler interface {
	CreateCustomer(context.Context, *connect.Request[mbotpb.CreateCustomerRequest]) (*connect.Response[mbotpb.CreateCustomerResponse], error)
	GetCustomer(context.Context, *connect.Request[mbotpb.GetCustomerRequest]) (*connect.Response[mbotpb.GetCustomerResponse], error)
	GetCustomersAll(context.Context, *connect.Request[mbotpb.GetCustomersAllRequest]) (*connect.Response[mbotpb.GetCustomersAllResponse], error)
	UpdateCustomer(context.Context, *connect.Request[mbotpb.UpdateCustomerRequest]) (*connect.Response[mbotpb.UpdateCustomerResponse], error)
	DeleteCustomer(context.Context, *connect.Request[mbotpb.DeleteCustomerRequest]) (*connect.Response[mbotpb.DeleteCustomerResponse], error)
}

// NewMBotServerServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMBotServerServiceHandler(svc MBotServerServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	mBotServerServiceCreateCustomerHandler := connect.NewUnaryHandler(
		MBotServerServiceCreateCustomerProcedure,
		svc.CreateCustomer,
		connect.WithSchema(mBotServerServiceCreateCustomerMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	mBotServerServiceGetCustomerHandler := connect.NewUnaryHandler(
		MBotServerServiceGetCustomerProcedure,
		svc.GetCustomer,
		connect.WithSchema(mBotServerServiceGetCustomerMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	mBotServerServiceGetCustomersAllHandler := connect.NewUnaryHandler(
		MBotServerServiceGetCustomersAllProcedure,
		svc.GetCustomersAll,
		connect.WithSchema(mBotServerServiceGetCustomersAllMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	mBotServerServiceUpdateCustomerHandler := connect.NewUnaryHandler(
		MBotServerServiceUpdateCustomerProcedure,
		svc.UpdateCustomer,
		connect.WithSchema(mBotServerServiceUpdateCustomerMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	mBotServerServiceDeleteCustomerHandler := connect.NewUnaryHandler(
		MBotServerServiceDeleteCustomerProcedure,
		svc.DeleteCustomer,
		connect.WithSchema(mBotServerServiceDeleteCustomerMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/mbot.MBotServerService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MBotServerServiceCreateCustomerProcedure:
			mBotServerServiceCreateCustomerHandler.ServeHTTP(w, r)
		case MBotServerServiceGetCustomerProcedure:
			mBotServerServiceGetCustomerHandler.ServeHTTP(w, r)
		case MBotServerServiceGetCustomersAllProcedure:
			mBotServerServiceGetCustomersAllHandler.ServeHTTP(w, r)
		case MBotServerServiceUpdateCustomerProcedure:
			mBotServerServiceUpdateCustomerHandler.ServeHTTP(w, r)
		case MBotServerServiceDeleteCustomerProcedure:
			mBotServerServiceDeleteCustomerHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMBotServerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMBotServerServiceHandler struct{}

func (UnimplementedMBotServerServiceHandler) CreateCustomer(context.Context, *connect.Request[mbotpb.CreateCustomerRequest]) (*connect.Response[mbotpb.CreateCustomerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mbot.MBotServerService.CreateCustomer is not implemented"))
}

func (UnimplementedMBotServerServiceHandler) GetCustomer(context.Context, *connect.Request[mbotpb.GetCustomerRequest]) (*connect.Response[mbotpb.GetCustomerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mbot.MBotServerService.GetCustomer is not implemented"))
}

func (UnimplementedMBotServerServiceHandler) GetCustomersAll(context.Context, *connect.Request[mbotpb.GetCustomersAllRequest]) (*connect.Response[mbotpb.GetCustomersAllResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mbot.MBotServerService.GetCustomersAll is not implemented"))
}

func (UnimplementedMBotServerServiceHandler) UpdateCustomer(context.Context, *connect.Request[mbotpb.UpdateCustomerRequest]) (*connect.Response[mbotpb.UpdateCustomerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mbot.MBotServerService.UpdateCustomer is not implemented"))
}

func (UnimplementedMBotServerServiceHandler) DeleteCustomer(context.Context, *connect.Request[mbotpb.DeleteCustomerRequest]) (*connect.Response[mbotpb.DeleteCustomerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mbot.MBotServerService.DeleteCustomer is not implemented"))
}
