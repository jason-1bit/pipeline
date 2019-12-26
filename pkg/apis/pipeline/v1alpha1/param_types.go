/*
Copyright 2019 The Tekton Authors

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

package v1alpha1

import (
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha2"
)

// ParamSpec defines arbitrary parameters needed beyond typed inputs (such as
// resources). Parameter values are provided by users as inputs on a TaskRun
// or PipelineRun.
type ParamSpec = v1alpha2.ParamSpec

// ResourceParam declares a string value to use for the parameter called Name, and is used in
// the specific context of PipelineResources.
type ResourceParam = v1alpha2.ResourceParam

// Param declares an ArrayOrString to use for the parameter called name.
type Param = v1alpha2.Param

// ParamType indicates the type of an input parameter;
// Used to distinguish between a single string and an array of strings.
type ParamType = v1alpha2.ParamType

// Valid ParamTypes:
const (
	ParamTypeString ParamType = v1alpha2.ParamTypeString
	ParamTypeArray  ParamType = v1alpha2.ParamTypeArray
)

// AllParamTypes can be used for ParamType validation.
var AllParamTypes = v1alpha2.AllParamTypes

// ArrayOrString is modeled after IntOrString in kubernetes/apimachinery:

// ArrayOrString is a type that can hold a single string or string array.
// Used in JSON unmarshalling so that a single JSON field can accept
// either an individual string or an array of strings.
type ArrayOrString = v1alpha2.ArrayOrString
