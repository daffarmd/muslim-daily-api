CREATE TABLE IF NOT EXISTS duas (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    arabic TEXT NOT NULL,
    latin TEXT NOT NULL,
    translation TEXT NOT NULL,
    category VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_duas_category ON duas (category);

CREATE TABLE IF NOT EXISTS prayer_times (
    id BIGSERIAL PRIMARY KEY,
    city_slug VARCHAR(120) NOT NULL,
    city_name VARCHAR(120) NOT NULL,
    prayer_date DATE NOT NULL,
    fajr VARCHAR(5) NOT NULL,
    dhuhr VARCHAR(5) NOT NULL,
    asr VARCHAR(5) NOT NULL,
    maghrib VARCHAR(5) NOT NULL,
    isha VARCHAR(5) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_prayer_times_city_date UNIQUE (city_slug, prayer_date)
);

CREATE INDEX IF NOT EXISTS idx_prayer_times_city_slug ON prayer_times (city_slug);
CREATE INDEX IF NOT EXISTS idx_prayer_times_prayer_date ON prayer_times (prayer_date);
