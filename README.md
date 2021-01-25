# Apollo
Apollo is an open source initiative that automatically sync an ECS cluster's state to the config in GitHub using GitOps. 

Its a next-generation continuous delivery tool for deployments inside ECS, which means you don't need a separate continuous delivery tool. It monitors all relevant GitHub repositories, detects new commits, triggers deployments and updates the desired running configuration.

The benefits are:
- You don't need to grant your continuous integration tool access to the ECS cluster.
- Every change is atomic and transactional.
- Git has your audit log.

Each transaction either fails or succeeds cleanly. You're entirely code-centric and don't need new continuous delivery infrastructure.

How Apollo works
----------------
It all starts with a commit to your application config repo.
![Overview](images/overview.png?raw=true "Overview")

“Our goal with Apollo is to empower organizations to declaratively build and run cloud native applications and workflows on ECS/Fargate using GitOps,” 

# Community & Developer information
We welcome all kinds of contributions to Apollo, be it code, issues you found, documentation, external tools, help and support or anything else really.

Instances of abusive, harassing, or otherwise unacceptable behavior may be reported by contacting a project maintainer.

To familiarise yourself with the project and how things work, you might be interested in the following:

Licence
----------------

Released under the [MIT License](https://www.opensource.org/licenses/mit-license.php).

Maintainer
----------------
* Micheal Montpetit [@bldmgr](https://github.com/apollo-command-and-service-module)

Our contributions guidelines
----------------

Developer Notes
----------------

- default setting are located in service-module/service-module.go

GoLang
```
go run ./service-module/service-module.go
```

Docker
```
docker build -t sm . && docker run sm 
```


**Basic Auth**
````
export GITHUB_TOKEN=<github personal access tokens>
export GITHUB_USER=<github user>
````

#### The Apollo Guidance Configuration file
The AGC file is require and contains the source of each service repositories

```
repos:
- name: service-name
  url: https://github.com/apollo-command-and-service-module/orbit.git
  branch: main
  config: config.yaml
```

override local path and configuration file
````
export AGC_PATH=./agc
export AGC_FILE=agc.yaml
````


Build documentation
----------------

Release documentation
----------------

Architectural Overview
----------------

![Architectural](images/architectural.png?raw=true "Architectural Overview")