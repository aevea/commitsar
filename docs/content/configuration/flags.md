---
id: flags
title: Flags
---

For more advanced usage a set of flags are provided.

| Name    | Flag  | Required | Default | Description                                                                        |
| ------- | ----- | -------- | ------- | ---------------------------------------------------------------------------------- |
| Verbose | --v   | false    | false   | Debug output into console                                                          |
| Strict  | --s   | false    | true    | Strict check of category types                                                     |
| All     | --all | false    | false   | Whether to check all commits on given branch. **Takes precedence over LIMIT flag** |
| Config path     | --config-path | false    | current directory | Path where .commitsar.yaml file is |

On top of that a single argument is allowed:

`commitsar <from commit>...<to commit>`

e.g. `commitsar 7dbf3e7db93ae2e02902cae9d2f1de1b1e5c8c92...d0240d3ed34685d0a5329b185e120d3e8c205be4`

If only one commit hash is used then commitsar will assume it to be the TO commit.