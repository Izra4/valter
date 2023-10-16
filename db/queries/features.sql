-- name: FeatDummy :execresult
INSERT INTO features(title,description,productId)
    VALUES ('Admin','Robot asisten kami dapat membantu dalam tugas-tugas administratif rumah sakit seperti pengiriman pesan, pengarsipan, dan pencatatan data.',
            '1');

-- name: FeatDummy2 :execresult
INSERT INTO features(title,description,productId)
    VALUES ('Perawatan','Robot kami dilengkapi dengan teknologi canggih yang memungkinkan mereka memberikan perawatan dasar kepada pasien.',
            '1');

-- name: FeatDummy3 :execresult
INSERT INTO features(title,description,productId)
    VALUES('Pemantauan',' Robot kami dapat memantau kondisi pasien secara real-time dan memberikan laporan kepada tim medis.',
           '1');

-- name: FeatDummy4 :execresult
INSERT INTO features(title,description,productId)
    VALUES('Tepat','Diatur untuk mematuhi prosedur dan protokol dengan ketat, yang meningkatkan keselamatan pasien dan memastikan bahwa perawatan sesuai dengan standar tertinggi.',
           '1');

-- name: FeatDummy5 :execresult
INSERT INTO features(title,description,productId)
    VALUES('Cepat','Teknologi kami dilengkapi dengan algoritma cerdas yang memungkinkan diagnosis penyakit dalam waktu singkat dan dengan akurasi yang tinggi. ',
           '2');

-- name: FeatDummy6 :execresult
INSERT INTO features(title,description,productId)
    VALUES('Manajemen',' Solusi kami menghadirkan sistem manajemen data medis yang aman dan efisien. Ini memungkinkan penyimpanan data medis yang terstruktur.','2');

-- name: FeatDummy7 :execresult
INSERT INTO features(title,description,productId)
    VALUES('Identifikasi','AI kami dapat menganalisis data medis pasien untuk mengidentifikasi pola dan tren yang unik. Ini membantu dalam merancang rencana perawatan yang lebih disesuaikan dengan kebutuhan pasien, meningkatkan hasil perawatan mereka.','2');

-- name: FeatDummy8 :execresult
INSERT INTO features(title,description,productId)
    VALUES('Analisis','Solusi kami juga dilengkapi dengan alat analitik yang kuat yang memungkinkan penyedia perawatan. Ini membantu dalam mengidentifikasi pola epidemiologi, risiko pasien, dan tren perawatan.','2');

-- name: ShowFeatures :many
SELECT
    p.id AS product_id,
    p.createdAt AS product_createdAt,
    p.updatedAt AS product_updatedAt,
    p.deletedAt AS product_deletedAt,
    p.name AS product_name,
    p.description AS product_description,
    p.link AS product_link,
    p.book_pict AS product_book_pict,
    f.id AS feature_id,
    f.createdAt AS feature_createdAt,
    f.updatedAt AS feature_updatedAt,
    f.deletedAt AS feature_deletedAt,
    f.title AS feature_title,
    f.description AS feature_description
FROM
    product p
        LEFT JOIN
    features f ON p.id = f.productId
WHERE
        p.id = ?;

-- name: GetFeatures :many
select * from features
    where  productId = ?;
