FROM maven:3.6.3-openjdk-14-slim AS build

COPY settings.xml /usr/share/maven/conf/

COPY pom.xml pom.xml
COPY eso-api/pom.xml eso-api/pom.xml
COPY eso-model/pom.xml eso-model/pom.xml
COPY eso-base/pom.xml eso-base/pom.xml
COPY eso-database/pom.xml eso-database/pom.xml

RUN mvn dependency:go-offline package -B

COPY eso-api/src eso-api/src
COPY eso-model/src eso-model/src
COPY eso-base/src eso-base/src
COPY eso-database/src eso-database/src

RUN mvn install

FROM openjdk:14-ea-jdk-alpine
USER root

RUN mkdir service

COPY --from=build /eso-base/target/ /service/

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.5.0/wait /wait

RUN chmod +x /wait

ENV JAVA_TOOL_OPTIONS -agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=*:5005

EXPOSE 5005

CMD /wait && java --enable-preview -jar /service/eso-base-1.0-SNAPSHOT.jar -Xdebug