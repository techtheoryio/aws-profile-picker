<p align="center">
  <img width="600" src="./images/example.svg">
  <br>
  <img alt="GitHub release (latest SemVer)" src="https://img.shields.io/github/v/release/techtheoryio/aws-profile-picker?color=gree">
  <a href="https://goreportcard.com/report/github.com/techtheoryio/aws-profile-picker"><img src="https://goreportcard.com/badge/github.com/techtheoryio/aws-profile-picker" alt="Version"></a>
  <a href="#"><img src="https://img.shields.io/github/go-mod/go-version/techtheoryio/aws-profile-picker" alt="Version"></a>
</p>

# aws-profile-picker

`aws-profile-picker` is a simple CLI built to assist switching AWS profiles. It includes:

- Dropdown list of all configured AWS profiles in `$HOME/.aws/config`
- Fuzzy search for all profiles
- No additional shell is being created - the environment variable `AWS_PROFILE` is set in the current shell

## Installation

1. Download [the latest release](https://github.com/techtheoryio/aws-profile-picker/releases).
2. Unpack the archive
3. `cd scripts`
4. `sh ./install.sh`
5. Add an alias to your shell configuration (.bshrc, .zshrc, etc.)
   ```bash
   alias awsp="source _awsp"
   ```
   This is required so no additional shell has to be created.

## Usage

Supplying no argument will open the interactive
prompt.

```bash
$ awsp
```

Alternatively you can also directly pass a
profile to search for as an argument. The profile
that matches the argument the most using the Levinsthein
algorithm will be set.
If no profile is found, the prompt opens.

```bash
$ awsp client-proj
```

## Uninstall

In the release archive, there's an uninstall script

1. `cd scripts`
2. `sh ./uninstall.sh`
3. Remove the alias from your shell configuration file.

### Special thanks

The `alias awsp="source _awsp"` idea comes straight from [johnnyopao/awsp](https://github.com/johnnyopao/awsp). It's a very similar tool, however we really wanted fuzzy search :)
