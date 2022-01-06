# Go messenger

## Glossary

### Bus

The bus is used to dispatch messages. The behavior of the bus is in its ordered middleware stack. The component comes
with a set of middleware that you can use.

### Envelope

Envelope is a wrapper for messages dispatched into the bus. Envelopes can be wrapped into more envelopes each of them
additional information to it.

### Matcher

Matchers are used by middlewares to filter envelopes that should be handled by them.

### Middleware

Middleware can access the message and its wrapper (the envelope) while it is dispatched through the bus. Literally
"the software in the middle", those are not about core concerns (business logic) of an application. Instead, they are
crosscutting concerns applicable throughout the application and affecting the entire message bus. For instance: logging,
validating a message, starting a transaction, ... They are also responsible for calling the next middleware in the
chain, which means they can tweak the envelope, by adding headers to it or even replacing it, as well as interrupt the
middleware chain. Middleware are called both when a message is originally dispatched and again later when a message is
received from transport.

### Handler

Responsible for handling messages using the business logic applicable to the messages.

### Sender

Responsible for sending messages to something. This something can be a message broker or a third party API for example.

### Receiver

Responsible for retrieving messages from external source. This can be a message queue puller or an API endpoint for
example. Receivers also responsible for acking or rejecting messages received by them.

### Bridge

Is responsible for connecting between a receiver and a bus. They get messages from receivers and dispatch them into the
bus.
