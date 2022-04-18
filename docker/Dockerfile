FROM centos:7
ADD entrypoint.sh /srv/
ADD demoservice /srv/demoservice
EXPOSE 8080
ENTRYPOINT ["/srv/entrypoint.sh"]
