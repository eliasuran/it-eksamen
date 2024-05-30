---
title: Rask start 
description: Få programmet opp og kjørende så raskt som mulig
---

### Forutsetninger

Det er noen forutseninger for å bruke Wolt API. Wolt API er fortsatt veldig nybegynner vennlig, og du trenger ikke vite mye om teknologiene som er brukt.

* git lastet ned
* docker lastet ned
* en database satt opp (helst postgres)

### Klone fra Github

Det første vi må gjøre er å få programmet lokalt. Dette kan vi gjøre ved å klone wolt api repositoriet fra github.

Kjør denne kommandoen i terminalen for å klone prosjektet fra github
```bash
git clone https://github.com/eliasuran/it-eksamen.git
```

Nå har du programmet lokalt på maskinen i et folder som heter it-eksamen.

### Docker compose

Den beste og raskeste måten å sette opp wolt api lokalt eller på en server, er med docker compose.

For å gjøre dette må vi lage en docker-compose fil i rooten av prosjektet

~/docker-compose.yaml
```yaml
name: it-eksamen # navnet på programmet

services: # her defineres de ulike programmene som skal kjøre
  api:
    environment: # her defineres environment variabler
      - DATABASE_URL=din_database_url # database url er det eneste som trengs av environment variabler og burde være satt til din database sin connection string
      - PORT=8080 # porten api-en vil kjøre på
    build: ./api # linker til et folder hvor det er en Dockerfile
    ports:
      - "8080:8080" # mapper porten som exposes av api-en til en på maskinen
    depends_on:
      - scraper # api-en kjøres etter scraper er ferdig

  scraper:
    environment:
      - DATABASE_URL=din_database_url
    build: ./scraper 

  # eksempel nettsiden som ligger i /website
  website:
    environment:
      - DATABASE_URL=din_database_url
    build: ./website # linker til et folder hvor det er en Dockerfile
```

Nå som vi har en docker-compose fil, kan vi kjøre hele programmet med:

```bash
sudo docker compose up
```
