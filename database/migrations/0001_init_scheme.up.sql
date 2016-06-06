CREATE TABLE player(
  id uuid PRIMARY KEY,
  name text UNIQUE NOT NULL
);
CREATE TYPE player_status_type AS ENUM ('joined', 'activated', 'deactivated', 'banned');
CREATE TABLE player_status(
  id uuid PRIMARY KEY,
  player uuid REFERENCES player NOT NULL,
  status player_status_type NOT NULL,
  changed timestamp NOT NULL
);

CREATE TABLE league(
  id uuid PRIMARY KEY,
  name text UNIQUE NOT NULL,
  url_name text UNIQUE NOT NULL
);
CREATE TYPE league_status_type AS ENUM ('started', 'dissolved');
CREATE TABLE league_status(
  id uuid PRIMARY KEY,
  league uuid REFERENCES league NOT NULL,
  status league_status_type NOT NULL,
  changed timestamp NOT NULL
);

CREATE TABLE member(
  id uuid PRIMARY KEY,
  player uuid REFERENCES player NOT NULL,
  league uuid REFERENCES league NOT NULL
);
CREATE TYPE member_role_type AS ENUM ('member', 'banned', 'admin', 'owner');
CREATE TABLE member_role(
  id uuid PRIMARY KEY,
  member uuid REFERENCES member NOT NULL,
  role member_role_type NOT NULL,
  changed timestamp NOT NULL
);
