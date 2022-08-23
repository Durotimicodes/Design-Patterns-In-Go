# factory design pattern motivation

1. used when object logic creation becomes to convoluted.
2. struct has too many fields want to initialize them properly.
3. It is an wholesale object creation, not a piece-wise, unlike builder, and 
can be outsourced to a seperate factory function aka constructor. They may also
exist in a seperate struct (Factory) basically for the sake of organization.

# factory is a component responsible for the wholesale creation of objects.
# A factory function is a free standing function which returns an instance of the object you want to creat

