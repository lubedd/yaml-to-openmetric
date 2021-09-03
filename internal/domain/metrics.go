package domain

type Currencies struct {
	Currencies []struct {
		Name  string `yaml:"name"`
		Value float32    `yaml:"value"`
	} `yaml:"currencies"`
}