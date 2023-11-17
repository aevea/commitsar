---
id: config-file
title: Configuration File
---

**The configuration file is still under development and is subject to changes**

**Name:** `.commitsar.yml`

In order to make configuration easier than through flags we provide configuration file support. Most up-to-date examples can be found in <https://github.com/aevea/commitsar/tree/master/config/testdata>.

By default the current working directory is used to scan for the file. However this can be overridden by specifying `COMMITSAR_CONFIG_PATH` environment variable. Accepts relative or absolute paths.

Example: `COMMITSAR_CONFIG_PATH=./testdata` will scan for `.commitsar.yaml` in the `testdata` folder.

## Global configuration

These are settings that get used across all runs of commitsar.

```yaml
version: 1
verbose: false
```

| Name    | Default Value | Description                                                                         | Available from |
| ------- | ------------- | ----------------------------------------------------------------------------------- | -------------- |
| version | 1             | Currently not in use. Might be used in the future in case of incompatible upgrades. | v0.14.0        |
| verbose | false         | Turns on debug logging of commitsar. Useful if you want to submit an issue.         | v0.14.0        |

## Commit style settings

```yaml
commits:
  disabled: false
  strict: true
  limit: 0
  all: false
  upstreamBranch: origin/master
```

| Name           | Default Value | Description                                                                                             | Available from |
| -------------- | ------------- | ------------------------------------------------------------------------------------------------------- | -------------- |
| disabled       | false         | Disables checking commits. Useful if you want to use commitsar only for PR titles.                    | v0.14.0        |
| strict         | true          | Enforces strict category enforcement.                                                                   | v0.14.0        |
| limit          | none          | Makes commitsar check only the last x commits. Useful if you want to run commitsar on master.          | v0.14.0        |
| all            | false         | Makes commitsar check all the commits in history. **Overrides the `limit` flag**                    | v0.14.0        |
| upstreamBranch | origin/master | Makes commitsar check against specific branch (e.g. use `origin/main` if `main` is your default branch) | v0.17.0        |

## Pull Request style settings

**Pull Request pipeline is still in early stages. Please report any bugs**

#### Conventional style

```yaml
pull_request:
  conventional: true
```

Setting `conventional` to true will enable the pipeline. This is useful for teams that use squash commits and don't care about having all of the commits in the PR compliant with conventional commits.

| Name         | Default Value | Description                                                             | Available from |
| ------------ | ------------- | ----------------------------------------------------------------------- | -------------- |
| conventional | false         | Turns on the pipeline and will check for a conventional commit PR title | v0.17.0        |

#### JIRA style

```yaml
pull_request:
  jira_title: true
  jira_keys:
    - TEST
    - TSLA
```

Setting `jira_title` to true will enable the pipeline. By default commitsar will use a basic regex to check for any JIRA-like references. Further scoping can be done using the `jira_keys` setting.

| Name       | Default Value | Description                                                           | Available from |
| ---------- | ------------- | --------------------------------------------------------------------- | -------------- |
| jira_title | false         | Turns on the pipeline and will check for JIRA issues in the PR title. | v0.15.0        |
| jira_keys  | none          | Array of string project keys from JIRA.                               | v0.15.0        |
