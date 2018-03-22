FROM scratch
ADD ec2-metadata-proxy /
CMD ["/ec2-metadata-proxy"]
