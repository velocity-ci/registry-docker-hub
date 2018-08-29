# registry.docker-hub
Provides Docker Hub registry authentication for Velocity CI

## Usage

```
docker:
  registries:
    - use: https://github.com/velocity-ci/registry.docker-hub/releases/download/0.1.0/docker-hub
      arguments:
        username: ${docker_username}
        password: ${docker_password}
```
