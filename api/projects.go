package api

import (
	"github.com/hayashiki/mcfly-clone/api/apidata"
	"github.com/hayashiki/mcfly-clone/models"
)

func (handlers *Handlers) PostProject(r *Responder, ctx *RequestContext) {
	reqData := ctx.RequestData.(*apidata.ProjectReq)

	sourceProvider := *ctx.SourceProvider
	projectData, err := sourceProvider.GetProjectData("token", reqData.Handle)
	if err != nil {
		return
	}

	project := models.Project{
		Handle:         reqData.Handle,
		SourceUrl:      projectData.Url,
		SourceProvider: reqData.Provider,
	}

	r.Respond(&apidata.ProjectResp{
		project.Handle,
		project.SourceUrl,
		project.SourceProvider,
	})
}

func (handlers *Handlers) GetProjects(r *Responder, ctx *RequestContext) {
	projects, err := handlers.DB.GetUserProjects(ctx.CurrentUser)
	if err != nil {
		r.RespondWithServerError(err)
	}
	projectsResp := make([]apidata.ProjectResp, len(projects))
	for i, p := range projects {
		projectsResp[i] = apidata.ProjectResp{p.Handle, p.SourceUrl, p.SourceProvider}
	}
	r.Respond(projectsResp)
	// r.RespondWithSuccess()
}
