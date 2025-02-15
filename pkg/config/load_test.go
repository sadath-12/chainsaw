package config

import (
	"reflect"
	"testing"
	"time"

	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    *v1alpha1.Configuration
		wantErr bool
	}{
		{
			name:    "confimap",
			path:    "../../testdata/config/configmap.yaml",
			wantErr: true,
		},
		{
			name:    "not found",
			path:    "../../testdata/config/not-found.yaml",
			wantErr: true,
		},
		{
			name:    "empty",
			path:    "../../testdata/config/empty.yaml",
			wantErr: true,
		},
		{
			name: "default",
			path: "../../testdata/config/default.yaml",
			want: &v1alpha1.Configuration{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "chainsaw.kyverno.io/v1alpha1",
					Kind:       "Configuration",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "default",
				},
				Spec: v1alpha1.ConfigurationSpec{
					SkipDelete:       false,
					FailFast:         false,
					ReportFormat:     "",
					ReportName:       "chainsaw-report",
					FullName:         false,
					IncludeTestRegex: "",
					ExcludeTestRegex: "",
				},
			},
		},
		{
			name: "custom-config",
			path: "../../testdata/config/custom-config.yaml",
			want: &v1alpha1.Configuration{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "chainsaw.kyverno.io/v1alpha1",
					Kind:       "Configuration",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "custom-config",
				},
				Spec: v1alpha1.ConfigurationSpec{
					Timeouts: v1alpha1.Timeouts{
						Apply:   &metav1.Duration{Duration: 5 * time.Second},
						Assert:  &metav1.Duration{Duration: 10 * time.Second},
						Error:   &metav1.Duration{Duration: 10 * time.Second},
						Delete:  &metav1.Duration{Duration: 5 * time.Second},
						Cleanup: &metav1.Duration{Duration: 5 * time.Second},
						Exec:    &metav1.Duration{Duration: 10 * time.Second},
					},
					SkipDelete:       true,
					FailFast:         true,
					Parallel:         ptr.To(4),
					ReportFormat:     "JSON",
					ReportName:       "custom-report",
					FullName:         true,
					IncludeTestRegex: "include-*",
					ExcludeTestRegex: "exclude-*",
				},
			},
		},
		{
			name:    "multiple",
			path:    "../../testdata/config/multiple.yaml",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Load(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
