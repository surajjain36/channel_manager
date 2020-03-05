Channel Manager
----------------------

**Steps to build the application using Docker**

- If required edit config.yml file
    - Enter required configurations.
- Install Docker
- Run following command from root folder to build Docker image
	- ```docker build -t channel_manager  .```
- Run following command from root folder to run Docker image		
	- ```docker run -it -p 3000:3000 channel_manager```


**To run in local machine**

Compile the code by running:

```$ make build```

Start the server by running:

```$ make run```


**API Documentation**

- API to create go routine:

    - API: [localhost:3000/v1/_create?start=2&step=3](localhost:3000/v1/_create?start=2&step=3)
    - Response will have ```ChannelID```

- API to list all go routines:

    - API: [localhost:3000/v1/_check](localhost:3000/v1/_check)

- API to find a go routines by ID:

    - API: [localhost:3000/v1/_check?id=<ChannelID>](localhost:3000/v1/_check?id=<ChannelID>)
    - Input: ```id = ChannelID```

- API to pause a go routine:

    - API: [localhost:3000/v1/_pause?id=<ChannelID>](localhost:3000/v1/_pause?id=<ChannelID>)
    - Input: ```id = ChannelID```

- API to clear a go routine:

    - API: [localhost:3000/v1/_clear?id=<ChannelID>](localhost:3000/v1/_clear?id=<ChannelID>)
    - Input: ```id = ChannelID```