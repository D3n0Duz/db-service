FROM scratch
COPY . /app
WORKDIR /app 
EXPOSE 3000

CMD ["/main"]

