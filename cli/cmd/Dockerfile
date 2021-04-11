#---------------------------------------------------------------------------------------------------
# hello-world [dockerfile]
#---------------------------------------------------------------------------------------------------

#build from alpine container with golang
FROM alpine:latest AS build

#set the working directory to cmd, so we can build easily
WORKDIR /app/

#copy the binary
COPY ./hello-world ./sample_read_only.txt /app/

#tar the binary
RUN tar -czvf hello-world.tar.gz ./hello-world ./sample_read_only.txt 

#build from build image
FROM alpine:latest

#set working directory to app
WORKDIR /app

#copy delve and the app
COPY --from=build /app/hello-world.tar.gz /app/hello-world.tar.gz

# At Start, de-compress tar file and run app
CMD tar -xzvf hello-world.tar.gz && ./hello-world
