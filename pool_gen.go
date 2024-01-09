//
// This file is generated. To change the content of this file, please do not
// apply the change to this file because it will get overwritten. Instead,
// change xenapi.go and execute 'go generate'.
//

package xenapi

import (
	"fmt"
	"reflect"
	"strconv"
	
	"time"
)

var _ = fmt.Errorf
var _ = reflect.TypeOf
var _ = strconv.Atoi
var _ = time.UTC

type PoolAllowedOperations string

const (
	// Indicates this pool is in the process of enabling HA
	PoolAllowedOperationsHaEnable PoolAllowedOperations = "ha_enable"
	// Indicates this pool is in the process of disabling HA
	PoolAllowedOperationsHaDisable PoolAllowedOperations = "ha_disable"
	// Indicates this pool is in the process of creating a cluster
	PoolAllowedOperationsClusterCreate PoolAllowedOperations = "cluster_create"
	// Indicates this pool is in the process of changing master
	PoolAllowedOperationsDesignateNewMaster PoolAllowedOperations = "designate_new_master"
	// Indicates this pool is in the process of configuring repositories
	PoolAllowedOperationsConfigureRepositories PoolAllowedOperations = "configure_repositories"
	// Indicates this pool is in the process of syncing updates
	PoolAllowedOperationsSyncUpdates PoolAllowedOperations = "sync_updates"
	// Indicates this pool is in the process of getting updates
	PoolAllowedOperationsGetUpdates PoolAllowedOperations = "get_updates"
	// Indicates this pool is in the process of applying updates
	PoolAllowedOperationsApplyUpdates PoolAllowedOperations = "apply_updates"
	// Indicates this pool is in the process of enabling TLS verification
	PoolAllowedOperationsTLSVerificationEnable PoolAllowedOperations = "tls_verification_enable"
	// A certificate refresh and distribution is in progress
	PoolAllowedOperationsCertRefresh PoolAllowedOperations = "cert_refresh"
	// Indicates this pool is exchanging internal certificates with a new joiner
	PoolAllowedOperationsExchangeCertificatesOnJoin PoolAllowedOperations = "exchange_certificates_on_join"
	// Indicates this pool is exchanging ca certificates with a new joiner
	PoolAllowedOperationsExchangeCaCertificatesOnJoin PoolAllowedOperations = "exchange_ca_certificates_on_join"
	// Indicates the primary host is sending its certificates to another host
	PoolAllowedOperationsCopyPrimaryHostCerts PoolAllowedOperations = "copy_primary_host_certs"
)

type TelemetryFrequency string

const (
	// Run telemetry task daily
	TelemetryFrequencyDaily TelemetryFrequency = "daily"
	// Run telemetry task weekly
	TelemetryFrequencyWeekly TelemetryFrequency = "weekly"
	// Run telemetry task monthly
	TelemetryFrequencyMonthly TelemetryFrequency = "monthly"
)

type PoolRecord struct {
	// Unique identifier/object reference
	UUID string
	// Short name
	NameLabel string
	// Description
	NameDescription string
	// The host that is pool master
	Master HostRef
	// Default SR for VDIs
	DefaultSR SRRef
	// The SR in which VDIs for suspend images are created
	SuspendImageSR SRRef
	// The SR in which VDIs for crash dumps are created
	CrashDumpSR SRRef
	// additional configuration
	OtherConfig map[string]string
	// true if HA is enabled on the pool, false otherwise
	HaEnabled bool
	// The current HA configuration
	HaConfiguration map[string]string
	// HA statefile VDIs in use
	HaStatefiles []string
	// Number of host failures to tolerate before the Pool is declared to be overcommitted
	HaHostFailuresToTolerate int
	// Number of future host failures we have managed to find a plan for. Once this reaches zero any future host failures will cause the failure of protected VMs.
	HaPlanExistsFor int
	// If set to false then operations which would cause the Pool to become overcommitted will be blocked.
	HaAllowOvercommit bool
	// True if the Pool is considered to be overcommitted i.e. if there exist insufficient physical resources to tolerate the configured number of host failures
	HaOvercommitted bool
	// Binary blobs associated with this pool
	Blobs map[string]BlobRef
	// user-specified tags for categorization purposes
	Tags []string
	// gui-specific configuration for pool
	GuiConfig map[string]string
	// Configuration for the automatic health check feature
	HealthCheckConfig map[string]string
	// Url for the configured workload balancing host
	WlbURL string
	// Username for accessing the workload balancing host
	WlbUsername string
	// true if workload balancing is enabled on the pool, false otherwise
	WlbEnabled bool
	// true if communication with the WLB server should enforce TLS certificate verification.
	WlbVerifyCert bool
	// true a redo-log is to be used other than when HA is enabled, false otherwise
	RedoLogEnabled bool
	// indicates the VDI to use for the redo-log other than when HA is enabled
	RedoLogVdi VDIRef
	// address of the vswitch controller
	VswitchController string
	// Pool-wide restrictions currently in effect
	Restrictions map[string]string
	// The set of currently known metadata VDIs for this pool
	MetadataVDIs []VDIRef
	// The HA cluster stack that is currently in use. Only valid when HA is enabled.
	HaClusterStack string
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []PoolAllowedOperations
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]PoolAllowedOperations
	// Pool-wide guest agent configuration information
	GuestAgentConfig map[string]string
	// Details about the physical CPUs on the pool
	CPUInfo map[string]string
	// The pool-wide policy for clients on whether to use the vendor device or not on newly created VMs. This field will also be consulted if the 'has_vendor_device' field is not specified in the VM.create call.
	PolicyNoVendorDevice bool
	// The pool-wide flag to show if the live patching feauture is disabled or not.
	LivePatchingDisabled bool
	// true if IGMP snooping is enabled in the pool, false otherwise.
	IgmpSnoopingEnabled bool
	// The UEFI certificates allowing Secure Boot
	UefiCertificates string
	// True if either a PSR is running or we are waiting for a PSR to be re-run
	IsPsrPending bool
	// True iff TLS certificate verification is enabled
	TLSVerificationEnabled bool
	// The set of currently enabled repositories
	Repositories []RepositoryRef
	// True if authentication by TLS client certificates is enabled
	ClientCertificateAuthEnabled bool
	// The name (CN/SAN) that an incoming client certificate must have to allow authentication
	ClientCertificateAuthName string
	// Url of the proxy used in syncing with the enabled repositories
	RepositoryProxyURL string
	// Username for the authentication of the proxy used in syncing with the enabled repositories
	RepositoryProxyUsername string
	// Password for the authentication of the proxy used in syncing with the enabled repositories
	RepositoryProxyPassword SecretRef
	// Default behaviour during migration, True if stream compression should be used
	MigrationCompression bool
	// true if bias against pool master when scheduling vms is enabled, false otherwise
	CoordinatorBias bool
	// The UUID of the pool for identification of telemetry data
	TelemetryUUID SecretRef
	// How often the telemetry collection will be carried out
	TelemetryFrequency TelemetryFrequency
	// The earliest timestamp (in UTC) when the next round of telemetry collection can be carried out
	TelemetryNextCollection time.Time
}

