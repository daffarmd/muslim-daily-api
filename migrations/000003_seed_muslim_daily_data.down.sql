DELETE FROM prayer_times
WHERE (city_slug, prayer_date) IN (
    ('jakarta', '2026-03-23'),
    ('jakarta', '2026-03-24'),
    ('bandung', '2026-03-23'),
    ('bandung', '2026-03-24'),
    ('surabaya', '2026-03-23'),
    ('surabaya', '2026-03-24'),
    ('yogyakarta', '2026-03-23'),
    ('yogyakarta', '2026-03-24')
);

DELETE FROM duas
WHERE title IN (
    'Doa Sebelum Makan',
    'Doa Sesudah Makan',
    'Doa Bangun Tidur',
    'Doa Masuk Masjid',
    'Doa Keluar Rumah'
);
