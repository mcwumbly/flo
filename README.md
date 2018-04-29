# ~ ~ ~ flo ~ ~ ~

A declarative thing doer

## Usage

Create a directory with any input files and a config file:
```sh
rm -rf /tmp/flo-test && mkdir -p /tmp/flo-test/some-input-dir
echo "some-input" > /tmp/flo-test/some-input-dir/input.txt
cp example-config.yml /tmp/flo-test/config.yml
```

Execute flo with the config:
```sh
flo -config /tmp/flo-test/config.yml
```

Check the output:
```
cat /tmp/flo-test/some-output-dir/output.txt
# some-input
```

Example config looks something like:
```yaml
---
name: some-job
tasks:
- name: some-task
  command:
    name: cp
    args:
    - some-input-dir/input.txt
    - some-output-dir/output.txt
  inputs:
  - some-input-dir
  outputs:
  - some-output-dir
```


### Contributing

Get started by cloning this repo and running the acceptance tests:

```
go get github.com/mcwumbly/flo
cd $GOPATH/src/github.com/mcwumbly/flo
ginkgo acceptance
```

If you have a question or an idea for an improvement, please open an issue.
