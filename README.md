# Versioner

Application version is usually stored in project files, such as pom.xml, package.json etc.

Sometimes there are more than 1 file that store the same version.

This tool helps to manage version and cut a release in such cases.

## Usage

### Operations

```shell
versioner release major
```

* `1.2.3-SNAPSHOT` to `2.0.0`
* `1.2.3` to `2.0.0`

---

```shell
versioner release minor
```

* `1.2.3-SNAPSHOT` to `1.3.0`
* `1.2.3` to `1.3.0`

---

```shell
versioner release patch
```

* `1.2.3-SNAPSHOT` to `1.2.4`

```shell
versioner story AAA-111
```

### Configuration

In order for versioner to know which files contain a version, a configution file should be created.

The configuration file is a `versioner.yaml` file by default which is taken from a current directory, or any custom file
passed with `--config` option

```yaml
profiles:
    default:
        files:
            - name: helm/*.yaml
              path: $.jxp.version
            - name: pom.xml
              path: $.project.version
            - name: package.json
              path: $.version
            - name: custom.properties
              regexp: "(?s).*app\.version=(.+)"
```

When it needs to read a version, the first file will be used. When it needs to write a version, every file will be
updated.

Two modes are supported: structured path and regexp for free form files.

The path format looks simlar to JsonPath:

* starts with root element - `$`
* node levels are separated with `.`
* supports json, yaml, xml based on file extension

The regexp format is helpful for unstructured text files.

### Profiles

The `default` profile is the one that is called by default without any additional parameters.

But any other profile could be chosen explicitly with `--profile` parameter:

```shell
versioner minor --profile custom
```
