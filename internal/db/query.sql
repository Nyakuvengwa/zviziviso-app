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

-- name: CreateUserDetails :one
INSERT INTO users (username, email, password_hash, first_name, last_name, role)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING user_id;

-- name: UpdateUserDetails :exec
UPDATE users
SET username = $2, first_name = $3, last_name = $4, role = $5
WHERE user_id = $1;

-- name: UpdateUserPassword :exec
UPDATE users
SET password_hash = $2
WHERE user_id = $1;

-- name: GetUserSummaryDetails :one
SELECT user_id, username, first_name, last_name, role
FROM users
WHERE user_id = $1;

-- name: CreateNewDeathNotice :one
INSERT INTO death_notices (full_name, date_of_death, age, cause_of_death, funeral_parlour_id, address_id, obituary, image_url)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING death_notice_id;

-- name: GetDeathNoticeById :one
SELECT *
FROM death_notices
WHERE death_notice_id = $1;

-- name: CreateNewFuneralParlour :one 
INSERT INTO funeral_parlours (name, address, contact_number, email)
VALUES ($1, $2, $3, $4)
RETURNING funeral_parlour_id;

-- name: GetFuneralParlourById :one
SELECT *
FROM funeral_parlours
WHERE funeral_parlour_id = $1;

--