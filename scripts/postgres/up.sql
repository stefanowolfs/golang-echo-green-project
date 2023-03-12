
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR ( 50 ) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS cost_centers (
                                            id SERIAL PRIMARY KEY,
                                            title VARCHAR ( 30 ) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS activities (
                                          id SERIAL PRIMARY KEY,
                                          title VARCHAR ( 30 ) NOT NULL,
                                          cost_center_id INTEGER REFERENCES cost_centers(id)
);

CREATE TABLE IF NOT EXISTS expenses (
                                        id SERIAL PRIMARY KEY,
                                        activity_id INTEGER REFERENCES activities(id),
                                        user_id INTEGER REFERENCES  users(id),
                                        description VARCHAR ( 100 ) NOT NULL,
                                        cost DECIMAL NOT NULL,
                                        billing_date DATE
);