-- Campaign Country Targeting Table
CREATE TABLE IF NOT EXISTS campaign_country_targeting (
    campaign_id VARCHAR(100),
    country_id INT,
    inclusion_flag BOOLEAN NOT NULL, -- 1 for include, 0 for exclude
    PRIMARY KEY (
        campaign_id,
        country_id,
        inclusion_flag
    ),
    FOREIGN KEY (campaign_id) REFERENCES campaign_detail (campaign_id),
    FOREIGN KEY (country_id) REFERENCES country (country_id)
);

-- Campaign OS Targeting Table
CREATE TABLE IF NOT EXISTS campaign_os_targeting (
    campaign_id VARCHAR(100),
    os_id INT,
    inclusion_flag BOOLEAN NOT NULL, -- 1 for include, 0 for exclude
    PRIMARY KEY (
        campaign_id,
        os_id,
        inclusion_flag
    ),
    FOREIGN KEY (campaign_id) REFERENCES campaign_detail (campaign_id),
    FOREIGN KEY (os_id) REFERENCES os (os_id)
);

-- Campaign App ID Targeting Table
CREATE TABLE IF NOT EXISTS campaign_app_id_targeting (
    campaign_id VARCHAR(100),
    app_id INT,
    inclusion_flag BOOLEAN NOT NULL, -- 1 for include, 0 for exclude
    PRIMARY KEY (
        campaign_id,
        app_id,
        inclusion_flag
    ),
    FOREIGN KEY (campaign_id) REFERENCES campaign_detail (campaign_id),
    FOREIGN KEY (app_id) REFERENCES app (app_id)
);