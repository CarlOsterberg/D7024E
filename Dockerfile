FROM golang:alpine
RUN mkdir /app 
ADD . /app/
WORKDIR /app/cmd
#RUN go build -o kademlia cmd/*.go
#RUN adduser -S -D -H -h /app appuser
#USER appuser
RUN go build -o /kadlab
RUN chmod +x /kadlab
#CMD ["./kademlia"]
#CMD [ "/kadlab" ]
RUN chmod +x /kadlab