UPDATE product_groups hapg, (
    SELECT * FROM product_groups cpg 
    WHERE cpg.platform_id = 2
    AND cpg.id IN (
        SELECT p.id FROM products p
        WHERE p.platform_id = 2
        AND p.code IN ('C CODE')
    ) as cpg
)
SET fields = 'values'
WHERE pg.platform_id = 1
AND pg.id IN (
    SELECT * FROM products p 
    WHERE p.platform_id = 1
    AND p.code IN ('HA CODE')
)


UPDATE product_groups hapg, product_groups cpg
SET fields = 'values'
WHERE hapg.platform_id = 1
AND cpg.platform_id = 2
AND (hapg.code, cpg.code) IN SELECT (('HA CODE', 'C CODE'))

INSERT omsth