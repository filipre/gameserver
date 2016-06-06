# Game server

Takes requests from smartphone clients and takes action or returns information

## Todos v0.1

- [ ] player package
  - [ ] `GET /player`
    - [ ] Think about parameters: `n`, `name`, ...
  - [x] `GET /player/:id`
  - [x] `POST /player`
  - [x] `POST /player/activate`
  - [x] `POST /player/deactivate`
  - [x] `POST /player/ban`
  - [ ] Overthink errors again
  - [ ] Godoc
  - [ ] Tests
- [ ] league package
  - [ ] `GET /league` and parameters
  - [ ] `GET /league/:id`
  - [ ] `POST /league`
  - [ ] `POST /league/dissolve`
- [ ] member package
  - [ ] Rethink API again (will propably not change)
  - [ ] `GET /league/:id/member`
  - [ ] `GET /league/:id/member/:id`
  - [ ] `POST /league/:id/member`
  - [ ] `POST /league/:id/member/:id/promote`
  - [ ] `POST /league/:id/member/:id/transfer`
  - [ ] `POST /league/:id/member/:id/ban`
- [ ] Testing
  - [ ] Unit tests
  - [ ] Integration tests
  - [ ] Manual tests with Postman
- [ ] Security & other stuff
  - [ ] Oauth token
  - [ ] HTTPS
  - [ ] API Version
  - [ ] Implement go logger to have at least some logging
- [ ] Deployment
 - [ ] Rewrite `.travis.yml` and `Supfile`
 - [ ] Make Docker volumes work. How?
 - [ ] Get Digitalocean account
 - [ ] Automate everything
