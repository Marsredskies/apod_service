# NASA's Picture Of The Day Downloader

A service that fetches a picture of the day from NASA's API with an ability to get an image by date and info about all images stored.

# Usage

Setup: 

- Create a `.env` file in the repo directory using env.example. Make sure you've set all the variables. Even tho the app itself has it's own default values in case of missing envs, docker-compose however needs all required values.

- If you are going to abuse nasa api somehow, create an API key at `https://api.nasa.gov/` to avoid rate limits. But the default key is set in the app just in case. 1 rph is a default behaviour.

- Run `make run`. Migration will apply automatically. After that the service will make an initial image download from NASA's api.

The app has two available endpoints:

- Get image by date:
    `http://localhost:8080/image?date={{YYYY-MM-DD}}`


    NOTE: the date should be the following format: `"YYYY-MM-DD"``, and there are obviously no images from the future.


- Get album of stored images:
    `http://localhost:8080/album`

    Shows a rendered html table with stored images info and links.

# Running tests

- Run `make db`
- In separate terminal window run `make test`
- Run `make stop` after tests finish

