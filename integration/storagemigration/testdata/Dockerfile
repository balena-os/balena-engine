FROM busybox:1.33.0
RUN mkdir /tmp/d1 && touch /tmp/d1/d1f1 && touch /tmp/f1 && touch /tmp/f2
RUN rm -R /tmp/d1 && mkdir /tmp/d1 && touch /tmp/d1/d1f2 && rm /tmp/f1
RUN ln -s /tmp/d1/d1f2 /tmp/slkn
RUN ln /tmp/d1/d1f2 /tmp/hlkn
