<!-- logo -->
<p align="center">
  <img width="300" src="logo.png">
</p>

<!-- short description -->
<h3 align="center">An all-purpose search engine for Craigslist free items</h3>
<h1 align="center">CL</h1>

<p align="center">
    <!-- license -->
    <img src="https://img.shields.io/github/license/starlightromero/cl" />
    <!-- code size  -->
    <img src="https://img.shields.io/github/languages/code-size/starlightromero/cl" />
    <!-- issues -->
    <img src="https://img.shields.io/github/issues/starlightromero/cl" />
    <!-- pull requests -->
    <img src="https://img.shields.io/github/issues-pr/starlightromero/cl" />
    <!-- number of commits per year -->
    <img src="https://img.shields.io/github/commit-activity/y/starlightromero/cl" />
    <!-- last commit -->
    <img src="https://img.shields.io/github/last-commit/starlightromero/cl" />
</p>


## Table of Contents

- [How to Run](#how-to-run)
- [Commands](#commands)


## How to Run

Clone the repo
```zsh
git clone git@github.com:starlightromero/cl.git
```

cd into the directory
```zsh
cd cl
```

Download dependencies
```zsh
go mod download
```

Run the application!
```zsh
go run .
```


## Commands

To list available commands, either run `cl` with no parameters or execute `cl help`:

```zsh
A all-purpose search engine for Craigslist free items

Commands:
  areas        Get a list of all Craigslist areas
  search       Search a Craigslist area for free items
  subregions   Get a list of all Craigslist subregions
  regions      Get a list of all Craigslist regions

Run 'cl COMMAND --help' for more information on a command.
```
