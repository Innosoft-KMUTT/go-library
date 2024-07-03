FROM askforanywork/dev-coder:latest

RUN wget https://download.docker.com/linux/static/stable/x86_64/docker-25.0.3.tgz \
	&& tar -xzvf docker-25.0.3.tgz \
	&& cp ./docker/* /usr/bin/ \
	&& rm -rf docker-25.0.3.tgz

COPY init-on-build.sh /tmp/init-on-build.sh
RUN chmod +x /tmp/init-on-build.sh \
	&& /tmp/init-on-build.sh

COPY init-on-create.sh /tmp/init-on-create.sh
RUN chmod 755 /tmp/init-on-create.sh

ARG USER=coder
RUN useradd --groups sudo --no-create-home --shell /usr/bin/zsh ${USER} \
	&& echo "${USER} ALL=(ALL) NOPASSWD:ALL" >/etc/sudoers.d/${USER} \
	&& chmod 0440 /etc/sudoers.d/${USER}
USER ${USER}

WORKDIR /home/${USER}