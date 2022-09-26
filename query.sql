SELECT
    count(id)
FROM
    room_restrictions
WHERE
    '1993-12-01' < end_date AND '2002-12-12' > start_date