IMAGE=12345saadiq/aniliststream

build:
	docker build -t $(IMAGE) .

push:
	docker push $(IMAGE)

up:
	docker compose up --build -d

down:
	docker compose down
