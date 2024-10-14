# Damga CLI

Damga CLI is a tool to communicate with Damga API to manage the resources.

## Install

Pick the desired version [from releases](https://github.com/devingen/damga/releases), download and extract the archive. 
You can use `damga` executable right away in the folder.
In order to access the `damga` command globally, move it to the machine's executable path (e.g.`/usr/local/bin`).

> Linux archives can be used on MacOS.

Commands to install `damga` command on a Linux server is:

```shell
wget https://github.com/devingen/damga/releases/download/damga-cli-v1.0.0/damga-cli-linux-amd64.tar.gz
tar xvfz damga-cli-*.tar.gz
sudo mv damga /usr/local/bin/
damga --version
```

## Usage

```shell
# Add a new target to a target group
damga tg add-target --api-key=ASD --id abcd --target 1.2.3.4

# Remove a target from a target group
damga tg remove-target --api-key=ASD --id abcd --target 1.2.3.4
```
