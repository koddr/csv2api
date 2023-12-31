# csv2api – Parse CSV with filtering and sending to API

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Wiki][wiki_img]][wiki_url]
[![License][license_img]][license_url]

The **csv2api** parser reads the CSV file with the raw data, filters the 
records, identifies fields to be changed, and sends a request to update the 
data to the specified endpoint of your REST API. 

All actions take place according to the settings in the configuration file.

Features:

- 100% **free** and **open source**.
- Works with any **size** of input CSV file.
- Use **any** configuration file format: JSON, YAML, TOML, or HCL (Terraform).
- Ability to keep a configuration file (_and, in the future, input data 
  file_) on a **remote server** with HTTP access, it will be read as if it 
  was in a folder on your local machine. 
- Configure **any** request body for the REST API endpoint directly in the 
  configuration file (in a clear declarative style).
- Extensive options for **filtering** incoming data from a CSV file.
- Provides extensive capabilities for constructing **multiple filters** to 
  accurately perform actions on selected fields.

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
  - Make sure that the first column name in the `columns_order` section is a 
    primary key (PK) for your process.
  - Set up your API (base URL, token, endpoints, etc) in the `api` section.
  - Set up the filter for your fields in the `filter_columns` section.
  - Set up fields to be updated in the `update_fields` section.
- In your input `data.csv`:
  - Make sure that the first line of your CSV file contains the correct field names.

> 💡 Note: See the repository's [Wiki][wiki_url] page to understand the
> structure of the config and input data files.

And now, run `csv2api` with options:

```console
csv2api \
  -c /path/to/config.yml \
  -d /path/to/data.csv \
  -e CONFIG
```

Done! 🎉 Your transactions have been performed:

``` console
Hello and welcome to csv2api! 👋
                                
– According to the settings in './config.yml', 5 transactions were filtered out of 10 to start the process.
– Only 3 transactions got into the final set of actions to be taken... Please wait!
                                                                                                                                
 ✓ Field 'tags' with values '[paid]' in the transaction '2' has been successfully updated (HTTP 200)
 ✓ Field 'tags' with values '[paid]' in the transaction '8' has been successfully updated (HTTP 200)
 ✓ Field 'tags' with values '[paid]' in the transaction '10' has been successfully updated (HTTP 200)
                                                                                                
– Saving filtered transactions to CSV file './filtered-1686993960.csv' in the current dir... OK!
                                
All done! 🎉 Time elapsed: 0.11s
```

### 🐳 Docker-way to quick start

If you don't want to physically install `csv2api` to your system, you feel
free to using our [official Docker image][docker_image_url] and run it from
isolated container:

```console
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} koddr/csv2api:latest [OPTIONS]
```

## 🧩 Options

| Option | Description                                                                        | Is required? | Default value |
|--------|------------------------------------------------------------------------------------|--------------|---------------|
| `-i`   | set to generate initial config (`config.yaml`) and example data (`data.csv`) files | no           | `false`       |
| `-c`   | set path to your config file                                                       | yes          | `""`          |
| `-d`   | set path to your CSV file with input data                                          | yes          | `""`          |
| `-e`   | set prefix used in your environment variables                                      | no           | `CONFIG`      |

## ✨ Solving case

In my work, I often have to work with **large amounts** of raw data in CSV format. 

Usually it goes like this:

1. Unload a file with data from one system.
2. Clean up this file from duplicates and unnecessary columns.
3. Make some changes in some columns of some rows.
4. Mapping the processed lines from CSV file to the database structure fields.
5. Write a function to bypass the CSV file and form the request body.
6. Write an HTTP client that will send requests to the REST API endpoint.
7. Send prepared request body to the REST API endpoint in other system 
   for specified DB records.

> And I'm not talking about the fact that the final REST API (where to send a 
request with the processed data) **do not** always have the same parameters for 
the request body.

To ease this whole process, I created this parser that takes absolutely any 
data file as input, does the conversions and filtering, and is set up in one 
single configuration file. 

Just prepare the data, set the configuration to your liking, run `csv2api` 
and wait a bit! Yes, it's that simple.

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
