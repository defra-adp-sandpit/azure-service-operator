// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601storage

import (
	v20210601s "github.com/Azure/azure-service-operator/v2/api/cdn/v1api20210601storage"
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
// Storage version of v1beta20210601.Profile
// Deprecated version of Profile. Use v1api20210601.Profile instead
type Profile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Profile_Spec   `json:"spec,omitempty"`
	Status            Profile_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Profile{}

// GetConditions returns the conditions of the resource
func (profile *Profile) GetConditions() conditions.Conditions {
	return profile.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (profile *Profile) SetConditions(conditions conditions.Conditions) {
	profile.Status.Conditions = conditions
}

var _ conversion.Convertible = &Profile{}

// ConvertFrom populates our Profile from the provided hub Profile
func (profile *Profile) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source v20210601s.Profile

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = profile.AssignProperties_From_Profile(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to profile")
	}

	return nil
}

// ConvertTo populates the provided hub Profile from our Profile
func (profile *Profile) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination v20210601s.Profile
	err := profile.AssignProperties_To_Profile(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from profile")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

var _ genruntime.KubernetesResource = &Profile{}

// AzureName returns the Azure name of the resource
func (profile *Profile) AzureName() string {
	return profile.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (profile Profile) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (profile *Profile) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (profile *Profile) GetSpec() genruntime.ConvertibleSpec {
	return &profile.Spec
}

// GetStatus returns the status of this resource
func (profile *Profile) GetStatus() genruntime.ConvertibleStatus {
	return &profile.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles"
func (profile *Profile) GetType() string {
	return "Microsoft.Cdn/profiles"
}

// NewEmptyStatus returns a new empty (blank) status
func (profile *Profile) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Profile_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (profile *Profile) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(profile.Spec)
	return profile.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (profile *Profile) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Profile_STATUS); ok {
		profile.Status = *st
		return nil
	}

	// Convert status to required version
	var st Profile_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	profile.Status = st
	return nil
}

// AssignProperties_From_Profile populates our Profile from the provided source Profile
func (profile *Profile) AssignProperties_From_Profile(source *v20210601s.Profile) error {

	// ObjectMeta
	profile.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Profile_Spec
	err := spec.AssignProperties_From_Profile_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Profile_Spec() to populate field Spec")
	}
	profile.Spec = spec

	// Status
	var status Profile_STATUS
	err = status.AssignProperties_From_Profile_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Profile_STATUS() to populate field Status")
	}
	profile.Status = status

	// Invoke the augmentConversionForProfile interface (if implemented) to customize the conversion
	var profileAsAny any = profile
	if augmentedProfile, ok := profileAsAny.(augmentConversionForProfile); ok {
		err := augmentedProfile.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Profile populates the provided destination Profile from our Profile
func (profile *Profile) AssignProperties_To_Profile(destination *v20210601s.Profile) error {

	// ObjectMeta
	destination.ObjectMeta = *profile.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210601s.Profile_Spec
	err := profile.Spec.AssignProperties_To_Profile_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Profile_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210601s.Profile_STATUS
	err = profile.Status.AssignProperties_To_Profile_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Profile_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForProfile interface (if implemented) to customize the conversion
	var profileAsAny any = profile
	if augmentedProfile, ok := profileAsAny.(augmentConversionForProfile); ok {
		err := augmentedProfile.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (profile *Profile) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: profile.Spec.OriginalVersion,
		Kind:    "Profile",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210601.Profile
// Deprecated version of Profile. Use v1api20210601.Profile instead
type ProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Profile `json:"items"`
}

// Storage version of v1beta20210601.APIVersion
// Deprecated version of APIVersion. Use v1api20210601.APIVersion instead
// +kubebuilder:validation:Enum={"2021-06-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2021-06-01")

type augmentConversionForProfile interface {
	AssignPropertiesFrom(src *v20210601s.Profile) error
	AssignPropertiesTo(dst *v20210601s.Profile) error
}

// Storage version of v1beta20210601.Profile_Spec
type Profile_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                    string  `json:"azureName,omitempty"`
	Location                     *string `json:"location,omitempty"`
	OriginResponseTimeoutSeconds *int    `json:"originResponseTimeoutSeconds,omitempty"`
	OriginalVersion              string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Sku         *Sku                               `json:"sku,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Profile_Spec{}

// ConvertSpecFrom populates our Profile_Spec from the provided source
func (profile *Profile_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210601s.Profile_Spec)
	if ok {
		// Populate our instance from source
		return profile.AssignProperties_From_Profile_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210601s.Profile_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = profile.AssignProperties_From_Profile_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Profile_Spec
func (profile *Profile_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210601s.Profile_Spec)
	if ok {
		// Populate destination from our instance
		return profile.AssignProperties_To_Profile_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210601s.Profile_Spec{}
	err := profile.AssignProperties_To_Profile_Spec(dst)
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

// AssignProperties_From_Profile_Spec populates our Profile_Spec from the provided source Profile_Spec
func (profile *Profile_Spec) AssignProperties_From_Profile_Spec(source *v20210601s.Profile_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	profile.AzureName = source.AzureName

	// Location
	profile.Location = genruntime.ClonePointerToString(source.Location)

	// OriginResponseTimeoutSeconds
	profile.OriginResponseTimeoutSeconds = genruntime.ClonePointerToInt(source.OriginResponseTimeoutSeconds)

	// OriginalVersion
	profile.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		profile.Owner = &owner
	} else {
		profile.Owner = nil
	}

	// Sku
	if source.Sku != nil {
		var sku Sku
		err := sku.AssignProperties_From_Sku(source.Sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_Sku() to populate field Sku")
		}
		profile.Sku = &sku
	} else {
		profile.Sku = nil
	}

	// Tags
	profile.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		profile.PropertyBag = propertyBag
	} else {
		profile.PropertyBag = nil
	}

	// Invoke the augmentConversionForProfile_Spec interface (if implemented) to customize the conversion
	var profileAsAny any = profile
	if augmentedProfile, ok := profileAsAny.(augmentConversionForProfile_Spec); ok {
		err := augmentedProfile.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Profile_Spec populates the provided destination Profile_Spec from our Profile_Spec
func (profile *Profile_Spec) AssignProperties_To_Profile_Spec(destination *v20210601s.Profile_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(profile.PropertyBag)

	// AzureName
	destination.AzureName = profile.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(profile.Location)

	// OriginResponseTimeoutSeconds
	destination.OriginResponseTimeoutSeconds = genruntime.ClonePointerToInt(profile.OriginResponseTimeoutSeconds)

	// OriginalVersion
	destination.OriginalVersion = profile.OriginalVersion

	// Owner
	if profile.Owner != nil {
		owner := profile.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Sku
	if profile.Sku != nil {
		var sku v20210601s.Sku
		err := profile.Sku.AssignProperties_To_Sku(&sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_Sku() to populate field Sku")
		}
		destination.Sku = &sku
	} else {
		destination.Sku = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(profile.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForProfile_Spec interface (if implemented) to customize the conversion
	var profileAsAny any = profile
	if augmentedProfile, ok := profileAsAny.(augmentConversionForProfile_Spec); ok {
		err := augmentedProfile.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1beta20210601.Profile_STATUS
// Deprecated version of Profile_STATUS. Use v1api20210601.Profile_STATUS instead
type Profile_STATUS struct {
	Conditions                   []conditions.Condition `json:"conditions,omitempty"`
	FrontDoorId                  *string                `json:"frontDoorId,omitempty"`
	Id                           *string                `json:"id,omitempty"`
	Kind                         *string                `json:"kind,omitempty"`
	Location                     *string                `json:"location,omitempty"`
	Name                         *string                `json:"name,omitempty"`
	OriginResponseTimeoutSeconds *int                   `json:"originResponseTimeoutSeconds,omitempty"`
	PropertyBag                  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState            *string                `json:"provisioningState,omitempty"`
	ResourceState                *string                `json:"resourceState,omitempty"`
	Sku                          *Sku_STATUS            `json:"sku,omitempty"`
	SystemData                   *SystemData_STATUS     `json:"systemData,omitempty"`
	Tags                         map[string]string      `json:"tags,omitempty"`
	Type                         *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Profile_STATUS{}

// ConvertStatusFrom populates our Profile_STATUS from the provided source
func (profile *Profile_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210601s.Profile_STATUS)
	if ok {
		// Populate our instance from source
		return profile.AssignProperties_From_Profile_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210601s.Profile_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = profile.AssignProperties_From_Profile_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Profile_STATUS
func (profile *Profile_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210601s.Profile_STATUS)
	if ok {
		// Populate destination from our instance
		return profile.AssignProperties_To_Profile_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210601s.Profile_STATUS{}
	err := profile.AssignProperties_To_Profile_STATUS(dst)
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

// AssignProperties_From_Profile_STATUS populates our Profile_STATUS from the provided source Profile_STATUS
func (profile *Profile_STATUS) AssignProperties_From_Profile_STATUS(source *v20210601s.Profile_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	profile.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// FrontDoorId
	profile.FrontDoorId = genruntime.ClonePointerToString(source.FrontDoorId)

	// Id
	profile.Id = genruntime.ClonePointerToString(source.Id)

	// Kind
	profile.Kind = genruntime.ClonePointerToString(source.Kind)

	// Location
	profile.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	profile.Name = genruntime.ClonePointerToString(source.Name)

	// OriginResponseTimeoutSeconds
	profile.OriginResponseTimeoutSeconds = genruntime.ClonePointerToInt(source.OriginResponseTimeoutSeconds)

	// ProvisioningState
	profile.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// ResourceState
	profile.ResourceState = genruntime.ClonePointerToString(source.ResourceState)

	// Sku
	if source.Sku != nil {
		var sku Sku_STATUS
		err := sku.AssignProperties_From_Sku_STATUS(source.Sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_Sku_STATUS() to populate field Sku")
		}
		profile.Sku = &sku
	} else {
		profile.Sku = nil
	}

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		profile.SystemData = &systemDatum
	} else {
		profile.SystemData = nil
	}

	// Tags
	profile.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	profile.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		profile.PropertyBag = propertyBag
	} else {
		profile.PropertyBag = nil
	}

	// Invoke the augmentConversionForProfile_STATUS interface (if implemented) to customize the conversion
	var profileAsAny any = profile
	if augmentedProfile, ok := profileAsAny.(augmentConversionForProfile_STATUS); ok {
		err := augmentedProfile.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Profile_STATUS populates the provided destination Profile_STATUS from our Profile_STATUS
func (profile *Profile_STATUS) AssignProperties_To_Profile_STATUS(destination *v20210601s.Profile_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(profile.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(profile.Conditions)

	// FrontDoorId
	destination.FrontDoorId = genruntime.ClonePointerToString(profile.FrontDoorId)

	// Id
	destination.Id = genruntime.ClonePointerToString(profile.Id)

	// Kind
	destination.Kind = genruntime.ClonePointerToString(profile.Kind)

	// Location
	destination.Location = genruntime.ClonePointerToString(profile.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(profile.Name)

	// OriginResponseTimeoutSeconds
	destination.OriginResponseTimeoutSeconds = genruntime.ClonePointerToInt(profile.OriginResponseTimeoutSeconds)

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(profile.ProvisioningState)

	// ResourceState
	destination.ResourceState = genruntime.ClonePointerToString(profile.ResourceState)

	// Sku
	if profile.Sku != nil {
		var sku v20210601s.Sku_STATUS
		err := profile.Sku.AssignProperties_To_Sku_STATUS(&sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_Sku_STATUS() to populate field Sku")
		}
		destination.Sku = &sku
	} else {
		destination.Sku = nil
	}

	// SystemData
	if profile.SystemData != nil {
		var systemDatum v20210601s.SystemData_STATUS
		err := profile.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(profile.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(profile.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForProfile_STATUS interface (if implemented) to customize the conversion
	var profileAsAny any = profile
	if augmentedProfile, ok := profileAsAny.(augmentConversionForProfile_STATUS); ok {
		err := augmentedProfile.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForProfile_Spec interface {
	AssignPropertiesFrom(src *v20210601s.Profile_Spec) error
	AssignPropertiesTo(dst *v20210601s.Profile_Spec) error
}

type augmentConversionForProfile_STATUS interface {
	AssignPropertiesFrom(src *v20210601s.Profile_STATUS) error
	AssignPropertiesTo(dst *v20210601s.Profile_STATUS) error
}

// Storage version of v1beta20210601.Sku
// Deprecated version of Sku. Use v1api20210601.Sku instead
type Sku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_Sku populates our Sku from the provided source Sku
func (sku *Sku) AssignProperties_From_Sku(source *v20210601s.Sku) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Name
	sku.Name = genruntime.ClonePointerToString(source.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		sku.PropertyBag = propertyBag
	} else {
		sku.PropertyBag = nil
	}

	// Invoke the augmentConversionForSku interface (if implemented) to customize the conversion
	var skuAsAny any = sku
	if augmentedSku, ok := skuAsAny.(augmentConversionForSku); ok {
		err := augmentedSku.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Sku populates the provided destination Sku from our Sku
func (sku *Sku) AssignProperties_To_Sku(destination *v20210601s.Sku) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(sku.PropertyBag)

	// Name
	destination.Name = genruntime.ClonePointerToString(sku.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSku interface (if implemented) to customize the conversion
	var skuAsAny any = sku
	if augmentedSku, ok := skuAsAny.(augmentConversionForSku); ok {
		err := augmentedSku.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1beta20210601.Sku_STATUS
// Deprecated version of Sku_STATUS. Use v1api20210601.Sku_STATUS instead
type Sku_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_Sku_STATUS populates our Sku_STATUS from the provided source Sku_STATUS
func (sku *Sku_STATUS) AssignProperties_From_Sku_STATUS(source *v20210601s.Sku_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Name
	sku.Name = genruntime.ClonePointerToString(source.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		sku.PropertyBag = propertyBag
	} else {
		sku.PropertyBag = nil
	}

	// Invoke the augmentConversionForSku_STATUS interface (if implemented) to customize the conversion
	var skuAsAny any = sku
	if augmentedSku, ok := skuAsAny.(augmentConversionForSku_STATUS); ok {
		err := augmentedSku.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Sku_STATUS populates the provided destination Sku_STATUS from our Sku_STATUS
func (sku *Sku_STATUS) AssignProperties_To_Sku_STATUS(destination *v20210601s.Sku_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(sku.PropertyBag)

	// Name
	destination.Name = genruntime.ClonePointerToString(sku.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSku_STATUS interface (if implemented) to customize the conversion
	var skuAsAny any = sku
	if augmentedSku, ok := skuAsAny.(augmentConversionForSku_STATUS); ok {
		err := augmentedSku.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1beta20210601.SystemData_STATUS
// Deprecated version of SystemData_STATUS. Use v1api20210601.SystemData_STATUS instead
type SystemData_STATUS struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_SystemData_STATUS populates our SystemData_STATUS from the provided source SystemData_STATUS
func (data *SystemData_STATUS) AssignProperties_From_SystemData_STATUS(source *v20210601s.SystemData_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// CreatedAt
	data.CreatedAt = genruntime.ClonePointerToString(source.CreatedAt)

	// CreatedBy
	data.CreatedBy = genruntime.ClonePointerToString(source.CreatedBy)

	// CreatedByType
	data.CreatedByType = genruntime.ClonePointerToString(source.CreatedByType)

	// LastModifiedAt
	data.LastModifiedAt = genruntime.ClonePointerToString(source.LastModifiedAt)

	// LastModifiedBy
	data.LastModifiedBy = genruntime.ClonePointerToString(source.LastModifiedBy)

	// LastModifiedByType
	data.LastModifiedByType = genruntime.ClonePointerToString(source.LastModifiedByType)

	// Update the property bag
	if len(propertyBag) > 0 {
		data.PropertyBag = propertyBag
	} else {
		data.PropertyBag = nil
	}

	// Invoke the augmentConversionForSystemData_STATUS interface (if implemented) to customize the conversion
	var dataAsAny any = data
	if augmentedData, ok := dataAsAny.(augmentConversionForSystemData_STATUS); ok {
		err := augmentedData.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SystemData_STATUS populates the provided destination SystemData_STATUS from our SystemData_STATUS
func (data *SystemData_STATUS) AssignProperties_To_SystemData_STATUS(destination *v20210601s.SystemData_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(data.PropertyBag)

	// CreatedAt
	destination.CreatedAt = genruntime.ClonePointerToString(data.CreatedAt)

	// CreatedBy
	destination.CreatedBy = genruntime.ClonePointerToString(data.CreatedBy)

	// CreatedByType
	destination.CreatedByType = genruntime.ClonePointerToString(data.CreatedByType)

	// LastModifiedAt
	destination.LastModifiedAt = genruntime.ClonePointerToString(data.LastModifiedAt)

	// LastModifiedBy
	destination.LastModifiedBy = genruntime.ClonePointerToString(data.LastModifiedBy)

	// LastModifiedByType
	destination.LastModifiedByType = genruntime.ClonePointerToString(data.LastModifiedByType)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSystemData_STATUS interface (if implemented) to customize the conversion
	var dataAsAny any = data
	if augmentedData, ok := dataAsAny.(augmentConversionForSystemData_STATUS); ok {
		err := augmentedData.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForSku interface {
	AssignPropertiesFrom(src *v20210601s.Sku) error
	AssignPropertiesTo(dst *v20210601s.Sku) error
}

type augmentConversionForSku_STATUS interface {
	AssignPropertiesFrom(src *v20210601s.Sku_STATUS) error
	AssignPropertiesTo(dst *v20210601s.Sku_STATUS) error
}

type augmentConversionForSystemData_STATUS interface {
	AssignPropertiesFrom(src *v20210601s.SystemData_STATUS) error
	AssignPropertiesTo(dst *v20210601s.SystemData_STATUS) error
}

func init() {
	SchemeBuilder.Register(&Profile{}, &ProfileList{})
}
