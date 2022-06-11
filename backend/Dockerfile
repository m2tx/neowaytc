FROM maven:3.8-openjdk-11 as build
WORKDIR /app
COPY mvnw .
COPY .mvn .mvn
COPY pom.xml .
COPY src src
RUN mvn install -DskipTests
FROM openjdk:11
WORKDIR /app
COPY --from=build /app/target/backend-0.0.1.jar /app
EXPOSE 8080
ENTRYPOINT ["sh", "-c"]
CMD ["java -jar backend-0.0.1.jar"] 