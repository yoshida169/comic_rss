CREATE TABLE comics (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    publisher VARCHAR(255) NOT NULL,
    latest_volume INT,
    rss_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO comics (title, author, publisher, latest_volume, rss_url) VALUES
    ('ワンピース', '尾田栄一郎', '集英社', 110, 'https://example.com/rss/onepiece'),
    ('呪術廻戦', '芥見下々', '集英社', 28, 'https://example.com/rss/jujutsukaisen'),
    ('チェンソーマン', '藤本タツキ', '集英社', 18, 'https://example.com/rss/chainsawman');
