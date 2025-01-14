package fs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscoverFolders(t *testing.T) {
	root := t.TempDir()
	dirs := []string{"dir1", "dir2", "dir2/subdir1", "dir3", "dir3/subdir1", "dir3/subdir1/subsubdir1"}
	files := []string{"file1", "dir2/file2", "dir3/subdir1/file3"}

	for _, dir := range dirs {
		assert.NoError(t, os.MkdirAll(filepath.Join(root, dir), os.ModePerm))
	}
	for _, file := range files {
		assert.NoError(t, os.WriteFile(filepath.Join(root, file), []byte("test"), os.ModePerm))
	}

	discovered, err := DiscoverFolders(root)
	assert.NoError(t, err)
	expectedDirs := []string{root}
	for _, dir := range dirs {
		expectedDirs = append(expectedDirs, filepath.Join(root, dir))
	}

	assert.ElementsMatch(t, expectedDirs, discovered)
}
