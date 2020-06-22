- In order for a variable to be echoed in the Dockerfile:
  1) It must be a build arg in the docker-compose.yml file. 
  2) It must also be defined as an argument in the dockerfile. 
  3) It doesn't matter if it's defined as an environmental variable in either file.

- A variable included under "environment:" or "env_file" will be accessible to the container. It doesn't matter if it's a build arg or included in the Dockerfile. 

- A variable only defined as a build arg will not be accessible to the container.
