class golang {

  package { "golang-go":
    ensure => "present"
  }

  exec { "setup-workspace":
    command => "/bin/echo 'export GOPATH=/vagrant' >>/home/vagrant/.profile",
    unless => "/bin/grep -q GOPATH /home/vagrant/.profile ; /usr/bin/test $? -eq 0"
  }

}
