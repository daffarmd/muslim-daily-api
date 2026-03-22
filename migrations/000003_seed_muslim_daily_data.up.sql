INSERT INTO duas (title, arabic, latin, translation, category)
VALUES
    (
        'Doa Sebelum Makan',
        'اللهم بارك لنا فيما رزقتنا',
        'Allahumma barik lana fima razaqtana',
        'Ya Allah, berkahilah rezeki yang Engkau berikan kepada kami.',
        'daily'
    ),
    (
        'Doa Sesudah Makan',
        'الحمد لله الذي أطعمنا وسقانا',
        'Alhamdulillahil ladzi ath''amana wa saqana',
        'Segala puji bagi Allah yang telah memberi kami makan dan minum.',
        'daily'
    ),
    (
        'Doa Bangun Tidur',
        'الحمد لله الذي أحيانا بعد ما أماتنا',
        'Alhamdulillahil ladzi ahyana ba''da ma amatana',
        'Segala puji bagi Allah yang telah menghidupkan kami setelah mematikan kami.',
        'morning'
    ),
    (
        'Doa Masuk Masjid',
        'اللهم افتح لي أبواب رحمتك',
        'Allahummaftah li abwaba rahmatik',
        'Ya Allah, bukakanlah untukku pintu-pintu rahmat-Mu.',
        'worship'
    ),
    (
        'Doa Keluar Rumah',
        'بسم الله توكلت على الله',
        'Bismillahi tawakkaltu ''alallah',
        'Dengan nama Allah, aku bertawakal kepada Allah.',
        'travel'
    )
ON CONFLICT DO NOTHING;

INSERT INTO prayer_times (city_slug, city_name, prayer_date, fajr, dhuhr, asr, maghrib, isha)
VALUES
    ('jakarta', 'Jakarta', '2026-03-23', '04:32', '12:01', '15:15', '18:05', '19:15'),
    ('jakarta', 'Jakarta', '2026-03-24', '04:31', '12:01', '15:15', '18:04', '19:14'),
    ('bandung', 'Bandung', '2026-03-23', '04:35', '12:03', '15:16', '18:07', '19:16'),
    ('bandung', 'Bandung', '2026-03-24', '04:34', '12:03', '15:16', '18:06', '19:16'),
    ('surabaya', 'Surabaya', '2026-03-23', '04:11', '11:33', '14:47', '17:37', '18:46'),
    ('surabaya', 'Surabaya', '2026-03-24', '04:10', '11:32', '14:46', '17:36', '18:45'),
    ('yogyakarta', 'Yogyakarta', '2026-03-23', '04:18', '11:42', '14:56', '17:46', '18:55'),
    ('yogyakarta', 'Yogyakarta', '2026-03-24', '04:17', '11:41', '14:55', '17:45', '18:54')
ON CONFLICT (city_slug, prayer_date) DO NOTHING;
