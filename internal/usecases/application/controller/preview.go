package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/aura-cd/backend/cmd/config"
	"github.com/aura-cd/backend/internal/models"
	"github.com/aura-cd/backend/internal/usecases/application/github"
	v1 "github.com/aura-cd/backend/proto"
)

// CreatePreviewEnvironmentParams はプレビュー環境を作成するためのパラメータです。
type CreatePreviewEnvironmentParams struct {
	Organization string
	Repository   string
	PrNumber     int
	Branch       string
	GitLink      string

	GhInstallID int64
	GhClient    github.GitHubClientInteractor
	Config      models.PullRequestConfig
}

func (ge *githubAppEvents) CreatePreviewEnvironment(ctx context.Context, param CreatePreviewEnvironmentParams) error {
	meta := github.MetaData{
		Organization: param.Organization,
		Repository:   param.Repository,
		Number:       param.PrNumber,
	}
	token, err := param.GhClient.GetToken(ctx, param.GhInstallID)
	if err != nil {
		_, _, err := param.GhClient.CreateReview(ctx, meta, fmt.Sprintf("[%v] ❌Failed to get token: %v", time.Now().Format(time.DateTime), err))
		return err
	}
	// TODO: image buildする
	image, err := ge.BuildImage(ctx, BuildImageParams{
		Organization:  param.Organization,
		Repository:    param.Repository,
		Branch:        param.Branch,
		PullRequestID: fmt.Sprintf("%d", param.PrNumber),
		GitLink:       param.GitLink,
		Config:        param.Config,
		Token:         token,
	})
	if err != nil {
		_, _, err := param.GhClient.CreateReview(ctx, meta, fmt.Sprintf("[%v] ❌Failed to build Docker image: %v", time.Now().Format(time.DateTime), err))
		return err
	}

	var configs []*v1.Config
	for _, c := range param.Config.ConfigMaps {
		configs = append(configs, &v1.Config{
			Name:  c.Name,
			Value: c.Value,
		})
	}

	if param.Config.Replicas == 0 {
		param.Config.Replicas = 1
	}

	// デプロイする
	dep, err := ge.customController.CreatePreview(ctx, &v1.CreatePreviewRequest{
		Organization:  param.Organization,
		Repository:    param.Repository,
		PullRequestId: fmt.Sprintf("%d", param.PrNumber),
		Branch:        param.Branch,
		Image:         *image,
		Replicas:      param.Config.Replicas,
		Configs:       configs,
		ExposePorts:   param.Config.ExposePort,
	})

	if err != nil {
		_, _, err = param.GhClient.CreateReview(ctx, meta, fmt.Sprintf("[%v] ❌Failed to create deployment and service: %v", time.Now().Format(time.DateTime), err))
		return err
	}

	var external string
	for _, p := range dep.GetExternal() {
		external += fmt.Sprintf("- https://%s\n", p)
	}

	_, _, err = param.GhClient.CreateReview(ctx, meta, fmt.Sprintf("[%v] ✔️Deployment and service created\n Exposing service on port ...\n %s", time.Now().Format(time.DateTime), external))
	if err != nil {
		return err
	}

	return nil
}

type DeletePreviewEnvironmentParams struct {
	Organization string
	Repository   string
	PrNumber     string
	Branch       string
}

func (ge *githubAppEvents) DeletePreviewEnvironment(ctx context.Context, param DeletePreviewEnvironmentParams) error {
	_, err := ge.customController.DeletePreview(ctx, &v1.DeletePreviewRequest{
		Organization:  param.Organization,
		Repository:    param.Repository,
		PullRequestId: param.PrNumber,
		Branch:        param.Branch,
	})

	return err
}

type UpdatePreviewEnvironmentParams struct {
	Organization string
	Repository   string
	PrNumber     int
	Branch       string
	GitLink      string
	GhInstallID  int64

	Config   models.PullRequestConfig
	GhClient github.GitHubClientInteractor
}

func (ge *githubAppEvents) UpdatePreviewEnvironment(ctx context.Context, param UpdatePreviewEnvironmentParams) error {
	meta := github.MetaData{
		Organization: param.Organization,
		Repository:   param.Repository,
		Number:       param.PrNumber,
	}

	token, err := param.GhClient.GetToken(ctx, param.GhInstallID)
	if err != nil {
		_, _, err := param.GhClient.CreateReview(ctx, meta, fmt.Sprintf("[%v] ❌Failed to get token: %v", time.Now().Format(time.DateTime), err))
		return err
	}

	image, err := ge.BuildImage(ctx, BuildImageParams{
		Organization:  param.Organization,
		Repository:    param.Repository,
		Branch:        param.Branch,
		PullRequestID: fmt.Sprintf("%d", param.PrNumber),
		GitLink:       param.GitLink,
		Config:        param.Config,
		Token:         token,
	})
	if err != nil {
		_, _, err := param.GhClient.CreateReview(ctx, meta, fmt.Sprintf("[%v] ❌Failed to build Docker image: %v", time.Now().Format(time.DateTime), err))
		return err
	}

	// デプロイする
	dep, err := ge.customController.UpdatePreview(ctx, &v1.CreatePreviewRequest{
		Organization:  param.Organization,
		Repository:    param.Repository,
		PullRequestId: fmt.Sprintf("%d", param.PrNumber),
		Branch:        param.Branch,
		Image:         *image,
	})

	if err != nil {
		_, _, err = param.GhClient.CreateReview(ctx, meta, fmt.Sprintf("[%v] ❌Failed to create deployment and service: %v", time.Now().Format(time.DateTime), err))
		return err
	}

	if dep.GetExternal() == nil {
		_, _, err = param.GhClient.CreateReview(ctx, meta, fmt.Sprintf("[%v] ✔️Deployment update successful ", time.Now().Format(time.DateTime)))
		if err != nil {
			return err
		}
		return nil
	}

	var external string
	for _, p := range dep.GetExternal() {
		external += fmt.Sprintf("- %s\n", p)
	}

	_, _, err = param.GhClient.CreateReview(ctx, meta, fmt.Sprintf("[%v] ✔️Deployment and service created\n Exposing service on port ...\n %s", time.Now().Format(time.DateTime), external))
	if err != nil {
		return err
	}

	return nil
}

type BuildImageParams struct {
	Organization  string
	Repository    string
	Branch        string
	PullRequestID string
	GitLink       string
	Token         string

	Config models.PullRequestConfig
}

func (ge *githubAppEvents) BuildImage(ctx context.Context, param BuildImageParams) (*string, error) {
	result, err := ge.customController.BuildImage(ctx, &v1.BuildImageRequest{
		Namespace:      param.Organization,
		Repository:     param.Repository,
		Branch:         param.Branch,
		PullRequestId:  param.PullRequestID,
		GitRepo:        param.GitLink,
		DockerfilePath: param.Config.BuildFilePath,
		DockerfileDir:  param.Config.BuildFileDir,
		ImageName:      fmt.Sprintf("%s/%s", config.Config.Registry.Host, config.Config.Application.ApplicationName),
		Token:          param.Token,
	})
	if err != nil {
		return nil, err
	}

	return &result.Image, nil

}
