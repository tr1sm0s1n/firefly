---
title: Contributing to Documentation
---

# Contributing to Documentation

This guide will walk you through setting up your machine for contributing to FireFly documentation. Documentation contributions are extremely valuable. If you discover something is missing in the docs, we would love to include your additions or clarifications to help the next person who has the same question.

This doc site is generated by a set of [Markdown](https://daringfireball.net/projects/markdown/) files in the main FireFly repository, under the `./doc-site` directory. You can browse the source for the current live site in GitHub here: [https://github.com/hyperledger/firefly/tree/main/doc-site](https://github.com/hyperledger/firefly/tree/main/doc-site)

---

## Process for updating documentation

The process for updating the documentation is really easy! You'll follow the same basic steps outlined in the same [steps outlined in the Contributor's guide](./index.md#-make-changes). Here are the detailed steps for contributing to the docs:

1. Fork [https://github.com/hyperledger/firefly](https://github.com/hyperledger/firefly)
2. Clone your fork locally to your computer
3. Follow the steps below to view your local copy of the docs in a browser
4. Make some improvements to the Markdown files
5. Verify that your changes look they way you want them to in your browser
6. Create a new git commit with your changes. Be sure to sign-off on your commit by using `git commit -s`!
7. Push your changes
8. [Open a Pull Request](https://github.com/hyperledger/firefly/compare) to incorporate your changes back into the hyperledger/firefly repo

---

This FireFly documentation site is based on the [Hyperledger documentation template](https://github.com/hyperledger-labs/documentation-template). The template utilizes MkDocs (documentation at [mkdocs.org](https://www.mkdocs.org)) and the theme Material for MkDocs (documentation at [Material for MkDocs](https://squidfunk.github.io/mkdocs-material/)). Material adds a number of extra features to MkDocs, and Hyperledger repositories can take advantage of the theme's [Insiders](https://squidfunk.github.io/mkdocs-material/insiders/) capabilities.

[Material for MkDocs]: https://squidfunk.github.io/mkdocs-material/
[Mike]: https://github.com/jimporter/mike

## Prerequisites

To test the documents and update the published site, the following tools are needed:

- A Bash shell
- git
- Python 3
- The [Material for Mkdocs] theme.
- The [Mike] MkDocs plugin for publishing versions to gh-pages.
  - Not used locally, but referenced in the `mkdocs.yml` file and needed for
    deploying the site to gh-pages.

### git

`git` can be installed locally, as described in the [Install Git Guide from GitHub](https://github.com/git-guides/install-git).

### Python 3

`Python 3` can be installed locally, as described in the [Python Getting Started guide](https://www.python.org/about/gettingstarted/).

### Virtual environment

It is recommended to install your Python dependencies in a virtual environment in case you have other conflicting Python installations on your machine. This also removes the need to install these packages globally on your computer.

```bash
cd doc-site
python3 -m venv venv
source venv/bin/activate
```

### Mkdocs

The Mkdocs-related items can be installed locally, as described in the [Material
for Mkdocs] installation instructions. The short, case-specific version of those
instructions follow:

```bash
pip install -r requirements.txt
```

### Verify Setup

To verify your setup, check that you can run `mkdocs` by running the command `mkdocs --help` to see the help text.

## Useful MkDocs Commands

The commands you will usually use with `mkdocs` are:

- `mkdocs serve` - Start the live-reloading docs server.
- `mkdocs build` - Build the documentation site.
- `mkdocs -h` - Print help message and exit.

## Directory layout

    mkdocs.yml     # The configuration file.
    docs/
        index.md   # The documentation homepage.
        SUMMARY.md # The main left nav
        ...        # Other markdown pages, images and other files.