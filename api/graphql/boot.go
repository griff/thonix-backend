package graphql

import (
	"github.com/griff/thonix/server"
	graphql "github.com/neelance/graphql-go"
	"regexp"
	"time"
)

type bootResolver struct{
	server.Boot
}

func (r *Resolver) Boot() (*bootResolver, error) {
	boot, err := r.Server.Boot()
	if err != nil {
		return nil, err
	}
	return &bootResolver{*boot}, nil
}

func (br *bootResolver) TotalSteps() int32 {
	return br.Boot.TotalSteps()
}

type bootStepResolver struct {
	name    string
	entries []*logEntryResolver
}

var thonixBlock = regexp.MustCompile(`^thonix:\s+(.*)$`)

func (br *bootResolver) Steps() ([]*bootStepResolver, error) {
	blocks, err := br.Boot.BlocksUntil(1*time.Second)
	if err != nil {
		return nil, err
	}

	result := make([]*bootStepResolver, 0)
	for _, block := range blocks {
		entries := make([]*logEntryResolver, 0)
		for _, entry := range block.Entries {
			entries = append(entries, &logEntryResolver{
				name:     "boot",
				LogEntry: entry,
			})
		}
		result = append(result, &bootStepResolver {
			name: block.Name,
			entries: entries,
		})
	}
	return result, nil;
}

func (s *bootStepResolver) Id() graphql.ID {
	return graphql.ID("bootstep:" + s.name)
}

func (s *bootStepResolver) Name() string {
	return s.name
}

func (s *bootStepResolver) Entries() []*logEntryResolver {
	return s.entries
}
