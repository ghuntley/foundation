// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package ops

import (
	"context"

	"google.golang.org/protobuf/proto"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace/source/protos"
)

type rnode struct {
	def *schema.SerializedInvocation
	reg *registration
	res *HandleResult
	err error // Error captured from a previous run.
}

type registration struct {
	key          string
	tmpl         proto.Message
	dispatcher   dispatcherFunc
	startSession startSessionFunc
	after        []string
}

type dispatcherFunc func(context.Context, *schema.SerializedInvocation, proto.Message) (*HandleResult, error)
type startSessionFunc func(context.Context) (dispatcherFunc, commitSessionFunc, error)
type commitSessionFunc func() error

var handlers = map[string]*registration{}

func Register[M proto.Message](mr Dispatcher[M]) {
	var startSession startSessionFunc
	if stateful, ok := mr.(BatchedDispatcher[M]); ok {
		startSession = func(ctx context.Context) (dispatcherFunc, commitSessionFunc, error) {
			st, err := stateful.StartSession(ctx)
			if err != nil {
				return nil, nil, err
			}
			return func(ctx context.Context, def *schema.SerializedInvocation, msg proto.Message) (*HandleResult, error) {
					return st.Handle(ctx, def, msg.(M))
				}, func() error {
					return st.Commit()
				}, nil
		}
	}

	register[M](func(ctx context.Context, def *schema.SerializedInvocation, msg proto.Message) (*HandleResult, error) {
		return mr.Handle(ctx, def, msg.(M))
	}, startSession)
}

func RegisterFunc[M proto.Message](mr func(ctx context.Context, def *schema.SerializedInvocation, m M) (*HandleResult, error)) {
	register[M](func(ctx context.Context, def *schema.SerializedInvocation, msg proto.Message) (*HandleResult, error) {
		return mr(ctx, def, msg.(M))
	}, nil)
}

func RunAfter(base, after proto.Message) {
	h := handlers[protos.TypeUrl(after)]
	h.after = append(h.after, protos.TypeUrl(base))
}

func register[M proto.Message](dispatcher dispatcherFunc, startSession startSessionFunc) {
	tmpl := protos.NewFromType[M]()
	reg := registration{
		key:          protos.TypeUrl(tmpl),
		tmpl:         tmpl,
		dispatcher:   dispatcher,
		startSession: startSession,
	}

	handlers[protos.TypeUrl(tmpl)] = &reg
}
