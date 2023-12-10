
<a name="readme-top"></a>

[![MIT License][license-shield]][license-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/Rahugg/CRM-system-go-microservices">
    <img src="pkg/crm_core/img/crm-icon.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">CRM-system-go-microservices</h3>

  <p align="center">
    Simple Customer Relationship Management helps to manage tasks, assign deadlines, control user roles, use of admin panel, feedback management, sales representation, deal management
  </p>
</div>



<!-- TABLE OF CONTENTS -->
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
        <li><a href="#tech-features">Tech features</a></li>
        <li><a href="#folder-structure">Folder Structure</a></li>
      </ul>
    </li>
     <li>
      <a href="#architecture-of-project">Architecture of project</a>
    </li>
    <li>
      <a href="#detailed">Detailed Architecture</a>
    </li>
    <li>
        <a href="#installation">Getting started</a>
        <ul>
          <li><a href="#launch">How to launch the service</a></li>
          <li><a href="#migrate">How to migrate the mock data</a></li>
        </ul>
    </li>
    <li>
     <a href="#contact">Contact</a>
    </li>
    <li>
      <a href="#reasons-why-i-used-these-technologies">Reasons of using these Technologies</a>
    </li>
    
  </ol>



<!-- ABOUT THE PROJECT -->
## About The Project
Customer Relationship Management 
### My application's features:
<ul>
  <li>Assign deadlines</li>
  <li>Contact management</li>
  <li>Control user roles</li>
  <li>Deal management</li>
  <li>Feedback management</li>
  <li>Helps to manage tasks</li>
  <li>Monitor the metrics</li>
  <li>Sales representation</li>
  <li>Use of admin panel</li>
</ul>



## Built With
* [![Golang][Golang-badge]][Golang-url] 
* [![Gin][Gin-badge]][Gin-url]
* [![gRPC][gRPC-badge]][gRPC-url]
* [![PostgreSQL][PostgreSQL-badge]][PostgreSQL-url] 
* [![Redis][Redis-badge]][Redis-url] 
* [![Kafka][Kafka-badge]][Kafka-url] 
* [![Docker][Docker-badge]][Docker-url] 
* [![Swagger][Swagger-badge]][Swagger-url] 
* [![Grafana][Grafana-badge]][Grafana-url]
* [![Prometheus][Prometheus-badge]][Prometheus-url]

## Tech features
* Concurrency
* Design Patterns
* Docker
* gRPC
* JWT Auth
* Kafka
* Linters
* Migrations
* Metrics Grafana/Prometheus
* Middleware
* PostgreSQL
* Redis
* RESTful APIs
* Swagger
* Viper Config

## Folder Structure

## Internal

### `app`
Main application logic and resource management.

### `controller`
HTTP and/or gRPC controllers for request handling.

### `docs`
Internal documentation.

### `entity`
Data structures representing domain entities.

### `repository`
Data access methods for persistence and retrieval.

### `service`
Encapsulated business logic.

### `storage`
Interaction with specific data storage systems.

## Cmd

### `auth`
Entry point for the authentication service.

### `crm_core`
Entry point for the core CRM service.

## Config

### `auth`
Configuration specific to the authentication service.

### `crm_core`
Configuration specific to the core CRM service.

## Data

### `kafka1`
Data for the Kafka service (topics, messages).

### `zookeeper`
Configuration files and state information for Zookeeper.

## Migrations

### `auth`
Database schema migration scripts for the authentication service.

### `crm_core`
Database schema migration scripts for the core CRM service.

## Pkg

### `auth`
Reusable functions and utilities specific to the authentication service.

### `crm_core`
Reusable functions and utilities specific to the core CRM service.
  
## Reasons why I used these technologies 
<summary>Technologies</summary>
  <ul>
    <li>
      <details>
        <summary>Golang</summary>
        I used Golang as a main backend language. And I wrote a backend by using Gin framework, It was easy to learn and write code in it, and Golang helped me to deal with concurrency. 
        Features where I used concurrency:
        * Graceful Shutdown
        * Kafka producer-consumer relation
        * User confirm.
        * In-memory caching
      </details>
    </li>
     <li>
      <details>
        <summary>PostgreSQL</summary>
        As a main database storage. Because of their open-source and availability, I preffered to use this database. Relational Database helped me to build relations among the entities, and it helped
        to build an application structured around a relationship between data tables.
      </details>
    </li>
     <li>
      <details>
        <summary>Redis</summary>
        As a NoSQL database, I used Redis to cache the most used and unchanged data, and this helped me to retrieve the data faster.
        This provides improved read performance (as requests can be split among the servers) and faster recovery when the primary server experiences an outage
      </details>
    </li>
      <li>
      <details>
        <summary>Kafka</summary>
        I used Kafka as a message broker. Because Kafka streams messages with very low latency and is suitable to analyze streaming data in real time. 
        It can be used as a monitoring service to raise alerts and etc.
        Kafka is suitable for my app that need to reanalyze the received data
      </details>
    </li>
    <li>
      <details>
        <summary>gRPC</summary>
         gRPC uses a binary format for data serialization and communication, which is much more efficient than traditional text-based formats such as JSON or XML. 
        This results in faster and more efficient communication between microservices.
      </details>
    </li>
    <li>
      <details>
        <summary>Docker</summary>
         Docker helps to containerize the application which can help to easy-sharing among the users and by installing some dependencies such as Redis and Kafka to project.
         Docker lets you build, test, and deploy applications quickly
         Because Docker containers encapsulate everything an application needs to run (and only those things), they allow applications to be shuttled easily between environments.
      </details>
    </li>
    <li>
      <details>
        <summary>Swagger</summary>
        Swagger allowed me to describe the structure of my APIs so that machines can read them. The ability of APIs to describe their own structure is the root of all awesomeness in Swagger.
      </details>
    </li>
    <li>
      <details>
        <summary>Grafana</summary>
         Grafana helps me to visualize the data and monitor the proccess of my app, I collect the metrics from Prometheus and visualizing them in Grafana on localhost:3000
      </details>
    </li>
     <li>
      <details>
        <summary>Prometheus</summary>
        Prometheus can collect and store metrics as time-series data, recording information with a timestamp, and I am visualizing them in Grafana
      </details>
    </li>
  </ul>
<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

## Installation

2. Clone the repo
   ```sh
   git clone https://github.com/Rahugg/CRM-system-go-microservices.git
   ```
3. Install go packages
   ```sh
   go mod tidy
   ```
4. Launch docker-compose 
   ```
   docker-compose up
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Launch
  1. To launch the auth-service: (make sure that the docker-compose is up) 
  ```
    make start-auth
  ```
  (Check makefile for other scripts)
  <br/>
  2. To launch the crm_core service: (make sure that the docker-compose is up)
  ```
    make start-crm
  ```
  (Check makefile for other scripts)
  
## Migrate
  1. To migrate the data and tables on services:(Check makefile for other scripts)
  ```
    make migrate-up
  ```
  2. To mock the database with mock data:(Check makefile for other scripts)
  ```
    make mock-data
  ```
  3. To drop all of the tables:
  ```
    make migrate-down
  ```
<p align="right">(<a href="#readme-top">back to top</a>)</p>


  
## [Architecture of project](pkg/crm_core/img/architecture-of-project.png)
## [Detailed](pkg/crm_core/img/detailed-architecture.png)

<!-- CONTACT -->
## Contact

Amanbek - [@telegram_handle](https://t.me/Rahuggg) - Rahuggg

Project Link: [https://github.com/Rahugg/CRM-system-go-microservices](https://github.com/Rahugg/CRM-system-go-microservices)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[forks-shield]: https://img.shields.io/github/forks/Rahugg/CRM-system-go-microservices.svg?style=for-the-badge
[forks-url]: https://github.com/Rahugg/CRM-system-go-microservices/network/members  
[stars-shield]: https://img.shields.io/github/stars/Rahugg/CRM-system-go-microservices.svg?style=for-the-badge
[stars-url]: https://github.com/Rahugg/CRM-system-go-microservices/stargazers
[license-shield]: https://img.shields.io/badge/license-MIT-blue
[license-url]: https://github.com/Rahugg/CRM-system-go-microservices/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/amanbek-faizolla
[Golang-badge]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Golang-url]: https://golang.org/
[Gin-badge]: https://img.shields.io/badge/Gin-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Gin-url]: https://gin-gonic.com/
[PostgreSQL-badge]: https://img.shields.io/badge/PostgreSQL-336791?style=for-the-badge&logo=postgresql&logoColor=white
[PostgreSQL-url]: https://www.postgresql.org/
[Redis-badge]: https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white
[Redis-url]: https://redis.io/
[Kafka-badge]: https://img.shields.io/badge/Apache%20Kafka-231F20?style=for-the-badge&logo=apache-kafka&logoColor=white
[Kafka-url]: https://kafka.apache.org/
[gRPC-badge]: https://img.shields.io/badge/gRPC-00ADD8?style=for-the-badge&logo=go&logoColor=white
[gRPC-url]: https://grpc.io/
[Docker-badge]: https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white
[Docker-url]: https://www.docker.com/
[Swagger-badge]: https://img.shields.io/badge/Swagger-85EA2D?style=for-the-badge&logo=swagger&logoColor=black
[Swagger-url]: https://swagger.io/
[Grafana-badge]: https://img.shields.io/badge/Grafana-F46800?style=for-the-badge&logo=grafana&logoColor=white
[Grafana-url]: https://grafana.com/
[Prometheus-badge]: https://img.shields.io/badge/Prometheus-E6522C?style=for-the-badge&logo=prometheus&logoColor=white
[Prometheus-url]: https://prometheus.io/
