FROM golang:1.8
ENV WEBAPP /go/src/server
RUN mkdir -p ${WEBAPP}
COPY . ${WEBAPP}
ENV PORT=8080 
WORKDIR ${WEBAPP}
RUN go install
RUN rm -rf ${WEBAPP} 
EXPOSE 8080
EXPOSE 8700
CMD server
