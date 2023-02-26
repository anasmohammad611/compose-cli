package types

type Template struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

type Service struct {
	Name      string              `json:"name"`
	Templates map[string]Template `json:"templates"`
}

type ServiceConfig struct {
	Services map[string]Service
}

func (s *ServiceConfig) Load(configJson string) error {
	// load all services from the config json string
	// and set it to the current config
	return nil
}

func (s ServiceConfig) Dump() string {
	// return a string which contains json dump for all the services
	return ""
}

func (s *ServiceConfig) GetTemplate(serviceName string, templateCode string) (Template, error) {
	// return a specific template object by searching it in the config
	return Template{}, nil
}
