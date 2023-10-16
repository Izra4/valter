-- name: Dummy :execresult
INSERT INTO product(name,description,link,book_pict)
    VALUES ('AI Assistant',
            'Asisten Kesehatan kami adalah teknologi yang memadukan kecerdasan buatan dengan robotika untuk memberikan dukungan dalam berbagai tugas perawatan kesehatan. Ini mencakup bantuan administratif, perawatan pasien, dan pemantauan kondisi pasien.',
            'assist',
            'assistBook.jpg');

-- name: Dummy2 :execresult
INSERT INTO product(name,description,link,book_pict)
    VALUES('AI Diagnostic','Produk AI kami adalah solusi canggih yang dirancang untuk mendukung berbagai aspek perawatan kesehatan. Ini mencakup diagnosa penyakit yang cepat dan akurat, manajemen data medis yang efisien, serta analisis data yang mendalam untuk perencanaan perawatan yang lebih baik.',
        'diagnostic','diagBook.jpg');

-- name: GetAllProducts :many
SELECT * FROM product;

-- name: GetProductsById :one
SELECT * FROM product
    WHERE id = ?;