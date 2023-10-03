CREATE EXTENSION PG_TRGM;
CREATE TABLE IF NOT EXISTS PESSOAS (
    ID UUID DEFAULT gen_random_uuid(),
    Apelido VARCHAR(32) UNIQUE,
    Nome VARCHAR(100),
    NASCIMENTO CHAR(10),
    STACK TEXT,
    BUSCA TEXT GENERATED ALWAYS AS (
        lower(apelido || ' ' || nome || ' ' || STACK)
    ) STORED
);

CREATE INDEX CONCURRENTLY IF NOT EXISTS IDX_BUSCA_TGRM ON PESSOAS USING GIST (BUSCA GIST_TRGM_OPS);