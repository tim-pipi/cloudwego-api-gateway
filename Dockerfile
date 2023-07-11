FROM quay.io/coreos/etcd:v3.5.0

CMD ["etcd", "--advertise-client-urls", "http://etcd:2379", "--listen-client-urls", "http://0.0.0.0:2379"]

EXPOSE 2379