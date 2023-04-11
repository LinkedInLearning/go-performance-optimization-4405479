SELECT
    sku
    -- TODO: More columns
FROM
    items
WHERE
    sku = $1
;