type PoolRef string

// Pool-wide information
type PoolClass struct {
	client *Client
}

// GetAllRecords Return a map of pool references to pool records for all pools known to the system.
func (_class PoolClass) GetAllRecords(sessionID SessionRef) (_retval map[PoolRef]PoolRecord, _err error) {
	_method := "pool.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolRefToPoolRecordMapToGo(_method + " -> ", _result)
	return
}

// GetAll Return a list of all the pools known to the system.
func (_class PoolClass) GetAll(sessionID SessionRef) (_retval []PoolRef, _err error) {
	_method := "pool.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolRefSetToGo(_method + " -> ", _result)
	return
}

// ResetTelemetryUUID Assign a new UUID to telemetry data.
func (_class PoolClass) ResetTelemetryUUID(sessionID SessionRef, self PoolRef) (_err error) {
	_method := "pool.reset_telemetry_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// SetTelemetryNextCollection Set the timestamp for the next telemetry data collection.
func (_class PoolClass) SetTelemetryNextCollection(sessionID SessionRef, self PoolRef, value time.Time) (_err error) {
	_method := "pool.set_telemetry_next_collection"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertTimeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetHTTPSOnly updates all the host firewalls in the pool to open or close port 80 depending on the value
func (_class PoolClass) SetHTTPSOnly(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	_method := "pool.set_https_only"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetUefiCertificates Sets the UEFI certificates for a pool and all its hosts
func (_class PoolClass) SetUefiCertificates(sessionID SessionRef, self PoolRef, value string) (_err error) {
	_method := "pool.set_uefi_certificates"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// DisableRepositoryProxy Disable the proxy for RPM package repositories.
func (_class PoolClass) DisableRepositoryProxy(sessionID SessionRef, self PoolRef) (_err error) {
	_method := "pool.disable_repository_proxy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// ConfigureRepositoryProxy Configure proxy for RPM package repositories.
func (_class PoolClass) ConfigureRepositoryProxy(sessionID SessionRef, self PoolRef, url string, username string, password string) (_err error) {
	_method := "pool.configure_repository_proxy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_urlArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "url"), url)
	if _err != nil {
		return
	}
	_usernameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "username"), username)
	if _err != nil {
		return
	}
	_passwordArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "password"), password)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _urlArg, _usernameArg, _passwordArg)
	return
}

// DisableClientCertificateAuth Disable client certificate authentication on the pool
func (_class PoolClass) DisableClientCertificateAuth(sessionID SessionRef, self PoolRef) (_err error) {
	_method := "pool.disable_client_certificate_auth"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// EnableClientCertificateAuth Enable client certificate authentication on the pool
func (_class PoolClass) EnableClientCertificateAuth(sessionID SessionRef, self PoolRef, name string) (_err error) {
	_method := "pool.enable_client_certificate_auth"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _nameArg)
	return
}

// CheckUpdateReadiness Check if the pool is ready to be updated. If not, report the reasons.
func (_class PoolClass) CheckUpdateReadiness(sessionID SessionRef, self PoolRef, requiresReboot bool) (_retval [][]string, _err error) {
	_method := "pool.check_update_readiness"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_requiresRebootArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "requires_reboot"), requiresReboot)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _requiresRebootArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetSetToGo(_method + " -> ", _result)
	return
}

// SyncUpdates Sync with the enabled repository
func (_class PoolClass) SyncUpdates(sessionID SessionRef, self PoolRef, force bool, token string, tokenID string) (_retval string, _err error) {
	_method := "pool.sync_updates"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_tokenArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "token"), token)
	if _err != nil {
		return
	}
	_tokenIDArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "token_id"), tokenID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _forceArg, _tokenArg, _tokenIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result)
	return
}

