// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.0

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type LogicHTTPServer interface {
	GetService(context.Context, *GetServiceReq) (*GetServiceReply, error)
	GroupRecall(context.Context, *GroupRecallRequest) (*SendReply, error)
	GroupSend(context.Context, *GroupSendRequest) (*SendReply, error)
	GroupSendMention(context.Context, *GroupSendMentionRequest) (*SendReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	RoomBroadcast(context.Context, *GroupSendRequest) (*SendReply, error)
	RoomSend(context.Context, *GroupSendRequest) (*SendReply, error)
	SingleRecall(context.Context, *SingleRecallRequest) (*SendReply, error)
	SingleSend(context.Context, *SingleSendRequest) (*SendReply, error)
}

func RegisterLogicHTTPServer(s *http.Server, srv LogicHTTPServer) {
	r := s.Route("/")
	r.POST("v1/single/send", _Logic_SingleSend0_HTTP_Handler(srv))
	r.POST("v1/single/recall", _Logic_SingleRecall0_HTTP_Handler(srv))
	r.POST("v1/group/send", _Logic_GroupSend0_HTTP_Handler(srv))
	r.POST("v1/group/send_mention", _Logic_GroupSendMention0_HTTP_Handler(srv))
	r.POST("v1/group/recall", _Logic_GroupRecall0_HTTP_Handler(srv))
	r.POST("v1/room/send", _Logic_RoomSend0_HTTP_Handler(srv))
	r.POST("v1/room/broadcast", _Logic_RoomBroadcast0_HTTP_Handler(srv))
	r.POST("v1/login", _Logic_Login0_HTTP_Handler(srv))
	r.GET("v1/service", _Logic_GetService0_HTTP_Handler(srv))
}

func _Logic_SingleSend0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SingleSendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.v1.Logic/SingleSend")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SingleSend(ctx, req.(*SingleSendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendReply)
		return ctx.Result(200, reply)
	}
}

func _Logic_SingleRecall0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SingleRecallRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.v1.Logic/SingleRecall")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SingleRecall(ctx, req.(*SingleRecallRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendReply)
		return ctx.Result(200, reply)
	}
}

func _Logic_GroupSend0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupSendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.v1.Logic/GroupSend")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GroupSend(ctx, req.(*GroupSendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendReply)
		return ctx.Result(200, reply)
	}
}

func _Logic_GroupSendMention0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupSendMentionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.v1.Logic/GroupSendMention")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GroupSendMention(ctx, req.(*GroupSendMentionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendReply)
		return ctx.Result(200, reply)
	}
}

func _Logic_GroupRecall0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupRecallRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.v1.Logic/GroupRecall")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GroupRecall(ctx, req.(*GroupRecallRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendReply)
		return ctx.Result(200, reply)
	}
}

func _Logic_RoomSend0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupSendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.v1.Logic/RoomSend")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoomSend(ctx, req.(*GroupSendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendReply)
		return ctx.Result(200, reply)
	}
}

func _Logic_RoomBroadcast0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupSendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.v1.Logic/RoomBroadcast")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoomBroadcast(ctx, req.(*GroupSendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendReply)
		return ctx.Result(200, reply)
	}
}

func _Logic_Login0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.v1.Logic/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Logic_GetService0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetServiceReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.v1.Logic/GetService")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetService(ctx, req.(*GetServiceReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetServiceReply)
		return ctx.Result(200, reply)
	}
}

type LogicHTTPClient interface {
	GetService(ctx context.Context, req *GetServiceReq, opts ...http.CallOption) (rsp *GetServiceReply, err error)
	GroupRecall(ctx context.Context, req *GroupRecallRequest, opts ...http.CallOption) (rsp *SendReply, err error)
	GroupSend(ctx context.Context, req *GroupSendRequest, opts ...http.CallOption) (rsp *SendReply, err error)
	GroupSendMention(ctx context.Context, req *GroupSendMentionRequest, opts ...http.CallOption) (rsp *SendReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	RoomBroadcast(ctx context.Context, req *GroupSendRequest, opts ...http.CallOption) (rsp *SendReply, err error)
	RoomSend(ctx context.Context, req *GroupSendRequest, opts ...http.CallOption) (rsp *SendReply, err error)
	SingleRecall(ctx context.Context, req *SingleRecallRequest, opts ...http.CallOption) (rsp *SendReply, err error)
	SingleSend(ctx context.Context, req *SingleSendRequest, opts ...http.CallOption) (rsp *SendReply, err error)
}

type LogicHTTPClientImpl struct {
	cc *http.Client
}

func NewLogicHTTPClient(client *http.Client) LogicHTTPClient {
	return &LogicHTTPClientImpl{client}
}

func (c *LogicHTTPClientImpl) GetService(ctx context.Context, in *GetServiceReq, opts ...http.CallOption) (*GetServiceReply, error) {
	var out GetServiceReply
	pattern := "v1/service"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.logic.v1.Logic/GetService"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) GroupRecall(ctx context.Context, in *GroupRecallRequest, opts ...http.CallOption) (*SendReply, error) {
	var out SendReply
	pattern := "v1/group/recall"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.v1.Logic/GroupRecall"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) GroupSend(ctx context.Context, in *GroupSendRequest, opts ...http.CallOption) (*SendReply, error) {
	var out SendReply
	pattern := "v1/group/send"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.v1.Logic/GroupSend"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) GroupSendMention(ctx context.Context, in *GroupSendMentionRequest, opts ...http.CallOption) (*SendReply, error) {
	var out SendReply
	pattern := "v1/group/send_mention"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.v1.Logic/GroupSendMention"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.v1.Logic/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) RoomBroadcast(ctx context.Context, in *GroupSendRequest, opts ...http.CallOption) (*SendReply, error) {
	var out SendReply
	pattern := "v1/room/broadcast"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.v1.Logic/RoomBroadcast"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) RoomSend(ctx context.Context, in *GroupSendRequest, opts ...http.CallOption) (*SendReply, error) {
	var out SendReply
	pattern := "v1/room/send"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.v1.Logic/RoomSend"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) SingleRecall(ctx context.Context, in *SingleRecallRequest, opts ...http.CallOption) (*SendReply, error) {
	var out SendReply
	pattern := "v1/single/recall"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.v1.Logic/SingleRecall"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) SingleSend(ctx context.Context, in *SingleSendRequest, opts ...http.CallOption) (*SendReply, error) {
	var out SendReply
	pattern := "v1/single/send"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.v1.Logic/SingleSend"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
