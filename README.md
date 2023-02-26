# compose-cli

## Interface

### CLI (Inputs parse karo sahi se)

```
$ cc init: 
    Args: None
    Task:Will initialise an empty docker-compose.yml file, as well as maybe a hidden config file.
$ cc add: 
    Args: ServiceName, TemplateCode[NO_AUTH]
    Flags: Pass filepath; 
    Task: Add stuff to config, and add stuff to docker-compose.yml
```

## Service (Context: Input validation, File I/O, Comm b/w presentation/core)
 
- Array of obj (ServiceName, TemplateName)
- Call Core.addConfig(oldConfig string, newItem) (newConfig, error)
- file management


## Core (Context: Yaml, Docker-compose, MasterConfig, Yaml/Docker related checks)

- Structure
- maintain config MASTER
    - service
    - templateArray
        - templateCode: [NO_AUTH]
        - description
        - templateContent
- Config Location
    - With Binary as json
    - hosted json
    - computed json from db(API)
        - @chara/postgres-master-slave-template
- yaml
- docker 


getMaster() struct 
    fs.read()
    http.get('staticconfig.json')

getByServiceAndTemplate(serviceName string, templateCode string) struct {
    http.get('/api/?service=&templateCode')
}


