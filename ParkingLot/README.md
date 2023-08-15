# Parking Lot Problem

See [website](https://adaickalavan.github.io/portfolio/parking_lot_problem/) for information.

## Problem Statement

Mr Zorro owns a multi-storey parking lot that can hold up to `n` vehicles at any given point in time. The parking slots are numbered, beginning at 1 and increases with increasing distance from the entry point in steps of one. Mr Zorro has requested your help to design an automated ticketing system for his parking lot.

When a vehicle enters the parking lot, its vehicle registration number (i.e., number plate) and colour are noted. Then, an available parking slot is allocated. Following are the rules of parking slot ticket issuance:

+ Each customer should be allocated the nearest available parking slot to the entry point.
+ Upon exiting the parking lot, the customer returns the ticket which marks their previously allocated lot as now available.
+ Due to government regulation, the system should provide the ability to determine:
    + Registration numbers of all cars of a particular colour.
    + Slot number in which a car with a given registration number is parked.
    + Slot numbers of all slots where a car of a particular colour is parked.

The ticketing system should be operable via two modes of input, namely, interactive commands and commands from a file. In other words, the ticketing system should be an executable which accepts:

1. Interactive commands from an interactive command prompt shell
2. A filename as an input argument at the command prompt and executes the commands from the given file

Example below includes all the commands which need to be supported.

**Example: File Input**

To run the code, it accepts input from a file:
```
$ bin/parking_lot file_inputs.txt
```
Input (contents of file):
```
create_parking_lot 6
park KA-01-HH-1234 White
park KA-01-HH-9999 White
park KA-01-BB-0001 Black
park KA-01-HH-7777 Red
park KA-01-HH-2701 Blue
park KA-01-HH-3141 Black
leave 4
status
park KA-01-P-333 White
park DL-12-AA-9999 White
registration_numbers_for_cars_with_colour White
slot_numbers_for_cars_with_colour White
slot_number_for_registration_number KA-01-HH-3141
slot_number_for_registration_number MH-04-AY-
```
Output (to STDOUT):
```
Created a parking lot with 6 slots
Allocated slot number: 1
Allocated slot number: 2
Allocated slot number: 3
Allocated slot number: 4
Allocated slot number: 5
Allocated slot number: 6
Slot number 4 is free
Slot No. Registration No Colour
1        KA-01-HH-1234   White
2        KA-01-HH-9999   White
3        KA-01-BB-0001   Black
5        KA-01-HH-2701   Blue
6        KA-01-HH-3141   Black
Allocated slot number: 4
Sorry, parking lot is full
KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333
1, 2, 4
6
Not found
```

**Example: Interactive**

To run the code, launch the shell, and to accept interactive input from the shell:
```
$ bin/parking_lot
```
Assuming a parking lot with `n=6` slots, the following commands should be run in sequence by typing them in at a prompt and should produce output as described below the command. Note that `exit` terminates the process and returns control to the shell.
```
$ create_parking_lot 6
Created a parking lot with 6 slots
$ park KA-01-HH-1234 White
Allocated slot number: 1
$ park KA-01-HH-9999 White
Allocated slot number: 2
$ park KA-01-BB-0001 Black
Allocated slot number: 3
$ park KA-01-HH-7777 Red
Allocated slot number: 4
$ park KA-01-HH-2701 Blue
Allocated slot number: 5
$ park KA-01-HH-3141 Black
Allocated slot number: 6
$ leave 4
Slot number 4 is free
$ status
Slot No. Registration No Colour
1        KA-01-HH-1234   White
2        KA-01-HH-9999   White
3        KA-01-BB-0001   Black
5        KA-01-HH-2701   Blue
6        KA-01-HH-3141   Black
$ park KA-01-P-333 White
Allocated slot number: 4
$ park DL-12-AA-9999 White
Sorry, parking lot is full
$ registration_numbers_for_cars_with_colour White
KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333
$ slot_numbers_for_cars_with_colour White
1, 2, 4
$ slot_number_for_registration_number KA-01-HH-3141
6
$ slot_number_for_registration_number MH-04-AY-1111
Not found
$ exit
```

## Learning Outcome

At the end of this project, we should be able to:

+ solve the Parking Lot problem with data structures optimized for complexity
+ write comprehensive unit tests in Golang
+ perform functional testing in Golang
+ utilize heap, hash map, and slice, data structures in Go
+ pretty print slices, arrays, and strings, using an interface
