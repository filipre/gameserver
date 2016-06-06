# Game server API

Request made from the smartphone client
Not available to the user
Security: Oauth or something like that...

## Version 0.1

- `PUT: /player`
- `GET: /player`
  - `n=:number`
- `GET: /player/:playerId`
- `POST: /player/:playerId`


- `PUT: /league`
- `GET: /league`
  - `n=:number`
  - `name=:slug`
- `GET: /league/:leagueId`
- `POST: /league/:leagueId`
- `DELETE: /league/:leagueId`


- `PUT: /league/:leagueId/member`
- `GET: /league/:leagueId/member`
  - `n=:number`
  - `name=:slug`
  - `rank=:number`
  - `phone=:string`
- `GET: /league/:leagueId/member/:memberId`
- `POST: /league/:leagueId/member/:memberId`
- `DELETE: /league/:leagueId/member/:memberId`

## Version 0.2

- `PUT: /league/:leagueId/stadium`
- `GET: /league/:leagueId/stadium`
  - `n=:number`
  - `name=:slug`
- `GET: /league/:leagueId/stadium/:stadiumId`
- `POST: /league/:leagueId/stadium/:stadiumId`
- `DELETE: /league/:leagueId/stadium/:stadiumId`


- `PUT: /league/:leagueId/match`
- `GET: /league/:leagueId/match`
  - `n=:number`
  - `from=:playerId`
  - `to=:playerId`
- `GET: /league/:leagueId/match/:matchId`
- `POST: /league/:leagueId/match/:matchId`
- `DELETE: /league/:leagueId/match/:matchId`


- `POST: /league/:leagueId/match/:matchId/accept`
- `POST: /league/:leagueId/match/:matchId/counter`
- `POST: /league/:leagueId/match/:matchId/decline`
- *`POST: /league/:leagueId/match/:matchId/verify`*


## Version 0.3+

<!--
- `PUT: /league/:leagueId/member/:memberId/friends`
- `GET: /league/:leagueId/member/:memberId/friends`
  - `detailed=:boolean`
- `DELETE: /league/:leagueId/member/:memberId/friends/:memberId`
-->


- `PUT: /league/:leagueId/member/:memberId/ignoring`
- `GET: /league/:leagueId/member/:memberId/ignoring`
  - `n=:number`
  - `detailed=:boolean`
- `DELETE: /league/:leagueId/member/:memberId/ignoring/:memberId`



- `PUT: /league/:leagueId/challenge`
- `GET: /league/:leagueId/challenge`
  - `n=:number`
  - `from=:playerId`
  - `to=:playerId`
- `GET: /league/:leagueId/challenge/:challengeId`
- `POST: /league/:leagueId/challenge/:challengeId`
- `DELETE: /league/:leagueId/challenge/:challengeId`


- `POST: /league/:leagueId/challenge/:challengeId/accept`
- `POST: /league/:leagueId/challenge/:challengeId/counter`
- `POST: /league/:leagueId/challenge/:challengeId/decline`


- `GET: /league/:leagueId/ranking`
  - `i=:index`
  - `n=:number`


- `PUT: /webhook`
- `GET: /webhook`
- `GET: /webhook/:webookId`
- `POST: /webhook/:webookId`
- `DELETE: /webhook`
