<p align="center">
  <img src="https://user-images.githubusercontent.com/39099403/211023529-5cfebcd6-f822-44e4-85b2-f47ee99f41c0.svg" width="550" alt="Dahl Logo" />
</p>

![CI](https://github.com/jesseduffield/lazygit/workflows/Continuous%20Integration/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/Spacelocust/dahl)](https://goreportcard.com/report/github.com/Spacelocust/dahl)
![GitHub repo size](https://img.shields.io/github/repo-size/Spacelocust/dahl)
[![GitHub Releases](https://img.shields.io/github/downloads/Spacelocust/dahl/total)](https://github.com/Spacelocust/dahl/releases)
[![GitHub tag](https://img.shields.io/github/tag/Spacelocust/dahl.svg)](https://github.com/Spacelocust/dahl/releases/latest)

  <p align="center">A progressive <a href="http://nodejs.org" target="_blank">Node.js</a> framework for building efficient and scalable server-side applications.</p>
    <p align="center">
<a href="https://www.npmjs.com/~nestjscore" target="_blank"><img src="https://img.shields.io/npm/v/@nestjs/core.svg" alt="NPM Version" /></a>
<a href="https://www.npmjs.com/~nestjscore" target="_blank"><img src="https://img.shields.io/npm/l/@nestjs/core.svg" alt="Package License" /></a>
<a href="https://www.npmjs.com/~nestjscore" target="_blank"><img src="https://img.shields.io/npm/dm/@nestjs/common.svg" alt="NPM Downloads" /></a>
<a href="https://circleci.com/gh/nestjs/nest" target="_blank"><img src="https://img.shields.io/circleci/build/github/nestjs/nest/master" alt="CircleCI" /></a>
<a href="https://coveralls.io/github/nestjs/nest?branch=master" target="_blank"><img src="https://coveralls.io/repos/github/nestjs/nest/badge.svg?branch=master#9" alt="Coverage" /></a>
<a href="https://discord.gg/G7Qnnhy" target="_blank"><img src="https://img.shields.io/badge/discord-online-brightgreen.svg" alt="Discord"/></a>
<a href="https://opencollective.com/nest#backer" target="_blank"><img src="https://opencollective.com/nest/backers/badge.svg" alt="Backers on Open Collective" /></a>
<a href="https://opencollective.com/nest#sponsor" target="_blank"><img src="https://opencollective.com/nest/sponsors/badge.svg" alt="Sponsors on Open Collective" /></a>
  <a href="https://paypal.me/kamilmysliwiec" target="_blank"><img src="https://img.shields.io/badge/Donate-PayPal-ff3f59.svg"/></a>
    <a href="https://opencollective.com/nest#sponsor"  target="_blank"><img src="https://img.shields.io/badge/Support%20us-Open%20Collective-41B883.svg" alt="Support us"></a>
  <a href="https://twitter.com/nestframework" target="_blank"><img src="https://img.shields.io/twitter/follow/nestframework.svg?style=social&label=Follow"></a>
</p>

## Sponsors

<p align="center">
 Maintainence of this project is made possible by all the <a href="https://github.com/Spacelocust/dahl/graphs/contributors">contributors</a> and <a href="https://github.com/sponsors/Spacelocust">sponsors</a>. If you'd like to sponsor this project and have your avatar or company logo appear below <a href="https://github.com/sponsors/Spacelocust">click here</a>. 💙
</p>

## Elevator Pitch
Hello fellow developer! 

Welcome to the Dahl repository! Dahl is a command line tool that helps developers create and use custom templates to generate files quickly and easily. With Dahl, you can save time and effort by using pre-defined templates to generate the files you need for your project, rather than starting from scratch every time.

This repository contains the source code for Dahl, as well as documentation and examples to help you get started. You can also share your own templates with the community by contributing to the repository.

Whether you're a seasoned developer or just starting out, Dahl is a valuable tool to have in your arsenal. Give it a try and see how it can streamline your workflow and save you time.

Thank you for checking out Dahl. We hope it becomes a valuable addition to your toolkit.

- [Requirements](https://github.com/jesseduffield/lazydocker#requirements)
- [Installation](https://github.com/jesseduffield/lazydocker#installation)
- [Usage](https://github.com/jesseduffield/lazydocker#usage)
- [Keybindings](/docs/keybindings)
- [Cool Features](https://github.com/jesseduffield/lazydocker#cool-features)
- [Contributing](https://github.com/jesseduffield/lazydocker#contributing)
- [Video Tutorial](https://youtu.be/NICqQPxwJWw)
- [Config Docs](/docs/Config.md)
- [Twitch Stream](https://www.twitch.tv/jesseduffield)
- [FAQ](https://github.com/jesseduffield/lazydocker#faq)

## Installation

### Binary Release (Linux/OSX/Windows)

You can manually download a binary release from [the release page](https://github.com/Spacelocust/dahl/releases).

Automated install/update, don't forget to always verify what you're piping into bash:

```sh
curl https://raw.githubusercontent.com/jesseduffield/lazydocker/master/scripts/install_update_linux.sh | bash
```
The script installs downloaded binary to `$HOME/.local/bin` directory by default, but it can be changed by setting `DIR` environment variable.

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

[![Docker Pulls](https://img.shields.io/docker/pulls/spacelocust/dahl.svg)](https://hub.docker.com/r/spacelocust/dahl)
[![Docker Stars](https://img.shields.io/docker/stars/spacelocust/dahl.svg)](https://hub.docker.com/r/spacelocust/dahl)
[![Docker Automated](https://img.shields.io/docker/cloud/automated/spacelocust/dahl.svg)](https://hub.docker.com/r/spacelocust/dahl)

1. <details><summary>Click if you have an ARM device</summary><p>

    - If you have a ARM 32 bit v6 architecture

        ```sh
        docker build -t lazyteam/lazydocker \
        --build-arg BASE_IMAGE_BUILDER=arm32v6/golang \
        --build-arg GOARCH=arm \
        --build-arg GOARM=6 \
        https://github.com/jesseduffield/lazydocker.git
        ```

    - If you have a ARM 32 bit v7 architecture

        ```sh
        docker build -t lazyteam/lazydocker \
        --build-arg BASE_IMAGE_BUILDER=arm32v7/golang \
        --build-arg GOARCH=arm \
        --build-arg GOARM=7 \
        https://github.com/jesseduffield/lazydocker.git
        ```

    - If you have a ARM 64 bit v8 architecture

        ```sh
        docker build -t lazyteam/lazydocker \
        --build-arg BASE_IMAGE_BUILDER=arm64v8/golang \
        --build-arg GOARCH=arm64 \
        https://github.com/jesseduffield/lazydocker.git
        ```

    </p></details>

1. Run the container

    ```sh
    docker run --rm -it -v \
    /var/run/docker.sock:/var/run/docker.sock \
    -v /yourpath:/.config/jesseduffield/lazydocker \
    lazyteam/lazydocker
    ```

    - Don't forget to change `/yourpath` to an actual path you created to store lazydocker's config
    - You can also use this [docker-compose.yml](https://github.com/jesseduffield/lazydocker/blob/master/docker-compose.yml)
    - You might want to create an alias, for example:

        ```sh
        echo "alias lzd='docker run --rm -it -v /var/run/docker.sock:/var/run/docker.sock -v /yourpath/config:/.config/jesseduffield/lazydocker lazyteam/lazydocker'" >> ~/.zshrc
        ```



For development, you can build the image using:

```sh
git clone https://github.com/jesseduffield/lazydocker.git
cd lazydocker
docker build -t lazyteam/lazydocker \
    --build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
    --build-arg VCS_REF=`git rev-parse --short HEAD` \
    --build-arg VERSION=`git describe --abbrev=0 --tag` \
    .
```

If you encounter a compatibility issue with Docker bundled binary, try rebuilding
the image with the build argument `--build-arg DOCKER_VERSION="v$(docker -v | cut -d" " -f3 | rev | cut -c 2- | rev)"`
so that the bundled docker binary matches your host docker binary version.


## Usage

Call `lazydocker` in your terminal. I personally use this a lot so I've made an alias for it like so:

```
echo "alias lzd='lazydocker'" >> ~/.zshrc
```

(you can substitute .zshrc for whatever rc file you're using)

- Basic video tutorial [here](https://youtu.be/NICqQPxwJWw).
- List of keybindings
  [here](/docs/keybindings).

## Cool features

everything is one keypress away (or one click away! Mouse support FTW):

- viewing the state of your docker or docker-compose container environment at a glance
- viewing logs for a container/service
- viewing ascii graphs of your containers' metrics so that you can not only feel but also look like a developer
- customising those graphs to measure nearly any metric you want
- attaching to a container/service
- restarting/removing/rebuilding containers/services
- viewing the ancestor layers of a given image
- pruning containers, images, or volumes that are hogging up disk space

## Contributing

There is still a lot of work to go! Please check out the [contributing guide](CONTRIBUTING.md).
For contributor discussion about things not better discussed here in the repo, join the discord channel

<a href="https://discord.gg/ehwFt2t4wt"><img src='/docs/resources/discord.png' width='75'></a>

## Social

If you want to see what I (Jesse) am up to in terms of development, follow me on
[twitter](https://twitter.com/DuffieldJesse) or watch me program on
[twitch](https://www.twitch.tv/jesseduffield)

## FAQ
