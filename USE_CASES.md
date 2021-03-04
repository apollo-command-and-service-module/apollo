## Use Cases to help you understand Apollo.

When we consider GitOps, the use cases that come immediately to mind are container orchestration and cluster management within Kubernetes. But not other cloud container orchestration and cluster management services - specifically with AWS ECS/Fargate.

GitOps helps teams be more autonomous and productive by enabling continuous deployment through the source control tools they work in every day. Changes are automatically applied to the cluster once a Pull Request is approved and merged. Changes are handled by reconcilers that look for discrepancies between what is described in Git (the desired state) with what is currently running in the clusters.

The idea of Apollo is that the desired state of the entire system is stored in version control via Git commits. A key feature of Apollo is that the application and environment’s ideal state is described using config files.

To make changes to the desired container orchestration and cluster management state, a developer issues a pull request which basically tells everyone about changes that you’ve published a new version of the configuration and instructs Apollo to update its state.

### Scenario

Team Green is responsible for the development and deployment of a simple application, and they want that whenever a new commit hits the master branch, a new CI Pipeline is executed, this CI Pipeline will take care of:
```
1. Download the new code
2. Run a code linter
3. Run the tests
4. Build the app on a container image
5. Push the container image to the container image repository
```

### Pre-requisites

Apollo has been deployed and configured to synchronize the service configuration repository.

### GitOps Workflow

Whenever a new commit hits stage or prod branches, the application is synchronized by Apollo.

```
1. Developers push new changes to the application repository
2. The changes are tested and a new image is built automatically by a CI Pipeline
3. Developers request the deployment of the new image by sending a PR to the stage branch that contains an updated apollo configuration
4. The PR is reviewed and merged
5. The new image version is deployed automatically by Apollo on the stage environment
6. After testing the application on the stage environment, a PR is sent to the production branch to deploy the application on the production environment
7. The PR is reviewed and merged
8. The new image version is deployed automatically by Apollo on production
```