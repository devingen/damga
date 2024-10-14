# Damga CLI

## Usage

```shell
export DAMGA_API_KEY=...

# Add a new target to the target group
damga tg add --id 67058cdd0d1b707123eac0e4 --target 1.2.3.4

# Remove a target from the target group
damga tg remove --id 67058cdd0d1b707123eac0e4 --target 1.2.3.4
```

## Development

To run a sample command: 

```shell
go run . tg add --id 67058cdd0d1b707123eac0e4 --target 1.2.3.4
```

To build and save into the project

```shell
go build -o ./bin/damga .
```

To build the command and put it into device executables, add 

```
export DEVINGEN_REPOS=~/Workspace/Devingen
export PATH=$DEVINGEN_REPOS/tools/bin:$PATH
```

to `~/.zshrc` and build with 

```shell
go build -o $DEVINGEN_REPOS/tools/bin/damga .
```

Then you'll be able to run the `damga` command.

### Releasing a new version

Create a git tag with the desired version and push the tag.

```
VERSION=1.0.0
make archive-release VERSION=$VERSION

# see tags
git tag --list

# create new tag
git tag -a damga-cli-v$VERSION -m ""

# push new tag
git push origin damga-cli-v$VERSION
```

