FROM v2fly/v2fly-core
# COPY geosite.dat /usr/local/share/v2ray/LoyalsoldierSite.dat
# COPY v2raya /usr/bin/v2raya
# RUN chmod +x /usr/bin/v2raya

RUN wget -O /tmp/v2ray.tar.gz https://github.com/wxyzZ/v2raya/releases/download/v1.0.0/v2raya_linux_x64_1.0.0.tar.gz && tar -xzf   /tmp/v2ray.tar.gz && chmod +x release/v2raya_linux_x64_1.0.0 && mv release/v2raya_linux_x64_1.0.0 /usr/bin/v2raya
RUN wget -O /usr/local/share/v2ray/LoyalsoldierSite.dat https://raw.githubusercontent.com/mzz2017/dist-v2ray-rules-dat/master/geosite.dat
RUN apk add --no-cache iptables ip6tables tzdata
LABEL org.opencontainers.image.source=https://github.com/v2rayA/v2rayA
EXPOSE 2017
VOLUME /etc/v2raya
ENTRYPOINT "v2raya --lite -a 0.0.0.0:2017 --log-level error"
