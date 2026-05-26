// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package admin

import (
	"errors"
	"net/http"
	"sync"

	"github.com/ava-labs/avalanchego/api"
	"github.com/ava-labs/avalanchego/api/server"
	"github.com/ava-labs/avalanchego/chains"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/profiler"
	"github.com/ava-labs/avalanchego/vms"
	"github.com/ava-labs/avalanchego/vms/registry"

	rpcdbpb "github.com/ava-labs/avalanchego/proto/pb/rpcdb"
)

const (
	maxAliasLength = 512

	// Name of file that stacktraces are written to
	stacktraceFile = "stacktrace.txt"
)

var (
	errAliasTooLong = errors.New("alias length is too long")
	errNoLogLevel   = errors.New("need to specify either displayLevel or logLevel")
)

type Config struct {
	Log          logging.Logger
	ProfileDir   string
	LogFactory   logging.Factory
	NodeConfig   interface{}
	DB           database.Database
	ChainManager chains.Manager
	HTTPServer   server.PathAdderWithReadLock
	VMRegistry   registry.VMRegistry
	VMManager    *vms.Manager
}

// Admin is the API service for node admin management
type Admin struct {
	Config
	lock     sync.RWMutex
	profiler profiler.Profiler
}

// NewService returns a new admin API service.
// All of the fields in [config] must be set.
func NewService(config Config) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

// StartCPUProfiler starts a cpu profile writing to the specified file
func (a *Admin) StartCPUProfiler(_ *http.Request, _ *struct{}, _ *api.EmptyReply) error {
	_ = "STUB: not implemented"
	return nil
}

// StopCPUProfiler stops the cpu profile
func (a *Admin) StopCPUProfiler(_ *http.Request, _ *struct{}, _ *api.EmptyReply) error {
	_ = "STUB: not implemented"
	return nil
}

// MemoryProfile runs a memory profile writing to the specified file
func (a *Admin) MemoryProfile(_ *http.Request, _ *struct{}, _ *api.EmptyReply) error {
	_ = "STUB: not implemented"
	return nil
}

// LockProfile runs a mutex profile writing to the specified file
func (a *Admin) LockProfile(_ *http.Request, _ *struct{}, _ *api.EmptyReply) error {
	_ = "STUB: not implemented"
	return nil
}

// AliasArgs are the arguments for calling Alias
type AliasArgs struct {
	Endpoint string `json:"endpoint"`
	Alias    string `json:"alias"`
}

// Alias attempts to alias an HTTP endpoint to a new name
func (a *Admin) Alias(_ *http.Request, args *AliasArgs, _ *api.EmptyReply) error {
	_ = "STUB: not implemented"
	return nil
}

// AliasChainArgs are the arguments for calling AliasChain
type AliasChainArgs struct {
	Chain string `json:"chain"`
	Alias string `json:"alias"`
}

// AliasChain attempts to alias a chain to a new name
func (a *Admin) AliasChain(_ *http.Request, args *AliasChainArgs, _ *api.EmptyReply) error {
	_ = "STUB: not implemented"
	return nil
}

// GetChainAliasesArgs are the arguments for calling GetChainAliases
type GetChainAliasesArgs struct {
	Chain string `json:"chain"`
}

// GetChainAliasesReply are the aliases of the given chain
type GetChainAliasesReply struct {
	Aliases []string `json:"aliases"`
}

// GetChainAliases returns the aliases of the chain
func (a *Admin) GetChainAliases(_ *http.Request, args *GetChainAliasesArgs, reply *GetChainAliasesReply) error {
	_ = "STUB: not implemented"
	return nil
}

// Stacktrace returns the current global stacktrace
func (a *Admin) Stacktrace(_ *http.Request, _ *struct{}, _ *api.EmptyReply) error {
	_ = "STUB: not implemented"
	return nil
}

type SetLoggerLevelArgs struct {
	LoggerName   string         `json:"loggerName"`
	LogLevel     *logging.Level `json:"logLevel"`
	DisplayLevel *logging.Level `json:"displayLevel"`
}

type LogAndDisplayLevels struct {
	LogLevel     logging.Level `json:"logLevel"`
	DisplayLevel logging.Level `json:"displayLevel"`
}

type LoggerLevelReply struct {
	LoggerLevels map[string]LogAndDisplayLevels `json:"loggerLevels"`
}

// SetLoggerLevel sets the log level and/or display level for loggers.
// If len([args.LoggerName]) == 0, sets the log/display level of all loggers.
// Otherwise, sets the log/display level of the loggers named in that argument.
// Sets the log level of these loggers to args.LogLevel.
// If args.LogLevel == nil, doesn't set the log level of these loggers.
// If args.LogLevel != nil, must be a valid string representation of a log level.
// Sets the display level of these loggers to args.LogLevel.
// If args.DisplayLevel == nil, doesn't set the display level of these loggers.
// If args.DisplayLevel != nil, must be a valid string representation of a log level.
func (a *Admin) SetLoggerLevel(_ *http.Request, args *SetLoggerLevelArgs, reply *LoggerLevelReply) error {
	_ = "STUB: not implemented"
	return nil
}

type GetLoggerLevelArgs struct {
	LoggerName string `json:"loggerName"`
}

// GetLoggerLevel returns the log level and display level of all loggers.
func (a *Admin) GetLoggerLevel(_ *http.Request, args *GetLoggerLevelArgs, reply *LoggerLevelReply) error {
	_ = "STUB: not implemented"
	return nil
}

// GetConfig returns the config that the node was started with.
func (a *Admin) GetConfig(_ *http.Request, _ *struct{}, reply *interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadVMsReply contains the response metadata for LoadVMs
type LoadVMsReply struct {
	// VMs and their aliases which were successfully loaded
	NewVMs map[ids.ID][]string `json:"newVMs"`
	// VMs that failed to be loaded and the error message
	FailedVMs map[ids.ID]string `json:"failedVMs,omitempty"`
}

// LoadVMs loads any new VMs available to the node and returns the added VMs.
func (a *Admin) LoadVMs(r *http.Request, _ *struct{}, reply *LoadVMsReply) error {
	_ = "STUB: not implemented"
	return nil
}

// extract the inner error messages

func (a *Admin) getLoggerNames(loggerName string) []string { _ = "STUB: not implemented"; return nil }

// Empty name means all loggers

func (a *Admin) getLogLevels(loggerNames []string) (map[string]LogAndDisplayLevels, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DBGetArgs struct {
	Key string `json:"key"`
}

type DBGetReply struct {
	Value     string        `json:"value"`
	ErrorCode rpcdbpb.Error `json:"errorCode"`
}

//nolint:staticcheck // renaming this method to DBGet would change the API method from "dbGet" to "dBGet"
func (a *Admin) DbGet(_ *http.Request, args *DBGetArgs, reply *DBGetReply) error {
	_ = "STUB: not implemented"
	return nil
}
