-- Campaign Details Table
CREATE TABLE IF NOT EXISTS campaign_detail (
    campaign_id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    image_url TEXT,
    cta TEXT,
    status ENUM('ACTIVE', 'INACTIVE') NOT NULL
);