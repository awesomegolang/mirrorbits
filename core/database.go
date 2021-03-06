package core

const (
	// RedisMinimumVersion contains the minimum redis version required to run the application
	RedisMinimumVersion = "3.2.0"
	// DBVersion represents the current DB format version
	DBVersion = 1
	// DBVersionKey contains the global redis key containing the DB version format
	DBVersionKey = "MIRRORBITS_DB_VERSION"
)
