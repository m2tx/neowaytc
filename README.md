# Neoway Technical Challenge
## How to run?
```
docker-compose up
```
## Access
Frontend: <http://localhost:80> <br/>
Backend: <http://localhost:8080>

## Frontend
### Tools
Typescript, Angular 14+, Material Angular

## Backend 
### Tools
Java 11+, Maven, Springboot 2.7

### Build / Package
```
mvn package
```
### Tests
```
mvn test
```
### Metrics / Health Endpoints  

| URI |
| ------------------ |
| <http://localhost:8080/actuator/health> |
| <http://localhost:8080/actuator/metrics> |
| <http://localhost:8080/actuator/metrics/http.server.requests> |
| <http://localhost:8080/actuator/metrics/application.ready.time> |

### API Endpoints

| Method | URI | Description |
| ------ | ------------------ | --------------------|
| Get | / | findAll |
| Get | /{id} | findById |
| Post | / | new |
| Put | /{id} | update |
| Delete | /{id} | delete |
| Post | /query/ | pageable |