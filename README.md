<h3 align = "center"> Dockerized Flutter(web) + Golang(Gin) + PostgreSQL(with GORM) + Redis --> Template  </h1>

<br>
<p align = "center">
  <i>
    This is inspired from the boilerplate at https://github.com/wzslr321/flutter_web_fullstack_boileplate

 *This repository is a fully dockerized application, which by default is a `<b>`web`</b>` app, but modyfing Dockerfiles lets to run it on `<b>` mobile `</b>` device too.
    Front-end is written in `<b>` Flutter `</b>`, Back-end in `<b>` Golang with Gin `</b>` framework. Databases are `<b>` PostgreSQL (with GORM) `</b>` and `<b>` Redis ( Redigo )`</b>`. Reverse proxy with `<b>` Nginx `</b>`.
  `</i>`*

*It does the following:*

- It contains simple CRUD functions for both, Postgresql and Redis.
- It follows the OpenAPI spec mentioned in the openapi.yaml file
- It has a seed function that fills the databases from the json files in seed folder only if the DBs are empty when it starts
- Redis stores the full categories/items and are only get with full id or set
- Postgres stores the users, categories, sub-categories, tags and items references and relations. It is called for getting lists of categories/tags/items without specifying IDs
- A background asynchronous microservice (via go channels) is called for the Postgres storage while Redis is synchronous
- WebSockets push notification implementation with Kafka
- Items have a status field and delete only changes state to inactive
- An approval workflow is added for Items

---

---

## *How to launch it locally?*

*`<br/>`*

*`<b>` 1. `</b>` Make sure that you have `docker` and `docker-compose` installed on your machine.*

*`<b>` 2. `</b>` Clone the repository to the folder of your choice, by executing: `git clone https://github.com/anshumandas/flutter_go.git` in the console.*

*`<b>` 3. `</b>` Navigate to  `"/"` path of cloned repository. To ensure that you are in a right place, run `ls`. Outpout should look like this:*

> *LICENSE  README.md  all_lint_rules.yaml  analysis_options.yaml  app  data  docker-compose.yml  env  nginx  server*

#### *4.  Run : `<b>` `docker-compose up --build` `</b>` in the console.*

*`<b>` 5. `</b>` Wait for the docker installation to complete. It probably is going to take some time, depending on your internet connection.*

*When the installation will be finished, you should see this output:
![img](https://user-images.githubusercontent.com/59893892/104651401-e1aac080-56b7-11eb-8c85-a449ee4cb6c0.png)*

#### *Now just head to the `http://localhost:8080/` in the browser too see the website!*

### *Project tree ( without meaningless files)*
