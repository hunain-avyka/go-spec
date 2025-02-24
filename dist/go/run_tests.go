package yaml

type Buildspec struct {
	Language             string    `json:"language,omitempty"`
	BuildTool            string    `json:"buildTool,omitempty"`
	Args                 string    `json:"args,omitempty"`
	RunOnlySelectedTests bool      `json:"runOnlySelectedTests,omitempty"`
	PreCommand           string    `json:"preCommand,omitempty"`
	PostCommand          string    `json:"postCommand,omitempty"`
	Reports              []*Report `json:"reports,omitempty"`
	EnableTestSplitting  bool      `json:"enableTestSplitting,omitempty"`
}

type Reports struct {
	Type string `json:"type,omitempty"`
	Spec *Spec  `json:"spec,omitempty"`
}

type Spec struct {
	Paths []string `json:"paths,omitempty"`
}
