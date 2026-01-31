# Informatikforløb om AI

Følgende er et repositorie, der indeholder både opgaver samt kildekoden til den bagvedliggende infrastruktur til at facilitere opgaverne.

## Indholdsfortegnelse

- [Opgaver](#opgaver)
- [Dokumentation af api](#dokumentation-af-api)
- [Infrastruktur](#infrastruktur)

## Opgaver

Undervejs i forløbet kan man finde en række interaktive `.html` filer under mappen `exercises/`.

Disse opgaver fungerer som små introduktioner til emner inden for programmering. Nogle af `html` dokumenterne forventer at kommunikere med serveren samt tilhørende MySQL-database.

## Dokumentation af api

API'en svarer i `JSON`-format og forventer også at eventuelle request-bodies er i `JSON`-format.

### Ping

Ping endpointet benyttes til at sikre, at man kan etablere en forbindelse til serveren.

#### `GET /ping`

En request til endpointet `/ping` kan bruges til at checke serverens status. Serveren besvarer med: `"pong"`.

### Posts

Serveren eksponerer en række endpoints, der relaterer sig til at kunne både læse og skrive *posts* fra en database.

#### `GET /posts`

En `GET`-request til dette endpoint returnerer en liste af alle *posts*.
```json
[
    { "id": 1, "content": "Sean Connery var den ringeste James Bond!" },
    { "id": 2, "content": "Daniel Craig – all the way!" }
]
```

#### `GET /posts/:id`

Et `id` på en post kan gives som URL-parameter. En `GET`-request til `/posts/1` returnerer eksempelvis den post, der har `id = 1`, såfremt denne eksisterer:
```json
{
    "id": 1,
    "content": "Sean Connery var den ringeste James Bond!"
}
```

#### `POST /posts`

En `POST`-request til `/posts`-endpointet gør det muligt at *poste* et post til serveren. Request-bodien skal kan eksempelvis være følgende:
```json
{
    "content": "Sean Connery var den ringeste James Bond!"
}
```

Serveren besvarer forespørgslen med statuskoden `201 Created` og returnerer id'et på den post, der netop er oprettet:
```json
{ "id": 1 }
```

### Completions

Completions endpointet bruges til at generere en *chat completion*. Med dette forståes et svar i næste led af en samtale med en sprogmodel.

#### `POST /completions`

En `POST`-request til `/completions`-endpointet generer et svar fra en sprogmodel. Et eksempel på en request-bodien er:

```json
{
    "model": "mistral-medium-latest",
    "messages": [
        { "role": "user", "content": "Hvad er meningen med livet" }
    ]
}
```

## Infrastruktur

*Infrastrukturen bag projektet er ikke relevant for elever, men man er selvfølgelig velkommen til at tage et kig på filerne.*

Under mappen `infrastructure` findes kildekoden til den bagvedliggende infrastruktur. Infrastrukturen består af en MySQL-database samt en api skrevet i programmeringssproget *golang*.

Der eksisterer en README.md fil i `infrastructure/` folderen, med yderligere information omkring infrastrukturen til den interesserede.
