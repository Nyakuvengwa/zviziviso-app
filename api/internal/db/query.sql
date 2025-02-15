-- name: ListCountries :many
SELECT id, iso_code3, name, dialing_code
FROM countries
ORDER BY name  
LIMIT $1 OFFSET $2;

-- name: GetCountry :one
SELECT id, iso_code3, name, dialing_code
FROM countries
WHERE id = $1;

