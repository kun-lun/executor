package storage

type State struct {
	Version        int    `json:"version"`
	KLVersion      string `json:"klVersion"`
	IAAS           string `json:"iaas"`
	ID             string `json:"id"`
	EnvID          string `json:"envID"`
	TFState        string `json:"tfState"`
	LatestTFOutput string `json:"latestTFOutput"`
}
