# Neoway Technical Challenge
## How to run?
```
docker-compose up
```
## Access
Frontend: <http://localhost:80> <br/>
Backend: <http://localhost:8081>

## Frontend
### Tools
Typescript, Angular 14+, Material Angular

## BackendGo 
### Tools
Go Language

### Build
```
go mod download
go build 
```
### Tests
```
go test -cover ./...
```
### Metrics / Health Endpoints  

| URI |
| ------------------ |
| <http://localhost:8081/health> |
| <http://localhost:8081/metrics> |

### Endpoints

| Method | URI | Description | 
| ------ | ------------------ | --------------------|
| GET | /api/identificationnumber/ | Get all identification numbers |
| GET | /api/identificationnumber/{id} | Get identification number by ID |
| POST | /api/identificationnumber/ | Register a new identification number |
| PUT | /api/identificationnumber/{id} | Update identification number |
| POST | /api/identificationnumber/query/?sort={column,direction}&page={page}&size={size} |  |
| GET | /graphql | |
| POST | /graphql | | 

### GraphQl Query

| Query | RequestString |
| ----- | ------------- |
| allIdentificationNumber | query { allIdentificationNumber { id } } |
| getIdentificationNumberByID | query { getIdentificationNumberByID (id:"98b72201-446f-461f-bed4-d8193eded5ea") { id,number } } |

```
curl -X POST http://localhost:8081/graphql -H 'Content-Type: application/json' -d 'query { allIdentificationNumber { id } }'
```
```
curl -X POST http://localhost:8081/graphql -H 'Content-Type: application/json' -d 'query { getIdentificationNumberByID (id:"98b72201-446f-461f-bed4-d8193eded5ea") { id,number } }'
```