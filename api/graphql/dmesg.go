package graphql

import (
	"fmt"
	"github.com/griff/thonix-backend/dmesg"
	graphql "github.com/neelance/graphql-go"
	"strings"
)

type Timestamp struct {
	dmesg.Timestamp
}

func (_ Timestamp) ImplementsGraphQLType(name string) bool {
	return name == "Timestamp"
}

func (t *Timestamp) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case dmesg.Timestamp:
		t.Timestamp = input
		return nil
	case int64:
		t.Timestamp = dmesg.Timestamp(input)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (*Resolver) Dmesg() []*logEntryResolver {
	return make([]*logEntryResolver, 0)
}

type logEntryResolver struct {
	dmesg.LogEntry
	name string
}

func (l *logEntryResolver) ID() graphql.ID {
	return graphql.ID(l.name + ":" + string(l.LogEntry.SeqNum))
}

func (l *logEntryResolver) Level() string {
	return strings.ToUpper(l.LogEntry.Level.String())
}

func (l *logEntryResolver) Facility() string {
	return strings.ToUpper(l.LogEntry.Facility.String())
}

func (l *logEntryResolver) SeqNum() int32 {
	return int32(l.LogEntry.SeqNum)
}

func (l *logEntryResolver) Timestamp() Timestamp {
	return Timestamp{l.LogEntry.Timestamp}
}

func (l *logEntryResolver) Message() string {
	return l.LogEntry.Message
}

func (l *logEntryResolver) Tags() []*tagResolver {
	ret := make([]*tagResolver, 0)
	for key, value := range l.LogEntry.Tags {
		ret = append(ret, &tagResolver{key, value})
	}
	return ret
}

func (l *logEntryResolver) TagByName(args *struct{ Name string }) *tagResolver {
	if ret, ok := l.LogEntry.Tags[args.Name]; ok {
		return &tagResolver{args.Name, ret}
	}
	return nil
}

type tagResolver struct {
	key   string
	value string
}

func (t *tagResolver) Key() string {
	return t.key
}

func (t *tagResolver) Value() string {
	return t.value
}
