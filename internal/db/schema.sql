CREATE TABLE Country (
    id SERIAL PRIMARY KEY,
    iso_code3 CHAR(3) UNIQUE NOT NULL,
    country_name VARCHAR(255) NOT NULL,
    dialing_code VARCHAR(10)
);

CREATE TABLE Province (
    id SERIAL PRIMARY KEY,
    country_id INTEGER REFERENCES countries(id) NOT NULL,
    province_name VARCHAR(255) NOT NULL,
    code VARCHAR(50) UNIQUE
);

-- DeathNotice Table (unchanged)
CREATE TABLE DeathNotice (
    DeathNoticeID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    FullName VARCHAR(255) NOT NULL,
    DateOfDeath DATE NOT NULL,
    Age INTEGER,
    CauseOfDeath TEXT,
    FuneralParlourID UUID REFERENCES FuneralParlour(FuneralParlourID),
    AddressID UUID REFERENCES Address(AddressID), -- Changed to AddressID
    Obituary TEXT,
    ImageURL VARCHAR(255),
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- FuneralParlour Table (unchanged)
CREATE TABLE FuneralParlour (
    FuneralParlourID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Name VARCHAR(255) NOT NULL,
    Address VARCHAR(255),
    ContactNumber VARCHAR(20),
    Email VARCHAR(255),
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Address Table (Church/Home Address combined and renamed)
CREATE TABLE Address (
    AddressID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    AddressType VARCHAR(50) CHECK (AddressType IN ('Church', 'Home', 'Event')),
    Address VARCHAR(255) NOT NULL,
    City VARCHAR(100),
    Province VARCHAR(100),
    PostalCode VARCHAR(20),
    ContactPerson VARCHAR(255),
    ContactNumber VARCHAR(20),
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Users Table (unchanged)
CREATE TABLE Users (
    UserID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Username VARCHAR(255) UNIQUE NOT NULL,
    Email VARCHAR(255) UNIQUE NOT NULL,
    PasswordHash VARCHAR(255) NOT NULL,
    FirstName VARCHAR(255),
    LastName VARCHAR(255),
    Role VARCHAR(50) CHECK (Role IN ('Admin', 'User')),
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- DeathNotice Comments Table (unchanged)
CREATE TABLE DeathNoticeComments(
    CommentID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    DeathNoticeID UUID REFERENCES DeathNotice(DeathNoticeID),
    UserID UUID REFERENCES Users(UserID),
    CommentText TEXT NOT NULL,
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- DeathNotice Flowers/Condolences table (unchanged)
CREATE TABLE DeathNoticeCondolences(
    CondolenceID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    DeathNoticeID UUID REFERENCES DeathNotice(DeathNoticeID),
    UserID UUID REFERENCES Users(UserID),
    CondolenceText TEXT NOT NULL,
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- DeathNotice Donations Table (unchanged)
CREATE TABLE DeathNoticeDonations(
    DonationID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    DeathNoticeID UUID REFERENCES DeathNotice(DeathNoticeID),
    UserID UUID REFERENCES Users(UserID),
    Amount DECIMAL(10,2) NOT NULL,
    DonationDate TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- DeathNotice Events Table (Updated to use AddressID)
CREATE TABLE DeathNoticeEvents(
    EventID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    DeathNoticeID UUID REFERENCES DeathNotice(DeathNoticeID),
    EventType VARCHAR(255) NOT NULL,
    EventDate TIMESTAMP WITH TIME ZONE NOT NULL,
    AddressID UUID REFERENCES Address(AddressID), -- Changed to use AddressID
    CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);