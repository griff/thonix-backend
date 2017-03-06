package graphql

import (
  "github.com/griff/thonix-backend/server"
	graphql "github.com/neelance/graphql-go"
)

const schema = `
  schema {
    query: Query
  }

  type Query {
    serverState: State!
    dmesg: [LogEntry]!
    boot: Boot!
  }
  
  enum State {
    booting
    bootFailed
    installing
    running
  }

  type LogConnection {
    totalCount: Int!
    edges: [LogEdge]
    entries: [LogEntry]
    pageInfo: PageInfo!
  }

  type PageInfo {
    startCursor: ID
    endCursor: ID
    hasNextPage: Boolean
  }

  type LogEdge {
    cursor: ID!
    node: LogEntry
  }

  type LogEntry {
    id: ID!
    level: LogLevel!
    facility: LogFacility!
    seqNum: Int!
    timestamp: Timestamp!
    message: String!
    tags: [Tag]!
    tagByName(name: String!): Tag
  }

  type Tag {
    key: String!
    value: String!
  }

  enum LogLevel {
    EMERG
    ALERT
    CRIT
    ERR
    WARNING
    NOTICE
    INFO
    DEBUG
  }

  enum LogFacility {
    KERN
    USER
    MAIL
    DAEMON
    AUTH
    SYSLOG
    LPR
    NEWS
    UUCP
    CRON
    AUTHPRIV
    FTP
  }

  scalar Timestamp

  type Boot {
    totalSteps: Int!
    steps: [LogBlock]!
  }

  type LogBlock {
    id: ID!
    name: String!
    entries: [LogEntry]!
  }
`

type Resolver struct{
  *server.Server
}

func Schema(s *server.Server) (*graphql.Schema, error) {
	return graphql.ParseSchema(schema, &Resolver{s})
}

func (r *Resolver) ServerState() string {
  return r.Server.Status().String()
}

