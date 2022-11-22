# velocity-blueprints

This repo contains examples of Velocity's blueprint annotations feature usage.

Folder structure:

* [/getting-started](getting-started) for a complete example of how to use the blueprint annotations feature when deploying your application.
* [/references](references) for a collection of snippets and references on how to customize your deployment to fit your application's needs.

Service Catalog:

* [references/kubernetes/database-containers/catalog.yaml] (catalog.yaml) - contains a list of the services, that the Service Catalog contains.
* [references/kubernetes/database-containers/] (database-containers) for the service deployment yamls files.

For adding new services to Service Catalog you need:

* Create service deployment yaml file and add under [references/kubernetes/database-containers/].
* Add new service item (with name and path to the service yaml) to the catalog.yaml. 