# Development Setup

You will need a few tools to clone and build the project.

## Required Tools

Git - https://www.linode.com/docs/development/version-control/how-to-install-git-on-linux-mac-and-windows/
Docker - https://docs.docker.com/docker-for-windows/install/

> Dev tools installs vary between systems, this documentation assumes you are using a Windows PC.

## Docker Build

Docker allows you to run the mud locally without having to install gcc or the other myriad of tools needed. This lets you create a local environment without the heavy lifting of installing every tool.

### Clone Repository

Note: **You will need to request access to the repository for the world (https://github.com/h3lls/TalesFromTheRimWorld) files.**

You can setup your repository by first cloning the project with:
```
c:
cd \
git config --global core.autocrlf false
git clone --recurse-submodules git@github.com:h3lls/TalesFromTheRim.git
cd TalesFromTheRim
```

### Give Docker Volume Permission

After cloning the repository you will need to configure your docker to allow the build process to mount your local drive as a volume. This lets you share your local folder with the docker image meaning changes you make in Windows will appear in your docker instance so you can edit and build as needed.


![Docker Settings](docker-settings.png ':size=200')
- Right click on the docker icon in the task bar and select (settings)  
![Docker Shared Drives](docker-shared.png ':size=400')



- Select the Shared Drives on the left and choose the drive where you cloned the repository.
