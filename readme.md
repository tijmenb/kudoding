## Robokudos (working title)

Robokudos is a Slackbot to facilitate positive reinforcement at H&F Amsterdam.

### Basic idea

When Rene cleans out the dishwasher, say in Slack in the format `kudos to <@username> <reason>`:

    kudos to @rene for doing the dishes!

Robokudos will respond:

    Kudo recorded for @rene! @rene has if now ranked #3 in kudos this month

TBD: figure out what actually to do with kudos once given.

### Server

The bot works by [setting up an outgoing webhook in Slack](https://hackersfounders.slack.com/services/new/outgoing-webhook). Whenever a message starts with `kudos to` it'll be sent to Robokudos, which runs on App engine.

### TBD:

* lets get a database running
* lets create category's with a couple of commands ( suggestions are a bit lower in this doc )
* lets create a API to read the results ( something in json, so it can be directly called from angular / backbone / polymer etc. )

## Commands (proposal):

What should the app be able to do?
* kudos to @tijmen for #beerrun -> add kudo's to a user, to a specific category
* kudos new #beerrun -> create a new category ( admin function? dont want to set it open i suppose )
* kudos list -> list all category's

Furthermore, id suspect we need to aggregate all the kudo's in a timespan, this is more for the output api. Proposal for the api:
+ /api/all                     -> lsit all kudo's for all #category's at all time
+ /api/all/20140101/now        -> list all kudo's for all #category's from 20140101 until now
+ /api/all/20140101/20141101   -> list all kudo's for all #category's from 20140101 until 20141101

and you should be able to replace all with a category ( just omit the hashtag )

### Running Tests/Local

The kudos bot uses a local Redis server to connect. The Slack API is accessed to be able to translate user ids into user names.

    go get github.com/tools/godep

    go get github.com/tijmenb/kudoding

    cd $GOPATH/src/github.com/tijmenb/kudoding

    git checkout develop

    export KUDOS_TOKEN=... # Please set the environment variable with a Slack API token

    godep go test

    godep go install

    kudoding

### Running on Heroku

The kudos bot uses a RedisToGo that exposes an `REDISTOGO_URL` environment variable.

    heroku git:remote -a kudos-hf

    heroku config:add BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go.git

    heroku config:set KUDOS_TOKEN=... # generate token in Slack API integrations

    git push heroku develop:master # deploy develop branch to master on heroku

    heroku logs
