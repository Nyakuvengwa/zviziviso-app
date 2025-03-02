-- Country Table
CREATE TABLE countries (
    id SERIAL PRIMARY KEY,
    iso_code3 CHAR(3) UNIQUE NOT NULL,
    country_name VARCHAR(255) NOT NULL,
    dialing_code VARCHAR(10)
);

-- Province Table
CREATE TABLE provinces (
    id SERIAL PRIMARY KEY,
    country_id INTEGER REFERENCES countries(id) NOT NULL,
    province_name VARCHAR(255) NOT NULL,
    code VARCHAR(50) UNIQUE
);

-- DeathNotice Table
CREATE TABLE death_notices (
    death_notice_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name VARCHAR(255) NOT NULL,
    date_of_death DATE NOT NULL,
    age INTEGER,
    cause_of_death TEXT,
    funeral_parlour_id UUID REFERENCES funeral_parlours(funeral_parlour_id),
    address_id UUID REFERENCES addresses(address_id),
    obituary TEXT,
    image_url VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- FuneralParlour Table
CREATE TABLE funeral_parlours (
    funeral_parlour_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    contact_number VARCHAR(20),
    email VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Address Table
CREATE TABLE addresses (
    address_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    address_type VARCHAR(50) CHECK (address_type IN ('Church', 'Home', 'Event')),
    address VARCHAR(255) NOT NULL,
    city VARCHAR(100),
    province VARCHAR(100),
    postal_code VARCHAR(20),
    contact_person VARCHAR(255),
    contact_number VARCHAR(20),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Users Table
CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    role VARCHAR(50) CHECK (role IN ('Admin', 'User')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- DeathNotice Comments Table
CREATE TABLE death_notice_comments (
    comment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    death_notice_id UUID REFERENCES death_notices(death_notice_id),
    user_id UUID REFERENCES users(user_id),
    comment_text TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- DeathNotice Flowers/Condolences table
CREATE TABLE death_notice_condolences (
    condolence_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    death_notice_id UUID REFERENCES death_notices(death_notice_id),
    user_id UUID REFERENCES users(user_id),
    condolence_text TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- DeathNotice Donations Table
CREATE TABLE death_notice_donations (
    donation_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    death_notice_id UUID REFERENCES death_notices(death_notice_id),
    user_id UUID REFERENCES users(user_id),
    amount DECIMAL(10,2) NOT NULL,
    donation_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- DeathNotice Events Table
CREATE TABLE death_notice_events (
    event_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    death_notice_id UUID REFERENCES death_notices(death_notice_id),
    event_type VARCHAR(255) NOT NULL,
    event_date TIMESTAMP WITH TIME ZONE NOT NULL,
    address_id UUID REFERENCES addresses(address_id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
