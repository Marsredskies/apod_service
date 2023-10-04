# NASA's Picture Of The Day Downloader

A service that fetches a picture of the day from NASA's API with an ability to get an image by date and info about all images stored.

# Usage

- Create a `.env` file in the repo directory using env.example. Make sure you've set all the variables. Even tho the app itself has it's own default values in case of missing envs, docker-compose however needs all required values.

- Run `make run`. Migration will apply automatically. After that the service will make an initial image download from NASA's api.

Two available endpoints:
- Get image by date:
    `http://localhost:8080/image?date=2023-10-04`


    NOTE: the date should be the following format: `"YYYY-MM-DD"``


- Get album of stored images:
    `http://localhost:8080/album`

    Shows a rendered html table with stored images info and links.

# Running tests

- Run `make db`
- In separate terminal window run `make test`
- Run `make stop`

