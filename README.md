# Versioner

Application version is usually stored in project files, such as pom.xml, package.json etc.

Sometimes there are more than 1 file that store the same version.

This tool helps to manage version and cut a release in such cases.

[![Release](https://img.shields.io/github/v/release/mih-kopylov/versioner?style=for-the-badge)](https://github.com/mih-kopylov/versioner/releases/latest)
[![GitHub license](https://img.shields.io/github/license/mih-kopylov/versioner?style=for-the-badge)](https://github.com/mih-kopylov/versioner/blob/master/LICENSE)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/mih-kopylov/versioner/build.yml?style=for-the-badge)](https://github.com/mih-kopylov/versioner/actions/workflows/build.yml)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](http://godoc.org/github.com/mih-kopylov/versioner)
[![Go Report Card](https://goreportcard.com/badge/github.com/mih-kopylov/versioner?style=for-the-badge)](https://goreportcard.com/report/github.com/mih-kopylov/versioner)

## Usage

### Operations

```shell
versioner bump major
```

* Changes `1.2.3-SNAPSHOT` to `2.0.0`
* Changes `1.2.3` to `2.0.0`

---

```shell
versioner bump minor
```

* Changes `1.2.3-SNAPSHOT` to `1.3.0`
* Changes `1.2.3` to `1.3.0`

---

```shell
versioner bump patch
```

* Changes `1.2.3-SNAPSHOT` to `1.2.4`
* Changes `1.2.3` to `1.2.4`

---

```shell
versioner release
```

* Changes `1.2.3-SNAPSHOT` to `1.2.3`
* Changes `1.2.3` to `1.2.3`

---

```shell
versioner snapshot
```

* Changes `1.2.3-SNAPSHOT` to `1.2.3-SNAPSHOT`
* Changes `1.2.3` to `1.2.3-SNAPSHOT`

---

```shell
versioner get
```
```shell
versioner get patch
```

* For `1.2.3-SNAPSHOT` prints `1.2.3-SNAPSHOT`

---

```shell
versioner get --release
```

* For `1.2.3-SNAPSHOT` prints `1.2.3`

---

```shell
versioner get minor --release
```

* For `1.2.3-SNAPSHOT` prints `1.2`

---

```shell
versioner get major --release
```

* For `1.2.3-SNAPSHOT` prints `1`

---


### Configuration

In order for versioner to know which files contain a version, a configuration file should be created.

The configuration file is a `versioner.yaml` file by default which is taken from a current directory

```yaml
files:
    - name: helm/*.yaml
      path: $.app.version
    - name: pom.xml
      path: $.project.version
    - name: package.json
      path: $.version
    - name: custom.properties
      regexp: "(?s).*app\.version=(.+)"
debug: false
```

When it needs to read a version, the first file will be used. When it needs to write a version, every file will be
updated.

Two modes are supported: structured path and regexp for free form files.

The path format looks simlar to JsonPath:

* starts with root element - `$`
* node levels are separated with `.`
* supports json, yaml, xml based on file extension

The regexp format is helpful for unstructured text files.

