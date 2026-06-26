package config

type ConfigPostgresMigration struct {
	Structure    bool   `mapstructure:"MIGRATE_POSTGRES_STRUCTURE"`
	Data         bool   `mapstructure:"MIGRATE_POSTGRES_DATA"`
	Drop         bool   `mapstructure:"MIGRATE_POSTGRES_DROP"`
	DropData     bool   `mapstructure:"MIGRATE_POSTGRES_DROP_DATA"`
	VersionTable string `mapstructure:"MIGRATE_POSTGRES_VERSION_TABLE"`
}
