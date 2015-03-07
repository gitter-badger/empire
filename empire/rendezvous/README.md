# Rendezvous

Rendezvous is a small TCP service that provides an access point for interactive
docker sessions. It's compatible with heroku's rendezvous service.

## Spec

* Clients connect over tls.
* When the client connects, it sends a secret followed by a `\r\n` sequence. The
  secret is the session identifier.
* `\0x03` and `\0x1C` represent SIGINT and SIGQUIT respectively. Rendezvous
  closes the tcp connection if these characters are received.
* Rendezvous will echo back what it receives from the connection, back to the
  sender.
