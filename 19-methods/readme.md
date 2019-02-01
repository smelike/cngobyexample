Methods

Go supports methods defined on struct types.


This area method has a receiver type of *rect.

Methods can be defined for either pointer or value receiver types. Here's an example of a value receiver.


Here we call the 2 methods defined for our struct.

Go automatically handles conversion between values and pointers for method calls. You may want to use a pointers 
receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.


Next we'll look at Go's mechanism for grouping and naming related sets of methods: interfaces.

