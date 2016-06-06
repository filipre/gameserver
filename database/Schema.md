# Database Scheme

**dprecated, please refer to the scheme.sql**

valid until I acutal implement the scheme

- Player
  - `id`: Uuid
  - `name`: String
  - `joined`: Date
  - `phone`: String
  - `picture`: Path
  - `email`: String
  - `password`: SHA
- PlayerStatus **[Log]**
  - `id`: Uuid
  - `player`: Player Uuid
  - `date`: Date
  - `status`: (activated, deactivated, banned)
- PlayerIgnoring
  - `id`: Uuid
  - `ignoring`: Player Uuid
- League
  - `id`: Uuid
  - `name`: String
  - `picture`: Path
- LeagueStatus **[Log]**
  - `id`: Uuid
  - `league`: League Uuid
  - `status`: (started, dissolved)
- Member
  - `id`: Uuid
  - `league`: League Uuid
  - `player`: Player Uuid
  - `altNickname`: String (overwrites Player.nickname)
- MemberRole **[Log]**
  - `id`: Uuid
  - `member`: Member Uuid
  - `date`: Date
  - `role`: (member, banned, admin, owner)
- MemberElo **[Log]**
  - `id`: Uuid
  - `member`: Member Uuid
  - `after`: Match Uuid
  - `elo`: Integer
- Stadium
  - `id`: Uuid
  - `name`: String
  - `picture`: Path
  - `latitude`: Float
  - `longitude`: Float
- Match
  - `id`: Uuid
  - `league`: League Uuid
  - `stadium`: Stadium Uuid
  - `verified`: Boolean
- MatchResult
  - `id`: Uuid
  - `match`: Match Uuid
  - `memberA`: Member Uuid
  - `memberB`: Member Uuid
  - `pointsAB`: Integer
  - `memberC`: Member Uuid
  - `memberD`: Member Uuid
  - `pointsCD`: Integer
  - `entered`: Date
- MatchConfirmation
  - `id`: Uuid
  - `matchResult`: MatchResult Uuid
  - `confirmer`: Member Uuid
  - `status`: (proposed, accepted, countered, declined)
- Challenge
  - `id`: Uuid
  - `league`: League Uuid
  - `challenger`: Member Uuid
  - `opponent`: Member Uuid
- StadiumBelt **[Log]**
  - `id`: Uuid
  - `stadium`: Stadium Uuid
  - `after`: Match Uuid
  - `stadiumBelt`: Member Uuid
- TheBelt **[Log]**
  - `id`: Uuid
  - `after`: Match Uuid
  - `theBelt`: Member Uuid
- InviteUrl
  - `id`: Uuid
  - `random`: String
- Webhook
  - `id`: Uuid
  - `action`: (user, league, ...)
  - `url`: Url

#### Example

|Match|Stadium|
|-|-|
|#1|BoulderHaus|

|MatchResult|PlayerA|PlayerB|PointsAB|PlayerC|PlayerD|PointsCD|
|-|-|-|-|-|-|-|
|#1|Ballack|Klose|10|Lahm|Özil|0

|MatchConfirmation|MatchResult|Confirmer|Status|
|-|-|-|-|
|#1|#1|Ballack|proposed|
|#2|#1|Lahm|accepted|
|#3|#1|Özil|countered|
|#4|#1|Klose|declined|
