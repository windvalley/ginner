#!/bin/bash
# service.sh
#
# Start gin and graceful stop or restart gin.


PROJECT_PATH="./"
BIN_NAME="ginner"
COMMAND="./$BIN_NAME -c conf/dev.config.toml"
PID_FILE=$PROJECT_PATH/logs/${BIN_NAME}.pid
PID=$(cat $PID_FILE 2>/dev/null)
STATUS=


usage(){
    [[ -z "$1" ]] || [[ "$#" != 1 ]] && {
        echo "Usage: $0 <start|stop|reload|restart|status>"
        exit 1
    }
}

status(){
    # shellcheck disable=SC2009
    ps -ef | grep -v grep | grep -q "$COMMAND" && {
        STATUS=running
        echo "$BIN_NAME is running, pid $PID" 
        return
    } 

    STATUS=stopped
    echo "$BIN_NAME was stopped"
}

start(){
    [[ "$STATUS" = "running" ]] && {
        echo "warning: $BIN_NAME is already running"
        exit 1
    }

    eval "$COMMAND" &
}

stop(){
    [[ "$STATUS" = "stopped" ]] && {
        echo "warning: $BIN_NAME was already stopped"
        return
    }

    kill -SIGTERM "$PID" && echo "success: graceful stopped"
}

reload(){
    [[ "$STATUS" = "stopped" ]] && {
        echo "warning: no $BIN_NAME process found"
        exit 1
    }

    kill -SIGHUP "$PID" && echo "success: graceful reload"
}

main(){
    usage "$@"

    cd $PROJECT_PATH || exit 1

    status &>/dev/null

    case $1 in
        start) start
            ;;
        stop) stop
            ;;
        reload) reload
            ;;
        restart) stop && sleep 2 && start
            ;;
        status) status
            ;;
        *) usage
            ;;
    esac
}


main "$@"
