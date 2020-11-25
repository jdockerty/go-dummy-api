## Dummy API

A dummy API that enables you to test and utilise a multitude of other services, such as Kubernetes, AWS ECS, and many others which require you to use an application to gain a working understanding of them.

Various use cases will be pushed into particular branches from my own usages.

The available routes are:
* `/health` is expected to return a 200 response and the ID of the currently running service, this is useful when managing multiple instances of the API, to see that you are directed to different endpoints.
* `/users` will return a list of all users, this piggybacks from the `jsonplaceholder` API to provide a list of random users attached to other data.
* `/user/<user_id>` will return the data for a single user, specified by ID.

The server listens for requests on port 8080.
