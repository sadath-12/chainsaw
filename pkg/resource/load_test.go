package resource

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestLoad(t *testing.T) {
	baseDir := filepath.Join("..", "..", "testdata", "resource")
	tests := []struct {
		fileName    string
		expectError bool
		expectedLen int
	}{
		{filepath.Join(baseDir, "valid.yaml"), false, 2},
		{filepath.Join(baseDir, "empty.yaml"), true, 0},
		{filepath.Join(baseDir, "invalid.yaml"), true, 0},
		{filepath.Join(baseDir, "nonexistent.yaml"), true, 0},
	}

	for _, tt := range tests {
		resources, err := Load(tt.fileName)
		if !tt.expectError {
			assert.NoError(t, err)
			assert.Len(t, resources, tt.expectedLen)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestParse(t *testing.T) {
	baseDir := filepath.Join("..", "..", "testdata", "resource")
	tests := []struct {
		fileName          string
		expectError       bool
		expectedLen       int
		expectedResources []unstructured.Unstructured
	}{
		{
			fileName:    filepath.Join(baseDir, "valid.yaml"),
			expectError: false,
			expectedLen: 2,
			expectedResources: []unstructured.Unstructured{
				{
					Object: map[string]interface{}{
						"apiVersion": "v1",
						"kind":       "Pod",
						"metadata": map[string]interface{}{
							"name": "test-pod",
						},
					},
				},
				{
					Object: map[string]interface{}{
						"apiVersion": "v1",
						"kind":       "Service",
						"metadata": map[string]interface{}{
							"name": "test-service",
						},
					},
				},
			},
		},
		{
			fileName:          filepath.Join(baseDir, "empty.yaml"),
			expectError:       false,
			expectedLen:       0,
			expectedResources: []unstructured.Unstructured{},
		},
		{
			fileName:    filepath.Join(baseDir, "invalid.yaml"),
			expectError: true,
			expectedLen: 0,
		},
	}

	for _, tt := range tests {
		content, readErr := os.ReadFile(tt.fileName)
		assert.NoError(t, readErr)

		resources, err := Parse(content)
		if !tt.expectError {
			assert.NoError(t, err)
			assert.Len(t, resources, tt.expectedLen)
			assert.ElementsMatch(t, tt.expectedResources, resources)
		} else {
			assert.Error(t, err)
		}
	}
}
