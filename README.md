# GoLang

A repository containing of all the codes I referred to and wrote during my learning of the Go Programming Language. Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.   
  
    
    
    
#The following are the pointers I made for myself to refer to during this journey of mine.   
/*** VARIABLES ***/

-> Lower case variables are scoped to the package
-> Upper case variables can be read outside

3 visibility scopes:
1) package level and lower case - any file in package can access the variables
2) package level and upper case first letter - exported from the package and globally visible
3) block scope - scoped to only the block

Naming conventions:
1) length should reflect life of variable
2) keep acronyms upper case -> var theURL string = "...."  (readability purposes)
3) use pascal or camel case

ALL DECLARED VARIABLES MUST BE USED, OTHERWISE WILL THROW AND ERROR

/*** TYPE CONVERSION ***/    

int -> float32 we must do it explicitly

no implicit type conversion in Go

destinationType(value)
use strconv for string conversions


/*** ALL ABOUT VARIABLES ***/     
1) Variable declaration
2) We cannot redeclare variables, but can shadow them
3) All declared variables must be used
4) Visibility rules




/*** PRIMITIVE TYPES IN GO ***/    
1) Boolean type

    true and false
    default value is false

2) Numeric type     
    a) int
        Signed integers: int8, int16, int 32   
        Unsigned integers: uint8, uint16, uint32      
        Arithmetic operations:
            cannot add two different integers of different types 

    b) floating point
        follow IEEE754 format
        Initializer syntax initalizes it to float64 and arithmetic operations will become complex
    
    c) complex  
        complex64 and complex128
        basic 4 arithmetic operations

3) Text type   
    a) string   
        any utf8 character
        strings are aliases for bytes, so whenever we access any character using [] notation, we must convert it to string
        strings are immutable
    b) rune
        any utf32 character
        they are true type aliases for int32




/*** CONSTANTS ***/    

They are replaced by the compiler at compile time
can be any variable type
follow conventions as variables
they can be shadowed

supports implicit conversion but only with untyped consts

1) Naming convention
2) Types constants
3) Untyped constants
4) Enumerated constants
    a) iota allows us to create related constants
    b) starts at 0 and increments by 1
    c) watch out for constant values that match zero values for variables
5) Enumeration expressions
    a) Arithmetic, bitwise and bitshifting operations are allowed


/*** MAPS ***/     
1) value can be anything   
2) keys must be testable by equality (slices, maps, etc are not testable)   
3) map is dynamic   
4) return order of a map is not fixed    
5) they are passed by reference  

/*** STRUCTS ***/   
1) naming convention is same as variables   
2) they are passed by value   

Go does not provide inheritance, but provides embedded structs by means of composition


/*** CONDITIONAL STATEMENTS ***/      

Always use curly braces even for single line
For the Initializer syntax, the pop variable is only available within the if block
Possible operations:    
 1) >   
 2) <   
 3) ==    
 4) <=    
 5) >=    
 6) !=     

Logical operators:
 1) ||    
 2) &&    


Short circuiting - if any or test passes, the subsequent tests are not executed.


No fall through property in switch statements, but we can group multiple cases together. No overlapping cases allowed. In Go, the break statement is implied.
The fallthrough keyword will execute the following case regardless of whether it is true or not.



/*** CONTROL FLOW ***/

1) defer - executes a statement after a function has finished executing all statements other than return statement. All defer statements are called in LIFO order. Allows to associate closing and opening of resources right next to each other, while allowing resources to be utilized. It takes in the argument at the time it is called, rather than later on in the function.

2) panic - like throwing an exception. It occurs when program does not know how to proceed. Panics stop the termination of the program. Panics will be executed after the execution of any defer statements. 

3) recover - used to recover from a panic.


/*** POINTERS ***/   

& -> 'address of' operator   
* -> dereferencing operator   

No pointer arithmetic allowed.   
"unsafe" package is used for such operations.   

The zero value of a pointer is <nil>   

Complex types (structs) are automatically dereferenced for us.   
We wrap *ms in parantheses to access fields as the * has a lower precedence than the . , so to prevent that from evaluating into *(ms.foo), we do (*ms).foo .   
Maps and slices are initialized with internal maps, so copying them will point to the same data.   


/*** FUNCTIONS IN GO ***/   

pascal/camel case
visibility rules are same as variable
Functions pass and call by value
If we use pointers, then they are called by reference
Maps and slices will be called by reference by default
Variatic functions - variatic parameters must be the last parameter.
Anonymous functions must be invoked immediately or assigned to a avariable. 
Functions are also treated as types in Go.
Methods are special functions. A method is a function executed in context of a type. They also receive parameters by value, unless we use pointers.


/*** INTERFACES IN GO ***/

Interfaces describe methods
With nested interfaces, if any one method requires pointer parameters, then all the methods must be implemented with pointers.

Best practices:
1) Use many, smaller interfaces rather than one large one
2) Single method interfaces have most power
3) Don't export interfaces for type that will be consumed
4) Do export interfaces for types that will be used by the package
5) Design functions and method to accept interfaces whenever possible

Embedding interfaces possible
Type conversion
Empty interface and type switches

