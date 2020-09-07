#!/bin/bash
# Start gin and graceful stop or restart gin.
#
# Mind that execute `export RUNENV=dev` first in development environment,
# or execute `export RUNENV=prod` first in production environment.
#
# Or execute as:
# RUNENV=<dev|prod> ./service.sh <start|stop|reload|restart|status>


PROJECT_PATH="./"
BIN_NAME="use-gin"
PID_FILE=$PROJECT_PATH/logs/${BIN_NAME}.pid
PID=$(cat $PID_FILE 2>/dev/null)


usage(){
    [[ -z "$1" ]] || [[ "$#" != 1 ]] && {
        echo "Usage: $0 <start|stop|reload|restart|status>"
        exit 1
    }
}

status(){
    [[ -z "$PID" ]] && {
        echo "$BIN_NAME was stopped"
        return
    }

    # shellcheck disable=SC2009
    ps -ef | grep -v grep | awk '{print $2}' | grep -wq "$PID" &&
        echo "$BIN_NAME is running, pid $PID" ||
           echo "$BIN_NAME was stopped"
}

start(){
    status | grep -q running && {
        echo "error: $BIN_NAME is already running, do not start again."
        exit 1
    }
    $PROJECT_PATH/$BIN_NAME &
}

stop(){
    status | grep -q stopped && {
        echo "error: $BIN_NAME was already stopped, no need to stop again."
        exit 1
    }
    kill -SIGTERM "$PID" && echo "success: graceful stopped"
}

reload(){
    status | grep -q stopped && {
        echo "error: no $BIN_NAME process found, please start first."
        exit 1
    }
    kill -SIGHUP "$PID" && echo "success: graceful reload"
}

main(){
    usage "$@"

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
