# heroku-test
![enter image description here](https://github.com/egonelbre/gophers/blob/master/sketch/friends/hugging-docker.png | width=200)

Just a Go *Hello World* service configured to run in Heroku trough Docker.

## Running local
```
make run-local
```
To see the logs, run:
```
make logs-local
```
## Deployment
This service is configured to run in Heroku trough Docker.
To deploy and release a new version, run:
```
make heroku-push
make heroku-release
```
To see the logs, run:
```
make heroku-logs
```
To check if it worked, open the app URL in Heroku:
**https://*[your-app-name]*.herokuapp.com/**
