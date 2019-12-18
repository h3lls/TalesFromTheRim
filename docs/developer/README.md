# Development Setup

You will need a few tools to clone and build the project.

## Required Tools

- Git - https://www.linode.com/docs/development/version-control/how-to-install-git-on-linux-mac-and-windows/
- Docker - https://docs.docker.com/docker-for-windows/install/

> Dev tools installs vary between systems, this documentation assumes you are using a Windows PC.

With git you should also setup an ssh key and add it to github.com https://github.com/settings/keys to get more details you can view https://help.github.com/articles/generating-an-ssh-key/ This will allow you to avoid having to enter username and password over and over again and also is quite a bit more secure.

## Docker Build

Docker allows you to run the mud locally without having to install gcc or the other myriad of tools needed. This lets you create a local environment without the heavy lifting of installing every tool.

### Clone Repository

> Note: **You will need to request access to the repository for the world (https://github.com/h3lls/TalesFromTheRimWorld) files.**

Open a command window and setup your repository by first cloning the project with the commands:  
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

### Create Docker Image

Open a command window and change to the TalesFromTheRim folder. Execute the startdocker.bat and pass in the rebuild parameter. This will only be needed when you need to update the image itself.  

```
startdocker.bat rebuild
```

After some time the build will finish and the mud will be running on port 5000. Open mudlet or your favorite telnet program and connect to **localhost** port **5000**.  

At this point you can edit any files you like in the src/ folder. To test you just run  

```startdocker.bat```

This will recompile the mud itself without the long image building process. You can also make lib changes without affecting the larger project.  

## Changes and Testing

After making any source code changes locally you may want to validate that everything is working on the remote server. There is a dev instance of the mud running on mudding.online port 5001 that will be built automatically whenever a pull request is created or updated.  

### Create a Working Branch

When you are starting work on a new feature or bug request you should create a new branch. First update your code on the master branch with:  

```
git pull
```

After pulling the latest data for master create a new branch with (replace <branch name> with a short name describing your feature or bug fix):  

```
git checkout -b <branch name>
```
Example:  
```
git checkout -b fix-protocol-error
```

### Making Changes

After you have made your changes you can call:  

```
git commit -m "Adding changes to fix protocol error"
```

After the commit you can do a push which will commit the changes to the remote (github) repository:  
```
git push
```

From this point if you feel that the code is tested and working properly and hasn't caused gross unexpected side-effects and bugs then you can create a pull request.  

### Creating a Pull Request

Open the repository at: https://github.com/h3lls/TalesFromTheRimWorld  

Select the branch you have been working on:  
![Select Branch](select-branch.png ':size=300')  
  
Click the "New Pull Request" button on your updated branch:  
![Select Branch](newpullrequest.png ':size=700')  
  
After entering a complete description and title along with a reference to any issue you are addressing, click the "Create pull request" button:  
![Select Branch](createpr.png ':size=600')  
  
This will start the build process in the background in a hook that runs on the server. Within a few seconds the updated code will be placed on the server and compiled for the dev instance. After that it will be restarted within a minute running the new code.  

> NOTE: *At this time there is no log access to the build process, please ask for the logs if things aren't building.*
  
To validate your changes login to mudding.online port 5001  

After you have tested in dev ask for a review and for the change to be merged in to master. After that at the next release your changes will be rolled in to the production system.  

