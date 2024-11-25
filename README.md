# Go API client for begetOpenapiVps

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1.6.1
- Package version: v1.6.1
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import begetOpenapiVps "github.com/LTD-Beget/openapi-vps-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), begetOpenapiVps.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), begetOpenapiVps.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), begetOpenapiVps.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), begetOpenapiVps.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.beget.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BackupServiceApi* | [**BackupServiceGetAvailableCopies**](docs/BackupServiceApi.md#backupservicegetavailablecopies) | **Get** /v1/vps/backup | 
*BackupServiceApi* | [**BackupServiceGetBackupFileList**](docs/BackupServiceApi.md#backupservicegetbackupfilelist) | **Get** /v1/vps/{id}/backup/{copy_id} | 
*BackupServiceApi* | [**BackupServiceGetOrders**](docs/BackupServiceApi.md#backupservicegetorders) | **Get** /v1/vps/backup/orders | 
*BackupServiceApi* | [**BackupServiceRestoreFile**](docs/BackupServiceApi.md#backupservicerestorefile) | **Post** /v1/vps/{id}/backup/{copy_id}/file | 
*BackupServiceApi* | [**BackupServiceRestoreServer**](docs/BackupServiceApi.md#backupservicerestoreserver) | **Post** /v1/vps/{id}/backup/{copy_id}/server | 
*ConfiguratorServiceApi* | [**ConfiguratorServiceGetCalculation**](docs/ConfiguratorServiceApi.md#configuratorservicegetcalculation) | **Get** /v1/vps/configurator/calculation | 
*ConfiguratorServiceApi* | [**ConfiguratorServiceGetConfiguratorInfo**](docs/ConfiguratorServiceApi.md#configuratorservicegetconfiguratorinfo) | **Get** /v1/vps/configurator/info | 
*ManageServiceApi* | [**ManageServiceAttachIpAddress**](docs/ManageServiceApi.md#manageserviceattachipaddress) | **Post** /v1/vps/{id}/network/{ip_address} | 
*ManageServiceApi* | [**ManageServiceAttachSshKey**](docs/ManageServiceApi.md#manageserviceattachsshkey) | **Post** /v1/vps/{id}/sshKey/{ssh_key_id} | 
*ManageServiceApi* | [**ManageServiceAttachToPrivateNetwork**](docs/ManageServiceApi.md#manageserviceattachtoprivatenetwork) | **Post** /v1/vps/{id}/private-network/{network_id} | 
*ManageServiceApi* | [**ManageServiceChangeConfiguration**](docs/ManageServiceApi.md#manageservicechangeconfiguration) | **Put** /v1/vps/server/{id}/configuration | 
*ManageServiceApi* | [**ManageServiceChangeSshAccess**](docs/ManageServiceApi.md#manageservicechangesshaccess) | **Put** /v1/vps/{id}/ssh/access | 
*ManageServiceApi* | [**ManageServiceCheckSoftwareRequirements**](docs/ManageServiceApi.md#manageservicechecksoftwarerequirements) | **Post** /v1/vps/software/requirements | 
*ManageServiceApi* | [**ManageServiceCreateVps**](docs/ManageServiceApi.md#manageservicecreatevps) | **Post** /v1/vps/server | 
*ManageServiceApi* | [**ManageServiceDetachFromPrivateNetwork**](docs/ManageServiceApi.md#manageservicedetachfromprivatenetwork) | **Delete** /v1/vps/{id}/private-network/{network_id} | 
*ManageServiceApi* | [**ManageServiceDetachIpAddress**](docs/ManageServiceApi.md#manageservicedetachipaddress) | **Delete** /v1/vps/network/detach/{ip_address} | 
*ManageServiceApi* | [**ManageServiceDetachSshKey**](docs/ManageServiceApi.md#manageservicedetachsshkey) | **Delete** /v1/vps/{id}/sshKey/{ssh_key_id} | 
*ManageServiceApi* | [**ManageServiceDisablePostInstallAlert**](docs/ManageServiceApi.md#manageservicedisablepostinstallalert) | **Delete** /v1/vps/{id}/software/post-install-alert | 
*ManageServiceApi* | [**ManageServiceGetAvailableConfiguration**](docs/ManageServiceApi.md#manageservicegetavailableconfiguration) | **Get** /v1/vps/configuration | 
*ManageServiceApi* | [**ManageServiceGetFileManagerSettings**](docs/ManageServiceApi.md#manageservicegetfilemanagersettings) | **Post** /v1/vps/{id}/fm | 
*ManageServiceApi* | [**ManageServiceGetHistory**](docs/ManageServiceApi.md#manageservicegethistory) | **Get** /v1/vps/{id}/history | 
*ManageServiceApi* | [**ManageServiceGetInfo**](docs/ManageServiceApi.md#manageservicegetinfo) | **Get** /v1/vps/server/{id} | 
*ManageServiceApi* | [**ManageServiceGetInstalledSoftware**](docs/ManageServiceApi.md#manageservicegetinstalledsoftware) | **Get** /v1/vps/{id}/software | 
*ManageServiceApi* | [**ManageServiceGetList**](docs/ManageServiceApi.md#manageservicegetlist) | **Get** /v1/vps/server/list | 
*ManageServiceApi* | [**ManageServiceGetRegionList**](docs/ManageServiceApi.md#manageservicegetregionlist) | **Get** /v1/vps/region | 
*ManageServiceApi* | [**ManageServiceGetStatuses**](docs/ManageServiceApi.md#manageservicegetstatuses) | **Get** /v1/vps/server/statuses | 
*ManageServiceApi* | [**ManageServiceRebootVps**](docs/ManageServiceApi.md#manageservicerebootvps) | **Post** /v1/vps/server/{id}/reboot | 
*ManageServiceApi* | [**ManageServiceReinstall**](docs/ManageServiceApi.md#manageservicereinstall) | **Post** /v1/vps/server/{id}/reinstall | 
*ManageServiceApi* | [**ManageServiceRemoveVps**](docs/ManageServiceApi.md#manageserviceremovevps) | **Post** /v1/vps/server/{id}/remove | 
*ManageServiceApi* | [**ManageServiceReserveVpsSubdomain**](docs/ManageServiceApi.md#manageservicereservevpssubdomain) | **Get** /v1/vps/subdomain/reserve | 
*ManageServiceApi* | [**ManageServiceResetPassword**](docs/ManageServiceApi.md#manageserviceresetpassword) | **Put** /v1/vps/{id}/password | 
*ManageServiceApi* | [**ManageServiceResetVps**](docs/ManageServiceApi.md#manageserviceresetvps) | **Post** /v1/vps/server/{id}/reset | 
*ManageServiceApi* | [**ManageServiceStartRescue**](docs/ManageServiceApi.md#manageservicestartrescue) | **Post** /v1/vps/server/{id}/rescue | 
*ManageServiceApi* | [**ManageServiceStartVps**](docs/ManageServiceApi.md#manageservicestartvps) | **Post** /v1/vps/server/{id}/start | 
*ManageServiceApi* | [**ManageServiceStopRescue**](docs/ManageServiceApi.md#manageservicestoprescue) | **Delete** /v1/vps/server/{id}/rescue | 
*ManageServiceApi* | [**ManageServiceStopVps**](docs/ManageServiceApi.md#manageservicestopvps) | **Post** /v1/vps/server/{id}/stop | 
*ManageServiceApi* | [**ManageServiceUnarchive**](docs/ManageServiceApi.md#manageserviceunarchive) | **Delete** /v1/vps/archive/{id} | 
*ManageServiceApi* | [**ManageServiceUpdateInfo**](docs/ManageServiceApi.md#manageserviceupdateinfo) | **Put** /v1/vps/server/{id}/info | 
*MarketplaceServiceApi* | [**MarketplaceServiceGetSoftwareInfo**](docs/MarketplaceServiceApi.md#marketplaceservicegetsoftwareinfo) | **Get** /v1/vps/marketplace/software/{name}/{version} | 
*MarketplaceServiceApi* | [**MarketplaceServiceGetSoftwareList**](docs/MarketplaceServiceApi.md#marketplaceservicegetsoftwarelist) | **Get** /v1/vps/marketplace/software/list | 
*NetworkServiceApi* | [**NetworkServiceCreatePrivateNetwork**](docs/NetworkServiceApi.md#networkservicecreateprivatenetwork) | **Post** /v1/vps/private-network | 
*NetworkServiceApi* | [**NetworkServiceGetNetworkInfo**](docs/NetworkServiceApi.md#networkservicegetnetworkinfo) | **Get** /v1/vps/network | 
*NetworkServiceApi* | [**NetworkServiceOrderIpAddress**](docs/NetworkServiceApi.md#networkserviceorderipaddress) | **Post** /v1/vps/network | 
*NetworkServiceApi* | [**NetworkServiceRemoveIpAddress**](docs/NetworkServiceApi.md#networkserviceremoveipaddress) | **Delete** /v1/vps/network/{ip_address} | 
*NetworkServiceApi* | [**NetworkServiceSuggestPrivateAddress**](docs/NetworkServiceApi.md#networkservicesuggestprivateaddress) | **Post** /v1/vps/private-network/{network_id}/suggested-address | 
*SnapshotServiceApi* | [**SnapshotServiceCreate**](docs/SnapshotServiceApi.md#snapshotservicecreate) | **Post** /v1/vps/snapshot | 
*SnapshotServiceApi* | [**SnapshotServiceCreateCalculator**](docs/SnapshotServiceApi.md#snapshotservicecreatecalculator) | **Post** /v1/vps/snapshot/calculator | 
*SnapshotServiceApi* | [**SnapshotServiceEdit**](docs/SnapshotServiceApi.md#snapshotserviceedit) | **Put** /v1/vps/snapshot/{id} | 
*SnapshotServiceApi* | [**SnapshotServiceGetAll**](docs/SnapshotServiceApi.md#snapshotservicegetall) | **Get** /v1/vps/snapshot | 
*SnapshotServiceApi* | [**SnapshotServiceGetAllRestores**](docs/SnapshotServiceApi.md#snapshotservicegetallrestores) | **Get** /v1/vps/snapshot/restore | 
*SnapshotServiceApi* | [**SnapshotServiceRemove**](docs/SnapshotServiceApi.md#snapshotserviceremove) | **Delete** /v1/vps/snapshot/{id} | 
*SnapshotServiceApi* | [**SnapshotServiceRestore**](docs/SnapshotServiceApi.md#snapshotservicerestore) | **Post** /v1/vps/snapshot/{id}/restore | 
*SoftwareLicenseServiceApi* | [**SoftwareLicenseServiceChangeLicensePlan**](docs/SoftwareLicenseServiceApi.md#softwarelicenseservicechangelicenseplan) | **Patch** /v1/vps/software/license/{vps_id} | 
*SoftwareLicenseServiceApi* | [**SoftwareLicenseServiceGetLicenseInfo**](docs/SoftwareLicenseServiceApi.md#softwarelicenseservicegetlicenseinfo) | **Get** /v1/vps/software/license | 
*SshKeyServiceApi* | [**SshKeyServiceAdd**](docs/SshKeyServiceApi.md#sshkeyserviceadd) | **Post** /v1/vps/sshKey | 
*SshKeyServiceApi* | [**SshKeyServiceGetAll**](docs/SshKeyServiceApi.md#sshkeyservicegetall) | **Get** /v1/vps/sshKey | 
*SshKeyServiceApi* | [**SshKeyServiceRemove**](docs/SshKeyServiceApi.md#sshkeyserviceremove) | **Delete** /v1/vps/sshKey/{id} | 
*StatisticServiceApi* | [**StatisticServiceGetCpu**](docs/StatisticServiceApi.md#statisticservicegetcpu) | **Get** /v1/vps/statistic/cpu/{id} | 
*StatisticServiceApi* | [**StatisticServiceGetCpuDetails**](docs/StatisticServiceApi.md#statisticservicegetcpudetails) | **Get** /v1/vps/statistic/cpu-details/{id} | 
*StatisticServiceApi* | [**StatisticServiceGetDisk**](docs/StatisticServiceApi.md#statisticservicegetdisk) | **Get** /v1/vps/statistic/disk/{id} | 
*StatisticServiceApi* | [**StatisticServiceGetDiskUsage**](docs/StatisticServiceApi.md#statisticservicegetdiskusage) | **Get** /v1/vps/statistic/disk-usage/{id} | 
*StatisticServiceApi* | [**StatisticServiceGetLoadAverage**](docs/StatisticServiceApi.md#statisticservicegetloadaverage) | **Get** /v1/vps/statistic/load-average/{id} | 
*StatisticServiceApi* | [**StatisticServiceGetMemory**](docs/StatisticServiceApi.md#statisticservicegetmemory) | **Get** /v1/vps/statistic/memory/{id} | 
*StatisticServiceApi* | [**StatisticServiceGetNetwork**](docs/StatisticServiceApi.md#statisticservicegetnetwork) | **Get** /v1/vps/statistic/network/{id} | 
*StatisticServiceApi* | [**StatisticServiceGetProcessList**](docs/StatisticServiceApi.md#statisticservicegetprocesslist) | **Get** /v1/vps/statistic/processes/{id} | 


## Documentation For Models

 - [BackupGetAvailableCopiesResponse](docs/BackupGetAvailableCopiesResponse.md)
 - [BackupGetBackupFileListResponse](docs/BackupGetBackupFileListResponse.md)
 - [BackupGetOrdersResponse](docs/BackupGetOrdersResponse.md)
 - [BackupRestoreFileRequest](docs/BackupRestoreFileRequest.md)
 - [BackupRestoreFileResponse](docs/BackupRestoreFileResponse.md)
 - [BackupRestoreFileResponseError](docs/BackupRestoreFileResponseError.md)
 - [BackupRestoreServerRequest](docs/BackupRestoreServerRequest.md)
 - [BackupRestoreServerResponse](docs/BackupRestoreServerResponse.md)
 - [BackupRestoreServerResponseError](docs/BackupRestoreServerResponseError.md)
 - [ConfiguratorConfiguratorSettings](docs/ConfiguratorConfiguratorSettings.md)
 - [ConfiguratorCpuSettings](docs/ConfiguratorCpuSettings.md)
 - [ConfiguratorDiskSettings](docs/ConfiguratorDiskSettings.md)
 - [ConfiguratorGetCalculationResponse](docs/ConfiguratorGetCalculationResponse.md)
 - [ConfiguratorGetCalculationResponseError](docs/ConfiguratorGetCalculationResponseError.md)
 - [ConfiguratorGetCalculationResponseSuccess](docs/ConfiguratorGetCalculationResponseSuccess.md)
 - [ConfiguratorGetConfiguratorInfoResponse](docs/ConfiguratorGetConfiguratorInfoResponse.md)
 - [ConfiguratorMemorySettings](docs/ConfiguratorMemorySettings.md)
 - [ConfiguratorRange](docs/ConfiguratorRange.md)
 - [ManageAttachIpAddressRequest](docs/ManageAttachIpAddressRequest.md)
 - [ManageAttachIpAddressResponse](docs/ManageAttachIpAddressResponse.md)
 - [ManageAttachIpAddressResponseError](docs/ManageAttachIpAddressResponseError.md)
 - [ManageAttachSshKeyResponse](docs/ManageAttachSshKeyResponse.md)
 - [ManageAttachSshKeyResponseError](docs/ManageAttachSshKeyResponseError.md)
 - [ManageAttachToPrivateNetworkRequest](docs/ManageAttachToPrivateNetworkRequest.md)
 - [ManageAttachToPrivateNetworkResponse](docs/ManageAttachToPrivateNetworkResponse.md)
 - [ManageAttachToPrivateNetworkResponseError](docs/ManageAttachToPrivateNetworkResponseError.md)
 - [ManageChangeConfigurationRequest](docs/ManageChangeConfigurationRequest.md)
 - [ManageChangeConfigurationResponse](docs/ManageChangeConfigurationResponse.md)
 - [ManageChangeConfigurationResponseError](docs/ManageChangeConfigurationResponseError.md)
 - [ManageChangeSshAccessRequest](docs/ManageChangeSshAccessRequest.md)
 - [ManageChangeSshAccessResponse](docs/ManageChangeSshAccessResponse.md)
 - [ManageChangeSshAccessResponseError](docs/ManageChangeSshAccessResponseError.md)
 - [ManageCheckSoftwareRequirementsRequest](docs/ManageCheckSoftwareRequirementsRequest.md)
 - [ManageCheckSoftwareRequirementsResponse](docs/ManageCheckSoftwareRequirementsResponse.md)
 - [ManageCheckSoftwareRequirementsResponseError](docs/ManageCheckSoftwareRequirementsResponseError.md)
 - [ManageConfigurationGroup](docs/ManageConfigurationGroup.md)
 - [ManageCreateVpsRequest](docs/ManageCreateVpsRequest.md)
 - [ManageCreateVpsResponse](docs/ManageCreateVpsResponse.md)
 - [ManageCreateVpsResponseError](docs/ManageCreateVpsResponseError.md)
 - [ManageCreateVpsResponseErrorInsufficientFundsError](docs/ManageCreateVpsResponseErrorInsufficientFundsError.md)
 - [ManageCreateVpsResponseErrorSoftwareVariableError](docs/ManageCreateVpsResponseErrorSoftwareVariableError.md)
 - [ManageCreateVpsResponseErrorSoftwareVariableErrorValueError](docs/ManageCreateVpsResponseErrorSoftwareVariableErrorValueError.md)
 - [ManageDetachFromPrivateNetworkResponse](docs/ManageDetachFromPrivateNetworkResponse.md)
 - [ManageDetachFromPrivateNetworkResponseError](docs/ManageDetachFromPrivateNetworkResponseError.md)
 - [ManageDetachIpAddressResponse](docs/ManageDetachIpAddressResponse.md)
 - [ManageDetachIpAddressResponseError](docs/ManageDetachIpAddressResponseError.md)
 - [ManageDetachSshKeyResponse](docs/ManageDetachSshKeyResponse.md)
 - [ManageDetachSshKeyResponseError](docs/ManageDetachSshKeyResponseError.md)
 - [ManageDisablePostInstallAlertResponse](docs/ManageDisablePostInstallAlertResponse.md)
 - [ManageGetAvailableConfigurationResponse](docs/ManageGetAvailableConfigurationResponse.md)
 - [ManageGetFileManagerSettingsResponse](docs/ManageGetFileManagerSettingsResponse.md)
 - [ManageGetFileManagerSettingsResponseCredentials](docs/ManageGetFileManagerSettingsResponseCredentials.md)
 - [ManageGetFileManagerSettingsResponseError](docs/ManageGetFileManagerSettingsResponseError.md)
 - [ManageGetHistoryResponse](docs/ManageGetHistoryResponse.md)
 - [ManageGetInfoResponse](docs/ManageGetInfoResponse.md)
 - [ManageGetInstalledSoftwareResponse](docs/ManageGetInstalledSoftwareResponse.md)
 - [ManageGetListResponse](docs/ManageGetListResponse.md)
 - [ManageGetRegionListResponse](docs/ManageGetRegionListResponse.md)
 - [ManageGetStatusesResponse](docs/ManageGetStatusesResponse.md)
 - [ManageGetStatusesResponseStatusInfo](docs/ManageGetStatusesResponseStatusInfo.md)
 - [ManageHistoryItem](docs/ManageHistoryItem.md)
 - [ManagePrivateNetworkInfo](docs/ManagePrivateNetworkInfo.md)
 - [ManageRebootVpsResponse](docs/ManageRebootVpsResponse.md)
 - [ManageRebootVpsResponseError](docs/ManageRebootVpsResponseError.md)
 - [ManageReinstallRequest](docs/ManageReinstallRequest.md)
 - [ManageReinstallResponse](docs/ManageReinstallResponse.md)
 - [ManageReinstallResponseError](docs/ManageReinstallResponseError.md)
 - [ManageReinstallResponseErrorInsufficientFundsError](docs/ManageReinstallResponseErrorInsufficientFundsError.md)
 - [ManageReinstallResponseErrorSoftwareVariableError](docs/ManageReinstallResponseErrorSoftwareVariableError.md)
 - [ManageReinstallResponseErrorSoftwareVariableErrorValueError](docs/ManageReinstallResponseErrorSoftwareVariableErrorValueError.md)
 - [ManageRemoveVpsResponse](docs/ManageRemoveVpsResponse.md)
 - [ManageRemoveVpsResponseError](docs/ManageRemoveVpsResponseError.md)
 - [ManageReserveVpsSubdomainResponse](docs/ManageReserveVpsSubdomainResponse.md)
 - [ManageReserveVpsSubdomainResponseError](docs/ManageReserveVpsSubdomainResponseError.md)
 - [ManageResetPasswordResponse](docs/ManageResetPasswordResponse.md)
 - [ManageResetPasswordResponseError](docs/ManageResetPasswordResponseError.md)
 - [ManageResetVpsResponse](docs/ManageResetVpsResponse.md)
 - [ManageResetVpsResponseError](docs/ManageResetVpsResponseError.md)
 - [ManageSoftwareInstallInfo](docs/ManageSoftwareInstallInfo.md)
 - [ManageStartRescueResponse](docs/ManageStartRescueResponse.md)
 - [ManageStartRescueResponseError](docs/ManageStartRescueResponseError.md)
 - [ManageStartVpsResponse](docs/ManageStartVpsResponse.md)
 - [ManageStartVpsResponseError](docs/ManageStartVpsResponseError.md)
 - [ManageStopRescueResponse](docs/ManageStopRescueResponse.md)
 - [ManageStopRescueResponseError](docs/ManageStopRescueResponseError.md)
 - [ManageStopVpsResponse](docs/ManageStopVpsResponse.md)
 - [ManageStopVpsResponseError](docs/ManageStopVpsResponseError.md)
 - [ManageUnarchiveResponse](docs/ManageUnarchiveResponse.md)
 - [ManageUpdateInfoRequest](docs/ManageUpdateInfoRequest.md)
 - [ManageUpdateInfoResponse](docs/ManageUpdateInfoResponse.md)
 - [ManageUpdateInfoResponseError](docs/ManageUpdateInfoResponseError.md)
 - [ManageVpsConfiguration](docs/ManageVpsConfiguration.md)
 - [ManageVpsInfo](docs/ManageVpsInfo.md)
 - [MarketplaceDomainField](docs/MarketplaceDomainField.md)
 - [MarketplaceEmailField](docs/MarketplaceEmailField.md)
 - [MarketplaceFieldCommon](docs/MarketplaceFieldCommon.md)
 - [MarketplaceFieldDesc](docs/MarketplaceFieldDesc.md)
 - [MarketplaceGetSoftwareInfoResponse](docs/MarketplaceGetSoftwareInfoResponse.md)
 - [MarketplaceGetSoftwareListResponse](docs/MarketplaceGetSoftwareListResponse.md)
 - [MarketplacePasswordField](docs/MarketplacePasswordField.md)
 - [MarketplaceSoftwareInfo](docs/MarketplaceSoftwareInfo.md)
 - [MarketplaceSoftwareInfoRequirements](docs/MarketplaceSoftwareInfoRequirements.md)
 - [MarketplaceTextField](docs/MarketplaceTextField.md)
 - [NetworkCreatePrivateNetworkRequest](docs/NetworkCreatePrivateNetworkRequest.md)
 - [NetworkCreatePrivateNetworkResponse](docs/NetworkCreatePrivateNetworkResponse.md)
 - [NetworkCreatePrivateNetworkResponseError](docs/NetworkCreatePrivateNetworkResponseError.md)
 - [NetworkGetNetworkInfoResponse](docs/NetworkGetNetworkInfoResponse.md)
 - [NetworkOrderIpAddressRequest](docs/NetworkOrderIpAddressRequest.md)
 - [NetworkOrderIpAddressResponse](docs/NetworkOrderIpAddressResponse.md)
 - [NetworkOrderIpAddressResponseError](docs/NetworkOrderIpAddressResponseError.md)
 - [NetworkRemoveIpAddressResponse](docs/NetworkRemoveIpAddressResponse.md)
 - [NetworkRemoveIpAddressResponseError](docs/NetworkRemoveIpAddressResponseError.md)
 - [NetworkSuggestPrivateAddressRequest](docs/NetworkSuggestPrivateAddressRequest.md)
 - [NetworkSuggestPrivateAddressResponse](docs/NetworkSuggestPrivateAddressResponse.md)
 - [SnapshotCreateCalculatorRequest](docs/SnapshotCreateCalculatorRequest.md)
 - [SnapshotCreateCalculatorResponse](docs/SnapshotCreateCalculatorResponse.md)
 - [SnapshotCreateRequest](docs/SnapshotCreateRequest.md)
 - [SnapshotCreateResponse](docs/SnapshotCreateResponse.md)
 - [SnapshotCreateResponseError](docs/SnapshotCreateResponseError.md)
 - [SnapshotEditRequest](docs/SnapshotEditRequest.md)
 - [SnapshotEditResponse](docs/SnapshotEditResponse.md)
 - [SnapshotGetAllResponse](docs/SnapshotGetAllResponse.md)
 - [SnapshotGetAllRestoresResponse](docs/SnapshotGetAllRestoresResponse.md)
 - [SnapshotRemoveResponse](docs/SnapshotRemoveResponse.md)
 - [SnapshotRemoveResponseError](docs/SnapshotRemoveResponseError.md)
 - [SnapshotRequiredConfiguration](docs/SnapshotRequiredConfiguration.md)
 - [SnapshotRestore](docs/SnapshotRestore.md)
 - [SnapshotRestoreRequest](docs/SnapshotRestoreRequest.md)
 - [SnapshotRestoreResponse](docs/SnapshotRestoreResponse.md)
 - [SnapshotRestoreResponseError](docs/SnapshotRestoreResponseError.md)
 - [SnapshotSnapshot](docs/SnapshotSnapshot.md)
 - [SoftwareLicenseChangeLicensePlanRequest](docs/SoftwareLicenseChangeLicensePlanRequest.md)
 - [SoftwareLicenseChangeLicensePlanResponse](docs/SoftwareLicenseChangeLicensePlanResponse.md)
 - [SoftwareLicenseChangeLicensePlanResponseError](docs/SoftwareLicenseChangeLicensePlanResponseError.md)
 - [SoftwareLicenseChangeLicensePlanResponseErrorInsufficientFundsError](docs/SoftwareLicenseChangeLicensePlanResponseErrorInsufficientFundsError.md)
 - [SoftwareLicenseGetLicenseInfoResponse](docs/SoftwareLicenseGetLicenseInfoResponse.md)
 - [SshKeyAddRequest](docs/SshKeyAddRequest.md)
 - [SshKeyAddResponse](docs/SshKeyAddResponse.md)
 - [SshKeyAddResponseError](docs/SshKeyAddResponseError.md)
 - [SshKeyGetAllResponse](docs/SshKeyGetAllResponse.md)
 - [SshKeyRemoveResponse](docs/SshKeyRemoveResponse.md)
 - [SshKeyRemoveResponseError](docs/SshKeyRemoveResponseError.md)
 - [StatisticGetCpuDetailsResponse](docs/StatisticGetCpuDetailsResponse.md)
 - [StatisticGetCpuResponse](docs/StatisticGetCpuResponse.md)
 - [StatisticGetDiskResponse](docs/StatisticGetDiskResponse.md)
 - [StatisticGetDiskUsageResponse](docs/StatisticGetDiskUsageResponse.md)
 - [StatisticGetLoadAverageResponse](docs/StatisticGetLoadAverageResponse.md)
 - [StatisticGetMemoryResponse](docs/StatisticGetMemoryResponse.md)
 - [StatisticGetNetworkResponse](docs/StatisticGetNetworkResponse.md)
 - [StatisticGetProcessListResponse](docs/StatisticGetProcessListResponse.md)
 - [StatisticGetProcessListResponseError](docs/StatisticGetProcessListResponseError.md)
 - [StatisticGetProcessListResponseProcessList](docs/StatisticGetProcessListResponseProcessList.md)
 - [StatisticGetProcessListResponseProcessListProcessInfo](docs/StatisticGetProcessListResponseProcessListProcessInfo.md)
 - [StatisticSeriesData](docs/StatisticSeriesData.md)
 - [StructuresAdditionalIpInfo](docs/StructuresAdditionalIpInfo.md)
 - [StructuresAttachedPrivateNetwork](docs/StructuresAttachedPrivateNetwork.md)
 - [StructuresConfigurationParams](docs/StructuresConfigurationParams.md)
 - [StructuresCopyInfo](docs/StructuresCopyInfo.md)
 - [StructuresCopyInfoConfiguration](docs/StructuresCopyInfoConfiguration.md)
 - [StructuresFileInfo](docs/StructuresFileInfo.md)
 - [StructuresInstalledSoftwareInfo](docs/StructuresInstalledSoftwareInfo.md)
 - [StructuresInstalledSoftwareInfoFieldValue](docs/StructuresInstalledSoftwareInfoFieldValue.md)
 - [StructuresIpInfo](docs/StructuresIpInfo.md)
 - [StructuresIssuedSoftwareLicense](docs/StructuresIssuedSoftwareLicense.md)
 - [StructuresOrderInfo](docs/StructuresOrderInfo.md)
 - [StructuresOrderInfoErrorDetails](docs/StructuresOrderInfoErrorDetails.md)
 - [StructuresOrderInfoErrorDetailsFileError](docs/StructuresOrderInfoErrorDetailsFileError.md)
 - [StructuresPrivateNetwork](docs/StructuresPrivateNetwork.md)
 - [StructuresRegionInfo](docs/StructuresRegionInfo.md)
 - [StructuresSoftwareCategory](docs/StructuresSoftwareCategory.md)
 - [StructuresSoftwareLicense](docs/StructuresSoftwareLicense.md)
 - [StructuresSoftwareLicenseBillingType](docs/StructuresSoftwareLicenseBillingType.md)
 - [StructuresSoftwareLicenseBillingTypeDaily](docs/StructuresSoftwareLicenseBillingTypeDaily.md)
 - [StructuresSoftwareLicenseBillingTypeMonthly](docs/StructuresSoftwareLicenseBillingTypeMonthly.md)
 - [StructuresSoftwareMetadata](docs/StructuresSoftwareMetadata.md)
 - [StructuresSshKeyInfo](docs/StructuresSshKeyInfo.md)


## Documentation For Authorization



### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



