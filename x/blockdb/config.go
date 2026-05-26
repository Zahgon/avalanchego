// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blockdb

// DefaultMaxDataFileSize is the default maximum size of the data block file in bytes (500GB).
const DefaultMaxDataFileSize = 500 * 1024 * 1024 * 1024

// DefaultMaxDataFiles is the default maximum number of data files descriptors cached.
const DefaultMaxDataFiles = 10

// DefaultBlockCacheSize is the default size of the block cache.
const DefaultBlockCacheSize uint16 = 256

// DatabaseConfig contains configuration parameters for BlockDB.
type DatabaseConfig struct {
	// IndexDir is the directory where the index file is stored.
	IndexDir string

	// DataDir is the directory where the data files are stored.
	DataDir string

	// MinimumHeight is the lowest block height tracked by the database.
	MinimumHeight uint64

	// MaxDataFileSize sets the maximum size of the data block file in bytes.
	MaxDataFileSize uint64

	// MaxDataFiles is the maximum number of data files descriptors cached.
	MaxDataFiles int

	// BlockCacheSize is the size of the block cache (default: 256).
	BlockCacheSize uint16

	// CheckpointInterval defines how frequently (in blocks) the index file header is updated (default: 1024).
	CheckpointInterval uint64

	// SyncToDisk determines if fsync is called after each write for durability.
	SyncToDisk bool
}

// DefaultConfig returns the default options for BlockDB.
func DefaultConfig() DatabaseConfig { _ = "STUB: not implemented"; return *new(DatabaseConfig) }

// WithDir sets both IndexDir and DataDir to the given value.
func (c DatabaseConfig) WithDir(directory string) DatabaseConfig {
	_ = "STUB: not implemented"
	return *new(DatabaseConfig)
}

// WithIndexDir returns a copy of the config with IndexDir set to the given value.
func (c DatabaseConfig) WithIndexDir(indexDir string) DatabaseConfig {
	_ = "STUB: not implemented"
	return *new(DatabaseConfig)
}

// WithDataDir returns a copy of the config with DataDir set to the given value.
func (c DatabaseConfig) WithDataDir(dataDir string) DatabaseConfig {
	_ = "STUB: not implemented"
	return *new(DatabaseConfig)
}

// WithSyncToDisk returns a copy of the config with SyncToDisk set to the given value.
func (c DatabaseConfig) WithSyncToDisk(syncToDisk bool) DatabaseConfig {
	_ = "STUB: not implemented"
	return *new(DatabaseConfig)
}

// WithMinimumHeight returns a copy of the config with MinimumHeight set to the given value.
func (c DatabaseConfig) WithMinimumHeight(minHeight uint64) DatabaseConfig {
	_ = "STUB: not implemented"
	return *new(DatabaseConfig)
}

// WithMaxDataFileSize returns a copy of the config with MaxDataFileSize set to the given value.
func (c DatabaseConfig) WithMaxDataFileSize(maxSize uint64) DatabaseConfig {
	_ = "STUB: not implemented"
	return *new(DatabaseConfig)
}

// WithMaxDataFiles returns a copy of the config with MaxDataFiles set to the given value.
func (c DatabaseConfig) WithMaxDataFiles(maxFiles int) DatabaseConfig {
	_ = "STUB: not implemented"
	return *new(DatabaseConfig)
}

// WithBlockCacheSize returns a copy of the config with BlockCacheSize set to the given value.
func (c DatabaseConfig) WithBlockCacheSize(size uint16) DatabaseConfig {
	_ = "STUB: not implemented"
	return *new(DatabaseConfig)
}

// WithCheckpointInterval returns a copy of the config with CheckpointInterval set to the given value.
func (c DatabaseConfig) WithCheckpointInterval(interval uint64) DatabaseConfig {
	_ = "STUB: not implemented"
	return *new(DatabaseConfig)
}

// Validate checks if the store options are valid.
func (c DatabaseConfig) Validate() error { _ = "STUB: not implemented"; return nil }
