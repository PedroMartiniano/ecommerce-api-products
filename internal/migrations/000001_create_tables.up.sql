CREATE TABLE IF NOT EXISTS categories(
    id VARCHAR(36),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    CONSTRAINT pk_categories_id PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS products(
    id VARCHAR(36),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price FLOAT,
    category_id VARCHAR(36),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    CONSTRAINT pk_products_id PRIMARY KEY (id),
    CONSTRAINT fk_products_category_id FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE IF NOT EXISTS stocks(
    id VARCHAR(36),
    product_id VARCHAR(36) UNIQUE,
    quantity INT,
    updated_at TIMESTAMPTZ,
    CONSTRAINT pk_stocks_id PRIMARY KEY (id),
    CONSTRAINT fk_stocks_product_id FOREIGN KEY (product_id) REFERENCES products(id)
);