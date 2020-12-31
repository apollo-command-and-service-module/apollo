# Apollo
Apollo is an open source initiative that automatically sync an ECS cluster's state matches to the config in GitHub using GitOps. 

Its a next-generation continuous delivery tool for deployments inside ECS, which means you don't need a separate continuous delivery tool. It monitors all relevant image repositories, detects new images, triggers deployments and updates the desired running configuration based on the task config.

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

Build documentation
----------------

Release documentation
----------------

Architectural Overview
----------------

![Architectural](images/architectural.png?raw=true "Architectural Overview")