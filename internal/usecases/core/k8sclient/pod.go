package k8sclient

import (
	"context"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *k8sClient) GetPods(ctx context.Context, namespace, prefix string) ([]*corev1.Pod, error) {
	pods, err := k.client.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var result []*corev1.Pod
	for _, pod := range pods.Items {
		if prefix != "" && !strings.HasPrefix(pod.Name, prefix) {
			continue
		}

		result = append(result, &pod)
	}

	return result, nil
}

func (k *k8sClient) CreatePod(ctx context.Context, pod *corev1.Pod, opts metav1.CreateOptions) (*corev1.Pod, error) {
	return k.client.CoreV1().Pods(pod.Namespace).Create(ctx, pod, opts)
}

func (k *k8sClient) UpdatePod(ctx context.Context, pod *corev1.Pod, opts metav1.UpdateOptions) (*corev1.Pod, error) {
	return k.client.CoreV1().Pods(pod.Namespace).Update(ctx, pod, opts)
}

func (k *k8sClient) DeletePod(ctx context.Context, namespace, name string, opts metav1.DeleteOptions) error {
	return k.client.CoreV1().Pods(namespace).Delete(ctx, name, opts)
}
