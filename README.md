<p align="center">
  <img src="https://user-images.githubusercontent.com/39099403/211157601-cab972ed-085e-4ff6-bb8a-99e4babb9e3c.svg" width="550" alt="Dahl Logo" />
</p>

<p align="center">A Simple CLI made with <a href="https://github.com/spf13/cobra" target="_blank">Cobra</a> package for generating files.</p>
<p align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/spacelocust/dahl" alt="GitHub go.mod Go version">
  <img src="https://goreportcard.com/badge/github.com/Spacelocust/dahl" alt="Go Report Card" />
  <img src="https://img.shields.io/github/actions/workflow/status/spacelocust/dahl/cd.yml" alt="GitHub Workflow Status" />
  <img src="https://img.shields.io/github/license/spacelocust/dahl" alt="Package License" />
  <img src="https://img.shields.io/github/repo-size/spacelocust/dahl" alt="GitHub repo size">
  <img src="https://img.shields.io/github/v/tag/spacelocust/dahl" alt="GitHub Tag" />
  <img src="https://img.shields.io/github/downloads/spacelocust/dahl/total" alt="GitHub Releases">
</p>

# Sponsors

<p align="center">
 Maintainence of this project is made possible by all the <a href="https://github.com/Spacelocust/dahl/graphs/contributors">contributors</a> and <a href="https://github.com/sponsors/Spacelocust">sponsors</a>. If you'd like to sponsor this project and have your avatar or company logo appear below <a href="https://github.com/sponsors/Spacelocust">click here</a>. 💙
</p>

# Overview
Hello fellow developer! 

Welcome to the Dahl repository! Dahl is a command line tool that helps developers create and use custom templates to generate files quickly and easily. With Dahl, you can save time and effort by using pre-defined templates to generate the files you need for your project, rather than starting from scratch every time.

This repository contains the source code for Dahl, as well as documentation and examples to help you get started. You can also share your own templates with the community by contributing to the repository.

Whether you're a seasoned developer or just starting out, Dahl is a valuable tool to have in your arsenal. Give it a try and see how it can streamline your workflow and save you time.

Thank you for checking out Dahl. We hope it becomes a valuable addition to your toolkit.

- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Commands](#commands)
    - [Examples](#examples)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [Support](#support)
- [License](#license)
- [Acknowledgements](#acknowledgements)

# Getting Started
## Installation

### Binary Release (Linux/OSX/Windows)

You can manually download a binary release from [the release page](https://github.com/Spacelocust/dahl/releases).

Automated install/update, don't forget to always verify what you're piping into bash:

```sh
curl https://raw.githubusercontent.com/Spacelocust/dahl/main/scripts/install.sh | bash
```
The script installs downloaded binary to `$HOME/.local/bin` directory by default, but it can be changed by setting `DIR` environment variable.

I already made an alias for zsh if you want:

```
echo "alias dahl='~/.local/bin/dahl'" >> ~/.zshrc
```
(you can substitute .zshrc for whatever rc file you're using)

### Go

Required Go Version >= **1.16**

```sh
go install github.com/Spacelocust/dahl@latest
```

Required Go version >= **1.8**, <= **1.17**

```sh
go get github.com/Spacelocust/dahl
```

### Docker

[![Docker Pulls](https://img.shields.io/docker/pulls/spacelocust/dahl)](https://hub.docker.com/r/spacelocust/dahl)
[![Docker Stars](https://img.shields.io/docker/stars/spacelocust/dahl)](https://hub.docker.com/r/spacelocust/dahl)
[![Docker Automated](https://img.shields.io/docker/cloud/automated/spacelocust/dahl)](https://hub.docker.com/r/spacelocust/dahl)
[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/spacelocust/dahl?sort=semver)](https://hub.docker.com/repository/docker/spacelocust/dahl/tags?page=1&ordering=last_updated)

```sh
docker run --rm -it -v /your/project/path:/app -w /app spacelocust/dahl:latest
```


## Usage
In your root project, use the following commande for generating a dahl-config.yml and .dahl directory
```sh
dahl init
```
Use **hello-world** flag if you want a example template and configuration
```sh
dahl init --hello-world
```

The dahl-config.yml will grouping all of your template configuration
```yml
## dahl-config.yml
templates:

  ## list template
  [template name]:
    to: # destination path of generating file
    from: # location path of template file
    filename: # filename of generating file
    prefix: # prefix filename
    suffix: # suffix filename
    
    ## list props (custom keys)
    props:
      [prop name]:
```

### Commands

```sh
## base command
dahl [command] [template] [flags]
```

You can use flags to overwrite some value from the config file, like filename, destination...

```sh
Flags:
    -n, --filename   (string)
    --from           (string)  
    --to             (string)
    --props          (json)
    -p, --prefix     (string)
    -s, --suffix     (string)
    -e, --extension  (string)
    -f, --force
    -y, --yes
```
**"Filename, From, To"** value are required, you must specify them with the cli or config file.

**Info**: the config file isn't required, you can use templating with the cli this way:

```sh
dahl run 
  --from /my/template/path 
  --to /my/destination/path
  --filename my-file
```

### Template

| **pattern**    | **description**      |
|----------------|----------------------|
|  \_{ value }\_ | related to file keys |
| \_{ @value }\_ | custom props keys    |

### Examples

```sh
dahl run express-controller -n User
```

```yml
## dahl-config.yml
templates:

  express-controller:
    to: "./src/controller"
    from: "/express/controller/my-controller.dahl"
    extension: ".js"
    props:
      http:
        status: 
          ok: 200
```

```javascript
// my-controller.dahl

export const get_{ filename }_ = (req, res) => {
  // stuff
  return res.status(_{ @http.status.ok }_).send()
}

export const update_{ filename }_ = (req, res) => {
  // stuff
  return res.status(_{ @http.status.ok }_).send()
}
```

# Roadmap

See the [open issues](https://github.com/Spacelocust/dahl/issues) for a list of proposed features (and known issues).

- [Top Feature Requests](https://github.com/Spacelocust/dahl/issues?q=label%3Aenhancement+is%3Aopen+sort%3Areactions-%2B1-desc) (Add your votes using the 👍 reaction)
- [Top Bugs](https://github.com/Spacelocust/dahl/issues?q=is%3Aissue+is%3Aopen+label%3Abug+sort%3Areactions-%2B1-desc) (Add your votes using the 👍 reaction)
- [Newest Bugs](https://github.com/Spacelocust/dahl/issues?q=is%3Aopen+is%3Aissue+label%3Abug)

## Contributing

There is still a lot of work to go! Please check out the [contributing guide](CONTRIBUTING.md).
For contributor discussion about things not better discussed here in the repo, join the discord channel

## Support

Reach out to the maintainer at one of the following places:

- [GitHub discussions](https://github.com/Spacelocust/dahl/discussions)
- The email which is located [in GitHub profile](https://github.com/Spacelocust)

## License

This project is licensed under the **MIT license**.

See [LICENSE](LICENSE) for more information.

## Acknowledgements

 - [lazydocker](https://github.com/jesseduffield/lazydocker) - File structuration
 - [docker](https://docs.docker.com/engine/reference/commandline/docker) - Command-line architecture
