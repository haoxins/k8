package k8

import corev1 "k8s.io/api/core/v1"

// Attach any non-duplicate volume mounts and env vars to each specified container
func AttachContainer(container corev1.Container, volumeMounts []corev1.VolumeMount, envVars []corev1.EnvVar) corev1.Container {
	container.VolumeMounts = AppendVolumeMounts(container.VolumeMounts, volumeMounts...)
	container.Env = AppendEnvVars(container.Env, envVars...)

	return container
}

// Attach any non-duplicate volume mounts and env vars to the specified containers
func AttachContainers(containers []corev1.Container, volumeMounts []corev1.VolumeMount, envVars []corev1.EnvVar) []corev1.Container {
	var updatedContainers = []corev1.Container{}
	for _, container := range containers {
		updatedContainer := AttachContainer(container, volumeMounts, envVars)
		updatedContainers = append(updatedContainers, updatedContainer)
	}
	return updatedContainers
}