// RemoveRepository Remove a repository from the enabled set
func (_class PoolClass) RemoveRepository(sessionID SessionRef, self PoolRef, value RepositoryRef) (_err error) {
	_method := "pool.remove_repository"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertRepositoryRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// AddRepository Add a repository to the enabled set
func (_class PoolClass) AddRepository(sessionID SessionRef, self PoolRef, value RepositoryRef) (_err error) {
	_method := "pool.add_repository"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertRepositoryRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetRepositories Set enabled set of repositories
func (_class PoolClass) SetRepositories(sessionID SessionRef, self PoolRef, value []RepositoryRef) (_err error) {
	_method := "pool.set_repositories"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertRepositoryRefSetToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// RotateSecret 
//
// Errors:
//  INTERNAL_ERROR - The server failed to handle your request, due to an internal error. The given message may give details useful for debugging the problem.
//  HOST_IS_SLAVE - You cannot make regular API calls directly on a supporter. Please pass API calls via the coordinator host.
//  CANNOT_CONTACT_HOST - Cannot forward messages because the server cannot be contacted. The server may be switched off or there may be network connectivity problems.
//  HA_IS_ENABLED - The operation could not be performed because HA is enabled on the Pool
//  NOT_SUPPORTED_DURING_UPGRADE - This operation is not supported during an upgrade.
func (_class PoolClass) RotateSecret(sessionID SessionRef) (_err error) {
	_method := "pool.rotate_secret"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

// RemoveFromGuestAgentConfig Remove a key-value pair from the pool-wide guest agent configuration
func (_class PoolClass) RemoveFromGuestAgentConfig(sessionID SessionRef, self PoolRef, key string) (_err error) {
	_method := "pool.remove_from_guest_agent_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}

// AddToGuestAgentConfig Add a key-value pair to the pool-wide guest agent configuration
func (_class PoolClass) AddToGuestAgentConfig(sessionID SessionRef, self PoolRef, key string, value string) (_err error) {
	_method := "pool.add_to_guest_agent_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}

// HasExtension Return true if the extension is available on the pool
func (_class PoolClass) HasExtension(sessionID SessionRef, self PoolRef, name string) (_retval bool, _err error) {
	_method := "pool.has_extension"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _nameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result)
	return
}

// SetIgmpSnoopingEnabled Enable or disable IGMP Snooping on the pool.
func (_class PoolClass) SetIgmpSnoopingEnabled(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	_method := "pool.set_igmp_snooping_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// DisableSslLegacy Sets ssl_legacy false on each host, pool-master last. See Host.ssl_legacy and Host.set_ssl_legacy.
func (_class PoolClass) DisableSslLegacy(sessionID SessionRef, self PoolRef) (_err error) {
	_method := "pool.disable_ssl_legacy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// EnableSslLegacy Sets ssl_legacy true on each host, pool-master last. See Host.ssl_legacy and Host.set_ssl_legacy.
func (_class PoolClass) EnableSslLegacy(sessionID SessionRef, self PoolRef) (_err error) {
	_method := "pool.enable_ssl_legacy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// ApplyEdition Apply an edition to all hosts in the pool
func (_class PoolClass) ApplyEdition(sessionID SessionRef, self PoolRef, edition string) (_err error) {
	_method := "pool.apply_edition"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_editionArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "edition"), edition)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _editionArg)
	return
}

// GetLicenseState This call returns the license state for the pool
func (_class PoolClass) GetLicenseState(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	_method := "pool.get_license_state"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result)
	return
}

// DisableLocalStorageCaching This call disables pool-wide local storage caching
func (_class PoolClass) DisableLocalStorageCaching(sessionID SessionRef, self PoolRef) (_err error) {
	_method := "pool.disable_local_storage_caching"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// EnableLocalStorageCaching This call attempts to enable pool-wide local storage caching
func (_class PoolClass) EnableLocalStorageCaching(sessionID SessionRef, self PoolRef) (_err error) {
	_method := "pool.enable_local_storage_caching"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// TestArchiveTarget This call tests if a location is valid
func (_class PoolClass) TestArchiveTarget(sessionID SessionRef, self PoolRef, config map[string]string) (_retval string, _err error) {
	_method := "pool.test_archive_target"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_configArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "config"), config)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _configArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result)
	return
}

// SetVswitchController Set the IP address of the vswitch controller.
func (_class PoolClass) SetVswitchController(sessionID SessionRef, address string) (_err error) {
	_method := "pool.set_vswitch_controller"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_addressArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "address"), address)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _addressArg)
	return
}

// DisableRedoLog Disable the redo log if in use, unless HA is enabled.
func (_class PoolClass) DisableRedoLog(sessionID SessionRef) (_err error) {
	_method := "pool.disable_redo_log"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

// EnableRedoLog Enable the redo log on the given SR and start using it, unless HA is enabled.
func (_class PoolClass) EnableRedoLog(sessionID SessionRef, sr SRRef) (_err error) {
	_method := "pool.enable_redo_log"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _srArg)
	return
}

// EnableTLSVerification Enable TLS server certificate verification
func (_class PoolClass) EnableTLSVerification(sessionID SessionRef) (_err error) {
	_method := "pool.enable_tls_verification"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

// CertificateSync Copy the TLS CA certificates and CRLs of the master to all slaves.
func (_class PoolClass) CertificateSync(sessionID SessionRef) (_err error) {
	_method := "pool.certificate_sync"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

// CrlList List the names of all installed TLS CA-issued Certificate Revocation Lists.
func (_class PoolClass) CrlList(sessionID SessionRef) (_retval []string, _err error) {
	_method := "pool.crl_list"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result)
	return
}

// CrlUninstall Remove a pool-wide TLS CA-issued Certificate Revocation List.
func (_class PoolClass) CrlUninstall(sessionID SessionRef, name string) (_err error) {
	_method := "pool.crl_uninstall"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _nameArg)
	return
}

// CrlInstall Install a TLS CA-issued Certificate Revocation List, pool-wide.
func (_class PoolClass) CrlInstall(sessionID SessionRef, name string, cert string) (_err error) {
	_method := "pool.crl_install"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_certArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "cert"), cert)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _nameArg, _certArg)
	return
}

// UninstallCaCertificate Remove a pool-wide TLS CA certificate.
func (_class PoolClass) UninstallCaCertificate(sessionID SessionRef, name string) (_err error) {
	_method := "pool.uninstall_ca_certificate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _nameArg)
	return
}

// InstallCaCertificate Install a TLS CA certificate, pool-wide.
func (_class PoolClass) InstallCaCertificate(sessionID SessionRef, name string, cert string) (_err error) {
	_method := "pool.install_ca_certificate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_certArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "cert"), cert)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _nameArg, _certArg)
	return
}

// CertificateList List the names of all installed TLS CA certificates.
func (_class PoolClass) CertificateList(sessionID SessionRef) (_retval []string, _err error) {
	_method := "pool.certificate_list"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result)
	return
}

// CertificateUninstall Remove a pool-wide TLS CA certificate.
func (_class PoolClass) CertificateUninstall(sessionID SessionRef, name string) (_err error) {
	_method := "pool.certificate_uninstall"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _nameArg)
	return
}

// CertificateInstall Install a TLS CA certificate, pool-wide.
func (_class PoolClass) CertificateInstall(sessionID SessionRef, name string, cert string) (_err error) {
	_method := "pool.certificate_install"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_certArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "cert"), cert)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _nameArg, _certArg)
	return
}

