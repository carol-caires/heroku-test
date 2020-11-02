heroku-push:
	heroku container:push web

heroku-release:
	heroku container:release web

heroku-logs:
	heroku logs --tail
