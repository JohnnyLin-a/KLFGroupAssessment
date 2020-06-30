docker run \
    -itd \
    --rm \
    --name klffrontend \
    -v "$(pwd)":/app \
    -p 3000:3000 \
    -e CHOKIDAR_USEPOLLING=true \
    klffrontend:dev