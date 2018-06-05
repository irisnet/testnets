FROM alpine:edge


ENV PACKAGES go make git libc-dev bash

ENV GOPATH       /root/go
ENV BASE_PATH    $GOPATH/src/github.com/irisnet/testnets
ENV REPO_PATH    $BASE_PATH/test-faucet
ENV RUN_PATH     /test-faucet
ENV LOG_DIR      $RUN_PATH/log
ENV PATH         $GOPATH/bin:$PATH

RUN mkdir -p $LOG_DIR $GOPATH/pkg $GOPATH/bin $BASE_PATH $REPO_PATH

COPY . $REPO_PATH

RUN apk add --no-cache $PACKAGES && \
    cd $REPO_PATH && make all && \
    cp $REPO_PATH/build/faucet $GOPATH/bin/ && \
    cp $REPO_PATH/config.yml $RUN_PATH/ && \
    apk del $PACKAGES && \
    rm -fr $GOPATH/pkg $GOPATH/src $GOPATH/bin/dep /root/.cache

EXPOSE 3434

WORKDIR $RUN_PATH
CMD faucet > $LOG_DIR/app.log && tail -f $LOG_DIR/app.log