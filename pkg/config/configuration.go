package config

type Config struct {
	Server struct {
		Port string `yaml:"port" env:"SERVER_PORT" env-default:"8080"`
		Host string `yaml:"host" env:"SERVER_HOST" env-default:"0.0.0.0"`
	} `yaml:"server"`
	Opists struct {
		PProfOn  bool `yaml:"pprof_on" env:"PPROF_ON" env-default:"false"`
		Security struct {
			Enabled             bool   `yaml:"enabled" env:"SECURITY_ENABLED" env-default:"true"`
			EnableUserLogging   bool   `yaml:"enabled_user_logging" env:"ENABLED_USER_LOGGING" env-default:"false"`
			AuthzModelFilePath  string `yaml:"authz_model_file_path" env:"SECURITY_AUTHZ_MODEL_FILE_PATH" env-default:"authz_model.conf"`
			AuthzPolicyFilePath string `yaml:"authz_policy_file_path" env:"SECURITY_AUTHZ_POLICY_FILE_PATH" env-default:"authz_policy.csv"`
			AdminUsername       string `yaml:"admin_username" env:"ADMIN_USERNAME"`
			AdminPassword       string `yaml:"admin_password" env:"ADMIN_PASSWORD"`
			Users               []struct {
				Username string `yaml:"username"`
				Password string `yaml:"password"`
			}
		} `yaml:"security"`
	} `yaml:"opists"`
}
