# fury-core-go-template
Our Go template for new APIs

This template implements a server struct containing a Core (services) and a Router.
The Core type is where the services are registered. All services receive a repository as a dependency.
Repositories interact with the chosen storage solution (postgres by default).