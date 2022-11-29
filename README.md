# CodeQuality Converter
Convert report from SARIF to other format reports.

* [GitLab CodeQuality](https://docs.gitlab.com/ee/ci/testing/code_quality.html#implementing-a-custom-tool)
* [GitLab SAST](https://gitlab.com/gitlab-org/security-products/security-report-schemas/-/blob/master/dist/sast-report-format.json?_gl=1%2a1guihbz%2a_ga%2aOTc1NjM2NDI3LjE2NjY3MDc4NzI.%2a_ga_ENFH3X7M5Y%2aMTY2OTcwMjI3MC45LjEuMTY2OTcwMjMxMC4wLjAuMA..)
* HTML (powered by [SARIF Web Component](https://github.com/microsoft/sarif-web-component))


## Usage
Linux only!


### Install
```shell
$ wget -O sarif-converter https://gitlab.com/ignis-build/sarif-converter/-/releases/permalink/latest/downloads/bin/sarif-converter-linux
$ chmod +x sarif-converter
```


### Run
Run a static analysis tool such as [Semgrep](https://semgrep.dev/).

```shell
$ semgrep --config=auto --sarif --output=semgrep.sarif .
```

Convert to GitLab Code Quality json.

```shell
$ ./sarif-converter semgrep.sarif gl-code-quality-report.json
```

Conver to GitLab SAST json.

```
$ ./sarif-converter --type sast semgrep.sarif gl-sast-report.json
```

Convert to html report.

```
$ ./sarif-converter --type html semgrep.sarif semgrp-report.html
```


### Run in GitLab CI
```yaml
codequality:sast:
  image: $CI_TEMPLATE_REGISTRY_HOST/security-products/semgrep:3
  before_script:
    - wget -O sarif-converter https://gitlab.com/ignis-build/sarif-converter/-/releases/permalink/latest/downloads/bin/sarif-converter-linux
    - chmod +x sarif-converter
  script:
    - /analyzer run
    - ./sarif-converter semgrep.sarif gl-code-quality-report.json
  artifacts:
    reports:
      codequality: gl-code-quality-report.json
```

![](docs/gitlab-merge-request.png)


## License
MIT
