#!/bin/sh

run_oss="false"
run_disk="false"

mkdir -p /var/log/alicloud/
mkdir -p /host/etc/kubernetes/volumes/disk/uuid

ossfsVer="1.80.6.ack.1"
armfsVer="1.80.6"

## check which plugin is running
for item in $@;
do
  if [ "$item" = "--driver=ossplugin.csi.alibabacloud.com" ]; then
      echo "Running oss plugin...."
      run_oss="true"
      mkdir -p /var/lib/kubelet/csi-plugins/ossplugin.csi.alibabacloud.com
      rm -rf /var/lib/kubelet/plugins/ossplugin.csi.alibabacloud.com/csi.sock
      /usr/bin/nsenter yum install -y fuse-devel
      if [ ! `/usr/bin/nsenter --mount=/proc/1/ns/mnt which ossfs` ]; then
          echo "First install ossfs...."
          cp /usr/bin/ossfs /host/usr/bin/
          echo "cp result -- `/usr/bin/nsenter --mount=/proc/1/ns/mnt which ossfs` --"
      else
          echo "ossfs is already on host"
      fi
  elif [ "$item" = "--driver=diskplugin.csi.alibabacloud.com" ]; then
      echo "Running disk plugin...."
			run_disk="true"
      mkdir -p /var/lib/kubelet/csi-plugins/diskplugin.csi.alibabacloud.com
      rm -rf /var/lib/kubelet/plugins/diskplugin.csi.alibabacloud.com/csi.sock
  elif [ "$item" = "--driver=nasplugin.csi.alibabacloud.com" ]; then
      echo "Running nas plugin...."
      mkdir -p /var/lib/kubelet/csi-plugins/nasplugin.csi.alibabacloud.com
      rm -rf /var/lib/kubelet/plugins/nasplugin.csi.alibabacloud.com/csi.sock
  elif [[ $item==*--driver=* ]]; then
      tmp=${item}
      driver_types=${tmp#*--driver=}
      driver_type_array=(${driver_types//,/ })
      for driver_type in ${driver_type_array[@]};
      do
          if [ "$driver_type" = "oss" ]; then
              echo "Running oss plugin...."
              run_oss="true"
              mkdir -p /var/lib/kubelet/csi-plugins/ossplugin.csi.alibabacloud.com
              rm -rf /var/lib/kubelet/plugins/ossplugin.csi.alibabacloud.com/csi.sock
              /usr/bin/nsenter yum install -y fuse-devel
          elif [ "$driver_type" = "disk" ]; then
              echo "Running disk plugin...."
							run_disk="true"
              mkdir -p /var/lib/kubelet/csi-plugins/diskplugin.csi.alibabacloud.com
              rm -rf /var/lib/kubelet/plugins/diskplugin.csi.alibabacloud.com/csi.sock
          elif [ "$driver_type" = "nas" ]; then
              echo "Running nas plugin...."
              mkdir -p /var/lib/kubelet/csi-plugins/nasplugin.csi.alibabacloud.com
              rm -rf /var/lib/kubelet/plugins/nasplugin.csi.alibabacloud.com/csi.sock
          fi
      done
  fi
done


## OSS plugin setup
if [ "$run_oss" = "true" ]; then
    echo "Starting deploy oss csi-plugin...."

    systemdDir="/host/usr/lib/systemd/system"
    if [[ ${host_os} == "lifsea" ]]; then
        systemdDir="/host/etc/systemd/system"
    fi
    # install OSSFS
    mkdir -p /host/etc/csi-tool/
    if [ ! `/nsenter --mount=/proc/1/ns/mnt which ossfs` ]; then
        echo "First install ossfs...."
        /nsenter --mount=/proc/1/ns/mnt yum install -y ossfs
        /nsenter --mount=/proc/1/ns/mnt ln -s /usr/bin/ossfs /usr/local/bin/ossfs
    # update OSSFS
    else
        echo "Check ossfs Version...."
        oss_info=`/nsenter --mount=/proc/1/ns/mnt ossfs --version | grep -E -o "V[0-9.a-z]+" | cut -d"V" -f2`
        if [ "$oss_info" != "$ossfsVer" ] || [ "$oss_info" != "$armfsVer" ]; then
            echo "Upgrade ossfs...."
            /nsenter --mount=/proc/1/ns/mnt yum remove -y ossfs
            /nsenter --mount=/proc/1/ns/mnt yum install -y ossfs
            /nsenter --mount=/proc/1/ns/mnt ln -s /usr/bin/ossfs /usr/local/bin/ossfs
        fi
    fi
fi

if [ "$run_disk" = "true" ] || [ "$run_oss" = "true"]; then
    updateConnector="true"
    if [ ! -f "/host/etc/csi-tool/csiplugin-connector" ]; then
      mkdir -p /host/etc/csi-tool/
			echo "mkdir /etc/csi-tool/ directory ...."
		else
			oldmd5=`md5sum /host/etc/csi-tool/csiplugin-connector | awk '{print $1}'`
			newmd5=`md5sum /bin/csiplugin-connector | awk '{print $1}'`
			if [ "$oldmd5" = "$newmd5" ]; then
					updateConnector="false"
			else
					rm -rf /host/etc/csi-tool/
					rm -rf /host/etc/csi-tool/connector.sock
					rm -rf /var/log/alicloud/connector.pid
					mkdir -p /host/etc/csi-tool/
			fi
		fi
		cp /freezefs.sh /host/etc/csi-tool/freezefs.sh
    if [ "$updateConnector" = "true" ]; then
        echo "Install csiplugin-connector...."
        cp /bin/csiplugin-connector /host/etc/csi-tool/csiplugin-connector
        chmod 755 /host/etc/csi-tool/csiplugin-connector
    fi


    # install/update csiplugin connector service
    updateConnectorService="true"
    if [[ ! -z "${PLUGINS_SOCKETS}" ]];then
        sed -i 's/Restart=always/Restart=on-failure/g' /bin/csiplugin-connector.service
        sed -i '/^\[Service\]/a Environment=\"WATCHDOG_SOCKETS_PATH='"${PLUGINS_SOCKETS}"'\"' /bin/csiplugin-connector.service
        sed -i '/ExecStop=\/bin\/kill -s QUIT $MAINPID/d' /bin/csiplugin-connector.service
        sed -i '/^\[Service\]/a ExecStop=sh -xc "if [ x$MAINPID != x ]; then /bin/kill -s QUIT $MAINPID; fi"' /bin/csiplugin-connector.service
    fi
    if [ -f "$systemdDir/csiplugin-connector.service" ];then
        echo "Check csiplugin-connector.service...."
        oldmd5=`md5sum $systemdDir/csiplugin-connector.service | awk '{print $1}'`
        newmd5=`md5sum /bin/csiplugin-connector.service | awk '{print $1}'`
        if [ "$oldmd5" = "$newmd5" ]; then
            updateConnectorService="false"
        else
            rm -rf $systemdDir/csiplugin-connector.service
        fi
    fi

    if [ "$updateConnectorService" = "true" ]; then
        echo "Install csiplugin connector service...."
        cp /bin/csiplugin-connector.service $systemdDir/csiplugin-connector.service
        /nsenter --mount=/proc/1/ns/mnt systemctl daemon-reload
    fi

    rm -rf /var/log/alicloud/connector.pid
    /nsenter --mount=/proc/1/ns/mnt systemctl enable csiplugin-connector.service
    /nsenter --mount=/proc/1/ns/mnt systemctl restart csiplugin-connector.service

fi


# start daemon
/bin/plugin.csi.alibabacloud.com $@
