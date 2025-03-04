# Go API client for CockroachDB Cloud API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2023-04-10
- Package version: 1.1.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://support.cockroachlabs.com](https://support.cockroachlabs.com)

## Installation

Install the package:

```shell
go get github.com/cockroachdb/cockroach-cloud-sdk-go
```

Import the package:

```golang
import "github.com/cockroachdb/pkg/client"
```

## Example Client Usage

```golang 
// Create a default Configuration object.
clientConfig := client.NewConfiguration(apiKey)

// Create a new client.
client := client.NewClient(clientConfig)

// Create a new service.
service := client.NewService(client)

requiredPathParam = "a_cluster_id"
requiredRequest := &client.OperationRequest{...}
optionalParams := &client.OperationOptions{...}

// Execute the request
apiResponse, httpResponse, err := := service.Operation(
	context.Background(), 
	requiredPathParam, 
	requiredRequest, 
	optionalParams,
)
if err != nil {
	// Process error
}

if apiResponse != nil {
	respJSON, err := apiResponse.MarshalJSON()
	if err != nil {
		// Process error
	}
	fmt.Println(string(respJSON))
}
```

## Documentation for API Endpoints

All URIs are relative to *https://cockroachlabs.cloud*

API | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
[AuditLogsApi](docs/AuditLogsApi.md) | **ListAuditLogs** | **Get** /api/v1/auditlogevents | List audit logs
[BillingApi](docs/BillingApi.md) | **GetInvoice** | **Get** /api/v1/invoices/{invoice_id} | Get a specific invoice for an organization
[BillingApi](docs/BillingApi.md) | **ListInvoices** | **Get** /api/v1/invoices | List invoices for a given organization
[ClientCACertificatesApi](docs/ClientCACertificatesApi.md) | **DeleteClientCACert** | **Delete** /api/v1/clusters/{cluster_id}/client-ca-cert | Delete Client CA Cert for a cluster
[ClientCACertificatesApi](docs/ClientCACertificatesApi.md) | **GetClientCACert** | **Get** /api/v1/clusters/{cluster_id}/client-ca-cert | Get Client CA Cert information for a cluster
[ClientCACertificatesApi](docs/ClientCACertificatesApi.md) | **SetClientCACert** | **Post** /api/v1/clusters/{cluster_id}/client-ca-cert | Set Client CA Cert for a cluster
[ClientCACertificatesApi](docs/ClientCACertificatesApi.md) | **UpdateClientCACert** | **Patch** /api/v1/clusters/{cluster_id}/client-ca-cert | Update Client CA Cert for a cluster
[ClustersApi](docs/ClustersApi.md) | **CreateCluster** | **Post** /api/v1/clusters | Create and initialize a new cluster
[ClustersApi](docs/ClustersApi.md) | **DeleteCluster** | **Delete** /api/v1/clusters/{cluster_id} | Delete a cluster and all of its data
[ClustersApi](docs/ClustersApi.md) | **GetCluster** | **Get** /api/v1/clusters/{cluster_id} | Get extended information about a cluster
[ClustersApi](docs/ClustersApi.md) | **GetConnectionString** | **Get** /api/v1/clusters/{cluster_id}/connection-string | Get a formatted generic connection string for a cluster
[ClustersApi](docs/ClustersApi.md) | **ListAvailableRegions** | **Get** /api/v1/clusters/available-regions | List the regions available for new clusters and nodes
[ClustersApi](docs/ClustersApi.md) | **ListClusterNodes** | **Get** /api/v1/clusters/{cluster_id}/nodes | List nodes for a cluster
[ClustersApi](docs/ClustersApi.md) | **ListClusters** | **Get** /api/v1/clusters | List clusters owned by an organization
[ClustersApi](docs/ClustersApi.md) | **ListMajorClusterVersions** | **Get** /api/v1/cluster-versions | List available major cluster versions
[ClustersApi](docs/ClustersApi.md) | **UpdateCluster** | **Patch** /api/v1/clusters/{cluster_id} | Scale or edit a cluster
[CustomerManagedEncryptionKeysApi](docs/CustomerManagedEncryptionKeysApi.md) | **EnableCMEKSpec** | **Post** /api/v1/clusters/{cluster_id}/cmek | Enable CMEK for a cluster
[CustomerManagedEncryptionKeysApi](docs/CustomerManagedEncryptionKeysApi.md) | **GetCMEKClusterInfo** | **Get** /api/v1/clusters/{cluster_id}/cmek | Get CMEK-related information for a cluster
[CustomerManagedEncryptionKeysApi](docs/CustomerManagedEncryptionKeysApi.md) | **UpdateCMEKSpec** | **Put** /api/v1/clusters/{cluster_id}/cmek | Enable or update the CMEK spec for a cluster
[CustomerManagedEncryptionKeysApi](docs/CustomerManagedEncryptionKeysApi.md) | **UpdateCMEKStatus** | **Patch** /api/v1/clusters/{cluster_id}/cmek | Update the CMEK-related status for a cluster
[DatabasesApi](docs/DatabasesApi.md) | **CreateDatabase** | **Post** /api/v1/clusters/{cluster_id}/databases | Create a new database
[DatabasesApi](docs/DatabasesApi.md) | **DeleteDatabase** | **Delete** /api/v1/clusters/{cluster_id}/databases/{name} | Delete a database
[DatabasesApi](docs/DatabasesApi.md) | **EditDatabase** | **Patch** /api/v1/clusters/{cluster_id}/databases/{name} | Update a database
[DatabasesApi](docs/DatabasesApi.md) | **EditDatabase2** | **Patch** /api/v1/clusters/{cluster_id}/databases | Update a database
[DatabasesApi](docs/DatabasesApi.md) | **ListDatabases** | **Get** /api/v1/clusters/{cluster_id}/databases | List databases for a cluster
[DefaultApi](docs/DefaultApi.md) | **GetClusterVersionDeferral** | **Get** /api/v1/clusters/{cluster_id}/version-deferral | 
[DefaultApi](docs/DefaultApi.md) | **SetClusterVersionDeferral** | **Put** /api/v1/clusters/{cluster_id}/version-deferral | 
[EgressRulesApi](docs/EgressRulesApi.md) | **AddEgressRule** | **Post** /api/v1/clusters/{cluster_id}/networking/egress-rules | Add an egress rule
[EgressRulesApi](docs/EgressRulesApi.md) | **DeleteEgressRule** | **Delete** /api/v1/clusters/{cluster_id}/networking/egress-rules/{rule_id} | Delete an existing egress rule
[EgressRulesApi](docs/EgressRulesApi.md) | **EditEgressRule** | **Patch** /api/v1/clusters/{cluster_id}/networking/egress-rules/{rule_id} | Edit an existing egress rule
[EgressRulesApi](docs/EgressRulesApi.md) | **GetEgressRule** | **Get** /api/v1/clusters/{cluster_id}/networking/egress-rules/{rule_id} | Get an existing egress rule
[EgressRulesApi](docs/EgressRulesApi.md) | **ListEgressRules** | **Get** /api/v1/clusters/{cluster_id}/networking/egress-rules | List all egress rules associates with a cluster
[EgressRulesApi](docs/EgressRulesApi.md) | **SetEgressTrafficPolicy** | **Post** /api/v1/clusters/{cluster_id}/networking/egress-rules/egress-traffic-policy | Outbound traffic management
[IPAllowlistsApi](docs/IPAllowlistsApi.md) | **AddAllowlistEntry** | **Post** /api/v1/clusters/{cluster_id}/networking/allowlist | Add a new CIDR address to the IP allowlist
[IPAllowlistsApi](docs/IPAllowlistsApi.md) | **AddAllowlistEntry2** | **Put** /api/v1/clusters/{cluster_id}/networking/allowlist/{entry.cidr_ip}/{entry.cidr_mask} | Add a new CIDR address to the IP allowlist
[IPAllowlistsApi](docs/IPAllowlistsApi.md) | **DeleteAllowlistEntry** | **Delete** /api/v1/clusters/{cluster_id}/networking/allowlist/{cidr_ip}/{cidr_mask} | Delete an IP allowlist entry
[IPAllowlistsApi](docs/IPAllowlistsApi.md) | **ListAllowlistEntries** | **Get** /api/v1/clusters/{cluster_id}/networking/allowlist | Get the IP allowlist and propagation status for a cluster
[IPAllowlistsApi](docs/IPAllowlistsApi.md) | **UpdateAllowlistEntry** | **Patch** /api/v1/clusters/{cluster_id}/networking/allowlist/{entry.cidr_ip}/{entry.cidr_mask} | Update an IP allowlist entry
[LogExportApi](docs/LogExportApi.md) | **DeleteLogExport** | **Delete** /api/v1/clusters/{cluster_id}/logexport | Delete the Log Export configuration for a cluster
[LogExportApi](docs/LogExportApi.md) | **EnableLogExport** | **Post** /api/v1/clusters/{cluster_id}/logexport | Create or update the Log Export configuration for a cluster
[LogExportApi](docs/LogExportApi.md) | **GetLogExportInfo** | **Get** /api/v1/clusters/{cluster_id}/logexport | Get the Log Export configuration for a cluster
[MaintenanceWindowsApi](docs/MaintenanceWindowsApi.md) | **DeleteMaintenanceWindow** | **Delete** /api/v1/clusters/{cluster_id}/maintenance-window | 
[MaintenanceWindowsApi](docs/MaintenanceWindowsApi.md) | **GetMaintenanceWindow** | **Get** /api/v1/clusters/{cluster_id}/maintenance-window | 
[MaintenanceWindowsApi](docs/MaintenanceWindowsApi.md) | **SetMaintenanceWindow** | **Put** /api/v1/clusters/{cluster_id}/maintenance-window | 
[MetricExportApi](docs/MetricExportApi.md) | **DeleteCloudWatchMetricExport** | **Delete** /api/v1/clusters/{cluster_id}/metricexport/cloudwatch | Delete the CloudWatch Metric Export configuration for a cluster
[MetricExportApi](docs/MetricExportApi.md) | **DeleteDatadogMetricExport** | **Delete** /api/v1/clusters/{cluster_id}/metricexport/datadog | Delete the Datadog Metric Export configuration for a cluster
[MetricExportApi](docs/MetricExportApi.md) | **EnableCloudWatchMetricExport** | **Post** /api/v1/clusters/{cluster_id}/metricexport/cloudwatch | Create or update the CloudWatch Metric Export configuration for a cluster
[MetricExportApi](docs/MetricExportApi.md) | **EnableDatadogMetricExport** | **Post** /api/v1/clusters/{cluster_id}/metricexport/datadog | Create or update the Datadog Metric Export configuration for a cluster
[MetricExportApi](docs/MetricExportApi.md) | **GetCloudWatchMetricExportInfo** | **Get** /api/v1/clusters/{cluster_id}/metricexport/cloudwatch | Get the CloudWatch Metric Export configuration for a cluster
[MetricExportApi](docs/MetricExportApi.md) | **GetDatadogMetricExportInfo** | **Get** /api/v1/clusters/{cluster_id}/metricexport/datadog | Get the Datadog Metric Export configuration for a cluster
[OrganizationsApi](docs/OrganizationsApi.md) | **GetOrganizationInfo** | **Get** /api/v1/organization | Get information about the caller&#39;s organization
[PrivateEndpointServicesApi](docs/PrivateEndpointServicesApi.md) | **CreatePrivateEndpointServices** | **Post** /api/v1/clusters/{cluster_id}/networking/private-endpoint-services | Create all PrivateEndpointServices for a cluster
[PrivateEndpointServicesApi](docs/PrivateEndpointServicesApi.md) | **ListAwsEndpointConnections** | **Get** /api/v1/clusters/{cluster_id}/networking/aws-endpoint-connections | List all AwsEndpointConnections for a cluster
[PrivateEndpointServicesApi](docs/PrivateEndpointServicesApi.md) | **ListPrivateEndpointServices** | **Get** /api/v1/clusters/{cluster_id}/networking/private-endpoint-services | List all PrivateEndpointServices for a cluster
[PrivateEndpointServicesApi](docs/PrivateEndpointServicesApi.md) | **SetAwsEndpointConnectionState** | **Patch** /api/v1/clusters/{cluster_id}/networking/aws-endpoint-connections/{endpoint_id} | Set the AWS Endpoint Connection state
[RoleManagementApi](docs/RoleManagementApi.md) | **AddUserToRole** | **Post** /api/v1/roles/{user_id}/{resource_type}/{resource_id}/{role_name} | Add the user to the given role
[RoleManagementApi](docs/RoleManagementApi.md) | **GetAllRolesForUser** | **Get** /api/v1/roles/{user_id} | Get all Role Grants for a user
[RoleManagementApi](docs/RoleManagementApi.md) | **GetPersonUsersByEmail** | **Get** /api/v1/users/persons-by-email | Search person users by email address
[RoleManagementApi](docs/RoleManagementApi.md) | **ListRoleGrants** | **Get** /api/v1/roles | List all RoleGrants
[RoleManagementApi](docs/RoleManagementApi.md) | **RemoveUserFromRole** | **Delete** /api/v1/roles/{user_id}/{resource_type}/{resource_id}/{role_name} | Remove the user from the given role
[RoleManagementApi](docs/RoleManagementApi.md) | **SetRolesForUser** | **Put** /api/v1/roles/{user_id} | Make a user&#39;s roles exactly those provided
[SCIMApi](docs/SCIMApi.md) | **CreateGroup** | **Post** /api/scim/v2/Groups | Create a group
[SCIMApi](docs/SCIMApi.md) | **CreateUser** | **Post** /api/scim/v2/Users | Create a user
[SCIMApi](docs/SCIMApi.md) | **DeleteGroup** | **Delete** /api/scim/v2/Groups/{id} | Delete a group based on ID
[SCIMApi](docs/SCIMApi.md) | **DeleteUser** | **Delete** /api/scim/v2/Users/{id} | Delete a user based on ID
[SCIMApi](docs/SCIMApi.md) | **GetGroup** | **Get** /api/scim/v2/Groups/{id} | Get a group based on ID
[SCIMApi](docs/SCIMApi.md) | **GetGroup2** | **Put** /api/scim/v2/Groups/{id}/.search | Get a group based on ID
[SCIMApi](docs/SCIMApi.md) | **GetGroups** | **Get** /api/scim/v2/Groups | Get groups based on query parameters
[SCIMApi](docs/SCIMApi.md) | **GetGroups2** | **Put** /api/scim/v2/Groups/.search | Get groups based on query parameters
[SCIMApi](docs/SCIMApi.md) | **GetResourceType** | **Get** /api/scim/v2/ResourceTypes/{resourceId} | 
[SCIMApi](docs/SCIMApi.md) | **GetResourceTypes** | **Get** /api/scim/v2/ResourceTypes | 
[SCIMApi](docs/SCIMApi.md) | **GetSchema** | **Get** /api/scim/v2/Schemas/{schemaId} | 
[SCIMApi](docs/SCIMApi.md) | **GetSchemas** | **Get** /api/scim/v2/Schemas | 
[SCIMApi](docs/SCIMApi.md) | **GetServiceProviderConfig** | **Get** /api/scim/v2/ServiceProviderConfig | Return our SCIM configuration
[SCIMApi](docs/SCIMApi.md) | **GetUser** | **Get** /api/scim/v2/Users/{id} | Get a user based on ID
[SCIMApi](docs/SCIMApi.md) | **GetUser2** | **Put** /api/scim/v2/Users/{id}/.search | Get a user based on ID
[SCIMApi](docs/SCIMApi.md) | **GetUsers** | **Get** /api/scim/v2/Users | Get Users based on query parameters
[SCIMApi](docs/SCIMApi.md) | **GetUsers2** | **Put** /api/scim/v2/Users/.search | Get Users based on query parameters
[SCIMApi](docs/SCIMApi.md) | **UpdateGroup** | **Put** /api/scim/v2/Groups/{id} | Update a group by supplying all values of the user object
[SCIMApi](docs/SCIMApi.md) | **UpdateUser** | **Put** /api/scim/v2/Users/{id} | Update a user by supplying all values of the user object
[SQLUsersApi](docs/SQLUsersApi.md) | **CreateSQLUser** | **Post** /api/v1/clusters/{cluster_id}/sql-users | Create a new SQL user
[SQLUsersApi](docs/SQLUsersApi.md) | **DeleteSQLUser** | **Delete** /api/v1/clusters/{cluster_id}/sql-users/{name} | Delete a SQL user
[SQLUsersApi](docs/SQLUsersApi.md) | **ListSQLUsers** | **Get** /api/v1/clusters/{cluster_id}/sql-users | List SQL users for a cluster
[SQLUsersApi](docs/SQLUsersApi.md) | **UpdateSQLUserPassword** | **Put** /api/v1/clusters/{cluster_id}/sql-users/{name}/password | Update a SQL user&#39;s password


## Author
support@cockroachlabs.com
