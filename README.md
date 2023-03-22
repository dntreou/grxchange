
# Grxchange

This is a follow-up implementation of the Grxchange-api, aiming to apply clean architecture and SOLID principles, creating a Graphql API.


## Architecture

For the architecture of this API I have chosen clean architecture introduced by Robert C. Martin (“Uncle Bob”)
this architecture similarly to other such as hexagonal, onion etc. 
aim to create modular application so that they can be easily maintained, concurrently developed and have a clear logical separation between the layers and the responsibilities of each accordingly.

![Clean Architecture](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

### graphql/
 contains the controllers/resolvers in relation to the figure above
### internal/ 
  This directory has a specal meaning in go, it hides the details below, meaning that what ever lives within the internal directory is only accessed by its parent
and here also live all the packages that contain the Use Cases or else the logic i.e internal/campaign.
### internal/domain
  In this directory lives the domain or else the Entities following the figure above and all the methods that validate set or get the various fields or the entity in its entirety.

## Methodology

We are going to utilize the [12th-Factor-App](https://12factor.net/) methodology in this project that provides us with guidelines and principles
based on years of experience from the writers, that help us tackle known enterprise software issues helping keep our product up to date, modular
and reliable.


## Design Patterns

## Coding Principles

## Choosing the right packages

### Graphql 

  #### Queries (read only)
  #### Mutation (read and then write)
  #### Subscriptions (open a socket connection)

### 4 Major Scalar types

  #### Int, String, Float, Boolean

#### Introspection support that answer questions about the API schema begin with __(double underscore) i.e __type, __schema, __typename.
##### Metafields are implicit 
##### type ahead graphqls feature 

### Graphql is a service
#### Structure: defined by the schema, that is like a catalog of all the operation a Graphql API can handle 
#### Behaviour: Implemented via functions called resolvers, each field in the Graphql schema is backed bya resolver function. A resolver function defines what data to fetch for its field

