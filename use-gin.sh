#!/bin/bash
# start gin and graceful stop or restart gin


[[ -z "$1" ]] || [[ "$#" != 1 ]] && {
    echo "Usage: $0 <start|stop|restart|status>"
}


PROJECT_PATH="./"
BIN_NAME="use-gin"
CONFIG_FILE="./dev.config.toml"
PID_FILE=$PROJECT_PATH/logs/${BIN_NAME}.pid
PID=$(cat $PID_FILE)


status(){
    ps -ef | grep -v grep | awk '{print $2}' | grep -wq $PID &&
        echo "$BIN_NAME is running, pid $PID" ||
           echo "$BIN_NAME was stopped"
}

case $1 in
    start)
        status | grep -q running && {
            echo "error: $BIN_NAME is already running, do not start again."
            exit 1
        }
        $PROJECT_PATH/$BIN_NAME -c $CONFIG_FILE &
        ;;
    stop)
        status | grep -q stopped && {
            echo "error: $BIN_NAME was already stopped, no need to stop again."
            exit 1
        }
        kill -SIGTERM $PID && echo "success: graceful stopped"
        ;;
    restart)
        status | grep -q stopped && {
            echo "error: no $BIN_NAME process found, please start first."
            exit 1
        }
        kill -SIGHUP $PID && echo "success: graceful restart"
        ;;
    status)
        status
        ;;
esac


exit 0
