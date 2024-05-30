---
title: Introduksjon til Wolt API
description: En introduksjon til prosjektet
---

### Hva er Wolt API

Wolt API er et program som lar deg kjøre en scraper og API av Wolt Norge på din egen maskin!

Dette gjør at du kan utvikle din egen API som bare du har tilgang til, som du da kan videre utvikle slik du vil.

### Egenskaper

Programmet har 2 hoved-deler. Api-en og scraper-en, begge skrevet i Go. Det er dette som henter dataen og gir tilgang til den.

Det er også lurt å ha en postgres database å sende dataen til, eller så kan man endre på koden for å få sitt eget oppsett. 
For å se hvordan du kan bruke noe annet enn postgres, se <a href="/introduksjon/03-oppsett-beskrivelse#egen-database">oppsett beskrivelse</a>.

I tillegg til api-en og scraper-en, har programmet en eksempel nettside som henter data fra api-et når du kjører det. 
Om du kjører wolt api med docker compose, vil denne nettsiden kjøre av seg selv med docker compose.

### Oppsett

Les <a href="/introduksjon/02-raskstart">rask start</a> for å raskt sette opp programmet med docker og begynne å få data! 

Du kan også lese <a href="/introduksjon/03-oppsett-beskrivelse">oppsett beskrivelse</a> for en mere detaljert beskrivelse av programmet.
