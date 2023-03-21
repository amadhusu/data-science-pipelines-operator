//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServer) DeepCopyInto(out *APIServer) {
	*out = *in
	if in.ArtifactScriptConfigMap != nil {
		in, out := &in.ArtifactScriptConfigMap, &out.ArtifactScriptConfigMap
		*out = new(ArtifactScriptConfigMap)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServer.
func (in *APIServer) DeepCopy() *APIServer {
	if in == nil {
		return nil
	}
	out := new(APIServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArtifactScriptConfigMap) DeepCopyInto(out *ArtifactScriptConfigMap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArtifactScriptConfigMap.
func (in *ArtifactScriptConfigMap) DeepCopy() *ArtifactScriptConfigMap {
	if in == nil {
		return nil
	}
	out := new(ArtifactScriptConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DSPASpec) DeepCopyInto(out *DSPASpec) {
	*out = *in
	if in.APIServer != nil {
		in, out := &in.APIServer, &out.APIServer
		*out = new(APIServer)
		(*in).DeepCopyInto(*out)
	}
	if in.PersistenceAgent != nil {
		in, out := &in.PersistenceAgent, &out.PersistenceAgent
		*out = new(PersistenceAgent)
		(*in).DeepCopyInto(*out)
	}
	if in.ScheduledWorkflow != nil {
		in, out := &in.ScheduledWorkflow, &out.ScheduledWorkflow
		*out = new(ScheduledWorkflow)
		(*in).DeepCopyInto(*out)
	}
	if in.ViewerCRD != nil {
		in, out := &in.ViewerCRD, &out.ViewerCRD
		*out = new(ViewerCRD)
		(*in).DeepCopyInto(*out)
	}
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(Database)
		(*in).DeepCopyInto(*out)
	}
	if in.MlPipelineUI != nil {
		in, out := &in.MlPipelineUI, &out.MlPipelineUI
		*out = new(MlPipelineUI)
		(*in).DeepCopyInto(*out)
	}
	if in.ObjectStorage != nil {
		in, out := &in.ObjectStorage, &out.ObjectStorage
		*out = new(ObjectStorage)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DSPASpec.
func (in *DSPASpec) DeepCopy() *DSPASpec {
	if in == nil {
		return nil
	}
	out := new(DSPASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DSPAStatus) DeepCopyInto(out *DSPAStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DSPAStatus.
func (in *DSPAStatus) DeepCopy() *DSPAStatus {
	if in == nil {
		return nil
	}
	out := new(DSPAStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataSciencePipelinesApplication) DeepCopyInto(out *DataSciencePipelinesApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataSciencePipelinesApplication.
func (in *DataSciencePipelinesApplication) DeepCopy() *DataSciencePipelinesApplication {
	if in == nil {
		return nil
	}
	out := new(DataSciencePipelinesApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataSciencePipelinesApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataSciencePipelinesApplicationList) DeepCopyInto(out *DataSciencePipelinesApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataSciencePipelinesApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataSciencePipelinesApplicationList.
func (in *DataSciencePipelinesApplicationList) DeepCopy() *DataSciencePipelinesApplicationList {
	if in == nil {
		return nil
	}
	out := new(DataSciencePipelinesApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataSciencePipelinesApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database) DeepCopyInto(out *Database) {
	*out = *in
	if in.MariaDB != nil {
		in, out := &in.MariaDB, &out.MariaDB
		*out = new(MariaDB)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalDB != nil {
		in, out := &in.ExternalDB, &out.ExternalDB
		*out = new(ExternalDB)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database.
func (in *Database) DeepCopy() *Database {
	if in == nil {
		return nil
	}
	out := new(Database)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDB) DeepCopyInto(out *ExternalDB) {
	*out = *in
	if in.PasswordSecret != nil {
		in, out := &in.PasswordSecret, &out.PasswordSecret
		*out = new(SecretKeyValue)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDB.
func (in *ExternalDB) DeepCopy() *ExternalDB {
	if in == nil {
		return nil
	}
	out := new(ExternalDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalStorage) DeepCopyInto(out *ExternalStorage) {
	*out = *in
	if in.S3CredentialSecret != nil {
		in, out := &in.S3CredentialSecret, &out.S3CredentialSecret
		*out = new(S3CredentialSecret)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalStorage.
func (in *ExternalStorage) DeepCopy() *ExternalStorage {
	if in == nil {
		return nil
	}
	out := new(ExternalStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MariaDB) DeepCopyInto(out *MariaDB) {
	*out = *in
	if in.PasswordSecret != nil {
		in, out := &in.PasswordSecret, &out.PasswordSecret
		*out = new(SecretKeyValue)
		**out = **in
	}
	out.PVCSize = in.PVCSize.DeepCopy()
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MariaDB.
func (in *MariaDB) DeepCopy() *MariaDB {
	if in == nil {
		return nil
	}
	out := new(MariaDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Minio) DeepCopyInto(out *Minio) {
	*out = *in
	if in.S3CredentialSecret != nil {
		in, out := &in.S3CredentialSecret, &out.S3CredentialSecret
		*out = new(S3CredentialSecret)
		**out = **in
	}
	out.PVCSize = in.PVCSize.DeepCopy()
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Minio.
func (in *Minio) DeepCopy() *Minio {
	if in == nil {
		return nil
	}
	out := new(Minio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MlPipelineUI) DeepCopyInto(out *MlPipelineUI) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MlPipelineUI.
func (in *MlPipelineUI) DeepCopy() *MlPipelineUI {
	if in == nil {
		return nil
	}
	out := new(MlPipelineUI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorage) DeepCopyInto(out *ObjectStorage) {
	*out = *in
	if in.Minio != nil {
		in, out := &in.Minio, &out.Minio
		*out = new(Minio)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalStorage != nil {
		in, out := &in.ExternalStorage, &out.ExternalStorage
		*out = new(ExternalStorage)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorage.
func (in *ObjectStorage) DeepCopy() *ObjectStorage {
	if in == nil {
		return nil
	}
	out := new(ObjectStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceAgent) DeepCopyInto(out *PersistenceAgent) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceAgent.
func (in *PersistenceAgent) DeepCopy() *PersistenceAgent {
	if in == nil {
		return nil
	}
	out := new(PersistenceAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequirements) DeepCopyInto(out *ResourceRequirements) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = new(Resources)
		(*in).DeepCopyInto(*out)
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = new(Resources)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequirements.
func (in *ResourceRequirements) DeepCopy() *ResourceRequirements {
	if in == nil {
		return nil
	}
	out := new(ResourceRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	out.CPU = in.CPU.DeepCopy()
	out.Memory = in.Memory.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3CredentialSecret) DeepCopyInto(out *S3CredentialSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3CredentialSecret.
func (in *S3CredentialSecret) DeepCopy() *S3CredentialSecret {
	if in == nil {
		return nil
	}
	out := new(S3CredentialSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledWorkflow) DeepCopyInto(out *ScheduledWorkflow) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledWorkflow.
func (in *ScheduledWorkflow) DeepCopy() *ScheduledWorkflow {
	if in == nil {
		return nil
	}
	out := new(ScheduledWorkflow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyValue) DeepCopyInto(out *SecretKeyValue) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyValue.
func (in *SecretKeyValue) DeepCopy() *SecretKeyValue {
	if in == nil {
		return nil
	}
	out := new(SecretKeyValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ViewerCRD) DeepCopyInto(out *ViewerCRD) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ViewerCRD.
func (in *ViewerCRD) DeepCopy() *ViewerCRD {
	if in == nil {
		return nil
	}
	out := new(ViewerCRD)
	in.DeepCopyInto(out)
	return out
}