// SendTestPost Send the given body to the given host and port, using HTTPS, and print the response.  This is used for debugging the SSL layer.
func (_class PoolClass) SendTestPost(sessionID SessionRef, host string, port int, body string) (_retval string, _err error) {
	_method := "pool.send_test_post"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_portArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "port"), port)
	if _err != nil {
		return
	}
	_bodyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "body"), body)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg, _portArg, _bodyArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result)
	return
}

// RetrieveWlbRecommendations Retrieves vm migrate recommendations for the pool from the workload balancing server
func (_class PoolClass) RetrieveWlbRecommendations(sessionID SessionRef) (_retval map[VMRef][]string, _err error) {
	_method := "pool.retrieve_wlb_recommendations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToStringSetMapToGo(_method + " -> ", _result)
	return
}

// RetrieveWlbConfiguration Retrieves the pool optimization criteria from the workload balancing server
func (_class PoolClass) RetrieveWlbConfiguration(sessionID SessionRef) (_retval map[string]string, _err error) {
	_method := "pool.retrieve_wlb_configuration"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result)
	return
}

// SendWlbConfiguration Sets the pool optimization criteria for the workload balancing server
func (_class PoolClass) SendWlbConfiguration(sessionID SessionRef, config map[string]string) (_err error) {
	_method := "pool.send_wlb_configuration"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_configArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "config"), config)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _configArg)
	return
}

// DeconfigureWlb Permanently deconfigures workload balancing monitoring on this pool
func (_class PoolClass) DeconfigureWlb(sessionID SessionRef) (_err error) {
	_method := "pool.deconfigure_wlb"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

// InitializeWlb Initializes workload balancing monitoring on this pool with the specified wlb server
func (_class PoolClass) InitializeWlb(sessionID SessionRef, wlbURL string, wlbUsername string, wlbPassword string, xenserverUsername string, xenserverPassword string) (_err error) {
	_method := "pool.initialize_wlb"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_wlbURLArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "wlb_url"), wlbURL)
	if _err != nil {
		return
	}
	_wlbUsernameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "wlb_username"), wlbUsername)
	if _err != nil {
		return
	}
	_wlbPasswordArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "wlb_password"), wlbPassword)
	if _err != nil {
		return
	}
	_xenserverUsernameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "xenserver_username"), xenserverUsername)
	if _err != nil {
		return
	}
	_xenserverPasswordArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "xenserver_password"), xenserverPassword)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _wlbURLArg, _wlbUsernameArg, _wlbPasswordArg, _xenserverUsernameArg, _xenserverPasswordArg)
	return
}

// DetectNonhomogeneousExternalAuth This call asynchronously detects if the external authentication configuration in any slave is different from that in the master and raises appropriate alerts
func (_class PoolClass) DetectNonhomogeneousExternalAuth(sessionID SessionRef, pool PoolRef) (_err error) {
	_method := "pool.detect_nonhomogeneous_external_auth"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_poolArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "pool"), pool)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _poolArg)
	return
}

// DisableExternalAuth This call disables external authentication on all the hosts of the pool
func (_class PoolClass) DisableExternalAuth(sessionID SessionRef, pool PoolRef, config map[string]string) (_err error) {
	_method := "pool.disable_external_auth"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_poolArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "pool"), pool)
	if _err != nil {
		return
	}
	_configArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "config"), config)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _poolArg, _configArg)
	return
}

// EnableExternalAuth This call enables external authentication on all the hosts of the pool
func (_class PoolClass) EnableExternalAuth(sessionID SessionRef, pool PoolRef, config map[string]string, serviceName string, authType string) (_err error) {
	_method := "pool.enable_external_auth"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_poolArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "pool"), pool)
	if _err != nil {
		return
	}
	_configArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "config"), config)
	if _err != nil {
		return
	}
	_serviceNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "service_name"), serviceName)
	if _err != nil {
		return
	}
	_authTypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "auth_type"), authType)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _poolArg, _configArg, _serviceNameArg, _authTypeArg)
	return
}

// CreateNewBlob Create a placeholder for a named binary blob of data that is associated with this pool
func (_class PoolClass) CreateNewBlob(sessionID SessionRef, pool PoolRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	_method := "pool.create_new_blob"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_poolArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "pool"), pool)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_mimeTypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "mime_type"), mimeType)
	if _err != nil {
		return
	}
	_publicArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "public"), public)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _poolArg, _nameArg, _mimeTypeArg, _publicArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBlobRefToGo(_method + " -> ", _result)
	return
}

// SetHaHostFailuresToTolerate Set the maximum number of host failures to consider in the HA VM restart planner
func (_class PoolClass) SetHaHostFailuresToTolerate(sessionID SessionRef, self PoolRef, value int) (_err error) {
	_method := "pool.set_ha_host_failures_to_tolerate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// HaComputeVMFailoverPlan Return a VM failover plan assuming a given subset of hosts fail
func (_class PoolClass) HaComputeVMFailoverPlan(sessionID SessionRef, failedHosts []HostRef, failedVms []VMRef) (_retval map[VMRef]map[string]string, _err error) {
	_method := "pool.ha_compute_vm_failover_plan"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_failedHostsArg, _err := convertHostRefSetToXen(fmt.Sprintf("%s(%s)", _method, "failed_hosts"), failedHosts)
	if _err != nil {
		return
	}
	_failedVmsArg, _err := convertVMRefSetToXen(fmt.Sprintf("%s(%s)", _method, "failed_vms"), failedVms)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _failedHostsArg, _failedVmsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToStringToStringMapMapToGo(_method + " -> ", _result)
	return
}

// HaComputeHypotheticalMaxHostFailuresToTolerate Returns the maximum number of host failures we could tolerate before we would be unable to restart the provided VMs
func (_class PoolClass) HaComputeHypotheticalMaxHostFailuresToTolerate(sessionID SessionRef, configuration map[VMRef]string) (_retval int, _err error) {
	_method := "pool.ha_compute_hypothetical_max_host_failures_to_tolerate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_configurationArg, _err := convertVMRefToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "configuration"), configuration)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _configurationArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result)
	return
}

// HaComputeMaxHostFailuresToTolerate Returns the maximum number of host failures we could tolerate before we would be unable to restart configured VMs
func (_class PoolClass) HaComputeMaxHostFailuresToTolerate(sessionID SessionRef) (_retval int, _err error) {
	_method := "pool.ha_compute_max_host_failures_to_tolerate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result)
	return
}

// HaFailoverPlanExists Returns true if a VM failover plan exists for up to 'n' host failures
func (_class PoolClass) HaFailoverPlanExists(sessionID SessionRef, n int) (_retval bool, _err error) {
	_method := "pool.ha_failover_plan_exists"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "n"), n)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _nArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result)
	return
}

