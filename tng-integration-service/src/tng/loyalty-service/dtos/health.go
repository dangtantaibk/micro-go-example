package dtos

// GetAppMetaResponse represents response of GetAppMeta.
type GetAppMetaResponse struct {
	Meta Meta           `json:"meta"`
	Data GetAppMetaData `json:"data"`
}

// GetAppMetaData represents application meta data.
type GetAppMetaData struct {
	FrameworkVersion   string `json:"framework_version,omitempty"`
	AppVersionName     string `json:"app_version_name,omitempty"`
	AppVersionCode     string `json:"app_version_code,omitempty"`
	AppBuildDate       string `json:"app_build_date,omitempty"`
	AppBuildCommitHash string `json:"app_build_commit_hash,omitempty"`
}
