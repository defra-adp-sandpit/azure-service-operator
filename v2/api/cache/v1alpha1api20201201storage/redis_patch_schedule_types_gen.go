// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201storage

import (
	"fmt"
	v20201201s "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20201201storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1alpha1api20201201.RedisPatchSchedule
// Deprecated version of RedisPatchSchedule. Use v1beta20201201.RedisPatchSchedule instead
type RedisPatchSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisPatchSchedules_Spec  `json:"spec,omitempty"`
	Status            RedisPatchSchedule_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisPatchSchedule{}

// GetConditions returns the conditions of the resource
func (schedule *RedisPatchSchedule) GetConditions() conditions.Conditions {
	return schedule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (schedule *RedisPatchSchedule) SetConditions(conditions conditions.Conditions) {
	schedule.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisPatchSchedule{}

// ConvertFrom populates our RedisPatchSchedule from the provided hub RedisPatchSchedule
func (schedule *RedisPatchSchedule) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20201201s.RedisPatchSchedule)
	if !ok {
		return fmt.Errorf("expected cache/v1beta20201201storage/RedisPatchSchedule but received %T instead", hub)
	}

	return schedule.AssignPropertiesFromRedisPatchSchedule(source)
}

// ConvertTo populates the provided hub RedisPatchSchedule from our RedisPatchSchedule
func (schedule *RedisPatchSchedule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20201201s.RedisPatchSchedule)
	if !ok {
		return fmt.Errorf("expected cache/v1beta20201201storage/RedisPatchSchedule but received %T instead", hub)
	}

	return schedule.AssignPropertiesToRedisPatchSchedule(destination)
}

var _ genruntime.KubernetesResource = &RedisPatchSchedule{}

