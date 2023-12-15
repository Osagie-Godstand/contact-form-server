DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'contacts_received') THEN
-- Create the "contacts_received" table
CREATE TABLE contacts_received (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    phonenumber VARCHAR(20) NOT NULL,
    message TEXT NOT NULL
);
  END IF;
END $$;

