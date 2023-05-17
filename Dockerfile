FROM golang:1.19.0

WORKDIR /usr/src/app

COPY . .

#adding ssh keys
ADD id_rsa /root/.ssh/id_rsa
RUN chmod 700 /root/.ssh/id_rsa

#setting git configs
RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

RUN go mod tidy

RUN rm id_rsa /root/.ssh/id_rsa

