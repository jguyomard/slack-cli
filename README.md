# Slack CLI

This package offers some CLI commands to interact with Slack.


## Get Started

### 1. Build Slack-CLI

```
go get github.com/jguyomard/slack-cli
```

Note: You need to set GOPATH correctly before running this.


### 2. [Create a new app on Slack](https://api.slack.com/apps)

Once you create this application, add `files:read` and `files:write:user` permissions.

Then, install this app: you will get an OAuth Access Token. Create `config.yaml` and edit it to add this token:

```
cp $GOPATH/src/github.com/jguyomard/slack-cli/config.yaml.sample /etc/slack-cli/config.yaml
vi /etc/slack-cli/config.yaml
```


### 3. Run this app

if `$PATH` contains `$GOPATH/bin`, `slack-cli` command is now available:

```
slack-cli --config-file=/etc/slack-cli/config.yaml
```


## Usage


### Purge old files

This command delete files older than 1 month from Slack.

```
slack-cli purge-files --dry-run
```


## Testing

To run the test suite:

```
make test
```

To run the linter:

```
make lint
```


## Issues

If you have any problems with or questions about this Service Provider, please contact me through a [GitHub issue](https://github.com/jguyomard/slack-cli/issues).


## Contributing

You are invited to contribute new features, fixes or updates to this container, through a [Github Pull Request](https://github.com/jguyomard/slack-cli/pulls).
