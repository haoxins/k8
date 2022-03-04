package k8

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

func AppendVolumeMounts(volumeMounts []corev1.VolumeMount, newVolumeMounts ...corev1.VolumeMount) []corev1.VolumeMount {
	for _, mounts := range newVolumeMounts {
		var conflict = false
		for _, mount := range volumeMounts {
			if mounts.MountPath == mount.MountPath {
				conflict = true
				break
			}
		}
		if !conflict {
			fmt.Println("Skipping append because of conflict")
			volumeMounts = append(volumeMounts, mounts)
		}
	}

	return volumeMounts
}

func AppendEnvVars(envVars []corev1.EnvVar, newEnvVars ...corev1.EnvVar) []corev1.EnvVar {
	for _, envVar := range newEnvVars {
		var conflict = false
		for _, env := range envVars {
			if envVar.Name == env.Name {
				conflict = true
				break
			}
		}
		if !conflict {
			fmt.Println("Skipping append because of conflict")
			envVars = append(envVars, envVar)
		}
	}

	return envVars
}
