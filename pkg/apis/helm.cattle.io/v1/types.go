package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HelmChart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HelmChartSpec   `json:"spec,omitempty"`
	Status HelmChartStatus `json:"status,omitempty"`
}

type HelmChartSpec struct {
	TargetNamespace string                        `json:"targetNamespace,omitempty"`
	Chart           string                        `json:"chart,omitempty"`
	Version         string                        `json:"version,omitempty"`
	Repo            string                        `json:"repo,omitempty"`
	RepoCA          string                        `json:"repoCA,omitempty"`
	Set             map[string]intstr.IntOrString `json:"set,omitempty"`
	ValuesContent   string                        `json:"valuesContent,omitempty"`
	HelmVersion     string                        `json:"helmVersion,omitempty"`
	Bootstrap       bool                          `json:"bootstrap,omitempty"`
	ChartContent    string                        `json:"chartContent,omitempty"`
	JobImage        string                        `json:"jobImage,omitempty"`
	Timeout         *metav1.Duration              `json:"timeout,omitempty"`
	FailurePolicy   string                        `json:"failurePolicy,omitempty"`
}

type HelmChartStatus struct {
	JobName string `json:"jobName,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HelmChartConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec HelmChartConfigSpec `json:"spec,omitempty"`
}

type HelmChartConfigSpec struct {
	ValuesContent string `json:"valuesContent,omitempty"`
	FailurePolicy string `json:"failurePolicy,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HelmRelease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec HelmReleaseSpec `json:"spec,omitempty"`
}

type Status string

// Info describes release information.
type Info struct {
	// FirstDeployed is when the release was first deployed.
	FirstDeployed metav1.Time `json:"first_deployed,omitempty"`
	// LastDeployed is when the release was last deployed.
	LastDeployed metav1.Time `json:"last_deployed,omitempty"`
	// Deleted tracks when this object was deleted.
	Deleted metav1.Time `json:"deleted"`
	// Description is human-friendly "log entry" about this release.
	Description string `json:"description,omitempty"`
	// Status is the current state of the release
	Status Status `json:"status,omitempty"`
	// Contains the rendered templates/NOTES.txt if available
	Notes string `json:"notes,omitempty"`
}

type HelmReleaseSpec struct {
	// Name is the name of the release
	Name string `json:"name,omitempty"`
	// Info provides information about a release
	Info *Info `json:"info,omitempty"`
	// Chart is the chart that was released.
	Chart *HelmChartSpec `json:"chart,omitempty"`
	// Config is the set of extra Values added to the chart.
	// These values override the default values inside of the chart.
	Config map[string]string `json:"config,omitempty"`
	// Manifest is the string representation of the rendered template.
	Manifest string `json:"manifest,omitempty"`
	// Hooks are all of the hooks declared for this release.
	// Hooks []*Hook `json:"hooks,omitempty"`
	// Version is an int which represents the revision of the release.
	Version int `json:"version,omitempty"`
	// Namespace is the kubernetes namespace of the release.
	Namespace string `json:"namespace,omitempty"`
	// Labels of the release.
	// Disabled encoding into Json cause labels are stored in storage driver metadata field.
	Labels map[string]string `json:"-"`
}
