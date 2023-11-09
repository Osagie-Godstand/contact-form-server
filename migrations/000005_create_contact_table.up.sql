DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'contacts_recieved') THEN
-- Create the "contacts_received" table
CREATE TABLE contacts_received (
    id UUID PRIMARY KEY,
    email VARCHAR(255),
    name VARCHAR(255),
    phonenumber VARCHAR(20),
    message VARCHAR(255)
);
  END IF;
END $$;

