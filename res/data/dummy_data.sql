-- Dummy data for adserving API tests

-- Apps
INSERT INTO app (app_id, app_identifier) VALUES
  (1, 'com.abc.xyz'),
  (2, 'com.gametion.ludokinggame'),
  (3, 'com.gametion.subwaysurfer'),
  (4, 'com.gametion.candycrush'),
  (5, 'com.gametion.clashofclans'),
  (6, 'com.gametion.pubg'),
  (7, 'com.gametion.pubgmobile'),
  (8, 'com.gametion.pubgmobilelite');

-- Countries
INSERT INTO country (country_id, country_name) VALUES
  (1, 'germany'),
  (2, 'us'),
  (3, 'india'),
  (4, 'uk'),
  (5, 'canada'),
  (6, 'australia'),
  (7, 'france'),
  (8, 'italy'),
  (9, 'spain'),
  (10, 'portugal');

-- OS
INSERT INTO os (os_id, os_name) VALUES
  (1, 'android'),
  (2, 'ios'),
  (3, 'windows'),
  (4, 'macos'),
  (5, 'linux');

-- Campaigns
INSERT INTO campaign_detail (campaign_id, name, image_url, cta, status)
VALUES
  ('1', 'duolingo', 'https://duolingo.com', 'Install', 'ACTIVE'),
  ('2', 'spotify', 'https://spotify.com', 'Download', 'ACTIVE'),
  ('3', 'subwaysurfer', 'https://subwaysurfer.com', 'Play', 'ACTIVE'),
  ('4', 'Nike', 'https://nike.com', 'Play', 'ACTIVE'),
  ('5', 'Adidas', 'https://adidas.com', 'Play', 'ACTIVE'),
  ('6', 'Puma', 'https://puma.com', 'Play', 'ACTIVE'),
  ('7', 'Reebok', 'https://reebok.com', 'Play', 'ACTIVE'),
  ('8', 'BMW', 'https://bmw.com', 'Play', 'ACTIVE');

-- Targeting: duolingo for com.abc.xyz in germany on android
INSERT INTO campaign_app_id_targeting (campaign_id, app_id, inclusion_flag) VALUES
  ('1', 1, 1);
INSERT INTO campaign_country_targeting (campaign_id, country_id, inclusion_flag) VALUES
  ('1', 1, 1);
INSERT INTO campaign_os_targeting (campaign_id, os_id, inclusion_flag) VALUES
  ('1', 1, 1);

-- Targeting: spotify for com.gametion.ludokinggame in us on android
INSERT INTO campaign_app_id_targeting (campaign_id, app_id, inclusion_flag) VALUES
  ('2', 2, 1);
INSERT INTO campaign_country_targeting (campaign_id, country_id, inclusion_flag) VALUES
  ('2', 2, 1);
INSERT INTO campaign_os_targeting (campaign_id, os_id, inclusion_flag) VALUES
  ('2', 1, 1);

-- Targeting: subwaysurfer for com.gametion.ludokinggame in us on android
INSERT INTO campaign_app_id_targeting (campaign_id, app_id, inclusion_flag) VALUES
  ('3', 2, 1);
INSERT INTO campaign_country_targeting (campaign_id, country_id, inclusion_flag) VALUES
  ('3', 2, 1);
INSERT INTO campaign_os_targeting (campaign_id, os_id, inclusion_flag) VALUES
  ('3', 1, 1); 

-- Targeting: Nike for com.gametion.ludokinggame in india on ios
INSERT INTO campaign_app_id_targeting (campaign_id, app_id, inclusion_flag) VALUES
  ('4', 2, 1);
INSERT INTO campaign_country_targeting (campaign_id, country_id, inclusion_flag) VALUES
  ('4', 3, 1);
INSERT INTO campaign_os_targeting (campaign_id, os_id, inclusion_flag) VALUES
  ('4', 2, 1); 

-- Targeting: Puma for com.gametion.candycrush in uk on windows
INSERT INTO campaign_app_id_targeting (campaign_id, app_id, inclusion_flag) VALUES
  ('5', 2, 1);
INSERT INTO campaign_country_targeting (campaign_id, country_id, inclusion_flag) VALUES
  ('5', 2, 1);
INSERT INTO campaign_os_targeting (campaign_id, os_id, inclusion_flag) VALUES
  ('5', 3, 1); 

-- Targeting: Puma for com.gametion.clashofclans in portugal on linux
INSERT INTO campaign_app_id_targeting (campaign_id, app_id, inclusion_flag) VALUES
  ('6', 2, 1);
INSERT INTO campaign_country_targeting (campaign_id, country_id, inclusion_flag) VALUES
  ('6', 2, 1);
INSERT INTO campaign_os_targeting (campaign_id, os_id, inclusion_flag) VALUES
  ('6', 3, 1); 

-- Targeting: Puma for com.gametion.pubg in australia on macos
INSERT INTO campaign_app_id_targeting (campaign_id, app_id, inclusion_flag) VALUES
  ('7', 2, 1);
INSERT INTO campaign_country_targeting (campaign_id, country_id, inclusion_flag) VALUES
  ('7', 6, 1);
INSERT INTO campaign_os_targeting (campaign_id, os_id, inclusion_flag) VALUES
  ('7', 5, 1); 

-- Targeting: Puma for com.gametion.pubgmobile in canada on windows
INSERT INTO campaign_app_id_targeting (campaign_id, app_id, inclusion_flag) VALUES
  ('8', 2, 1);
INSERT INTO campaign_country_targeting (campaign_id, country_id, inclusion_flag) VALUES
  ('8', 5, 1);
INSERT INTO campaign_os_targeting (campaign_id, os_id, inclusion_flag) VALUES
  ('8', 3, 1); 

-- Targeting: BMW for com.gametion.pubgmobilelite in canada on windows
INSERT INTO campaign_app_id_targeting (campaign_id, app_id, inclusion_flag) VALUES
  ('8', 8, 1);
INSERT INTO campaign_country_targeting (campaign_id, country_id, inclusion_flag) VALUES
  ('8', 5, 1);
INSERT INTO campaign_os_targeting (campaign_id, os_id, inclusion_flag) VALUES
  ('8', 3, 1); 


