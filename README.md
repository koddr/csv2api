# csv2api – 

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
![Code coverage][go_code_coverage_img]
[![Wiki][wiki_img]][wiki_url]
[![License][license_img]][license_url]

**csv2api** description ...

Features:

- 100% **free** and **open source**
- ...
- ...

## ⚡️ Quick start

First, [download][go_download] and install **Go**. Version `1.20` (or higher)
is required.

Installation is done by using the [`go install`][go_install] command:

```console
go install github.com/koddr/csv2api@latest
```

> 💡 Note: See the repository's [Release page][repo_releases_url], if you want
> to download a ready-made `deb`, `rpm`, `apk` or `Arch Linux` package.

GNU/Linux and macOS users available way to install via [Homebrew][brew_url]:

```console
# Tap a new formula:
brew tap koddr/tap

# Installation:
brew install koddr/tap/csv2api
```

Next, run `csv2api` with `-i` option to generate initial `config.yml` and 
`data.csv` files in the current dir:

```console
csv2api -i
```

Prepare your config and input data files:

- In your `config.yml`:
  - ...
- In your input `data.csv`:
  - ...

> 💡 Note: See the repository's [Wiki][wiki_url] page to understand the
> structure of the config and input data files.

And now, run `csv2api` with options:

```console
csv2api -i
```

Done! 🎉 Your tasks have been executed:

``` console
2023/06/15 15:51:19 👋 The processing and filtering of 10 transactions begins.
2023/06/15 15:51:19 – According to the settings, 1 transactions were filtered to update...
2023/06/15 15:51:19 – Starting the process of updating filtered transactions! This may take some time...
2023/06/15 15:51:19 ✓ Field 'tags' with values '[second_touch]' in the transaction '1' has been successfully updated (HTTP 200)
2023/06/15 15:51:20 – Saving filtered transactions to CSV file ('./filtered-1686833479.csv') in the current dir... OK!
2023/06/15 15:51:20 ✌️ Done! All transactions from the filtered list have been successfully updated (elapsed time: 0.29s).
```

### 🐳 Docker-way to quick start

If you don't want to physically install `csv2api` to your system, you feel
free to using our [official Docker image][docker_image_url] and run it from
isolated container:

```console
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} koddr/csv2api:latest [COMMANDS]
```

## ✨ Usage

Basic usage and full code examples of all functions of the `csv2api` 
package, you can find on the [pkg.go.dev][go_dev_url] page.

## 🧩 Options

| Option | Description                                                                                                                  | Is required? | Default value  |
|--------|------------------------------------------------------------------------------------------------------------------------------|--------------|----------------|
| `-i`   | set to generate initial config (`config.yaml`) and example data (`data.csv`) files                                           | no           | `false`        |
| `-c`   | set path to your config file                                                                                                 | yes          | `./config.yml` |
| `-d`   | set path to your CSV file with input data                                                                                    | yes          | `./data.csv`   |
| `-e`   | set prefix used in your environment variables                                                                                | no           | `CONFIG`       |

## ✨ Solving case

...

## 💡 Motivation

...

## 🏆 A win-win cooperation

And now, I invite you to participate in this project! Let's work **together** to
create the **most useful** tool for developers on the web today.

- [Issues][repo_issues_url]: ask questions and submit your features.
- [Pull requests][repo_pull_request_url]: send your improvements to the current.

Your PRs & issues are welcome! Thank you 😘

## ⚠️ License

[`csv2api`][repo_url] is free and open-source software licensed 
under the [Apache 2.0 License][license_url], created and supported with 🩵 
for people and robots by [Vic Shóstak][author].

[go_download]: https://golang.org/dl/
[go_install]: https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies
[go_version_img]: https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none
[go_report_url]: https://goreportcard.com/report/github.com/koddr/csv2api
[go_code_coverage_img]: https://img.shields.io/badge/code_coverage-0%25-success?style=for-the-badge&logo=none
[go_dev_url]: https://pkg.go.dev/github.com/koddr/csv2api
[docker_image_url]: https://hub.docker.com/repository/docker/koddr/csv2api
[brew_url]: https://brew.sh
[wiki_img]: https://img.shields.io/badge/docs-wiki_page-blue?style=for-the-badge&logo=none
[wiki_url]: https://github.com/koddr/csv2api/wiki
[license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[license_url]: https://github.com/koddr/csv2api/blob/main/LICENSE
[repo_url]: https://github.com/koddr/csv2api
[repo_releases_url]: https://github.com/koddr/csv2api/releases
[repo_issues_url]: https://github.com/koddr/csv2api/issues
[repo_pull_request_url]: https://github.com/koddr/csv2api/pulls
[author]: https://github.com/koddr
