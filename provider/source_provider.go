package provider

type ProjectData struct {
	Url    string `json:"url"`
	Handle string `json:"handle"`
}

type SourceProvider interface {
	AuthProvider

	// GetProjects return all projects on a source provider owned by a given user
	GetProjectData(string, string) (*ProjectData, error)
}
