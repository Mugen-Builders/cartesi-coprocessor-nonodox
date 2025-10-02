<br>
<p align="center">
    <img src="https://github.com/Mugen-Builders/.github/assets/153661799/7ed08d4c-89f4-4bde-a635-0b332affbd5d" align="center" width="20%">
</p>
<br>
<div align="center">
    <i>A tool for developing Cartesi Coprocessor applications</i>
</div>
<div align="center">
<b>This tool aims to be an iterative development environment for Cartesi Coprocessor applications</b>
</div>
<br>
<p align="center">
<a href="https://github.com/Mugen-Builders/cartesi-coprocessor-nonodox/blob/main/LICENSE">	
	<img src="https://img.shields.io/github/license/Mugen-Builders/cartesi-coprocessor-nonodox?style=default&logo=opensourceinitiative&logoColor=white&color=00ADD8" alt="license">
	<img src="https://img.shields.io/github/last-commit/Mugen-Builders/cartesi-coprocessor-nonodox?style=default&logo=git&logoColor=white&color=000000" alt="last-commit">
</p>

## Table of Contents

- [Overview](#overview)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running](#running)

## Overview

<div align="justify">
This is an iterative tool designed to accelerate the "debugging" and "development" process of applications using the Cartesi Coprocessor infrastructure, providing a faster path to the production environment.
</div>

## Getting Started

### Prerequisites

1. [Foundry](https://book.getfoundry.sh/getting-started/installation)
2. [Nonodo](https://github.com/Calindra/nonodo?tab=readme-ov-file#installation)

### Installation

There are two installation options for the tool: one involves installing the binary according to the architectures available on the releases page, and the other uses Golang to install the package. Choose one of these options to proceed.

- Install the binary:

1. Go to latest [release page](https://github.com/Mugen-Builders/cartesi-coprocessor-nonodox/releases) and download the archive for your host platform;
2. Extract the archive;
3. Add the binary's path to the system PATH so that it can be initialized using just the `nonodox` command;

- Install the package with golang:

```sh
go install github.com/Mugen-Builders/cartesi-coprocessor-nonodox/cmd/nonodox@latest
```

> [!WARNING]
> The command above installs NoNodoX into the `bin` directory inside the directory defined by the `GOPATH` environment variable.
> If you don't set the `GOPATH` variable, the default install location is `$HOME/go/bin`.
> So, to call the `nonodox` command directly, you should add it to the `PATH` variable.
> The command below exemplifies that.
> 
> ```sh
> export PATH="$HOME/go/bin:$PATH"
> ```

### Running

1. Download the state file (.json):

```bash
curl -O https://raw.githubusercontent.com/Mugen-Builders/cartesi-coprocessor-nonodox/refs/heads/main/anvil_state.json
```

2. Start the anvil instance:

```sh
anvil --load-state anvil_state.json
```

> [!CAUTION]
> Before running the command below, please make sure that you have deployed the CoprocesorAdapter instance, passing `0x9A9f2CCfdE556A7E9Ff0848998Aa4a0CFD8863AE` as the coprocessor address to its constructor

3. Running the tool:

```sh
nonodox
```

> [!NOTE]
> If you wish to make any customizations to the tool's execution environment, you can use the `--config` flag along with the path to a `.toml` file, which should contain the following variables:
>
> ```toml
> [anvil]
> http_url = "http://127.0.0.1:8545"
> ws_url = "ws://127.0.0.1:8545"
> private_key = "<private-key-without-0x>" 
> input_box_block = "7"
> ```
