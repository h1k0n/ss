#
sudo service sshd start
sudo rc-update add sshd

echo 'http://mirrors.aliyun.com/alpine/v3.16/main1
http://mirrors.aliyun.com/alpine/v3.16/community1
' >> /etc/apk/repositories 

sudo apk add curl
sudo apk add font-terminus font-inconsolata font-dejavu font-noto font-noto-cjk font-awesome font-noto-extra
