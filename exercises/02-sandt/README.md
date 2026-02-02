# 🇩🇰 Sandt 🇩🇰

Sandt er et top-moderne Javascript framework til dig, der ønsker at tilgå platformen: 

**🇩🇰 Sandhedens Stemme 🇩🇰 **

## 📖 Indholdsfortegnelse

- [Sandhedens Stemme](#sandhedens-stemme)
- [Frameworket Sandt](#frameworket-sandt)
- [Læs Sandheden](#læs-sandheden)
- [Skriv Sandheden](#skriv-sandheden)
- [Slet Usandheder](#slet-usandheder)
- [Virker Sandheden?](#virker-sandheden?)
- [Biddrag til Sandheden](#biddrag-til-sandheden)

## 🇩🇰 Sandhedens Stemme 🇩🇰

Sandhedens stemme er et højreekstremistisk, nationalistisk socialt socialmedieplatform uden censur – *lad sandheden runge!*

Budgettet på den yderste højrefløj er desværre ikke hvad det har været. Der har derfor ikke været ressourcer til at bygge et front-end interface til sandhedens medie.

Operationerne på det sociale medie udelukkende gennem kommunikation til sandhedens stemme's servere. 

## 🇩🇰 Frameworket Sandt

Sandt er et lille Javascript framework, der tillader brugere at interagere med Sandhedens Stemme's server.

Specifikt faciliterer frameworket kommunikation med serveren på addressen:

`https://informatik.mads-studsgaard.com`

### 🫵 Er Sandt noget for dig?

Sandt har en lang række af funktioner til dig, den erfarne informatik elev, der er klar på en udfordring!

Sandt er dit direkte adgangskort til det frie, højreekstremistiske medie: *Sandhedens Stemme*.

Du kan både læse, skrive og slette posts med Sandt som dit Javascript framework.

## 🚀 Læs Sandheden

Først of fremmest er det vigtigt, at kunne læse sandheder, direkte fra kilden: vores database.

### 🔭 `Sandt.listPosts(): Promise<Posts[]>`

Funktionen `Sandt.listPosts()` er til dig, der ønsker at se det hele - uden filter.

Funktionen sender en `GET`-request til `/posts` på serveren.

**BEMÆRK** at funktionen er `async` – derfor kaldes den med et `await` nøgleord:

```js
const posts = await Sandt.listPosts();
```

### 🔍 `Sandt.getPost(id: number): Promise<Post>`

Funktionen `Sandt.getPost(id)` er til den selektive bruger, der ønsker at se én bestemt sandhed ad gangen.

Funktionen sender en `GET`-request til `/posts/:id` *(`:id` erstattes af det `id`, der gives som parameter til funktionen).*

**BEMÆRK** at funktionen er `async` – derfor kaldes den med et `await` nøgleord:

```js
const post = await Sandt.getPost(1);
```

## 🖋️ Skriv Sandheden

Det er ikke nok, kun at lytte til sandheden – man må tage opgaven om, at sprede det sande budskab i egen hånd!

### 🎤 `Sandt.createPost(content: string): Promise<{ id: number }>`

Funktionen `Sandt.createPost(content)` er til den, der ønsker at sprede det sande budskab.

Funktionen sender en `POST`-request til `/posts`, hvormed indholdet specificeret i `content` parameteren uploades til serveren.

**BEMÆRK** at funktionen er `async` – derfor kaldes den med et `await` nøgleord:

```js
await Sandt.createPost("Danmark til Ejderen!");
```

## 🧹 Slet Usandheder

En gang imellem skal usandheder slettes...

### 👎 `Sandt.deletePost(id: number) Promise<void>`

Funktionen `Sandt.deletePost(id)` er til dig, der tager opgaven om at holde platformens usandheder nede i egen hånd!

Funktionen sender en `DELETE`-request til `/posts/:id`, hvormed en post med det angivne `id` slettes.

**BEMÆRK** at funktionen er `async` – derfor kaldes den med et `await` nøgleord:

```js
await Sandt.deletePost(1)
```

## 📡 Virker Sandheden?

Nogle gange er det nødvendigt at undersøge, om sandheden kører som den skal.

### `Sandt.ping(): Promise<boolean>`

Funktionen `Sandt.ping()` er til dig, der vil vide, om systemet kører.

Funktionen sender en `GET`-request til `/ping` og returnerer `Promise<boolean>`

**BEMÆRK** at funktionen er `async` – derfor kaldes den med et `await` nøgleord:

```js
const isRunning = await Sandt.ping()
```

## 📟 Biddrag til Sandheden

Ønsker du at biddrage til Sandt, er dette meget let – vi tester ikke den kode, vi skriver!

Lav en PR til dette repositorie og ændringerne vil blive aktualiseret!

🇩🇰 FOR FÆDRELANDET 🇩🇰
