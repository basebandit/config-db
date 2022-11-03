package v1

import "github.com/flanksource/kommons"

type Kubernetes struct {
	BaseScraper     `json:",inline"`
	Namespace       string          `json:"namespace,omitempty"`
	UseCache        bool            `json:"useCache,omitempty"`
	AllowIncomplete bool            `json:"allowIncomplete,omitempty"`
	Scope           string          `json:"scope,omitempty"`
	Since           string          `json:"since,omitempty"`
	Selector        string          `json:"selector,omitempty"`
	FieldSelector   string          `json:"fieldSelector,omitempty"`
	MaxInflight     int64           `json:"maxInflight,omitempty"`
	Exclusions      []string        `json:"exclusions,omitempty"`
	Kubeconfig      *kommons.EnvVar `json:"kubeconfig,omitempty"`
}

type KubernetesFile struct {
	BaseScraper `json:",inline"`
	Selector    ResourceSelector `json:"selector,inline"`
	Container   string           `json:"container,omitempty"`
	Files       []PodFile        `json:"files,omitempty"`
}

type PodFile struct {
	Path   []string `json:"path,omitempty"`
	Format string   `json:"format,omitempty"`
}

type ResourceSelector struct {
	Namespace     string `json:"namespace,omitempty"`
	Kind          string `json:"kind,omitempty"`
	Name          string `yaml:"name,omitempty" json:"name,omitempty"`
	LabelSelector string `json:"labelSelector,omitempty" yaml:"labelSelector,omitempty"`
	FieldSelector string `json:"fieldSelector,omitempty" yaml:"fieldSelector,omitempty"`
}

func (r ResourceSelector) IsEmpty() bool {
	return r.Name == "" && r.LabelSelector == "" && r.FieldSelector == ""
}

func (r ResourceSelector) String() string {
	if r.Name != "" {
		return r.Name
	}
	if r.LabelSelector != "" {
		return r.LabelSelector
	}
	return r.FieldSelector
}