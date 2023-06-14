// Package env contains reused environment variables.
package env

const (
	// BLOB provider configuration.
	BLOBProviderClientID     = "BLOBPROVIDER_CLIENTID"
	BLOBProviderClientSecret = "BLOBPROVIDER_CLIENTSECRET"
	BLOBProviderClientToken  = "BLOBPROVIDER_CLIENTTOKEN"
	BLOBProviderContainer    = "BLOBPROVIDER_CONTAINER"
	BLOBProviderRegion       = "BLOBPROVIDER_REGION"
	BLOBProviderType         = "BLOBPROVIDER_TYPE"
	BLOBProviderURI          = "BLOBPROVIDER_URI"

	// Builds.
	BuildDate = "BUILD_DATE"
	CommitID  = "COMMIT_ID"

	// Data provider configuration.
	DataProviderClientID      = "DATAPROVIDER_CLIENTID"
	DataProviderClientSecret  = "DATAPROVIDER_CLIENTSECRET"
	DataProviderClientToken   = "DATAPROVIDER_CLIENTTOKEN"
	DataProviderContainer     = "DATAPROVIDER_CONTAINER"
	DataProviderDatabase      = "DATAPROVIDER_DATABASE"
	DataProviderPassword      = "DATAPROVIDER_PASSWORD"
	DataProviderRegion        = "DATAPROVIDER_REGION"
	DataProviderSASLMechanism = "DATAPROVIDER_SASLMECHANISM"
	DataProviderType          = "DATAPROVIDER_TYPE"
	DataProviderURI           = "DATAPROVIDER_URI"
	DataProviderUsername      = "DATAPROVIDER_USERNAME"

	// Debug mode.
	Debug = "DEBUG"

	// Hostname.
	Hostname = "HOSTNAME"

	// Static HTML path.
	HTMLPath = "HTML_PATH"

	// Logging.
	LoggingLevel = "LOGGING_LEVEL"

	// Message broker configuration.
	MessageBrokerURI = "MESSAGEBROKER_URI"

	// OpenShift.
	OpenShiftBuildCommit = "OPENSHIFT_BUILD_COMMIT"

	// Path.
	Path = "PATH"

	// Pool.
	PoolID                 = "POOL_ID"
	PoolSecretProviderType = "POOL_SECRETPROVIDERTYPE"

	// Search engine configuration.
	SearchEnginePassword = "SEARCHENGINE_PASSWORD"
	SearchEngineSniff    = "SEARCHENGINE_SNIFF"
	SearchEngineURI      = "SEARCHENGINE_URI"
	SearchEngineUsername = "SEARCHENGINE_USERNAME"

	// Secret provider configuration.
	SecretProviderClientID     = "SECRETPROVIDER_CLIENTID"     // #nosec G101
	SecretProviderClientSecret = "SECRETPROVIDER_CLIENTSECRET" // #nosec G101
	SecretProviderClientToken  = "SECRETPROVIDER_CLIENTTOKEN"  // #nosec G101
	SecretProviderContainer    = "SECRETPROVIDER_CONTAINER"    // #nosec G101
	SecretProviderRegion       = "SECRETPROVIDER_REGION"       // #nosec G101
	SecretProviderType         = "SECRETPROVIDER_TYPE"         // #nosec G101
	SecretProviderURI          = "SECRETPROVIDER_URI"          // #nosec G101

	// Security headers.
	SecurityHeadersConnectSrc    = "SECURITYHEADERS_CONNECTSRC"
	SecurityHeadersDefaultSrc    = "SECURITYHEADERS_DEFAULTSRC"
	SecurityHeadersFeaturePolicy = "SECURITYHEADERS_FEATUREPOLICY"
	SecurityHeadersFontSrc       = "SECURITYHEADERS_FONTSRC"
	SecurityHeadersFrameSrc      = "SECURITYHEADERS_FRAMESRC"
	SecurityHeadersImgSrc        = "SECURITYHEADERS_IMGSRC"
	SecurityHeadersObjectSrc     = "SECURITYHEADERS_OBJECTSRC"
	SecurityHeadersScriptSrc     = "SECURITYHEADERS_SCRIPTSRC"
	SecurityHeadersStyleSrc      = "SECURITYHEADERS_STYLESRC"
	SecurityHeadersTelemetryType = "SECURITYHEADERS_TELEMETRYTYPE"

	// SystemRoot.
	SystemRoot = "SYSTEMROOT"

	// Telemetry.
	TelemetryClientID     = "TELEMETRY_CLIENTID"
	TelemetryClientSecret = "TELEMETRY_CLIENTSECRET" // #nosec G101
	TelemetryDaemonURI    = "TELEMETRY_DAEMONURI"
	TelemetryLoggingLevel = "TELEMETRY_LOGGINGLEVEL"
	TelemetryType         = "TELEMETRY_TYPE"
	TelemetryVersion      = "TELEMETRY_VERSION"

	// Test.
	TestPersistData = "TEST_PERSIST_DATA"

	// Vault.
	VaultAddr         = "VAULT_ADDR"
	VaultSkipVerify   = "VAULT_SKIP_VERIFY"
	VaultToken        = "VAULT_TOKEN"
	VaultUnsealShards = "VAULT_UNSEAL_SHARDS"

	// Worker.
	WorkerID = "WORKER_ID"
)
