-- name: ListCountries :many
SELECT id, iso_code3, country_name, dialing_code
FROM Country
ORDER BY country_name;

-- name: GetCountry :one
SELECT id, iso_code3, country_name, dialing_code
FROM Country
WHERE id = $1;

-- name: GetProvincesByCountryId :many
SELECT * 
FROM Province
WHERE country_id = $1

-- name: GetProvincesById :one
SELECT * 
FROM Province
WHERE id = $1
