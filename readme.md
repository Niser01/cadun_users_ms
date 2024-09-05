docker build -t cadun_users_db .
docker run -d -t -i -p 3306:3306 --name cadun_users_db cadun_users_db


docker build -t cadun_users_ms .
docker run --name cadun_users_ms -p 8080:8080 -e DB_HOST=host.docker.internal -e DB_PORT=3306 -e DB_USER=CADUN -e DB_PASSWORD=2024 -e DB_NAME=cadun_users_db -e URL=0.0.0.0:8080 cadun_users_ms

docker run --name cadun_users_ms -p 8080:8080 -e DB_HOST=dpg-crcup62j1k6c73cu0lb0-a.ohio-postgres.render.com -e DB_PORT=5432 -e DB_USER=cadun_users_db_r3on_user -e DB_PASSWORD=3yZ19qokcHXVYc4w48J5wPw5v1GFqxny -e DB_NAME=cadun_users_db_r3on -e URL=0.0.0.0:8080 cadun_users_ms