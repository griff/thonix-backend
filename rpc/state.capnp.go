package rpc

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

// Constants defined in state.capnp.
const (
	Server_socketPath      = "/var/run/thonix.sock"
	ServerAdmin_socketPath = "/var/run/thonix-admin.sock"
)

type Status uint16

// Values of Status.
const (
	Status_booting    Status = 0
	Status_bootFailed Status = 1
	Status_installing Status = 2
	Status_running    Status = 3
)

// String returns the enum's constant name.
func (c Status) String() string {
	switch c {
	case Status_booting:
		return "booting"
	case Status_bootFailed:
		return "bootFailed"
	case Status_installing:
		return "installing"
	case Status_running:
		return "running"

	default:
		return ""
	}
}

// StatusFromString returns the enum value with a name,
// or the zero value if there's no such value.
func StatusFromString(c string) Status {
	switch c {
	case "booting":
		return Status_booting
	case "bootFailed":
		return Status_bootFailed
	case "installing":
		return Status_installing
	case "running":
		return Status_running

	default:
		return 0
	}
}

type Status_List struct{ capnp.List }

func NewStatus_List(s *capnp.Segment, sz int32) (Status_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return Status_List{l.List}, err
}

func (l Status_List) At(i int) Status {
	ul := capnp.UInt16List{List: l.List}
	return Status(ul.At(i))
}

func (l Status_List) Set(i int, v Status) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type Server struct{ Client capnp.Client }

func (c Server) State(ctx context.Context, params func(Server_state_Params) error, opts ...capnp.CallOption) Server_state_Results_Promise {
	if c.Client == nil {
		return Server_state_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xf94cc7457609fe37,
			MethodID:      0,
			InterfaceName: "rpc/state.capnp:Server",
			MethodName:    "state",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Server_state_Params{Struct: s}) }
	}
	return Server_state_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Server_Server interface {
	State(Server_state) error
}

func Server_ServerToClient(s Server_Server) Server {
	c, _ := s.(server.Closer)
	return Server{Client: server.New(Server_Methods(nil, s), c)}
}

func Server_Methods(methods []server.Method, s Server_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xf94cc7457609fe37,
			MethodID:      0,
			InterfaceName: "rpc/state.capnp:Server",
			MethodName:    "state",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Server_state{c, opts, Server_state_Params{Struct: p}, Server_state_Results{Struct: r}}
			return s.State(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	return methods
}

// Server_state holds the arguments for a server call to Server.state.
type Server_state struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Server_state_Params
	Results Server_state_Results
}

type Server_state_Params struct{ capnp.Struct }

// Server_state_Params_TypeID is the unique identifier for the type Server_state_Params.
const Server_state_Params_TypeID = 0xb00fb38278d90499

func NewServer_state_Params(s *capnp.Segment) (Server_state_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Server_state_Params{st}, err
}

func NewRootServer_state_Params(s *capnp.Segment) (Server_state_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Server_state_Params{st}, err
}

func ReadRootServer_state_Params(msg *capnp.Message) (Server_state_Params, error) {
	root, err := msg.RootPtr()
	return Server_state_Params{root.Struct()}, err
}

func (s Server_state_Params) String() string {
	str, _ := text.Marshal(0xb00fb38278d90499, s.Struct)
	return str
}

// Server_state_Params_List is a list of Server_state_Params.
type Server_state_Params_List struct{ capnp.List }

// NewServer_state_Params creates a new list of Server_state_Params.
func NewServer_state_Params_List(s *capnp.Segment, sz int32) (Server_state_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Server_state_Params_List{l}, err
}

func (s Server_state_Params_List) At(i int) Server_state_Params {
	return Server_state_Params{s.List.Struct(i)}
}

