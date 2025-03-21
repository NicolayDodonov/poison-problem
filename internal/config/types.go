package config

type Config struct {
	Logger     `yaml:"logger"`
	Simulation `yaml:"simulation"`
}

type Logger struct {
	Type string `yaml:"type" env-required:"true"`
	Path string `yaml:"type" env-required:"true"`
}

type Simulation struct {
	Type       string `yaml:"type" env-required:"true"`
	TargetAge  int    `yaml:"targetAge"`
	StartAgent int    `yaml:"startCountAgent"`
	EndAgent   int    `yaml:"endCountAgent"`
	MaxEpoch   int    `yaml:"maxCountEpoch"`
	LoadSing   string `yaml:"pathLoadSing"`
	SaveSing   string `yaml:"pathSaveSing"`
	SaveStat   string `yaml:"pathSaveStat"`
}
