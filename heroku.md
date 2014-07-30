# Heroku Deployment

  heroku git:remote -a kudos-hf
  git push heroku develop
  heroku config:add BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go.git
  git push heroku develop:master
  heroku config:set KUDOS_TOKEN=... # generate token in Slack API integrations
  heroku logs
