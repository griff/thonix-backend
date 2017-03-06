package server

import (
  "github.com/griff/thonix/dmesg"
  "log"
  "regexp"
  "strings"
  "time"
)

type Boot struct {
  totalSteps int32
  blocks chan LogBlock
  logger dmesg.LogReader
  ch <-chan dmesg.LogEntry
  errch <-chan error
}

type LogBlock struct {
  Name string
  Entries chan dmesg.LogEntry
}

type ReadLogBlock struct {
  Name string
  Entries []dmesg.LogEntry
}

func NewBoot(steps int32) (*Boot, error) {
  log, err := dmesg.OpenDmesg()
  if err != nil {
    return nil, err
  }
  ch, errch := dmesg.Channel(log)
  ret := &Boot {
    totalSteps: steps,
    logger: log,
    ch: ch,
    errch: errch,
    blocks: make(chan LogBlock),
  }
  go ret.blockReader()
  return ret, nil
}

func (b *Boot) TotalSteps() int32 {
  return b.totalSteps
}

func (b *Boot) Blocks() <-chan LogBlock {
  return b.blocks
}

var thonixBlock = regexp.MustCompile(`^thonix:\s+(.*)$`)

func (b *Boot) blockReader() {
  current := LogBlock{
    Name:    "booting",
    Entries: make(chan dmesg.LogEntry),
  }
  firstEntry := true

  for entry := range b.ch {
    m := thonixBlock.FindStringSubmatch(entry.Message)
    if len(m) == 2 {
      close(current.Entries)
      block := strings.TrimSpace(m[1])
      current = LogBlock{
        Name:    block,
        Entries: make(chan dmesg.LogEntry),
      }
      log.Printf("Found new block '%v'", block)
      b.blocks <- current
    } else {
      if firstEntry {
        b.blocks <- current
      }
      current.Entries <- entry
    }
    firstEntry = false
  }
  close(b.blocks)
}

func (b *Boot) BlocksUntil(after time.Duration) ([]ReadLogBlock, error) {
  ret := make([]ReadLogBlock, 0)
  block := ReadLogBlock{}
  entries := make(chan dmesg.LogEntry)
  var err error
  for {
    select {
    case logBlock := <-b.blocks:
      block := ReadLogBlock {
        Name: logBlock.Name,
        Entries: make([]dmesg.LogEntry, 0),
      }
      ret = append(ret, block)
      entries = logBlock.Entries

    case entry := <-entries:
      block = ret[len(ret)-1]
      block.Entries = append(block.Entries, entry)      
      ret[len(ret)-1] = block

    case err := <-b.errch:
      return ret, err

    case <-time.After(after):
      return ret, err
    }
  }
}