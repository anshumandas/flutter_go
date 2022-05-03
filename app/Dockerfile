FROM cirrusci/flutter:dev AS build

WORKDIR /app

RUN flutter config --enable-web \
    && sudo apt-get update \
    && sudo apt-get install -y nginx \
    && sudo apt-get install -y chromium-browser \
    && sudo rm -rf /var/lib/apt/lists/*

COPY . .



