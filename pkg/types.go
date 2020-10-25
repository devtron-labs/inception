/*
Copyright 2020 Devtron Labs Pvt Ltd.

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

package pkg

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ManifestResponse struct {
	Manifest string `protobuf:"bytes,1,req,name=manifest" json:"manifest,omitempty"`
}

type GetRequest struct {
	Name             string                  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Namespace        string                  `protobuf:"bytes,2,req,name=namespace" json:"namespace,omitempty"`
	GroupVersionKind schema.GroupVersionKind `protobuf:"bytes,3,req,name=groupVersionKind" json:"groupVersionKind,omitempty"`
}

type DeleteRequest struct {
	Name             string                  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Namespace        string                  `protobuf:"bytes,2,req,name=namespace" json:"namespace,omitempty"`
	GroupVersionKind schema.GroupVersionKind `protobuf:"bytes,3,req,name=groupVersionKind" json:"groupVersionKind,omitempty"`
	Force            *bool                   `protobuf:"bytes,4,req,name=force" json:"force,omitempty"`
}

type PatchRequest struct {
	Name             string                  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Namespace        string                  `protobuf:"bytes,2,req,name=namespace" json:"namespace,omitempty"`
	GroupVersionKind schema.GroupVersionKind `protobuf:"bytes,3,req,name=groupVersionKind" json:"groupVersionKind,omitempty"`
	Patch            string                  `protobuf:"bytes,4,req,name=patch" json:"patch,omitempty"`
	PatchType        string                  `protobuf:"bytes,5,req,name=patchType" json:"patchType,omitempty"`
}

type ApplyRequest struct {
	Manifest  string `protobuf:"bytes,1,req,name=manifest" json:"manifest,omitempty"`
	Namespace string `protobuf:"bytes,2,req,name=namespace" json:"namespace,omitempty"`
	Force     *bool  `protobuf:"bytes,3,req,name=force" json:"force,omitempty"`
	Validate  *bool  `protobuf:"bytes,4,req,name=validate" json:"validate,omitempty"`
}

type ApplyResponse struct {
	Message          string                  `protobuf:"bytes,1,req,name=message" json:"message,omitempty"`
	Err              string                  `protobuf:"bytes,2,req,name=err" json:"err,omitempty"`
	Name             string                  `protobuf:"bytes,3,req,name=name" json:"name,omitempty"`
	Namespace        string                  `protobuf:"bytes,4,req,name=namespace" json:"namespace,omitempty"`
	GroupVersionKind schema.GroupVersionKind `protobuf:"bytes,5,req,name=groupVersionKind" json:"groupVersionKind,omitempty"`
}