// HaPreventRestartsFor When this call returns the VM restart logic will not run for the requested number of seconds. If the argument is zero then the restart thread is immediately unblocked
func (_class PoolClass) HaPreventRestartsFor(sessionID SessionRef, seconds int) (_err error) {
	_method := "pool.ha_prevent_restarts_for"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_secondsArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "seconds"), seconds)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _secondsArg)
	return
}

// DesignateNewMaster Perform an orderly handover of the role of master to the referenced host.
func (_class PoolClass) DesignateNewMaster(sessionID SessionRef, host HostRef) (_err error) {
	_method := "pool.designate_new_master"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}

// SyncDatabase Forcibly synchronise the database now
func (_class PoolClass) SyncDatabase(sessionID SessionRef) (_err error) {
	_method := "pool.sync_database"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

// DisableHa Turn off High Availability mode
func (_class PoolClass) DisableHa(sessionID SessionRef) (_err error) {
	_method := "pool.disable_ha"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

// EnableHa Turn on High Availability mode
func (_class PoolClass) EnableHa(sessionID SessionRef, heartbeatSrs []SRRef, configuration map[string]string) (_err error) {
	_method := "pool.enable_ha"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_heartbeatSrsArg, _err := convertSRRefSetToXen(fmt.Sprintf("%s(%s)", _method, "heartbeat_srs"), heartbeatSrs)
	if _err != nil {
		return
	}
	_configurationArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "configuration"), configuration)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _heartbeatSrsArg, _configurationArg)
	return
}

// CreateVLANFromPIF Create a pool-wide VLAN by taking the PIF.
//
// Errors:
//  VLAN_TAG_INVALID - You tried to create a VLAN, but the tag you gave was invalid -- it must be between 0 and 4094. The parameter echoes the VLAN tag you gave.
func (_class PoolClass) CreateVLANFromPIF(sessionID SessionRef, pif PIFRef, network NetworkRef, vlan int) (_retval []PIFRef, _err error) {
	_method := "pool.create_VLAN_from_PIF"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_pifArg, _err := convertPIFRefToXen(fmt.Sprintf("%s(%s)", _method, "pif"), pif)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_vlanArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "VLAN"), vlan)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _pifArg, _networkArg, _vlanArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefSetToGo(_method + " -> ", _result)
	return
}

// ManagementReconfigure Reconfigure the management network interface for all Hosts in the Pool
//
// Errors:
//  HA_IS_ENABLED - The operation could not be performed because HA is enabled on the Pool
//  PIF_NOT_PRESENT - This host has no PIF on the given network.
//  CANNOT_PLUG_BOND_SLAVE - This PIF is a bond member and cannot be plugged.
//  PIF_INCOMPATIBLE_PRIMARY_ADDRESS_TYPE - The primary address types are not compatible
//  PIF_HAS_NO_NETWORK_CONFIGURATION - PIF has no IP configuration (mode currently set to 'none')
//  PIF_HAS_NO_V6_NETWORK_CONFIGURATION - PIF has no IPv6 configuration (mode currently set to 'none')
func (_class PoolClass) ManagementReconfigure(sessionID SessionRef, network NetworkRef) (_err error) {
	_method := "pool.management_reconfigure"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _networkArg)
	return
}

// CreateVLAN Create PIFs, mapping a network to the same physical interface/VLAN on each host. This call is deprecated: use Pool.create_VLAN_from_PIF instead.
//
// Errors:
//  VLAN_TAG_INVALID - You tried to create a VLAN, but the tag you gave was invalid -- it must be between 0 and 4094. The parameter echoes the VLAN tag you gave.
func (_class PoolClass) CreateVLAN(sessionID SessionRef, device string, network NetworkRef, vlan int) (_retval []PIFRef, _err error) {
	_method := "pool.create_VLAN"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_deviceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "device"), device)
	if _err != nil {
		return
	}
	_networkArg, _err := convertNetworkRefToXen(fmt.Sprintf("%s(%s)", _method, "network"), network)
	if _err != nil {
		return
	}
	_vlanArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "VLAN"), vlan)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _deviceArg, _networkArg, _vlanArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPIFRefSetToGo(_method + " -> ", _result)
	return
}

// RecoverSlaves Instruct a pool master, M, to try and contact its slaves and, if slaves are in emergency mode, reset their master address to M.
func (_class PoolClass) RecoverSlaves(sessionID SessionRef) (_retval []HostRef, _err error) {
	_method := "pool.recover_slaves"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefSetToGo(_method + " -> ", _result)
	return
}

// EmergencyResetMaster Instruct a slave already in a pool that the master has changed
func (_class PoolClass) EmergencyResetMaster(sessionID SessionRef, masterAddress string) (_err error) {
	_method := "pool.emergency_reset_master"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_masterAddressArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_address"), masterAddress)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _masterAddressArg)
	return
}

// EmergencyTransitionToMaster Instruct host that's currently a slave to transition to being master
func (_class PoolClass) EmergencyTransitionToMaster(sessionID SessionRef) (_err error) {
	_method := "pool.emergency_transition_to_master"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

// Eject Instruct a pool master to eject a host from the pool
func (_class PoolClass) Eject(sessionID SessionRef, host HostRef) (_err error) {
	_method := "pool.eject"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _hostArg)
	return
}

