# velocity-blueprints

This repo contains examples of Velocity's blueprint annotations feature usage.

Folder structure:

* [/getting-started](getting-started) for a complete example of how to use the blueprint annotations feature when deploying your application.
* [/references](references) for a collection of snippets and references on how to customize your deployment to fit your application's needs.

Service Catalog:

* [catalog.yaml](catalog.yaml) - contains a list of the services, that the Service Catalog contains.

For adding new services to Service Catalog you need:

* Create service manifest yaml files and add it under [/references](references).
* Add new service item (with name and path to the service manifest yaml) to the [catalog.yaml](catalog.yaml) file. 