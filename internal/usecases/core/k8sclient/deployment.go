package k8sclient

import (
	"context"
	"fmt"

	coreErr "github.com/gantrycd/backend/internal/error"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type CreateDeploymentParams struct {
	Namespace string
	AppName   string
	Image     string
}

func (k *k8sClient) CreateDeployment(ctx context.Context, in CreateDeploymentParams, opts ...Option) (*appsv1.Deployment, error) {
	o := newOption()
	for _, opt := range opts {
		opt(o)
	}

	k.l.Info("creating deployment", "namespace", in.Namespace, "appName", in.AppName, "image", in.Image, CreatedByLabel, o.labelSelector[CreatedByLabel])

	o.labelSelector[AppLabel] = in.AppName

	return k.client.AppsV1().Deployments(in.Namespace).Create(ctx, &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: fmt.Sprintf("%s-", in.AppName),
			Labels:       o.labelSelector,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: o.replica,
			Selector: &metav1.LabelSelector{
				MatchLabels: o.labelSelector,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: o.labelSelector,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            in.AppName,
							Image:           in.Image,
							ImagePullPolicy: o.containerOption[in.Image].imagePullPolicy,
						},
					},
				},
			},
		},
	}, metav1.CreateOptions{})
}

type GetDeploymentParams struct {
	Namespace     string
	Repository    string
	PullRequestID string
	Branch        string
}

func (k *k8sClient) GetDeployment(ctx context.Context, param GetDeploymentParams) (*appsv1.Deployment, error) {
	deps, err := k.client.AppsV1().Deployments(param.Namespace).List(ctx, metav1.ListOptions{
		LabelSelector: labels.Set(map[string]string{
			RepositoryLabel: param.Repository,
			PullRequestID:   param.PullRequestID,
			BaseBranchLabel: param.Branch,
		}).String(),
	})

	for _, dep := range deps.Items {
		if dep.Labels[RepositoryLabel] == param.Repository && dep.Labels[PullRequestID] == param.PullRequestID && dep.Labels[BaseBranchLabel] == param.Branch {
			return &dep, nil
		}
	}

	if err != nil {
		return nil, err
	}

	return nil, coreErr.ErrDeploymentsNotFound
}

func (k *k8sClient) ListDeployments(ctx context.Context, namespace string, opts ...Option) (*appsv1.DeploymentList, error) {
	o := newOption()
	for _, opt := range opts {
		opt(o)
	}

	return k.client.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{LabelSelector: labels.Set(o.labelSelector).String()})
}

func (k *k8sClient) DeleteDeployment(ctx context.Context, namespace string, opt ...Option) error {
	o := newOption()

	for _, opt := range opt {
		opt(o)
	}

	deps, err := k.client.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: labels.Set(o.labelSelector).String(),
	})

	if err != nil {
		return err
	}

	for _, dep := range deps.Items {
		err := k.client.AppsV1().Deployments(namespace).Delete(ctx, dep.Name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}