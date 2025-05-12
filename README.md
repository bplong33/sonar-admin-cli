# Sonar Admin CLI

This tool was designed to make it easy to automate some routine actions
with [SonarQube Server](https://docs.sonarsource.com/sonarqube-server/latest/).
In my experience, bulk actions do not receive a lot of support within SonarQube's
UI. Simple tasks such as changing permissions on many projects or adding
permissions for a group to multiple projects take a bit longer than they really
need to.

This CLI tool is the solution to that problem, allowing the user to speed up
many of these routine tasks.

## Installation

...

## Configuration

The default configuration file is stored at `$HOME/.sonar-admin-cli.toml`. You
can set this configuration manually.

> [!NOTE]
> I'm working on a feature to allow setting these configuration via a `config` command
> option

```toml
# .sonar-admin-cli.toml
[sonar]
active_env = env1

[sonar.env1]
host = http://localhost:8080/api
token = squ_12341234
```

## ToDo

- [ ] Add `config` command
  - [ ] Add option to create additional environments
  - [ ] Add option to set an active environment
- [ ] Add command to add a group to a project with X permissions
