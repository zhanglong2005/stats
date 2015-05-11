#!/bin/sh

case "$1" in
	start)
		echo "Starting stats_server"
		if [ `ps ax | grep stats_server | grep -v 'go test ' | grep -v CompileDaemon | grep -v grep | wc -l` == 1 ]
		then
			echo "stats_server already running"
		else
			cd $GO/bin
			stats_server -root=../src/github.com/bborbe/stats/files -loglevel=DEBUG  &
		fi
	;;
	stop)
		echo "Stopping stats_server"
		if [ `ps ax | grep stats_server | grep -v 'go test ' | grep -v CompileDaemon | grep -v grep | wc -l` == 1 ]
		then
			killall stats_server
		else
			echo "stats_server not running"
		fi
	;;
	restart)
		echo "Restarting stats_server"
		$0 stop
		sleep 1
		$0 start
	;;
	graceful)
		echo "Restarting graceful"
		cd $GO/src && go install github.com/bborbe/stats/bin/stats_server
		cd $GO/src && go test `find github.com/bborbe/stats -name "*_test.go" | dirof | unique` && cd $GO/bin &&	killall -USR2 stats_server
	;;
	install)
		echo "Installing stats_server"
		cd $GO/src && go install github.com/bborbe/stats/bin/stats_server
	;;
	devstart)
		echo "Starting Dev-Mode"
		if [ `ps ax | grep 'CompileDaemon -recursive=true'| grep -v grep | wc -l` == 1 ]
		then
			echo "CompileDaemon already running"
		else
			cd $GO
			$GO/bin/CompileDaemon -recursive=true -pattern="(.+\\.go)$" -directory="src/github.com/bborbe/stats" -build="go install" -command="sh src/github.com/bborbe/stats/scripts/stats.sh graceful" &
		fi
		sleep 1
    $0 start
	;;
	devstop)
		echo "Stopping Dev-Mode"
		if [ `ps ax | grep 'CompileDaemon -recursive=true'| grep -v grep | wc -l` == 1 ]
		then
			killall CompileDaemon
		else
			echo "CompileDaemon not running"
		fi
		sleep 1
		$0 stop
	;;
	devrestart)
		echo "Restarting Dev-Mode"
  	$0 devstop
  	sleep 1
  	$0 devstart
	;;
	*)
		echo "Usage: $0 {start|stop|restart|graceful|install|devstart|devstop|devrestart}"
		exit 1
	;;
esac

exit 0