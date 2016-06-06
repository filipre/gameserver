-- test data for testing

-- 4 players join the app
INSERT INTO player VALUES
  ('f808c7a7-05e8-4313-8ac9-461a128b1599', 'klose'),
  ('f0c04522-20db-4d92-8b79-531e5f80f4d1', 'Schweinsteiger'),
  ('a35c1052-04d3-4645-8301-9eff134f697f', 'Michael'),
  ('132e9567-54e2-42d0-86f5-ed6023a51432', 'Oliver_Kahn'),
  ('dfde451e-eaaa-4b75-bf9a-5466431b2fa7', 'BannMePlease');

-- at the same time, we save the time when they joined
-- may 16, 2016 15:36:38 - 19:36:38
INSERT INTO player_status VALUES
  ('d3a5ea2a-f114-483c-8774-2690b073c614', 'f808c7a7-05e8-4313-8ac9-461a128b1599', 'joined', TIMESTAMP '2016-05-16 15:36:38'),
  ('3944291c-95fc-44c8-a8bd-0f6e729b6c68', 'f0c04522-20db-4d92-8b79-531e5f80f4d1', 'joined', TIMESTAMP '2016-05-16 16:36:38'),
  ('cc13009b-1421-49f0-b7cd-3581a1b9cc86', 'a35c1052-04d3-4645-8301-9eff134f697f', 'joined', TIMESTAMP '2016-05-16 17:36:38'),
  ('cd991de6-14c1-49be-a534-25e8cab8411f', '132e9567-54e2-42d0-86f5-ed6023a51432', 'joined', TIMESTAMP '2016-05-16 18:36:38'),
  ('25e81d20-3f1e-4f04-aab0-9a00e5be48a4', 'dfde451e-eaaa-4b75-bf9a-5466431b2fa7', 'joined', TIMESTAMP '2016-05-16 19:36:38');

-- 1h later, one of them gets banned for reasons
INSERT INTO player_status VALUES
  ('e34c6358-ddb2-42df-aecc-d3ad87541abb', 'dfde451e-eaaa-4b75-bf9a-5466431b2fa7', 'banned', TIMESTAMP '2016-05-16 20:36:38');

-- klose starts a league "Klose League" (../klose-league/..)
INSERT INTO league VALUES
  ('a9fcfeea-b73e-4ceb-94d9-9953eba7e630', 'Klose League', 'klose-league');
INSERT INTO league_status VALUES
  ('905140f0-6318-43ac-bc88-0f431a71dd92', 'a9fcfeea-b73e-4ceb-94d9-9953eba7e630', 'started', TIMESTAMP '2016-05-16 21:36:38');

-- obviously, klose is owner of the "Klose League"
INSERT INTO member VALUES
  ('52597ae4-0a32-4c0e-a2dd-a32b78a9c866', 'f808c7a7-05e8-4313-8ac9-461a128b1599', 'a9fcfeea-b73e-4ceb-94d9-9953eba7e630');
INSERT INTO member_role VALUES
  ('7cfa88bd-f2f3-4e58-b065-2c8316dd8873', '52597ae4-0a32-4c0e-a2dd-a32b78a9c866', 'owner', TIMESTAMP '2016-05-16 22:36:38');

-- Schweinsteiger, Michael and Oliver_Kahn join as well
-- all start as member of course
INSERT INTO member VALUES
  ('0b118cb1-f974-4435-a77b-e3d83b763a56', 'f0c04522-20db-4d92-8b79-531e5f80f4d1', 'a9fcfeea-b73e-4ceb-94d9-9953eba7e630'),
  ('3f2864c1-a3f2-41f9-9837-459aae3e0017', 'a35c1052-04d3-4645-8301-9eff134f697f', 'a9fcfeea-b73e-4ceb-94d9-9953eba7e630'),
  ('e55f6959-9e26-4113-90ef-edeeb079e449', '132e9567-54e2-42d0-86f5-ed6023a51432', 'a9fcfeea-b73e-4ceb-94d9-9953eba7e630');
INSERT INTO member_role VALUES
  ('ed216f0b-eea1-4fbe-b7c6-82271f0723ed', '0b118cb1-f974-4435-a77b-e3d83b763a56', 'member', TIMESTAMP '2016-05-16 23:36:38'),
  ('9f277099-b1cd-4974-b3f3-f373a6653733', '3f2864c1-a3f2-41f9-9837-459aae3e0017', 'member', TIMESTAMP '2016-05-17 00:36:38'),
  ('2e62a632-7a79-46ad-92ac-e44362feac14', 'e55f6959-9e26-4113-90ef-edeeb079e449', 'member', TIMESTAMP '2016-05-17 01:36:38');

-- Oliver_Kahn, however, gets promoted
-- And Michael gets banned
INSERT INTO member_role VALUES
  ('6191bd1b-6e27-4c24-b94f-3e5173f393ea', 'e55f6959-9e26-4113-90ef-edeeb079e449', 'admin', TIMESTAMP '2016-05-17 02:36:38'),
  ('00a3dcf1-aeb9-40f5-afaf-6e8e3e1ae207', '3f2864c1-a3f2-41f9-9837-459aae3e0017', 'banned', TIMESTAMP '2016-05-17 03:36:38');

-- Klose gives Oliver_Kahn Owner rights (i.e. Klose becomes admin)
INSERT INTO member_role VALUES
  ('5672e79c-f19d-4729-8fd9-1e1c8957df81', '52597ae4-0a32-4c0e-a2dd-a32b78a9c866', 'admin', TIMESTAMP '2016-05-17 04:36:38'),
  ('8890c374-189c-4a84-b423-55b352202242', 'e55f6959-9e26-4113-90ef-edeeb079e449', 'owner', TIMESTAMP '2016-05-17 04:36:38');
