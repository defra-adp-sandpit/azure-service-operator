// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210501

import "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

type Server_StatusARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Identity: The cmk identity for the server.
	Identity *Identity_StatusARM `json:"identity,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the server.
	Properties *ServerProperties_StatusARM `json:"properties,omitempty"`

	// Sku: The SKU (pricing tier) of the server.
	Sku *Sku_StatusARM `json:"sku,omitempty"`

	// SystemData: The system metadata relating to this resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type Identity_StatusARM struct {
	// PrincipalId: ObjectId from the KeyVault
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: TenantId from the KeyVault
	TenantId *string `json:"tenantId,omitempty"`

	// Type: Type of managed service identity.
	Type *IdentityStatusType `json:"type,omitempty"`

	// UserAssignedIdentities: Metadata of user assigned identity.
	UserAssignedIdentities map[string]v1.JSON `json:"userAssignedIdentities,omitempty"`
}

type ServerProperties_StatusARM struct {
	// AdministratorLogin: The administrator's login name of a server. Can only be specified when the server is being created
	// (and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AvailabilityZone: availability Zone information of the server.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	// Backup: Backup related properties of a server.
	Backup *Backup_StatusARM `json:"backup,omitempty"`

	// CreateMode: The mode to create a new MySQL server.
	CreateMode *ServerPropertiesStatusCreateMode `json:"createMode,omitempty"`

	// DataEncryption: The Data Encryption for CMK.
	DataEncryption *DataEncryption_StatusARM `json:"dataEncryption,omitempty"`

	// FullyQualifiedDomainName: The fully qualified domain name of a server.
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`

	// HighAvailability: High availability related properties of a server.
	HighAvailability *HighAvailability_StatusARM `json:"highAvailability,omitempty"`

	// MaintenanceWindow: Maintenance window of a server.
	MaintenanceWindow *MaintenanceWindow_StatusARM `json:"maintenanceWindow,omitempty"`

	// Network: Network related properties of a server.
	Network *Network_StatusARM `json:"network,omitempty"`

	// ReplicaCapacity: The maximum number of replicas that a primary server can have.
	ReplicaCapacity *int `json:"replicaCapacity,omitempty"`

	// ReplicationRole: The replication role.
	ReplicationRole *ReplicationRole_Status `json:"replicationRole,omitempty"`

	// RestorePointInTime: Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime *string `json:"restorePointInTime,omitempty"`

	// SourceServerResourceId: The source MySQL server id.
	SourceServerResourceId *string `json:"sourceServerResourceId,omitempty"`

	// State: The state of a server.
	State *ServerPropertiesStatusState `json:"state,omitempty"`

	// Storage: Storage related properties of a server.
	Storage *Storage_StatusARM `json:"storage,omitempty"`

	// Version: Server version.
	Version *ServerVersion_Status `json:"version,omitempty"`
}

type Sku_StatusARM struct {
	// Name: The name of the sku, e.g. Standard_D32s_v3.
	Name *string `json:"name,omitempty"`

	// Tier: The tier of the particular SKU, e.g. GeneralPurpose.
	Tier *SkuStatusTier `json:"tier,omitempty"`
}

type Backup_StatusARM struct {
	// BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	// EarliestRestoreDate: Earliest restore point creation time (ISO8601 format)
	EarliestRestoreDate *string `json:"earliestRestoreDate,omitempty"`

	// GeoRedundantBackup: Whether or not geo redundant backup is enabled.
	GeoRedundantBackup *EnableStatusEnum_Status `json:"geoRedundantBackup,omitempty"`
}

type DataEncryption_StatusARM struct {
	// GeoBackupKeyUri: Geo backup key uri as key vault can't cross region, need cmk in same region as geo backup
	GeoBackupKeyUri *string `json:"geoBackupKeyUri,omitempty"`

	// GeoBackupUserAssignedIdentityId: Geo backup user identity resource id as identity can't cross region, need identity in
	// same region as geo backup
	GeoBackupUserAssignedIdentityId *string `json:"geoBackupUserAssignedIdentityId,omitempty"`

	// PrimaryKeyUri: Primary key uri
	PrimaryKeyUri *string `json:"primaryKeyUri,omitempty"`

	// PrimaryUserAssignedIdentityId: Primary user identity resource id
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	// Type: The key type, AzureKeyVault for enable cmk, SystemManaged for disable cmk.
	Type *DataEncryptionStatusType `json:"type,omitempty"`
}

type HighAvailability_StatusARM struct {
	// Mode: High availability mode for a server.
	Mode *HighAvailabilityStatusMode `json:"mode,omitempty"`

	// StandbyAvailabilityZone: Availability zone of the standby server.
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty"`

	// State: The state of server high availability.
	State *HighAvailabilityStatusState `json:"state,omitempty"`
}

type IdentityStatusType string

const IdentityStatusTypeUserAssigned = IdentityStatusType("UserAssigned")

type MaintenanceWindow_StatusARM struct {
	// CustomWindow: indicates whether custom window is enabled or disabled
	CustomWindow *string `json:"customWindow,omitempty"`

	// DayOfWeek: day of week for maintenance window
	DayOfWeek *int `json:"dayOfWeek,omitempty"`

	// StartHour: start hour for maintenance window
	StartHour *int `json:"startHour,omitempty"`

	// StartMinute: start minute for maintenance window
	StartMinute *int `json:"startMinute,omitempty"`
}

type Network_StatusARM struct {
	// DelegatedSubnetResourceId: Delegated subnet resource id used to setup vnet for a server.
	DelegatedSubnetResourceId *string `json:"delegatedSubnetResourceId,omitempty"`

	// PrivateDnsZoneResourceId: Private DNS zone resource id.
	PrivateDnsZoneResourceId *string `json:"privateDnsZoneResourceId,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for this server. Value is 'Disabled' when server
	// has VNet integration.
	PublicNetworkAccess *EnableStatusEnum_Status `json:"publicNetworkAccess,omitempty"`
}

type SkuStatusTier string

const (
	SkuStatusTierBurstable       = SkuStatusTier("Burstable")
	SkuStatusTierGeneralPurpose  = SkuStatusTier("GeneralPurpose")
	SkuStatusTierMemoryOptimized = SkuStatusTier("MemoryOptimized")
)

type Storage_StatusARM struct {
	// AutoGrow: Enable Storage Auto Grow or not.
	AutoGrow *EnableStatusEnum_Status `json:"autoGrow,omitempty"`

	// Iops: Storage IOPS for a server.
	Iops *int `json:"iops,omitempty"`

	// StorageSizeGB: Max storage size allowed for a server.
	StorageSizeGB *int `json:"storageSizeGB,omitempty"`

	// StorageSku: The sku name of the server storage.
	StorageSku *string `json:"storageSku,omitempty"`
}
