-- name: ListCountries :many
SELECT id, iso_code3, country_name, dialing_code
FROM countries
ORDER BY country_name;

-- name: GetCountry :one
SELECT id, iso_code3, country_name, dialing_code
FROM countries
WHERE id = $1;

-- name: GetProvincesByCountryId :many
SELECT * 
FROM provinces
WHERE country_id = $1
ORDER BY province_name;

-- name: GetProvincesById :one
SELECT * 
FROM provinces
WHERE id = $1;

-- name: GetUserByEmailOrUsername :many 
SELECT *
FROM users
WHERE email = $1 OR username = $2;