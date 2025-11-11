package storage

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// FileLock represents a file lock
type FileLock struct {
	path string
	file *os.File
}

// LockFile acquires an advisory lock on a file
func LockFile(path string) (*FileLock, error) {
	// TODO: Implement cross-platform file locking
	// - Linux/macOS: Use golang.org/x/sys/unix Flock
	// - Windows: Use LockFileEx via syscall
	// For now, simple implementation using lock files

	lockPath := path + ".lock"
	lockFile, err := os.OpenFile(lockPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire lock: %w", err)
	}

	return &FileLock{
		path: lockPath,
		file: lockFile,
	}, nil
}

// UnlockFile releases a file lock
func UnlockFile(lock *FileLock) error {
	if lock == nil {
		return nil
	}

	lock.file.Close()
	return os.Remove(lock.path)
}

// AtomicWrite writes data atomically by writing to a temp file then renaming
func AtomicWrite(path string, data []byte) error {
	// TODO: Implement atomic write
	// 1. Write to {path}.tmp
	// 2. Rename to {path} (atomic operation)
	// 3. This prevents corruption if interrupted

	dir := filepath.Dir(path)
	tmpFile, err := ioutil.TempFile(dir, ".tmp-*")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tmpPath := tmpFile.Name()

	// Write data
	_, err = tmpFile.Write(data)
	if err != nil {
		tmpFile.Close()
		os.Remove(tmpPath)
		return fmt.Errorf("failed to write temp file: %w", err)
	}

	// Close temp file
	err = tmpFile.Close()
	if err != nil {
		os.Remove(tmpPath)
		return fmt.Errorf("failed to close temp file: %w", err)
	}

	// Atomic rename
	err = os.Rename(tmpPath, path)
	if err != nil {
		os.Remove(tmpPath)
		return fmt.Errorf("failed to rename temp file: %w", err)
	}

	return nil
}
