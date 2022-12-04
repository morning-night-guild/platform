// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/article/v1/article.proto

package articlev1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/morning-night-guild/platform/pkg/connect/proto/article/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ArticleServiceName is the fully-qualified name of the ArticleService service.
	ArticleServiceName = "proto.article.v1.ArticleService"
)

// ArticleServiceClient is a client for the proto.article.v1.ArticleService service.
type ArticleServiceClient interface {
	// 共有
	// Need X-Api-Key Header
	Share(context.Context, *connect_go.Request[v1.ShareRequest]) (*connect_go.Response[v1.ShareResponse], error)
	// 一覧
	List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error)
}

// NewArticleServiceClient constructs a client for the proto.article.v1.ArticleService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewArticleServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ArticleServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &articleServiceClient{
		share: connect_go.NewClient[v1.ShareRequest, v1.ShareResponse](
			httpClient,
			baseURL+"/proto.article.v1.ArticleService/Share",
			opts...,
		),
		list: connect_go.NewClient[v1.ListRequest, v1.ListResponse](
			httpClient,
			baseURL+"/proto.article.v1.ArticleService/List",
			opts...,
		),
	}
}

// articleServiceClient implements ArticleServiceClient.
type articleServiceClient struct {
	share *connect_go.Client[v1.ShareRequest, v1.ShareResponse]
	list  *connect_go.Client[v1.ListRequest, v1.ListResponse]
}

// Share calls proto.article.v1.ArticleService.Share.
func (c *articleServiceClient) Share(ctx context.Context, req *connect_go.Request[v1.ShareRequest]) (*connect_go.Response[v1.ShareResponse], error) {
	return c.share.CallUnary(ctx, req)
}

// List calls proto.article.v1.ArticleService.List.
func (c *articleServiceClient) List(ctx context.Context, req *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// ArticleServiceHandler is an implementation of the proto.article.v1.ArticleService service.
type ArticleServiceHandler interface {
	// 共有
	// Need X-Api-Key Header
	Share(context.Context, *connect_go.Request[v1.ShareRequest]) (*connect_go.Response[v1.ShareResponse], error)
	// 一覧
	List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error)
}

// NewArticleServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewArticleServiceHandler(svc ArticleServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/proto.article.v1.ArticleService/Share", connect_go.NewUnaryHandler(
		"/proto.article.v1.ArticleService/Share",
		svc.Share,
		opts...,
	))
	mux.Handle("/proto.article.v1.ArticleService/List", connect_go.NewUnaryHandler(
		"/proto.article.v1.ArticleService/List",
		svc.List,
		opts...,
	))
	return "/proto.article.v1.ArticleService/", mux
}

// UnimplementedArticleServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedArticleServiceHandler struct{}

func (UnimplementedArticleServiceHandler) Share(context.Context, *connect_go.Request[v1.ShareRequest]) (*connect_go.Response[v1.ShareResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.article.v1.ArticleService.Share is not implemented"))
}

func (UnimplementedArticleServiceHandler) List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.article.v1.ArticleService.List is not implemented"))
}
