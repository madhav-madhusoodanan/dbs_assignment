FROM mysql:latest

RUN sudo apt-get install mysql

COPY /backend /app
WORKDIR /app

RUN cargo install diesel_cli --no-default-features --features mysql && diesel setup

# CMD ["cargo", "install", "diesel_cli", "--no-default-features", "--features", "mysql"]

# CMD ["LIBRARY_PATH=/usr/lib/x86_64-linux-gnu:$LIBRARY_PATH"]
# CMD ["export", "LIBRARY_PATH"]
# CMD ["cargo", "run", "diesel", "setup"]
# CMD ["cargo", "run", "diesel", "migration", "generate", "create_posts"]

# expose the port publicly
# copy the file 