// JoinForce Instruct host to join a new pool
func (_class PoolClass) JoinForce(sessionID SessionRef, masterAddress string, masterUsername string, masterPassword string) (_err error) {
	_method := "pool.join_force"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_masterAddressArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_address"), masterAddress)
	if _err != nil {
		return
	}
	_masterUsernameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_username"), masterUsername)
	if _err != nil {
		return
	}
	_masterPasswordArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_password"), masterPassword)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _masterAddressArg, _masterUsernameArg, _masterPasswordArg)
	return
}

// Join Instruct host to join a new pool
//
// Errors:
//  JOINING_HOST_CANNOT_CONTAIN_SHARED_SRS - The server joining the pool cannot contain any shared storage.
func (_class PoolClass) Join(sessionID SessionRef, masterAddress string, masterUsername string, masterPassword string) (_err error) {
	_method := "pool.join"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_masterAddressArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_address"), masterAddress)
	if _err != nil {
		return
	}
	_masterUsernameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_username"), masterUsername)
	if _err != nil {
		return
	}
	_masterPasswordArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "master_password"), masterPassword)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _masterAddressArg, _masterUsernameArg, _masterPasswordArg)
	return
}

// SetCoordinatorBias Set the coordinator_bias field of the given pool.
func (_class PoolClass) SetCoordinatorBias(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	_method := "pool.set_coordinator_bias"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetMigrationCompression Set the migration_compression field of the given pool.
func (_class PoolClass) SetMigrationCompression(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	_method := "pool.set_migration_compression"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetIsPsrPending Set the is_psr_pending field of the given pool.
func (_class PoolClass) SetIsPsrPending(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	_method := "pool.set_is_psr_pending"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetLivePatchingDisabled Set the live_patching_disabled field of the given pool.
func (_class PoolClass) SetLivePatchingDisabled(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	_method := "pool.set_live_patching_disabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetPolicyNoVendorDevice Set the policy_no_vendor_device field of the given pool.
func (_class PoolClass) SetPolicyNoVendorDevice(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	_method := "pool.set_policy_no_vendor_device"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetWlbVerifyCert Set the wlb_verify_cert field of the given pool.
func (_class PoolClass) SetWlbVerifyCert(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	_method := "pool.set_wlb_verify_cert"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetWlbEnabled Set the wlb_enabled field of the given pool.
func (_class PoolClass) SetWlbEnabled(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	_method := "pool.set_wlb_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// RemoveFromHealthCheckConfig Remove the given key and its corresponding value from the health_check_config field of the given pool.  If the key is not in that Map, then do nothing.
func (_class PoolClass) RemoveFromHealthCheckConfig(sessionID SessionRef, self PoolRef, key string) (_err error) {
	_method := "pool.remove_from_health_check_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}

// AddToHealthCheckConfig Add the given key-value pair to the health_check_config field of the given pool.
func (_class PoolClass) AddToHealthCheckConfig(sessionID SessionRef, self PoolRef, key string, value string) (_err error) {
	_method := "pool.add_to_health_check_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}

// SetHealthCheckConfig Set the health_check_config field of the given pool.
func (_class PoolClass) SetHealthCheckConfig(sessionID SessionRef, self PoolRef, value map[string]string) (_err error) {
	_method := "pool.set_health_check_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// RemoveFromGuiConfig Remove the given key and its corresponding value from the gui_config field of the given pool.  If the key is not in that Map, then do nothing.
func (_class PoolClass) RemoveFromGuiConfig(sessionID SessionRef, self PoolRef, key string) (_err error) {
	_method := "pool.remove_from_gui_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}

// AddToGuiConfig Add the given key-value pair to the gui_config field of the given pool.
func (_class PoolClass) AddToGuiConfig(sessionID SessionRef, self PoolRef, key string, value string) (_err error) {
	_method := "pool.add_to_gui_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}

// SetGuiConfig Set the gui_config field of the given pool.
func (_class PoolClass) SetGuiConfig(sessionID SessionRef, self PoolRef, value map[string]string) (_err error) {
	_method := "pool.set_gui_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// RemoveTags Remove the given value from the tags field of the given pool.  If the value is not in that Set, then do nothing.
func (_class PoolClass) RemoveTags(sessionID SessionRef, self PoolRef, value string) (_err error) {
	_method := "pool.remove_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// AddTags Add the given value to the tags field of the given pool.  If the value is already in that Set, then do nothing.
func (_class PoolClass) AddTags(sessionID SessionRef, self PoolRef, value string) (_err error) {
	_method := "pool.add_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetTags Set the tags field of the given pool.
func (_class PoolClass) SetTags(sessionID SessionRef, self PoolRef, value []string) (_err error) {
	_method := "pool.set_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetHaAllowOvercommit Set the ha_allow_overcommit field of the given pool.
func (_class PoolClass) SetHaAllowOvercommit(sessionID SessionRef, self PoolRef, value bool) (_err error) {
	_method := "pool.set_ha_allow_overcommit"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// RemoveFromOtherConfig Remove the given key and its corresponding value from the other_config field of the given pool.  If the key is not in that Map, then do nothing.
func (_class PoolClass) RemoveFromOtherConfig(sessionID SessionRef, self PoolRef, key string) (_err error) {
	_method := "pool.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}

// AddToOtherConfig Add the given key-value pair to the other_config field of the given pool.
func (_class PoolClass) AddToOtherConfig(sessionID SessionRef, self PoolRef, key string, value string) (_err error) {
	_method := "pool.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}

// SetOtherConfig Set the other_config field of the given pool.
func (_class PoolClass) SetOtherConfig(sessionID SessionRef, self PoolRef, value map[string]string) (_err error) {
	_method := "pool.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetCrashDumpSR Set the crash_dump_SR field of the given pool.
func (_class PoolClass) SetCrashDumpSR(sessionID SessionRef, self PoolRef, value SRRef) (_err error) {
	_method := "pool.set_crash_dump_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetSuspendImageSR Set the suspend_image_SR field of the given pool.
func (_class PoolClass) SetSuspendImageSR(sessionID SessionRef, self PoolRef, value SRRef) (_err error) {
	_method := "pool.set_suspend_image_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetDefaultSR Set the default_SR field of the given pool.
func (_class PoolClass) SetDefaultSR(sessionID SessionRef, self PoolRef, value SRRef) (_err error) {
	_method := "pool.set_default_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetNameDescription Set the name_description field of the given pool.
func (_class PoolClass) SetNameDescription(sessionID SessionRef, self PoolRef, value string) (_err error) {
	_method := "pool.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetNameLabel Set the name_label field of the given pool.
func (_class PoolClass) SetNameLabel(sessionID SessionRef, self PoolRef, value string) (_err error) {
	_method := "pool.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// GetTelemetryNextCollection Get the telemetry_next_collection field of the given pool.
func (_class PoolClass) GetTelemetryNextCollection(sessionID SessionRef, self PoolRef) (_retval time.Time, _err error) {
	_method := "pool.get_telemetry_next_collection"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTimeToGo(_method + " -> ", _result)
	return
}

// GetTelemetryFrequency Get the telemetry_frequency field of the given pool.
func (_class PoolClass) GetTelemetryFrequency(sessionID SessionRef, self PoolRef) (_retval TelemetryFrequency, _err error) {
	_method := "pool.get_telemetry_frequency"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumTelemetryFrequencyToGo(_method + " -> ", _result)
	return
}

// GetTelemetryUUID Get the telemetry_uuid field of the given pool.
func (_class PoolClass) GetTelemetryUUID(sessionID SessionRef, self PoolRef) (_retval SecretRef, _err error) {
	_method := "pool.get_telemetry_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSecretRefToGo(_method + " -> ", _result)
	return
}

// GetCoordinatorBias Get the coordinator_bias field of the given pool.
func (_class PoolClass) GetCoordinatorBias(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_coordinator_bias"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result)
	return
}

// GetMigrationCompression Get the migration_compression field of the given pool.
func (_class PoolClass) GetMigrationCompression(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_migration_compression"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result)
	return
}

// GetRepositoryProxyPassword Get the repository_proxy_password field of the given pool.
func (_class PoolClass) GetRepositoryProxyPassword(sessionID SessionRef, self PoolRef) (_retval SecretRef, _err error) {
	_method := "pool.get_repository_proxy_password"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSecretRefToGo(_method + " -> ", _result)
	return
}

// GetRepositoryProxyUsername Get the repository_proxy_username field of the given pool.
func (_class PoolClass) GetRepositoryProxyUsername(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_repository_proxy_username"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result)
	return
}

// GetRepositoryProxyURL Get the repository_proxy_url field of the given pool.
func (_class PoolClass) GetRepositoryProxyURL(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_repository_proxy_url"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result)
	return
}

// GetClientCertificateAuthName Get the client_certificate_auth_name field of the given pool.
func (_class PoolClass) GetClientCertificateAuthName(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_client_certificate_auth_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result)
	return
}

// GetClientCertificateAuthEnabled Get the client_certificate_auth_enabled field of the given pool.
func (_class PoolClass) GetClientCertificateAuthEnabled(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_client_certificate_auth_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result)
	return
}

// GetRepositories Get the repositories field of the given pool.
func (_class PoolClass) GetRepositories(sessionID SessionRef, self PoolRef) (_retval []RepositoryRef, _err error) {
	_method := "pool.get_repositories"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertRepositoryRefSetToGo(_method + " -> ", _result)
	return
}

// GetTLSVerificationEnabled Get the tls_verification_enabled field of the given pool.
func (_class PoolClass) GetTLSVerificationEnabled(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_tls_verification_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result)
	return
}

// GetIsPsrPending Get the is_psr_pending field of the given pool.
func (_class PoolClass) GetIsPsrPending(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_is_psr_pending"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result)
	return
}

// GetUefiCertificates Get the uefi_certificates field of the given pool.
func (_class PoolClass) GetUefiCertificates(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_uefi_certificates"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result)
	return
}

// GetIgmpSnoopingEnabled Get the igmp_snooping_enabled field of the given pool.
func (_class PoolClass) GetIgmpSnoopingEnabled(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_igmp_snooping_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result)
	return
}

// GetLivePatchingDisabled Get the live_patching_disabled field of the given pool.
func (_class PoolClass) GetLivePatchingDisabled(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_live_patching_disabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result)
	return
}

// GetPolicyNoVendorDevice Get the policy_no_vendor_device field of the given pool.
func (_class PoolClass) GetPolicyNoVendorDevice(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_policy_no_vendor_device"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result)
	return
}

// GetCPUInfo Get the cpu_info field of the given pool.
func (_class PoolClass) GetCPUInfo(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	_method := "pool.get_cpu_info"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result)
	return
}

// GetGuestAgentConfig Get the guest_agent_config field of the given pool.
func (_class PoolClass) GetGuestAgentConfig(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	_method := "pool.get_guest_agent_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result)
	return
}

// GetCurrentOperations Get the current_operations field of the given pool.
func (_class PoolClass) GetCurrentOperations(sessionID SessionRef, self PoolRef) (_retval map[string]PoolAllowedOperations, _err error) {
	_method := "pool.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumPoolAllowedOperationsMapToGo(_method + " -> ", _result)
	return
}

// GetAllowedOperations Get the allowed_operations field of the given pool.
func (_class PoolClass) GetAllowedOperations(sessionID SessionRef, self PoolRef) (_retval []PoolAllowedOperations, _err error) {
	_method := "pool.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumPoolAllowedOperationsSetToGo(_method + " -> ", _result)
	return
}

// GetHaClusterStack Get the ha_cluster_stack field of the given pool.
func (_class PoolClass) GetHaClusterStack(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_ha_cluster_stack"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result)
	return
}

// GetMetadataVDIs Get the metadata_VDIs field of the given pool.
func (_class PoolClass) GetMetadataVDIs(sessionID SessionRef, self PoolRef) (_retval []VDIRef, _err error) {
	_method := "pool.get_metadata_VDIs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefSetToGo(_method + " -> ", _result)
	return
}

// GetRestrictions Get the restrictions field of the given pool.
func (_class PoolClass) GetRestrictions(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	_method := "pool.get_restrictions"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result)
	return
}

// GetVswitchController Get the vswitch_controller field of the given pool.
func (_class PoolClass) GetVswitchController(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_vswitch_controller"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result)
	return
}

// GetRedoLogVdi Get the redo_log_vdi field of the given pool.
func (_class PoolClass) GetRedoLogVdi(sessionID SessionRef, self PoolRef) (_retval VDIRef, _err error) {
	_method := "pool.get_redo_log_vdi"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result)
	return
}

// GetRedoLogEnabled Get the redo_log_enabled field of the given pool.
func (_class PoolClass) GetRedoLogEnabled(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_redo_log_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result)
	return
}

// GetWlbVerifyCert Get the wlb_verify_cert field of the given pool.
func (_class PoolClass) GetWlbVerifyCert(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_wlb_verify_cert"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result)
	return
}

// GetWlbEnabled Get the wlb_enabled field of the given pool.
func (_class PoolClass) GetWlbEnabled(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_wlb_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result)
	return
}

// GetWlbUsername Get the wlb_username field of the given pool.
func (_class PoolClass) GetWlbUsername(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_wlb_username"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result)
	return
}

// GetWlbURL Get the wlb_url field of the given pool.
func (_class PoolClass) GetWlbURL(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_wlb_url"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result)
	return
}

// GetHealthCheckConfig Get the health_check_config field of the given pool.
func (_class PoolClass) GetHealthCheckConfig(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	_method := "pool.get_health_check_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result)
	return
}

// GetGuiConfig Get the gui_config field of the given pool.
func (_class PoolClass) GetGuiConfig(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	_method := "pool.get_gui_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result)
	return
}

// GetTags Get the tags field of the given pool.
func (_class PoolClass) GetTags(sessionID SessionRef, self PoolRef) (_retval []string, _err error) {
	_method := "pool.get_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result)
	return
}

// GetBlobs Get the blobs field of the given pool.
func (_class PoolClass) GetBlobs(sessionID SessionRef, self PoolRef) (_retval map[string]BlobRef, _err error) {
	_method := "pool.get_blobs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToBlobRefMapToGo(_method + " -> ", _result)
	return
}

// GetHaOvercommitted Get the ha_overcommitted field of the given pool.
func (_class PoolClass) GetHaOvercommitted(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_ha_overcommitted"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result)
	return
}

// GetHaAllowOvercommit Get the ha_allow_overcommit field of the given pool.
func (_class PoolClass) GetHaAllowOvercommit(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_ha_allow_overcommit"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result)
	return
}

// GetHaPlanExistsFor Get the ha_plan_exists_for field of the given pool.
func (_class PoolClass) GetHaPlanExistsFor(sessionID SessionRef, self PoolRef) (_retval int, _err error) {
	_method := "pool.get_ha_plan_exists_for"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result)
	return
}

// GetHaHostFailuresToTolerate Get the ha_host_failures_to_tolerate field of the given pool.
func (_class PoolClass) GetHaHostFailuresToTolerate(sessionID SessionRef, self PoolRef) (_retval int, _err error) {
	_method := "pool.get_ha_host_failures_to_tolerate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result)
	return
}

// GetHaStatefiles Get the ha_statefiles field of the given pool.
func (_class PoolClass) GetHaStatefiles(sessionID SessionRef, self PoolRef) (_retval []string, _err error) {
	_method := "pool.get_ha_statefiles"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result)
	return
}

// GetHaConfiguration Get the ha_configuration field of the given pool.
func (_class PoolClass) GetHaConfiguration(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	_method := "pool.get_ha_configuration"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result)
	return
}

// GetHaEnabled Get the ha_enabled field of the given pool.
func (_class PoolClass) GetHaEnabled(sessionID SessionRef, self PoolRef) (_retval bool, _err error) {
	_method := "pool.get_ha_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result)
	return
}

// GetOtherConfig Get the other_config field of the given pool.
func (_class PoolClass) GetOtherConfig(sessionID SessionRef, self PoolRef) (_retval map[string]string, _err error) {
	_method := "pool.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result)
	return
}

// GetCrashDumpSR Get the crash_dump_SR field of the given pool.
func (_class PoolClass) GetCrashDumpSR(sessionID SessionRef, self PoolRef) (_retval SRRef, _err error) {
	_method := "pool.get_crash_dump_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefToGo(_method + " -> ", _result)
	return
}

// GetSuspendImageSR Get the suspend_image_SR field of the given pool.
func (_class PoolClass) GetSuspendImageSR(sessionID SessionRef, self PoolRef) (_retval SRRef, _err error) {
	_method := "pool.get_suspend_image_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefToGo(_method + " -> ", _result)
	return
}

// GetDefaultSR Get the default_SR field of the given pool.
func (_class PoolClass) GetDefaultSR(sessionID SessionRef, self PoolRef) (_retval SRRef, _err error) {
	_method := "pool.get_default_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefToGo(_method + " -> ", _result)
	return
}

// GetMaster Get the master field of the given pool.
func (_class PoolClass) GetMaster(sessionID SessionRef, self PoolRef) (_retval HostRef, _err error) {
	_method := "pool.get_master"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefToGo(_method + " -> ", _result)
	return
}

// GetNameDescription Get the name_description field of the given pool.
func (_class PoolClass) GetNameDescription(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result)
	return
}

// GetNameLabel Get the name_label field of the given pool.
func (_class PoolClass) GetNameLabel(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result)
	return
}

// GetUUID Get the uuid field of the given pool.
func (_class PoolClass) GetUUID(sessionID SessionRef, self PoolRef) (_retval string, _err error) {
	_method := "pool.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result)
	return
}

// GetByUUID Get a reference to the pool instance with the specified UUID.
func (_class PoolClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PoolRef, _err error) {
	_method := "pool.get_by_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_uuidArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "uuid"), uuid)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _uuidArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolRefToGo(_method + " -> ", _result)
	return
}

// GetRecord Get a record containing the current state of the given pool.
func (_class PoolClass) GetRecord(sessionID SessionRef, self PoolRef) (_retval PoolRecord, _err error) {
	_method := "pool.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolRecordToGo(_method + " -> ", _result)
	return
}
