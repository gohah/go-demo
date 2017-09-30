#!/usr/bin/env bash

CUR_PATH=$(cd `dirname $0`; pwd)
DAEMON_SH="daemonAgent.sh"

function startp(){
    stopp;
    $CUR_PATH/bin/$DAEMON_SH "$CUR_PATH/bin" 2>/dev/null 1>&2 &
}

function stopp() {
    num_dae=`ps -ef | grep "/logcollect/bin/daemonAgent.sh" | grep -v grep | wc -l`
    if [ $num_dae -ne 0 ]; then
        echo -n "Stop daemonAgent..."
        echo $num_dae
        ps -ef | grep "/logcollect/bin/daemonAgent.sh" | grep -v grep  | awk '{print $2}' | xargs kill -9
        sleep 1
    fi

    if [ -f /tmp/LogCollect.pid ]; then
        kill `cat /tmp/LogCollect.pid`
        sleep 1
        echo "Stop agent normal ... 1"
        if [ -f /tmp/LogCollect.pid ]; then
            rm /tmp/LogCollect.pid
        fi
    fi
    num_agent=`ps -ef | grep "/logcollect/bin/LogCollect" | grep -v grep | wc -l`
    if [ $num_agent -ne 0 ]; then
        echo -n "Stop agent use ps ..."
        echo $num_agent
        ps -ef | grep "/logcollect/bin/LogCollect" | grep -v grep  | awk '{print $2}' | xargs kill -9
        sleep 1
    fi
}

case $1 in
start)
    echo  "Starting LogCollect ..."
    startp;
    echo "Start Success"
    ;;
stop)
    echo  "Stoping LogCollect..."
    stopp;
    num=`ps -ef | grep "/logcollect/bin/daemonAgent.sh" | grep -v grep | wc -l`
    if [ $num -ne 0 ]; then
        echo "Stoped Failed"
    else
        echo "Stoped Success"
    fi
    ;;
upgrade)
    echo  "Upgrading LogCollect..."
    ;;
status)
    if [ `ps -ef | grep "/logcollect/bin/LogCollect" | grep -v grep | wc -l` -gt 0 ]; then
        echo "LogCollect is running"
    else
        echo "LogCollect is stoped"
    fi
    ;;
restart)
    echo  "Restarting LogCollect..."
    stopp;
    startp;
    echo "Restart Success"
    ;;
*)
    echo "Usage: $0 {start|stop|restart|status|upgrade}" >&2
esac
exit 0
