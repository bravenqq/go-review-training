1. 在你本机打一个 tunnel 2030 -> nuc10:2030

   ssh -L 2030:127.0.0.1:2030 nqqhk -N

2. set remote server thought socks5

cat ~/.ssh/config.d/local.config

    Host nuc10proxy
      Hostname 10.0.0.14
      port 22
      User bantana
      IdentityFile ~/.ssh/id_rsa

通过 ssh tunnel 在本机和 nuc10proxy 之间建立一个 socks5 服务器

    ssh -qND 1090 nuc10proxy

3. testing

   curl --socks5-hostname 127.0.0.1:1090 http://10.0.0.14:9000

4. open mozilla

   manual proxy setting

   only set socks5

   127.0.0.1 1090
