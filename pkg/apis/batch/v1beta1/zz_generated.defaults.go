//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by defaulter-gen. DO NOT EDIT.

package v1beta1

import (
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&batchv1beta1.CronJob{}, func(obj interface{}) { SetObjectDefaults_CronJob(obj.(*batchv1beta1.CronJob)) })
	scheme.AddTypeDefaultingFunc(&batchv1beta1.CronJobList{}, func(obj interface{}) { SetObjectDefaults_CronJobList(obj.(*batchv1beta1.CronJobList)) })
	return nil
}

func SetObjectDefaults_CronJob(in *batchv1beta1.CronJob) {
	SetDefaults_CronJob(in)
	corev1.SetDefaults_PodSpec(&in.Spec.JobTemplate.Spec.Template.Spec)
	for i := range in.Spec.JobTemplate.Spec.Template.Spec.Volumes {
		a := &in.Spec.JobTemplate.Spec.Template.Spec.Volumes[i]
		corev1.SetDefaults_Volume(a)
		if a.VolumeSource.HostPath != nil {
			corev1.SetDefaults_HostPathVolumeSource(a.VolumeSource.HostPath)
		}
		if a.VolumeSource.Secret != nil {
			corev1.SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
		}
		if a.VolumeSource.ISCSI != nil {
			if a.VolumeSource.ISCSI.ISCSIInterface == "" {
				a.VolumeSource.ISCSI.ISCSIInterface = "default"
			}
		}
		if a.VolumeSource.RBD != nil {
			if a.VolumeSource.RBD.RBDPool == "" {
				a.VolumeSource.RBD.RBDPool = "rbd"
			}
			if a.VolumeSource.RBD.RadosUser == "" {
				a.VolumeSource.RBD.RadosUser = "admin"
			}
			if a.VolumeSource.RBD.Keyring == "" {
				a.VolumeSource.RBD.Keyring = "/etc/ceph/keyring"
			}
		}
		if a.VolumeSource.DownwardAPI != nil {
			corev1.SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
			for j := range a.VolumeSource.DownwardAPI.Items {
				b := &a.VolumeSource.DownwardAPI.Items[j]
				if b.FieldRef != nil {
					corev1.SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
		if a.VolumeSource.ConfigMap != nil {
			corev1.SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
		}
		if a.VolumeSource.AzureDisk != nil {
			if a.VolumeSource.AzureDisk.CachingMode == nil {
				ptrVar1 := v1.AzureDataDiskCachingMode(v1.AzureDataDiskCachingReadWrite)
				a.VolumeSource.AzureDisk.CachingMode = &ptrVar1
			}
			if a.VolumeSource.AzureDisk.FSType == nil {
				var ptrVar1 string = "ext4"
				a.VolumeSource.AzureDisk.FSType = &ptrVar1
			}
			if a.VolumeSource.AzureDisk.ReadOnly == nil {
				var ptrVar1 bool = false
				a.VolumeSource.AzureDisk.ReadOnly = &ptrVar1
			}
			if a.VolumeSource.AzureDisk.Kind == nil {
				ptrVar1 := v1.AzureDataDiskKind(v1.AzureSharedBlobDisk)
				a.VolumeSource.AzureDisk.Kind = &ptrVar1
			}
		}
		if a.VolumeSource.Projected != nil {
			corev1.SetDefaults_ProjectedVolumeSource(a.VolumeSource.Projected)
			for j := range a.VolumeSource.Projected.Sources {
				b := &a.VolumeSource.Projected.Sources[j]
				if b.DownwardAPI != nil {
					for k := range b.DownwardAPI.Items {
						c := &b.DownwardAPI.Items[k]
						if c.FieldRef != nil {
							corev1.SetDefaults_ObjectFieldSelector(c.FieldRef)
						}
					}
				}
				if b.ServiceAccountToken != nil {
					corev1.SetDefaults_ServiceAccountTokenProjection(b.ServiceAccountToken)
				}
			}
		}
		if a.VolumeSource.ScaleIO != nil {
			if a.VolumeSource.ScaleIO.StorageMode == "" {
				a.VolumeSource.ScaleIO.StorageMode = "ThinProvisioned"
			}
			if a.VolumeSource.ScaleIO.FSType == "" {
				a.VolumeSource.ScaleIO.FSType = "xfs"
			}
		}
		if a.VolumeSource.Ephemeral != nil {
			if a.VolumeSource.Ephemeral.VolumeClaimTemplate != nil {
				corev1.SetDefaults_PersistentVolumeClaimSpec(&a.VolumeSource.Ephemeral.VolumeClaimTemplate.Spec)
				corev1.SetDefaults_ResourceList(&a.VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.Resources.Limits)
				corev1.SetDefaults_ResourceList(&a.VolumeSource.Ephemeral.VolumeClaimTemplate.Spec.Resources.Requests)
			}
		}
	}
	for i := range in.Spec.JobTemplate.Spec.Template.Spec.InitContainers {
		a := &in.Spec.JobTemplate.Spec.Template.Spec.InitContainers[i]
		corev1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					corev1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		corev1.SetDefaults_ResourceList(&a.Resources.Limits)
		corev1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			corev1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.ProbeHandler.HTTPGet != nil {
				corev1.SetDefaults_HTTPGetAction(a.LivenessProbe.ProbeHandler.HTTPGet)
			}
			if a.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.ReadinessProbe != nil {
			corev1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.ProbeHandler.HTTPGet != nil {
				corev1.SetDefaults_HTTPGetAction(a.ReadinessProbe.ProbeHandler.HTTPGet)
			}
			if a.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.StartupProbe != nil {
			corev1.SetDefaults_Probe(a.StartupProbe)
			if a.StartupProbe.ProbeHandler.HTTPGet != nil {
				corev1.SetDefaults_HTTPGetAction(a.StartupProbe.ProbeHandler.HTTPGet)
			}
			if a.StartupProbe.ProbeHandler.GRPC != nil {
				if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					corev1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					corev1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Spec.JobTemplate.Spec.Template.Spec.Containers {
		a := &in.Spec.JobTemplate.Spec.Template.Spec.Containers[i]
		corev1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					corev1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		corev1.SetDefaults_ResourceList(&a.Resources.Limits)
		corev1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			corev1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.ProbeHandler.HTTPGet != nil {
				corev1.SetDefaults_HTTPGetAction(a.LivenessProbe.ProbeHandler.HTTPGet)
			}
			if a.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.ReadinessProbe != nil {
			corev1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.ProbeHandler.HTTPGet != nil {
				corev1.SetDefaults_HTTPGetAction(a.ReadinessProbe.ProbeHandler.HTTPGet)
			}
			if a.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.StartupProbe != nil {
			corev1.SetDefaults_Probe(a.StartupProbe)
			if a.StartupProbe.ProbeHandler.HTTPGet != nil {
				corev1.SetDefaults_HTTPGetAction(a.StartupProbe.ProbeHandler.HTTPGet)
			}
			if a.StartupProbe.ProbeHandler.GRPC != nil {
				if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					corev1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					corev1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Spec.JobTemplate.Spec.Template.Spec.EphemeralContainers {
		a := &in.Spec.JobTemplate.Spec.Template.Spec.EphemeralContainers[i]
		corev1.SetDefaults_EphemeralContainer(a)
		for j := range a.EphemeralContainerCommon.Ports {
			b := &a.EphemeralContainerCommon.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.EphemeralContainerCommon.Env {
			b := &a.EphemeralContainerCommon.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					corev1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		corev1.SetDefaults_ResourceList(&a.EphemeralContainerCommon.Resources.Limits)
		corev1.SetDefaults_ResourceList(&a.EphemeralContainerCommon.Resources.Requests)
		if a.EphemeralContainerCommon.LivenessProbe != nil {
			corev1.SetDefaults_Probe(a.EphemeralContainerCommon.LivenessProbe)
			if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.HTTPGet != nil {
				corev1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.HTTPGet)
			}
			if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.EphemeralContainerCommon.ReadinessProbe != nil {
			corev1.SetDefaults_Probe(a.EphemeralContainerCommon.ReadinessProbe)
			if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.HTTPGet != nil {
				corev1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.HTTPGet)
			}
			if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.EphemeralContainerCommon.StartupProbe != nil {
			corev1.SetDefaults_Probe(a.EphemeralContainerCommon.StartupProbe)
			if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.HTTPGet != nil {
				corev1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.StartupProbe.ProbeHandler.HTTPGet)
			}
			if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.EphemeralContainerCommon.Lifecycle != nil {
			if a.EphemeralContainerCommon.Lifecycle.PostStart != nil {
				if a.EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet != nil {
					corev1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.EphemeralContainerCommon.Lifecycle.PreStop != nil {
				if a.EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet != nil {
					corev1.SetDefaults_HTTPGetAction(a.EphemeralContainerCommon.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	corev1.SetDefaults_ResourceList(&in.Spec.JobTemplate.Spec.Template.Spec.Overhead)
}

func SetObjectDefaults_CronJobList(in *batchv1beta1.CronJobList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_CronJob(a)
	}
}
