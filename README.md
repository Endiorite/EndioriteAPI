# EndioriteAPI

**Endiorite Open Source Minecraft Server API**  
*Created by RemBog*

Welcome to the EndioriteAPI, your gateway to retrieving player statistics from the Endiorite Minecraft server. Currently, this API provides access to **PlayersStats**, with more features coming soon!

---

## Players Stats Endpoints

### Get All Players Stats

```
GET http://play.endiorite.fr:8080/playersStats/getAll
```

**Description:**  
Returns the statistics of all players in the following JSON format:

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

### Get Single Player Stats

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

## Future Updates

- Additional endpoints for server status, world data, and more will be added soon. Stay tuned!

---

If you have any questions or want to contribute, feel free to reach out or open an issue.

---

**Enjoy using EndioriteAPI!**