func (s Server_state_Params_List) Set(i int, v Server_state_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Server_state_Params_Promise is a wrapper for a Server_state_Params promised by a client call.
type Server_state_Params_Promise struct{ *capnp.Pipeline }

func (p Server_state_Params_Promise) Struct() (Server_state_Params, error) {
	s, err := p.Pipeline.Struct()
	return Server_state_Params{s}, err
}

type Server_state_Results struct{ capnp.Struct }

// Server_state_Results_TypeID is the unique identifier for the type Server_state_Results.
const Server_state_Results_TypeID = 0x99111c0510c08861

func NewServer_state_Results(s *capnp.Segment) (Server_state_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Server_state_Results{st}, err
}

func NewRootServer_state_Results(s *capnp.Segment) (Server_state_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Server_state_Results{st}, err
}

func ReadRootServer_state_Results(msg *capnp.Message) (Server_state_Results, error) {
	root, err := msg.RootPtr()
	return Server_state_Results{root.Struct()}, err
}

func (s Server_state_Results) String() string {
	str, _ := text.Marshal(0x99111c0510c08861, s.Struct)
	return str
}

func (s Server_state_Results) State() Status {
	return Status(s.Struct.Uint16(0))
}

func (s Server_state_Results) SetState(v Status) {
	s.Struct.SetUint16(0, uint16(v))
}

// Server_state_Results_List is a list of Server_state_Results.
type Server_state_Results_List struct{ capnp.List }

// NewServer_state_Results creates a new list of Server_state_Results.
func NewServer_state_Results_List(s *capnp.Segment, sz int32) (Server_state_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Server_state_Results_List{l}, err
}

func (s Server_state_Results_List) At(i int) Server_state_Results {
	return Server_state_Results{s.List.Struct(i)}
}

func (s Server_state_Results_List) Set(i int, v Server_state_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Server_state_Results_Promise is a wrapper for a Server_state_Results promised by a client call.
type Server_state_Results_Promise struct{ *capnp.Pipeline }

func (p Server_state_Results_Promise) Struct() (Server_state_Results, error) {
	s, err := p.Pipeline.Struct()
	return Server_state_Results{s}, err
}

type ServerAdmin struct{ Client capnp.Client }

func (c ServerAdmin) SetState(ctx context.Context, params func(ServerAdmin_setState_Params) error, opts ...capnp.CallOption) ServerAdmin_setState_Results_Promise {
	if c.Client == nil {
		return ServerAdmin_setState_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc4445c7546550365,
			MethodID:      0,
			InterfaceName: "rpc/state.capnp:ServerAdmin",
			MethodName:    "setState",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(ServerAdmin_setState_Params{Struct: s}) }
	}
	return ServerAdmin_setState_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c ServerAdmin) State(ctx context.Context, params func(Server_state_Params) error, opts ...capnp.CallOption) Server_state_Results_Promise {
	if c.Client == nil {
		return Server_state_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xf94cc7457609fe37,
			MethodID:      0,
			InterfaceName: "rpc/state.capnp:Server",
			MethodName:    "state",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Server_state_Params{Struct: s}) }
	}
	return Server_state_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type ServerAdmin_Server interface {
	SetState(ServerAdmin_setState) error

	State(Server_state) error
}

func ServerAdmin_ServerToClient(s ServerAdmin_Server) ServerAdmin {
	c, _ := s.(server.Closer)
	return ServerAdmin{Client: server.New(ServerAdmin_Methods(nil, s), c)}
}

