CREATE TABLE drivers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    license VARCHAR(50) NOT NULL,
    available BOOLEAN NOT NULL,
    location VARCHAR(255),
    rating DECIMAL(3, 2),
    car_model VARCHAR(255),
    car_number VARCHAR(50),
    phone_number VARCHAR(50),
    email VARCHAR(255),
    total_trips INT DEFAULT 0,
    language VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    last_online TIMESTAMP WITH TIME ZONE
);

INSERT INTO drivers (name, license, available, location, rating, car_model, car_number, phone_number, email, language) VALUES
('Иван Иванов', 'XYZ123', true, 'Москва, Россия', 4.5, 'Toyota Camry', 'A123BC', '+71234567890', 'ivan@example.com', 'Русский'),
('Анна Смирнова', 'ABC456', false, 'Санкт-Петербург, Россия', 4.7, 'Hyundai Solaris', 'B456CD', '+79876543210', 'anna@example.com', 'Русский'),
('Петр Петров', 'DEF789', true, 'Новосибирск, Россия', 4.8, 'Kia Rio', 'C789EF', '+71239876543', 'petr@example.com', 'Русский'),
('Ольга Сидорова', 'GHI012', false, 'Екатеринбург, Россия', 4.6, 'Volkswagen Polo', 'D012GH', '+79871234567', 'olga@example.com', 'Русский'),
('Сергей Сергеев', 'JKL345', true, 'Казань, Россия', 4.9, 'Skoda Rapid', 'E345IJ', '+71234567891', 'sergey@example.com', 'Русский'),
('Мария Иванова', 'MNO678', false, 'Нижний Новгород, Россия', 4.7, 'Renault Logan', 'F678MN', '+79876543211', 'maria@example.com', 'Русский'),
('Алексей Алексеев', 'PQR901', true, 'Самара, Россия', 4.8, 'Lada Vesta', 'G901QR', '+71239876544', 'alexey@example.com', 'Русский'),
('Наталья Петрова', 'STU234', true, 'Омск, Россия', 4.6, 'Ford Focus', 'H234ST', '+71239876545', 'natalia@example.com', 'Русский'),
('Дмитрий Смирнов', 'VWX567', false, 'Ростов-на-Дону, Россия', 4.7, 'Chevrolet Cruze', 'I567VW', '+79871234568', 'dmitry@example.com', 'Русский'),
('Елена Сергеева', 'YZA890', true, 'Уфа, Россия', 4.8, 'Mazda 3', 'J890YZ', '+71234567892', 'elena@example.com', 'Русский'),
('Андрей Андреев', 'BCD123', false, 'Волгоград, Россия', 4.9, 'Nissan Almera', 'K123BC', '+79876543212', 'andrey@example.com', 'Русский'),
('Татьяна Иванова', 'EFG456', true, 'Пермь, Россия', 4.7, 'Toyota Corolla', 'L456EF', '+71239876546', 'tatiana@example.com', 'Русский');
('Виктор Сидоров', 'HIJ789', true, 'Красноярск, Россия', 4.6, 'Hyundai Elantra', 'M789HI', '+71239876547', 'viktor@example.com', 'Русский'),
('Ирина Петрова', 'KLM012', false, 'Воронеж, Россия', 4.7, 'Skoda Octavia', 'N012KL', '+79871234569', 'irina@example.com', 'Русский'),
('Александр Смирнов', 'NOP345', true, 'Саратов, Россия', 4.8, 'Volkswagen Jetta', 'O345NP', '+71234567893', 'alexander@example.com', 'Русский'),
('Светлана Иванова', 'QRS678', false, 'Краснодар, Россия', 4.9, 'Ford Fiesta', 'P678QR', '+79876543213', 'svetlana@example.com', 'Русский'),
('Михаил Сергеев', 'TUV901', true, 'Тюмень, Россия', 4.7, 'Renault Sandero', 'Q901TU', '+71239876548', 'mikhail@example.com', 'Русский');
