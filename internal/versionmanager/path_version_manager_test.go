package versionmanager

import (
	"testing"

	"github.com/mih-kopylov/versioner/internal/pathfinder"
	"github.com/stretchr/testify/assert"
)

func TestXmlPathVersionManager(t *testing.T) {
	content := `
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.group</groupId>
    <artifactId>art-art</artifactId>
    <version>1.22.1-SNAPSHOT</version>

    <parent>
        <groupId>com.parent.group</groupId>
        <artifactId>service-parent</artifactId>
        <version>0.30.0</version>
    </parent>

</project>
`
	expectedContent := `
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.group</groupId>
    <artifactId>art-art</artifactId>
    <version>1.23.0</version>

    <parent>
        <groupId>com.parent.group</groupId>
        <artifactId>service-parent</artifactId>
        <version>0.30.0</version>
    </parent>

</project>
`
	versionManager := PathVersionManager{content, &pathfinder.XmlPathFinder{}, "$.project.version", "pom.xml"}
	version, err := versionManager.Read()
	if assert.NoError(t, err) {
		assert.Equal(t, "1.22.1-SNAPSHOT", version)
	}

	result, err := versionManager.Write("1.23.0")
	if assert.NoError(t, err) {
		assert.Equal(t, expectedContent, result)
	}
}

func TestJsonPathVersionManager(t *testing.T) {
	content := `
{
    "name": "web-web",
    "version": "1.22.1-SNAPSHOT",
    "private": true,
    "main": "index.js",
    "scripts": {
        "build": "run-s build:client build:server",
    },
    "engineStrict": true,
    "dependencies": {
        "@babel/runtime": "7.12.1",
        "zxcvbn": "4.4.2"
    }
}
`
	expectedContent := `
{
    "name": "web-web",
    "version": "1.23.0",
    "private": true,
    "main": "index.js",
    "scripts": {
        "build": "run-s build:client build:server",
    },
    "engineStrict": true,
    "dependencies": {
        "@babel/runtime": "7.12.1",
        "zxcvbn": "4.4.2"
    }
}
`
	versionManager := PathVersionManager{content, &pathfinder.JsonPathFinder{}, "$.version", "package.json"}
	version, err := versionManager.Read()
	if assert.NoError(t, err) {
		assert.Equal(t, "1.22.1-SNAPSHOT", version)
	}

	result, err := versionManager.Write("1.23.0")
	if assert.NoError(t, err) {
		assert.Equal(t, expectedContent, result)
	}
}

func TestYamlPathVersionManager(t *testing.T) {
	content := `
jxp:
  name: web-web
  image: 118522672572.dkr.ecr.eu-west-1.amazonaws.com/web-web
  version: 1.22.1-SNAPSHOT
ingress:
  enabled: true
deployment:
  replicaCount: 1
  annotations:
    ad.datadoghq.com/web-web.logs: |
      [
        {
          "source": "nodejs",
        }
      ]
  env:
  - name: ANALYTICS_COOKIE_PREFIX
    value: jga
`
	expectedContent := `
jxp:
  name: web-web
  image: 118522672572.dkr.ecr.eu-west-1.amazonaws.com/web-web
  version: 1.23.0
ingress:
  enabled: true
deployment:
  replicaCount: 1
  annotations:
    ad.datadoghq.com/web-web.logs: |
      [
        {
          "source": "nodejs",
        }
      ]
  env:
  - name: ANALYTICS_COOKIE_PREFIX
    value: jga
`
	versionManager := PathVersionManager{content, &pathfinder.YamlPathFinder{}, "$.jxp.version", "teams.yaml"}
	version, err := versionManager.Read()
	if assert.NoError(t, err) {
		assert.Equal(t, "1.22.1-SNAPSHOT", version)
	}

	result, err := versionManager.Write("1.23.0")
	if assert.NoError(t, err) {
		assert.Equal(t, expectedContent, result)
	}
}