func ServerAdmin_Methods(methods []server.Method, s ServerAdmin_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc4445c7546550365,
			MethodID:      0,
			InterfaceName: "rpc/state.capnp:ServerAdmin",
			MethodName:    "setState",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := ServerAdmin_setState{c, opts, ServerAdmin_setState_Params{Struct: p}, ServerAdmin_setState_Results{Struct: r}}
			return s.SetState(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xf94cc7457609fe37,
			MethodID:      0,
			InterfaceName: "rpc/state.capnp:Server",
			MethodName:    "state",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Server_state{c, opts, Server_state_Params{Struct: p}, Server_state_Results{Struct: r}}
			return s.State(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	return methods
}

// ServerAdmin_setState holds the arguments for a server call to ServerAdmin.setState.
type ServerAdmin_setState struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  ServerAdmin_setState_Params
	Results ServerAdmin_setState_Results
}

type ServerAdmin_setState_Params struct{ capnp.Struct }

// ServerAdmin_setState_Params_TypeID is the unique identifier for the type ServerAdmin_setState_Params.
const ServerAdmin_setState_Params_TypeID = 0xdaaaff1e6236b2d2

func NewServerAdmin_setState_Params(s *capnp.Segment) (ServerAdmin_setState_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return ServerAdmin_setState_Params{st}, err
}

func NewRootServerAdmin_setState_Params(s *capnp.Segment) (ServerAdmin_setState_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return ServerAdmin_setState_Params{st}, err
}

func ReadRootServerAdmin_setState_Params(msg *capnp.Message) (ServerAdmin_setState_Params, error) {
	root, err := msg.RootPtr()
	return ServerAdmin_setState_Params{root.Struct()}, err
}

func (s ServerAdmin_setState_Params) String() string {
	str, _ := text.Marshal(0xdaaaff1e6236b2d2, s.Struct)
	return str
}

func (s ServerAdmin_setState_Params) NewState() Status {
	return Status(s.Struct.Uint16(0))
}

func (s ServerAdmin_setState_Params) SetNewState(v Status) {
	s.Struct.SetUint16(0, uint16(v))
}

// ServerAdmin_setState_Params_List is a list of ServerAdmin_setState_Params.
type ServerAdmin_setState_Params_List struct{ capnp.List }

// NewServerAdmin_setState_Params creates a new list of ServerAdmin_setState_Params.
func NewServerAdmin_setState_Params_List(s *capnp.Segment, sz int32) (ServerAdmin_setState_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return ServerAdmin_setState_Params_List{l}, err
}

func (s ServerAdmin_setState_Params_List) At(i int) ServerAdmin_setState_Params {
	return ServerAdmin_setState_Params{s.List.Struct(i)}
}

func (s ServerAdmin_setState_Params_List) Set(i int, v ServerAdmin_setState_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// ServerAdmin_setState_Params_Promise is a wrapper for a ServerAdmin_setState_Params promised by a client call.
type ServerAdmin_setState_Params_Promise struct{ *capnp.Pipeline }

func (p ServerAdmin_setState_Params_Promise) Struct() (ServerAdmin_setState_Params, error) {
	s, err := p.Pipeline.Struct()
	return ServerAdmin_setState_Params{s}, err
}

type ServerAdmin_setState_Results struct{ capnp.Struct }

// ServerAdmin_setState_Results_TypeID is the unique identifier for the type ServerAdmin_setState_Results.
const ServerAdmin_setState_Results_TypeID = 0x98f99489bbc2e286

func NewServerAdmin_setState_Results(s *capnp.Segment) (ServerAdmin_setState_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ServerAdmin_setState_Results{st}, err
}

func NewRootServerAdmin_setState_Results(s *capnp.Segment) (ServerAdmin_setState_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ServerAdmin_setState_Results{st}, err
}

func ReadRootServerAdmin_setState_Results(msg *capnp.Message) (ServerAdmin_setState_Results, error) {
	root, err := msg.RootPtr()
	return ServerAdmin_setState_Results{root.Struct()}, err
}

func (s ServerAdmin_setState_Results) String() string {
	str, _ := text.Marshal(0x98f99489bbc2e286, s.Struct)
	return str
}

// ServerAdmin_setState_Results_List is a list of ServerAdmin_setState_Results.
type ServerAdmin_setState_Results_List struct{ capnp.List }

// NewServerAdmin_setState_Results creates a new list of ServerAdmin_setState_Results.
func NewServerAdmin_setState_Results_List(s *capnp.Segment, sz int32) (ServerAdmin_setState_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return ServerAdmin_setState_Results_List{l}, err
}

func (s ServerAdmin_setState_Results_List) At(i int) ServerAdmin_setState_Results {
	return ServerAdmin_setState_Results{s.List.Struct(i)}
}

func (s ServerAdmin_setState_Results_List) Set(i int, v ServerAdmin_setState_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// ServerAdmin_setState_Results_Promise is a wrapper for a ServerAdmin_setState_Results promised by a client call.
type ServerAdmin_setState_Results_Promise struct{ *capnp.Pipeline }

func (p ServerAdmin_setState_Results_Promise) Struct() (ServerAdmin_setState_Results, error) {
	s, err := p.Pipeline.Struct()
	return ServerAdmin_setState_Results{s}, err
}

const schema_b545ff514b234a75 = "x\xda\x8c\x93\xcdk\x13Q\x14\xc5\xcf\xbd31\x16+" +
	"\xe9s\x84XH\x15\xbf\x10\x85&\xda\xfa\x81\xdd\xb4\x96" +
	"~`U\xc83\xb8\xb0\xb8\x99\xa6C:&\x9d\x94\x99" +
	"\x97XWE\x17\xa2\x88\x0bA\xd4\x82\x0b\xd7UQ+" +
	".\xac\x1b\x11\xc5\xbd\xe0\xa2\x057\xdd\xfb\x07t\xe3\x93" +
	"I3iSju\xf5\x029s~\xe7p\xef=\x9e" +
	"\xa6>>\x11\x9b1\x01y2\xb6M\xdfY\xfe\xfc\xf1" +
	"\xde\xa3\x95\xa7\x10)\x02\xcc8\xd0}\x98\xaf\x13Lm" +
	"\xdf\xfd\xd4\x16K\x89Y\xc8$E\x7f\x09>F \xab" +
	"\x9d{A\xfa\xc9\x8ba\xf3\xf5\xcd\xea3\x88\x94\xa9\x1d" +
	"\xe3\xcaP\xe5\xda\xc0\x17\x80\xba\xcfr\x0fY\xe79\x0e" +
	"\xe4\x06\xd8\xa0\\\x96\x99\x00=k.N\xdf~\x97x" +
	"\x03\x91\x8c\xfc$\x1f\x08Q\x8f\xb9\xf8r\xfc\xfe\x91y" +
	"\x88\xa4\xa9\xcf\xfcn\xa9\x0e~\xbb\xb8\xb2\xea\xb4k3" +
	"\xa7\x06L\xb4\x19\xba2r\xf0\x82\xd4\x83\xef\x01\xb2." +
	"\xf1\xb2u\x95\x93\x80\xe5\xf0\xb0\xf50\xfc\xa5\xbf\xcf\x9f" +
	"\x1e\xdb\xab\xe7\x96 S\x8d\x1e\xb7x,\xec\xf1\xa0\xd6" +
	"\xa3\x01\xdc\xe8\xf6\x96\x17\xac\x0f5\xb7\xaf<l\xfd\x0a" +
	"c\xe8\x1d\x1d\xd5\xe7\xa7\xda_\x85bn\x12\xff\xe0\x05" +
	"\xebg(\xb1\x16\xb9\x80N\xedO\xe53\x81\xb2\x95\xe9" +
	"\xa4\xf3\xf6\x947\xd5\x93s\xfc\xaa\xe3\x9f\x1b\x9ft\xbd" +
	"t\xe0\xa8\x9c\xb2\x95s\xe8\xb2\x13$*%\x154\xe4" +
	"F\xb3<\x1d4\xc9\xa4i\x98\x80I\x80\xd8\xd9\x05\xc8" +
	"\xed\x06\xc9\xddL\xfbj*J\xac\xa5\x03Q\x02\xf47" +
	"\xd7z\x88r\xbe\xd8\xeb\xa8\xac\xad&\xb2D\xd4\x0a\xa6" +
	"V@\xd0\x92\xceTm?\xe3W<\xce\xa8\x89\xb2\xe7" +
	"Nw\xda\xf5\x0f\x8c|\xf1\x1fI\xb3\xb6\x1f\xb7'\xb7" +
	"\xe8S\xce\x17kLj\x86\xce5\xa0T\x87\xa6\x13\xa1" +
	"\xb6a\xc4\x1b+\xc4']O\x9a\xb4~\x0fiTG" +
	"\xfe0\xd4\x844\x8d\xd8\xba\x05\xa0h\xd7\x85\x18\x01\x8b" +
	"\x96\xb8\x8e\xc6\x00\xa0\x8fj^k\xcb\x87\xff\x9c`\xbd" +
	"\xf0\xfa\xc9\x8c\x00\xb2\xd5 \xb9\x87I{\xce\x8d\x08\xb1" +
	"\xc5|(B\xf4\xae2jY\xd6\xaeb\xf3^\xd1=" +
	"Qt\xa9Bt\x81E,\xbe\xba\x0e}\x94\xa5M\x09" +
	"\xcaV\x95 K$\xdb\x88\x01q\xb4?L#\xf6\x8f" +
	"\x02\xc4\xa2#|\x0c\xd1\xde\x0f\xcc\x8c\x95\xcb\xca\xf5\x0a" +
	":|\x87l\xb7\x04\xc3\x19\xd7\xae\x17(\xbbTra" +
	"x\x85\x19\xbf\xe2y\xaeW\xf8\x13\x00\x00\xff\xff\xeeQ" +
	"J2"

func init() {
	schemas.Register(schema_b545ff514b234a75,
		0x98f99489bbc2e286,
		0x99111c0510c08861,
		0x9c7679af0447ab97,
		0xb00fb38278d90499,
		0xb2278b64ac6b0296,
		0xc4445c7546550365,
		0xdaaaff1e6236b2d2,
		0xf94cc7457609fe37,
		0xf9ad1935a0761d0b)
}
