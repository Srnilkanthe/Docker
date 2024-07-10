Dockerfile
    Dockerfile is nothing but a simple text (*.txt) file which contains set of instractions which is to executed one by one 
so as to automate the process of  image creation .

Before we construct our Dockerfile, you need to understand what makes up the file. This will be a text file, 
named Dockerfile, that includes specific keywords that dictate how to build a specific image. The specific keywords
you can use in a file are:

FROM defines the base image used to start the build process.
MAINTAINER defines the full name and email address of the image creator.
ENV sets environment variables.
WORKDIR sets the path where the command, defined with CMD, is to be executed.
ADD copies the files from a source on the host into the container’s filesystem at the set destination.
COPY copy the files from a source on the host into the containe.
RUN is the central executing directive for Dockerfiles, also known as the run Dockerfile command.
WORKDIR sets the path where the command, defined with CMD, is to be executed.
EXPOSE associates a specific port to enable networking between the container and the outside world.
USER sets the UID (or username) which is to run the container.
VOLUME is used to enable access from the container to a directory on the host machine.
LABEL allows you to add a label to your docker image.
ENTRYPOINT sets a default application to be used every time a container is created with the image.
CMD can be used for executing a specific command within the container.

Not all keywords are required for a Dockerfile to function. Case in point, our example will only make use of FROM, MAINTAINER, and RUN.