// AzureName returns the Azure name of the resource (always "default")
func (schedule *RedisPatchSchedule) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (schedule RedisPatchSchedule) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (schedule *RedisPatchSchedule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (schedule *RedisPatchSchedule) GetSpec() genruntime.ConvertibleSpec {
	return &schedule.Spec
}

// GetStatus returns the status of this resource
func (schedule *RedisPatchSchedule) GetStatus() genruntime.ConvertibleStatus {
	return &schedule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/patchSchedules"
func (schedule *RedisPatchSchedule) GetType() string {
	return "Microsoft.Cache/redis/patchSchedules"
}

// NewEmptyStatus returns a new empty (blank) status
func (schedule *RedisPatchSchedule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisPatchSchedule_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (schedule *RedisPatchSchedule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(schedule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  schedule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (schedule *RedisPatchSchedule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RedisPatchSchedule_Status); ok {
		schedule.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisPatchSchedule_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	schedule.Status = st
	return nil
}

// AssignPropertiesFromRedisPatchSchedule populates our RedisPatchSchedule from the provided source RedisPatchSchedule
func (schedule *RedisPatchSchedule) AssignPropertiesFromRedisPatchSchedule(source *v20201201s.RedisPatchSchedule) error {

	// ObjectMeta
	schedule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RedisPatchSchedules_Spec
	err := spec.AssignPropertiesFromRedisPatchSchedulesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisPatchSchedulesSpec() to populate field Spec")
	}
	schedule.Spec = spec

	// Status
	var status RedisPatchSchedule_Status
	err = status.AssignPropertiesFromRedisPatchScheduleStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisPatchScheduleStatus() to populate field Status")
	}
	schedule.Status = status

	// No error
	return nil
}

// AssignPropertiesToRedisPatchSchedule populates the provided destination RedisPatchSchedule from our RedisPatchSchedule
func (schedule *RedisPatchSchedule) AssignPropertiesToRedisPatchSchedule(destination *v20201201s.RedisPatchSchedule) error {

	// ObjectMeta
	destination.ObjectMeta = *schedule.ObjectMeta.DeepCopy()

	// Spec
	var spec v20201201s.RedisPatchSchedules_Spec
	err := schedule.Spec.AssignPropertiesToRedisPatchSchedulesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisPatchSchedulesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20201201s.RedisPatchSchedule_Status
	err = schedule.Status.AssignPropertiesToRedisPatchScheduleStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisPatchScheduleStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (schedule *RedisPatchSchedule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: schedule.Spec.OriginalVersion,
		Kind:    "RedisPatchSchedule",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20201201.RedisPatchSchedule
// Deprecated version of RedisPatchSchedule. Use v1beta20201201.RedisPatchSchedule instead
type RedisPatchScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisPatchSchedule `json:"items"`
}

// Storage version of v1alpha1api20201201.RedisPatchSchedule_Status
// Deprecated version of RedisPatchSchedule_Status. Use v1beta20201201.RedisPatchSchedule_Status instead
type RedisPatchSchedule_Status struct {
	Conditions      []conditions.Condition `json:"conditions,omitempty"`
	Id              *string                `json:"id,omitempty"`
	Location        *string                `json:"location,omitempty"`
	Name            *string                `json:"name,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ScheduleEntries []ScheduleEntry_Status `json:"scheduleEntries,omitempty"`
	Type            *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisPatchSchedule_Status{}

// ConvertStatusFrom populates our RedisPatchSchedule_Status from the provided source
func (schedule *RedisPatchSchedule_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20201201s.RedisPatchSchedule_Status)
	if ok {
		// Populate our instance from source
		return schedule.AssignPropertiesFromRedisPatchScheduleStatus(src)
	}

	// Convert to an intermediate form
	src = &v20201201s.RedisPatchSchedule_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = schedule.AssignPropertiesFromRedisPatchScheduleStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RedisPatchSchedule_Status
func (schedule *RedisPatchSchedule_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20201201s.RedisPatchSchedule_Status)
	if ok {
		// Populate destination from our instance
		return schedule.AssignPropertiesToRedisPatchScheduleStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v20201201s.RedisPatchSchedule_Status{}
	err := schedule.AssignPropertiesToRedisPatchScheduleStatus(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignPropertiesFromRedisPatchScheduleStatus populates our RedisPatchSchedule_Status from the provided source RedisPatchSchedule_Status
func (schedule *RedisPatchSchedule_Status) AssignPropertiesFromRedisPatchScheduleStatus(source *v20201201s.RedisPatchSchedule_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	schedule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	schedule.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	schedule.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	schedule.Name = genruntime.ClonePointerToString(source.Name)

	// ScheduleEntries
	if source.ScheduleEntries != nil {
		scheduleEntryList := make([]ScheduleEntry_Status, len(source.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range source.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry ScheduleEntry_Status
			err := scheduleEntry.AssignPropertiesFromScheduleEntryStatus(&scheduleEntryItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromScheduleEntryStatus() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		schedule.ScheduleEntries = scheduleEntryList
	} else {
		schedule.ScheduleEntries = nil
	}

	// Type
	schedule.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		schedule.PropertyBag = propertyBag
	} else {
		schedule.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToRedisPatchScheduleStatus populates the provided destination RedisPatchSchedule_Status from our RedisPatchSchedule_Status
func (schedule *RedisPatchSchedule_Status) AssignPropertiesToRedisPatchScheduleStatus(destination *v20201201s.RedisPatchSchedule_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(schedule.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(schedule.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(schedule.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(schedule.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(schedule.Name)

	// ScheduleEntries
	if schedule.ScheduleEntries != nil {
		scheduleEntryList := make([]v20201201s.ScheduleEntry_Status, len(schedule.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range schedule.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry v20201201s.ScheduleEntry_Status
			err := scheduleEntryItem.AssignPropertiesToScheduleEntryStatus(&scheduleEntry)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToScheduleEntryStatus() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		destination.ScheduleEntries = scheduleEntryList
	} else {
		destination.ScheduleEntries = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(schedule.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20201201.RedisPatchSchedules_Spec
type RedisPatchSchedules_Spec struct {
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cache.azure.com/Redis resource
	Owner           *genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner,omitempty" kind:"Redis"`
	PropertyBag     genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ScheduleEntries []ScheduleEntry                    `json:"scheduleEntries,omitempty"`
	Tags            map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RedisPatchSchedules_Spec{}

// ConvertSpecFrom populates our RedisPatchSchedules_Spec from the provided source
func (schedules *RedisPatchSchedules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20201201s.RedisPatchSchedules_Spec)
	if ok {
		// Populate our instance from source
		return schedules.AssignPropertiesFromRedisPatchSchedulesSpec(src)
	}

	// Convert to an intermediate form
	src = &v20201201s.RedisPatchSchedules_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = schedules.AssignPropertiesFromRedisPatchSchedulesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RedisPatchSchedules_Spec
func (schedules *RedisPatchSchedules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20201201s.RedisPatchSchedules_Spec)
	if ok {
		// Populate destination from our instance
		return schedules.AssignPropertiesToRedisPatchSchedulesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20201201s.RedisPatchSchedules_Spec{}
	err := schedules.AssignPropertiesToRedisPatchSchedulesSpec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignPropertiesFromRedisPatchSchedulesSpec populates our RedisPatchSchedules_Spec from the provided source RedisPatchSchedules_Spec
func (schedules *RedisPatchSchedules_Spec) AssignPropertiesFromRedisPatchSchedulesSpec(source *v20201201s.RedisPatchSchedules_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Location
	schedules.Location = genruntime.ClonePointerToString(source.Location)

	// OriginalVersion
	schedules.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		schedules.Owner = &owner
	} else {
		schedules.Owner = nil
	}

	// ScheduleEntries
	if source.ScheduleEntries != nil {
		scheduleEntryList := make([]ScheduleEntry, len(source.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range source.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry ScheduleEntry
			err := scheduleEntry.AssignPropertiesFromScheduleEntry(&scheduleEntryItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromScheduleEntry() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		schedules.ScheduleEntries = scheduleEntryList
	} else {
		schedules.ScheduleEntries = nil
	}

	// Tags
	schedules.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		schedules.PropertyBag = propertyBag
	} else {
		schedules.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToRedisPatchSchedulesSpec populates the provided destination RedisPatchSchedules_Spec from our RedisPatchSchedules_Spec
func (schedules *RedisPatchSchedules_Spec) AssignPropertiesToRedisPatchSchedulesSpec(destination *v20201201s.RedisPatchSchedules_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(schedules.PropertyBag)

	// Location
	destination.Location = genruntime.ClonePointerToString(schedules.Location)

	// OriginalVersion
	destination.OriginalVersion = schedules.OriginalVersion

	// Owner
	if schedules.Owner != nil {
		owner := schedules.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// ScheduleEntries
	if schedules.ScheduleEntries != nil {
		scheduleEntryList := make([]v20201201s.ScheduleEntry, len(schedules.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range schedules.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry v20201201s.ScheduleEntry
			err := scheduleEntryItem.AssignPropertiesToScheduleEntry(&scheduleEntry)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToScheduleEntry() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		destination.ScheduleEntries = scheduleEntryList
	} else {
		destination.ScheduleEntries = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(schedules.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20201201.ScheduleEntry
// Deprecated version of ScheduleEntry. Use v1beta20201201.ScheduleEntry instead
type ScheduleEntry struct {
	DayOfWeek         *string                `json:"dayOfWeek,omitempty"`
	MaintenanceWindow *string                `json:"maintenanceWindow,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartHourUtc      *int                   `json:"startHourUtc,omitempty"`
}

// AssignPropertiesFromScheduleEntry populates our ScheduleEntry from the provided source ScheduleEntry
func (entry *ScheduleEntry) AssignPropertiesFromScheduleEntry(source *v20201201s.ScheduleEntry) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// DayOfWeek
	entry.DayOfWeek = genruntime.ClonePointerToString(source.DayOfWeek)

	// MaintenanceWindow
	entry.MaintenanceWindow = genruntime.ClonePointerToString(source.MaintenanceWindow)

	// StartHourUtc
	entry.StartHourUtc = genruntime.ClonePointerToInt(source.StartHourUtc)

	// Update the property bag
	if len(propertyBag) > 0 {
		entry.PropertyBag = propertyBag
	} else {
		entry.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToScheduleEntry populates the provided destination ScheduleEntry from our ScheduleEntry
func (entry *ScheduleEntry) AssignPropertiesToScheduleEntry(destination *v20201201s.ScheduleEntry) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(entry.PropertyBag)

	// DayOfWeek
	destination.DayOfWeek = genruntime.ClonePointerToString(entry.DayOfWeek)

	// MaintenanceWindow
	destination.MaintenanceWindow = genruntime.ClonePointerToString(entry.MaintenanceWindow)

	// StartHourUtc
	destination.StartHourUtc = genruntime.ClonePointerToInt(entry.StartHourUtc)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20201201.ScheduleEntry_Status
// Deprecated version of ScheduleEntry_Status. Use v1beta20201201.ScheduleEntry_Status instead
type ScheduleEntry_Status struct {
	DayOfWeek         *string                `json:"dayOfWeek,omitempty"`
	MaintenanceWindow *string                `json:"maintenanceWindow,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartHourUtc      *int                   `json:"startHourUtc,omitempty"`
}

// AssignPropertiesFromScheduleEntryStatus populates our ScheduleEntry_Status from the provided source ScheduleEntry_Status
func (entry *ScheduleEntry_Status) AssignPropertiesFromScheduleEntryStatus(source *v20201201s.ScheduleEntry_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// DayOfWeek
	entry.DayOfWeek = genruntime.ClonePointerToString(source.DayOfWeek)

	// MaintenanceWindow
	entry.MaintenanceWindow = genruntime.ClonePointerToString(source.MaintenanceWindow)

	// StartHourUtc
	entry.StartHourUtc = genruntime.ClonePointerToInt(source.StartHourUtc)

	// Update the property bag
	if len(propertyBag) > 0 {
		entry.PropertyBag = propertyBag
	} else {
		entry.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToScheduleEntryStatus populates the provided destination ScheduleEntry_Status from our ScheduleEntry_Status
func (entry *ScheduleEntry_Status) AssignPropertiesToScheduleEntryStatus(destination *v20201201s.ScheduleEntry_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(entry.PropertyBag)

	// DayOfWeek
	destination.DayOfWeek = genruntime.ClonePointerToString(entry.DayOfWeek)

	// MaintenanceWindow
	destination.MaintenanceWindow = genruntime.ClonePointerToString(entry.MaintenanceWindow)

	// StartHourUtc
	destination.StartHourUtc = genruntime.ClonePointerToInt(entry.StartHourUtc)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&RedisPatchSchedule{}, &RedisPatchScheduleList{})
}
