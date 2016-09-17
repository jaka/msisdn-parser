class msisdn-server {

#  file { "/etc/init.d/msisdn-server.sh":
#    source => 'puppet:///modules/msisdn-server/msisdn-server.sh',
#    owner => 'root',
#    group => 'root',
#    mode => '0755',
#  }

  file { "/etc/init/msisdn-server.conf":
    source => 'puppet:///modules/msisdn-server/msisdn-server.conf',
    owner => 'root',
    group => 'root',
    mode => '0644',
  }

#  exec { "add-to-startup":
#    command => "/bin/sed -i \\\$i\\ su\\ -c\\ /home/vagrant/start-msisdn-server.sh\\ vagrant\\ \\& /etc/rc.local"
#  }

  file { "/compiled":
    ensure => 'directory',
    owner => 'vagrant',
    group => 'vagrant',
    mode => '0755',
  }

  exec { "build-server":
    command => "/bin/su -c 'export GOPATH=/vagrant; go build -o /compiled/msisdn-server /vagrant/src/server/server.go' vagrant",
    require => File["/compiled"],
    before => Service["msisdn-server"]
  }

#  exec { "sysv startup":
#    command => "/usr/sbin/update-rc.d msisdn-server.sh defaults",
#    require => File["/etc/init.d/msisdn-server.sh"]
#  }

  service { "msisdn-server":
    ensure => "running",
    provider => "upstart",
    require => File[ "/etc/init/msisdn-server.conf" ]
  }

}
