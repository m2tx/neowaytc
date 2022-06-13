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
| GET | /api/identificationnumber/ | findAll |
| GET | /api/identificationnumber/{id} | findById |
| POST | /api/identificationnumber/ | new |
| PUT | /api/identificationnumber/{id} | update |
| DELETE | /api/identificationnumber/{id} | delete |
| POST | /api/identificationnumber/query/ | pageable |
