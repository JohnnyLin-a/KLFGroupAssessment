docker run \
    -itd \
    --rm \
    --name klf_frontend \
    -v "$(pwd)":/app \
    -v /app/node_modules \
    -p 3000:3000 \
    -e CHOKIDAR_USEPOLLING=true \
    klf_frontend:dev
