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

### Running locally

Install the [App Engine SDK](https://developers.google.com/appengine/docs/go/gettingstarted/devenvironment) and run:

    goapp serve robokudos.go


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
