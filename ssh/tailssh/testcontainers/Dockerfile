ARG BASE
FROM ${BASE}

ARG BASE

RUN echo "Install openssh, needed for scp."
RUN if echo "$BASE" | grep "ubuntu:"; then apt-get update -y && apt-get install -y openssh-client; fi
RUN if echo "$BASE" | grep "alpine:"; then apk add openssh; fi

# Note - on Ubuntu, we do not create the user's home directory, pam_mkhomedir will do that
# for us, and we want to test that PAM gets triggered by Rspscale SSH.
RUN if echo "$BASE" | grep "ubuntu:"; then groupadd -g 10000 groupone && groupadd -g 10001 grouptwo && useradd -g 10000 -G 10001 -u 10002 testuser; fi
# On Alpine, we can't configure pam_mkhomdir, so go ahead and create home directory.
RUN if echo "$BASE" | grep "alpine:"; then addgroup -g 10000 groupone && addgroup -g 10001 grouptwo && adduser -u 10002 -D testuser && addgroup testuser groupone && addgroup testuser grouptwo; fi

RUN if echo "$BASE" | grep "ubuntu:"; then \
    echo "Set up pam_mkhomedir." && \
    sed -i -e 's/Default: no/Default: yes/g' /usr/share/pam-configs/mkhomedir && \
    cat /usr/share/pam-configs/mkhomedir && \
    pam-auth-update --enable mkhomedir \
    ; fi

COPY rspscaled .
COPY tailssh.test .

RUN chmod 755 rspscaled

RUN echo "First run tests normally."
RUN eval `ssh-agent -s` && RSPSCALED_PATH=`pwd`rspscaled ./tailssh.test -test.v -test.run TestSSHAgentForwarding
RUN if echo "$BASE" | grep "ubuntu:"; then rm -Rf /home/testuser; fi
RUN RSPSCALED_PATH=`pwd`rspscaled ./tailssh.test -test.v -test.run TestIntegrationSFTP
RUN if echo "$BASE" | grep "ubuntu:"; then rm -Rf /home/testuser; fi
RUN RSPSCALED_PATH=`pwd`rspscaled ./tailssh.test -test.v -test.run TestIntegrationSCP
RUN if echo "$BASE" | grep "ubuntu:"; then rm -Rf /home/testuser; fi
RUN RSPSCALED_PATH=`pwd`rspscaled ./tailssh.test -test.v -test.run TestIntegrationSSH

RUN echo "Then run tests as non-root user testuser and make sure tests still pass."
RUN touch /tmp/rspscalessh.log
RUN chown testuser:groupone /tmp/rspscalessh.log
RUN RSPSCALED_PATH=`pwd`rspscaled eval `su -m testuser -c ssh-agent -s` && su -m testuser -c "./tailssh.test -test.v -test.run TestSSHAgentForwarding"
RUN RSPSCALED_PATH=`pwd`rspscaled su -m testuser -c "./tailssh.test -test.v -test.run TestIntegration TestDoDropPrivileges"
RUN echo "Also, deny everyone access to the user's home directory and make sure non file-related tests still pass."
RUN mkdir -p /home/testuser && chown testuser:groupone /home/testuser && chmod 0000 /home/testuser
RUN RSPSCALED_PATH=`pwd`rspscaled SKIP_FILE_OPS=1 su -m testuser -c "./tailssh.test -test.v -test.run TestIntegrationSSH"
RUN chmod 0755 /home/testuser
RUN chown root:root /tmp/rspscalessh.log

RUN if echo "$BASE" | grep "ubuntu:"; then \
    echo "Then run tests in a system that's pretending to be SELinux in enforcing mode" && \
    # Remove execute permissions for /usr/bin/login so that it fails.
    mv /usr/bin/login /tmp/login_orig && \
    # Use nonsense for /usr/bin/login so that it fails.
    # It's not the same failure mode as in SELinux, but failure is good enough for test.
    echo "adsfasdfasdf" > /usr/bin/login && \
    chmod 755 /usr/bin/login && \
    # Simulate getenforce command
    printf "#!/bin/bash\necho 'Enforcing'" > /usr/bin/getenforce && \
    chmod 755 /usr/bin/getenforce && \
    eval `ssh-agent -s` && RSPSCALED_PATH=`pwd`rspscaled ./tailssh.test -test.v -test.run TestSSHAgentForwarding && \
    RSPSCALED_PATH=`pwd`rspscaled ./tailssh.test -test.v -test.run TestIntegration && \
    mv /tmp/login_orig /usr/bin/login && \
    rm /usr/bin/getenforce \
    ; fi

RUN echo "Then remove the login command and make sure tests still pass."
RUN rm `which login`
RUN eval `ssh-agent -s` && RSPSCALED_PATH=`pwd`rspscaled ./tailssh.test -test.v -test.run TestSSHAgentForwarding
RUN if echo "$BASE" | grep "ubuntu:"; then rm -Rf /home/testuser; fi
RUN RSPSCALED_PATH=`pwd`rspscaled ./tailssh.test -test.v -test.run TestIntegrationSFTP
RUN if echo "$BASE" | grep "ubuntu:"; then rm -Rf /home/testuser; fi
RUN RSPSCALED_PATH=`pwd`rspscaled ./tailssh.test -test.v -test.run TestIntegrationSCP
RUN if echo "$BASE" | grep "ubuntu:"; then rm -Rf /home/testuser; fi
RUN RSPSCALED_PATH=`pwd`rspscaled ./tailssh.test -test.v -test.run TestIntegrationSSH

RUN echo "Then remove the su command and make sure tests still pass."
RUN chown root:root /tmp/rspscalessh.log
RUN rm `which su`
RUN eval `ssh-agent -s` && RSPSCALED_PATH=`pwd`rspscaled ./tailssh.test -test.v -test.run TestSSHAgentForwarding
RUN RSPSCALED_PATH=`pwd`rspscaled ./tailssh.test -test.v -test.run TestIntegration

RUN echo "Test doDropPrivileges"
RUN RSPSCALED_PATH=`pwd`rspscaled ./tailssh.test -test.v -test.run TestDoDropPrivileges
