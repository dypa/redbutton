FROM scratch
COPY redbutton /app/
COPY web/ /app/
WORKDIR /app
CMD ["/app/redbutton"]
