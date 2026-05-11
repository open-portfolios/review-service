# review

E-commerce comment review system.

## Prerequisites

### Container Engine

This project uses [Podman](https://podman.io) as the container engine, but any OCI-compatible container engine should work.

[Docker](https://www.docker.com) support is included.

### Taskfile

[Taskfile](https://taskfile.dev) is introduced as an alternative to Makefile. Instead of typing a cluster of long long commands, you can use `task <recipe>` to run a defined recipe.

| Recipe     | Effect                                                       |
| ---------- | ------------------------------------------------------------ |
| `install`  | Install CLI tools needed for development                     |
| `up`       | Compose up containers                                        |
| `down`     | Shut down containers                                         |
| `clean`    | Shut down containers and **remove all data** (be careful!)   |
| `database` | Connect to the interactive shell of the database             |
| `migrate`  | Create tables according to the SQL files under [sql/](./sql) |
| `serve`    | Run server                                                   |
| `all`      | Perform `conf`, `api` and `wire` tasks                       |
| `conf`     | Generate configuration protobuf                              |
| `api`      | Generate API protobuf                                        |
| `wire`     | Generate Dependency Injection code                           |
| `build`    | Build executables                                            |

Taskfile uses the YAML format, and you will find it familiar if you have read GitHub Actions workflows before.

### Go

*The language we Gophers love*. The [Go](https://go.dev) version of this project is 1.22.

## Deployment

This project uses two types of configuration files:

- [configs/config.yaml](configs/config.yaml) holds non‑sensitive, environment‑agnostic settings (e.g., service endpoints, logging levels, database connection pools). This file is safe to commit to version control and can be shared across team members.
- .env stores sensitive information such as API secrets, tokens, passwords, or private keys. It is never committed to the repository (see [.gitignore](.gitignore)). Each developer or deployment environment maintains its own `.env` file. An example dotenv file is [.env.example](.env.example).

Keeping secrets out of code and configuration repositories reduces the risk of accidental exposure. It also simplifies environment switching because only the .env file needs to change when moving between staging and production.

## License 

Copyright 2026 Open Portfolios Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.