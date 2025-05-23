/*
Copyright The Longhorn Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	longhornv1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
)

// VolumeRecurringJobSpecApplyConfiguration represents a declarative configuration of the VolumeRecurringJobSpec type for use
// with apply.
type VolumeRecurringJobSpecApplyConfiguration struct {
	Name        *string                           `json:"name,omitempty"`
	Groups      []string                          `json:"groups,omitempty"`
	Task        *longhornv1beta1.RecurringJobType `json:"task,omitempty"`
	Cron        *string                           `json:"cron,omitempty"`
	Retain      *int                              `json:"retain,omitempty"`
	Concurrency *int                              `json:"concurrency,omitempty"`
	Labels      map[string]string                 `json:"labels,omitempty"`
}

// VolumeRecurringJobSpecApplyConfiguration constructs a declarative configuration of the VolumeRecurringJobSpec type for use with
// apply.
func VolumeRecurringJobSpec() *VolumeRecurringJobSpecApplyConfiguration {
	return &VolumeRecurringJobSpecApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *VolumeRecurringJobSpecApplyConfiguration) WithName(value string) *VolumeRecurringJobSpecApplyConfiguration {
	b.Name = &value
	return b
}

// WithGroups adds the given value to the Groups field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Groups field.
func (b *VolumeRecurringJobSpecApplyConfiguration) WithGroups(values ...string) *VolumeRecurringJobSpecApplyConfiguration {
	for i := range values {
		b.Groups = append(b.Groups, values[i])
	}
	return b
}

// WithTask sets the Task field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Task field is set to the value of the last call.
func (b *VolumeRecurringJobSpecApplyConfiguration) WithTask(value longhornv1beta1.RecurringJobType) *VolumeRecurringJobSpecApplyConfiguration {
	b.Task = &value
	return b
}

// WithCron sets the Cron field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Cron field is set to the value of the last call.
func (b *VolumeRecurringJobSpecApplyConfiguration) WithCron(value string) *VolumeRecurringJobSpecApplyConfiguration {
	b.Cron = &value
	return b
}

// WithRetain sets the Retain field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Retain field is set to the value of the last call.
func (b *VolumeRecurringJobSpecApplyConfiguration) WithRetain(value int) *VolumeRecurringJobSpecApplyConfiguration {
	b.Retain = &value
	return b
}

// WithConcurrency sets the Concurrency field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Concurrency field is set to the value of the last call.
func (b *VolumeRecurringJobSpecApplyConfiguration) WithConcurrency(value int) *VolumeRecurringJobSpecApplyConfiguration {
	b.Concurrency = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *VolumeRecurringJobSpecApplyConfiguration) WithLabels(entries map[string]string) *VolumeRecurringJobSpecApplyConfiguration {
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}
