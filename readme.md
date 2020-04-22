# Access Log Service
#### This service is just to demonstrate GoLang usage, Clean Architecture, etc

<!-- About the project -->
## Project
Access-Log, is a GoLang Gin REST API service, objective of this demo application is to fetch/find duplicate users from the millions of access logs out there.

Duplicates are pairs of `user_id`s for which there are at least two *distinct* matching ip adresses in the access log.
`user_id` is considered a duplicate of itself.

### Requirements
- Each user can have multiple access records, and it is perfectly ok if there are many accesses from single ip address.
- There are no unique constraints in access log at all.
- Number of requests per user varies greatly from 1 to million or even more.
- Number of different IPs user uses are on other hand rather small - most users have 1 or 2 distinct IPs.
- Access log can be generated randomly in database or in plaintext file - it's up to you to decide.
- Log format is rougly like this: `create table conn_log (user_id bigint, ip_addr varchar(15), ts timestamp)`
- IPs are in a regular IPv4 format (4 octets in decimal delimited by dots).
- There should be no less than 10 millions of records in access log.
- Service response time should not exceed 5ms

### Example:

There are such records in conn_log:

```
1, 127.0.0.1, 17:51:59
2, 127.0.0.1, 17:52:59
1, 127.0.0.2, 17:53:59
2, 127.0.0.2, 17:54:59
2, 127.0.0.3, 17:55:59
3, 127.0.0.3, 17:55:59
3, 127.0.0.1, 17:56:59
4, 127.0.0.1, 17:57:59
```

Get request: http://localhost:12345/1/2
Response:
```json
{ "isDuplicate": true }
```


<!-- How to Run -->
## How to Run
- Clone this repo
- Replace DB details in .env
- Go to the root folder
- Run App using: go run server.go



<!-- Technologies Used -->
## Technologies Used
- GoLang
- Gin-Gonic/Gin: RESTFul Framework, we have chosen this because of its performance (than Mux and other frameworks)
- MongoDB (mgo.v2 package): NoSQL DB, the application does db-read more and also schema might require changes often, so have chosen NoSQL
- Go Testing: As a simple Unit testing framework
- Swagger: For API documentation
- Go Mod: As a package management tool
- Docker


<!-- Project structure -->
## Project structure

The project follows "Clean Architecture", you can see the folders are structured in the way so that it supports MVC separation of concerns. Project is loosely coupled, all structures are defined using interfaces. The idea is to add Inversion of Control package like sarulabs/di.

- `config`: App configuration goes here, as of now we have DB config, in future we can add Dependency Injection configs and other third party bootstraps.
- `controllers`: Contains all the Gin handlers, that deals with services basically based on request
- `middlewares`: All the middlewares that needs to be injected goes here
- `models`: Contains all the models and necessary helper methods specific to the model goes here
- `repositories`: Holds all the model manipulation codes, it gets the instance of mongodb & queries according to requirement
- `services`: All business logics required for the application goes here
- `static`: Launch page and Error page goes here
- `utils`: Utilities & helper functions required globally goes here
- `validators`: Model validation, request validations etc goes here, it can be enhanced by bringing 3rd party schema validators



<!-- Future Enhancements -->
## Future Enhancements

- Dependencies should be resolved using IoC packages like sarulabs/di etc
- Logger middleware can be replaced with some robust package
- Add Test coverage
- Go Packages dependencies can be gracefully handled using "dep"
- Segregate .env -> config based on environment variables


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Renish B - [@renishb10](https://twitter.com/renishb10) - renishb10@gmail.com

Project Link: [https://github.com/renishb10/accesslog-service](https://github.com/renishb10/accesslog-service)