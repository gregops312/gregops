oclif-hello-world
=================

oclif example Hello World CLI

[![oclif](https://img.shields.io/badge/cli-oclif-brightgreen.svg)](https://oclif.io)
[![Version](https://img.shields.io/npm/v/oclif-hello-world.svg)](https://npmjs.org/package/oclif-hello-world)
[![CircleCI](https://circleci.com/gh/oclif/hello-world/tree/main.svg?style=shield)](https://circleci.com/gh/oclif/hello-world/tree/main)
[![Downloads/week](https://img.shields.io/npm/dw/oclif-hello-world.svg)](https://npmjs.org/package/oclif-hello-world)
[![License](https://img.shields.io/npm/l/oclif-hello-world.svg)](https://github.com/oclif/hello-world/blob/main/package.json)

<!-- toc -->
* [Usage](#usage)
* [Commands](#commands)
<!-- tocstop -->
# Usage
<!-- usage -->
```sh-session
$ npm install -g gregops
$ gregops COMMAND
running command...
$ gregops (--version)
gregops/0.0.0 darwin-x64 node-v18.12.1
$ gregops --help [COMMAND]
USAGE
  $ gregops COMMAND
...
```
<!-- usagestop -->
# Commands
<!-- commands -->
* [`gregops hello PERSON`](#gregops-hello-person)
* [`gregops hello world`](#gregops-hello-world)
* [`gregops help [COMMAND]`](#gregops-help-command)
* [`gregops plugins`](#gregops-plugins)
* [`gregops plugins:install PLUGIN...`](#gregops-pluginsinstall-plugin)
* [`gregops plugins:inspect PLUGIN...`](#gregops-pluginsinspect-plugin)
* [`gregops plugins:install PLUGIN...`](#gregops-pluginsinstall-plugin-1)
* [`gregops plugins:link PLUGIN`](#gregops-pluginslink-plugin)
* [`gregops plugins:uninstall PLUGIN...`](#gregops-pluginsuninstall-plugin)
* [`gregops plugins:uninstall PLUGIN...`](#gregops-pluginsuninstall-plugin-1)
* [`gregops plugins:uninstall PLUGIN...`](#gregops-pluginsuninstall-plugin-2)
* [`gregops plugins update`](#gregops-plugins-update)

## `gregops hello PERSON`

Say hello

```
USAGE
  $ gregops hello [PERSON] -f <value>

ARGUMENTS
  PERSON  Person to say hello to

FLAGS
  -f, --from=<value>  (required) Who is saying hello

DESCRIPTION
  Say hello

EXAMPLES
  $ oex hello friend --from oclif
  hello friend from oclif! (./src/commands/hello/index.ts)
```

_See code: [dist/commands/hello/index.ts](https://github.com/gkman/gregops/blob/v0.0.0/dist/commands/hello/index.ts)_

## `gregops hello world`

Say hello world

```
USAGE
  $ gregops hello world

DESCRIPTION
  Say hello world

EXAMPLES
  $ gregops hello world
  hello world! (./src/commands/hello/world.ts)
```

## `gregops help [COMMAND]`

Display help for gregops.

```
USAGE
  $ gregops help [COMMAND] [-n]

ARGUMENTS
  COMMAND  Command to show help for.

FLAGS
  -n, --nested-commands  Include all nested commands in the output.

DESCRIPTION
  Display help for gregops.
```

_See code: [@oclif/plugin-help](https://github.com/oclif/plugin-help/blob/v5.1.22/src/commands/help.ts)_

## `gregops plugins`

List installed plugins.

```
USAGE
  $ gregops plugins [--core]

FLAGS
  --core  Show core plugins.

DESCRIPTION
  List installed plugins.

EXAMPLES
  $ gregops plugins
```

_See code: [@oclif/plugin-plugins](https://github.com/oclif/plugin-plugins/blob/v2.1.12/src/commands/plugins/index.ts)_

## `gregops plugins:install PLUGIN...`

Installs a plugin into the CLI.

```
USAGE
  $ gregops plugins:install PLUGIN...

ARGUMENTS
  PLUGIN  Plugin to install.

FLAGS
  -f, --force    Run yarn install with force flag.
  -h, --help     Show CLI help.
  -v, --verbose

DESCRIPTION
  Installs a plugin into the CLI.
  Can be installed from npm or a git url.

  Installation of a user-installed plugin will override a core plugin.

  e.g. If you have a core plugin that has a 'hello' command, installing a user-installed plugin with a 'hello' command
  will override the core plugin implementation. This is useful if a user needs to update core plugin functionality in
  the CLI without the need to patch and update the whole CLI.


ALIASES
  $ gregops plugins add

EXAMPLES
  $ gregops plugins:install myplugin 

  $ gregops plugins:install https://github.com/someuser/someplugin

  $ gregops plugins:install someuser/someplugin
```

## `gregops plugins:inspect PLUGIN...`

Displays installation properties of a plugin.

```
USAGE
  $ gregops plugins:inspect PLUGIN...

ARGUMENTS
  PLUGIN  [default: .] Plugin to inspect.

FLAGS
  -h, --help     Show CLI help.
  -v, --verbose

DESCRIPTION
  Displays installation properties of a plugin.

EXAMPLES
  $ gregops plugins:inspect myplugin
```

## `gregops plugins:install PLUGIN...`

Installs a plugin into the CLI.

```
USAGE
  $ gregops plugins:install PLUGIN...

ARGUMENTS
  PLUGIN  Plugin to install.

FLAGS
  -f, --force    Run yarn install with force flag.
  -h, --help     Show CLI help.
  -v, --verbose

DESCRIPTION
  Installs a plugin into the CLI.
  Can be installed from npm or a git url.

  Installation of a user-installed plugin will override a core plugin.

  e.g. If you have a core plugin that has a 'hello' command, installing a user-installed plugin with a 'hello' command
  will override the core plugin implementation. This is useful if a user needs to update core plugin functionality in
  the CLI without the need to patch and update the whole CLI.


ALIASES
  $ gregops plugins add

EXAMPLES
  $ gregops plugins:install myplugin 

  $ gregops plugins:install https://github.com/someuser/someplugin

  $ gregops plugins:install someuser/someplugin
```

## `gregops plugins:link PLUGIN`

Links a plugin into the CLI for development.

```
USAGE
  $ gregops plugins:link PLUGIN

ARGUMENTS
  PATH  [default: .] path to plugin

FLAGS
  -h, --help     Show CLI help.
  -v, --verbose

DESCRIPTION
  Links a plugin into the CLI for development.
  Installation of a linked plugin will override a user-installed or core plugin.

  e.g. If you have a user-installed or core plugin that has a 'hello' command, installing a linked plugin with a 'hello'
  command will override the user-installed or core plugin implementation. This is useful for development work.


EXAMPLES
  $ gregops plugins:link myplugin
```

## `gregops plugins:uninstall PLUGIN...`

Removes a plugin from the CLI.

```
USAGE
  $ gregops plugins:uninstall PLUGIN...

ARGUMENTS
  PLUGIN  plugin to uninstall

FLAGS
  -h, --help     Show CLI help.
  -v, --verbose

DESCRIPTION
  Removes a plugin from the CLI.

ALIASES
  $ gregops plugins unlink
  $ gregops plugins remove
```

## `gregops plugins:uninstall PLUGIN...`

Removes a plugin from the CLI.

```
USAGE
  $ gregops plugins:uninstall PLUGIN...

ARGUMENTS
  PLUGIN  plugin to uninstall

FLAGS
  -h, --help     Show CLI help.
  -v, --verbose

DESCRIPTION
  Removes a plugin from the CLI.

ALIASES
  $ gregops plugins unlink
  $ gregops plugins remove
```

## `gregops plugins:uninstall PLUGIN...`

Removes a plugin from the CLI.

```
USAGE
  $ gregops plugins:uninstall PLUGIN...

ARGUMENTS
  PLUGIN  plugin to uninstall

FLAGS
  -h, --help     Show CLI help.
  -v, --verbose

DESCRIPTION
  Removes a plugin from the CLI.

ALIASES
  $ gregops plugins unlink
  $ gregops plugins remove
```

## `gregops plugins update`

Update installed plugins.

```
USAGE
  $ gregops plugins update [-h] [-v]

FLAGS
  -h, --help     Show CLI help.
  -v, --verbose

DESCRIPTION
  Update installed plugins.
```
<!-- commandsstop -->
