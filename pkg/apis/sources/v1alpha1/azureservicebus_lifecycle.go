/*
Copyright 2023 TriggerMesh Inc.

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
	"context"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/triggermesh/triggermesh/pkg/apis/common/v1alpha1"
	"github.com/triggermesh/triggermesh/pkg/apis/sources"
)

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (s *AzureServiceBusSource) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("AzureServiceBusSource")
}

// GetConditionSet implements duckv1.KRShaped.
func (s *AzureServiceBusSource) GetConditionSet() apis.ConditionSet {
	return azureServiceBusSourceConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (s *AzureServiceBusSource) GetStatus() *duckv1.Status {
	return &s.Status.Status.Status
}

// GetSink implements EventSender.
func (s *AzureServiceBusSource) GetSink() *duckv1.Destination {
	return &s.Spec.Sink
}

// GetStatusManager implements Reconcilable.
func (s *AzureServiceBusSource) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: s.GetConditionSet(),
		Status:       &s.Status.Status,
	}
}

// AsEventSource implements EventSource.
func (s *AzureServiceBusSource) AsEventSource() string {
	if s.Spec.TopicID != nil {
		return s.Spec.TopicID.String()
	} else if s.Spec.QueueID != nil {
		return s.Spec.QueueID.String()
	}
	return ""
}

// GetEventTypes returns the event types generated by the source.
func (*AzureServiceBusSource) GetEventTypes() []string {
	return []string{
		AzureEventType(sources.AzureServiceServiceBus, AzureServiceBusGenericEventType),
	}
}

// GetAdapterOverrides implements AdapterConfigurable.
func (s *AzureServiceBusSource) GetAdapterOverrides() *v1alpha1.AdapterOverrides {
	return s.Spec.AdapterOverrides
}

// Status conditions
const (
	// AzureServiceBusConditionSubscribed has status True when the source has subscribed to a topic or queue.
	AzureServiceBusConditionSubscribed apis.ConditionType = "Subscribed"
)

// azureServiceBusSourceConditionSet is a set of conditions for
// AzureServiceBusSource objects.
var azureServiceBusSourceConditionSet = v1alpha1.NewConditionSet(
	AzureServiceBusConditionSubscribed,
)

// MarkSubscribed sets the Subscribed condition to True.
func (s *AzureServiceBusSourceStatus) MarkSubscribed() {
	azureServiceBusSourceConditionSet.Manage(s).MarkTrue(AzureServiceBusConditionSubscribed)
}

// MarkSubscribedWithReason sets the Subscribed condition to True with reason.
func (s *AzureServiceBusSourceStatus) MarkSubscribedWithReason(reason, msg string) {
	azureServiceBusSourceConditionSet.Manage(s).MarkTrueWithReason(AzureServiceBusConditionSubscribed, reason, msg)
}

// MarkNotSubscribed sets the Subscribed condition to False with the given
// reason and message.
func (s *AzureServiceBusSourceStatus) MarkNotSubscribed(reason, msg string) {
	azureServiceBusSourceConditionSet.Manage(s).MarkFalse(AzureServiceBusConditionSubscribed, reason, msg)
}

// SetDefaults implements apis.Defaultable
func (s *AzureServiceBusSource) SetDefaults(ctx context.Context) {
}

// Validate implements apis.Validatable
func (s *AzureServiceBusSource) Validate(ctx context.Context) *apis.FieldError {
	return nil
}
