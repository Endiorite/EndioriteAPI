# EndioriteAPI

**Endiorite Open Source Minecraft Server API**  
*Created by RemBog*

---
# Open access endpoints

---

## 1. Players Stats Endpoints

---

### Get All Players Stats

```
GET http://play.endiorite.fr:8080/playersStats/getAll
```

**Example Request:**  
`http://play.endiorite.fr:8080/playersStats/getAll`

**Response:**
```
[
  {
    "xuid": "2535435907623454",
    "username": "RemBog88",
    "kills": 6,
    "deaths": 1,
    "kill_streak": 4,
    "best_kill_streak": 4,
    "playing_time": 76962
  },
  {
    "xuid": "9101919276667612",
    "username": "RemBog89",
    "kills": 2,
    "deaths": 9,
    "kill_streak": 0,
    "best_kill_streak": 2,
    "playing_time": 18209
  }
]
```

---

### Get Player Stats

```
GET http://play.endiorite.fr:8080/playersStats/get?xuid={xuid}
GET http://play.endiorite.fr:8080/playersStats/get?username={username}
```

**Query Parameters:**  
- `xuid` (string) — Player’s unique identifier (optional if username provided)  
- `username` (string) — Player’s username (optional if xuid provided)

**Example Requests:**  
- `http://play.endiorite.fr:8080/playersStats/get?xuid=2535435907623454`  
- `http://play.endiorite.fr:8080/playersStats/get?username=RemBog88`

**Response:**
```
{
  "xuid": "2535435907623454",
  "username": "RemBog88",
  "kills": 6,
  "deaths": 1,
  "kill_streak": 4,
  "best_kill_streak": 4,
  "playing_time": 76962
}
```

---

### Get Top Players by Stat

```
GET http://play.endiorite.fr:8080/playersStats/top/{topType}?page={page}&limit={limit}
```

**Path Parameter:**  
- `{topType}` — The type of stat leaderboard. Possible values:  
  - `kills`  
  - `deaths`  
  - `killStreak`  
  - `bestKillStreak`  
  - `playingTime`

**Query Parameters (optional):**  
- `page` (integer, default: 1) — Pagination page number  
- `limit` (integer, default: 10) — Number of results per page

**Example Request:**  
`http://play.endiorite.fr:8080/playersStats/top/kills?page=10&limit=5`

**Response:**
```
[
  {"username":"RemBog88","stat":160},
  {"username":"RemBog88_1","stat":135},
  {"username":"RemBog88Pro","stat":128},
  {"username":"RemBog88X","stat":122},
  {"username":"RemBog_88","stat":122}
]
```

---

## 2. Players Money Endpoints

---

### Get All Players Money

```
GET http://play.endiorite.fr:8080/playersMoney/getAll
```

**Example Request:**  
`http://play.endiorite.fr:8080/playersMoney/getAll`

**Response:**
```
[
  {
    "username": "RemBog77",
    "money": 1029103.5
  },
  {
    "username": "RemBog88",
    "money": 1921
  },
  {
    "username": "RemBog99",
    "money": 200810.12
  }
]
```

---

### Get Player Money

```
GET http://play.endiorite.fr:8080/playersMoney/get/{username}
```

**Path Parameter:**
- `{username}` — Player's username

**Example Requests:**
- `http://play.endiorite.fr:8080/playersMoney/get/RemBog88`

**Response:**
```
{"money":8245.9}
```

---

### Get Top Players by Stat

```
GET http://play.endiorite.fr:8080/playersMoney/top?page={page}&limit={limit}
```

**Query Parameters (optional):**
- `page` (integer, default: 1) — Pagination page number
- `limit` (integer, default: 10) — Number of results per page

**Example Request:**  
`http://play.endiorite.fr:8080/playersMoney/top?page=4&limit=20`

**Response:**
```
[
  {"username":"RemBog88","money":96010200.19},
  {"username":"RemBog88_1","money":135201.71},
  {"username":"RemBog88Pro","money":42820.1},
  {"username":"RemBog88X","money":8221.06},
  {"username":"RemBog_88","money":39.2}
]
```

---

## 3. Players Cosmetics Endpoints
*The system is not optimal, but in the future it will be possible to obtain information about cosmetics such as the display name, image, body part it is applied to, rarity, and whether it is a cape or costume.*

---

### Get the list of cosmetics for a username, specifying whether the player owns them or not

```
GET http://play.endiorite.fr:8080/playersCosmetics/getList/{username}
```

**Path Parameter:**
- `{username}` — Player's username

**Example Request:**  
`http://play.endiorite.fr:8080/playersCosmetics/getList/RemBog88`

**Response:**
```
{
  "Chapeaux": {
    "arch_crown": true,
    "barrel": false,
    "black_beard": true,
    ...
  },
  "Dos": {
    "angel_wings": fase,
    "arrow_quiver": false,
    "backpack_character": true,
    ...
  },
  "Main Gauche": {
    "glass_magnifying": true,
    "gun_peacemaker": true,
    "ice_septer": false,
    ...
  }
}
```

---

### Get the list of cosmetics equipped for a username

```
GET http://play.endiorite.fr:8080/playersCosmetics/getEquippedList/{username}
```

**Path Parameter:**
- `{username}` — Player's username

**Example Request:**  
`http://play.endiorite.fr:8080/playersCosmetics/getEquippedList/RemBog88`

**Response:**
```
{
  "cape": null,
  "costume": {
    "back": {
      "category": "Dos",
      "cosmetic": "angel_wings"
    },
    "head": {
      "category": "Chapeaux",
      "cosmetic": "nether_beetle_hat"
    },
    "other": null
  }
}
```

---
# Restricted access endpoints
**You must obtain authorisation from RemBog or Yass to access the following endpoints.**

---

### Check if the Discord user is linked to an in-game account

```
GET http://play.endiorite.fr:8080/userLink/check/{userId}
```

**Path Parameter:**
- `{userId}` — The Discord user ID

**Example Request:**  
```
fetch('http://play.endiorite.fr:8080/userLink/check/629679896670765073', {
    method: 'GET',
    headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer SECRET'
    }
})
```

**Response:**
```
{isLinked: true, username: RemBog88}
```

---

### Create a request to link a Discord user to an in-game account

```
POST http://play.endiorite.fr:8080/userLink/link
```

**Request Body:**
```
{
  userId: {userId}
  username: {username}
  code: {code}
}
```

**Body Parameters:**
- `{userId}` — The Discord user ID
- `{username}` — The user's in-game username
- `{code}` — The randomly generated 6-digit code

**Example Request:**  
```
fetch('http://play.endiorite.fr:8080/userLink/link', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer SECRET'
    },
    body: JSON.stringify({
        userId: "629679896670765073",
        username: "RemBog88",
        code: 123456
    })
})
```

**Response:**
```
"success": true, "message": "Link created successfully."
```

---

### Remove a Discord user's in-game account link

```
POST http://play.endiorite.fr:8080/userLink/unlink/{userId}
```

**Path Parameter:**
- `{userId}` — The Discord user ID

**Example Request:**
```
fetch('http://play.endiorite.fr:8080/userLink/unlink/629679896670765073', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer SECRET'
    }
})
```

**Response:**
```
{"success": true, "message": "Link deleted successfully."}
```

---

## Future Updates

- Additional endpoints for server status, world data, and more will be added soon. Stay tuned!

---

## Contribution
API developed free of charge, if you wish to contribute you can send via PayPal to the link below.

- https://paypal.me/rembog
