# BigNum - Arbitrarily precise math and numbers.

BigNum is a library to create and calculate numbers of arbitrary size and precision. 

## Why? Doesn't Go have `big` already?

This is mostly a learning experiment for me and is not currently ready to be used for anything mission critical. The goal was to try and make my implementation of [strtod](https://github.com/CrazyInfin8/StrToD) more accurate as the exponent gets further from 0. I plan to port this to C and potentially other langauges in the future.

## Simple(-ish)

As this is a learning experiment for me, I tried to keep things (relatively) simple. Each operation is also not optimized and their implementations are mainly based on grade school math. Almost all operations also create new objects for the results instead of performing the operation in place, which could consume more memory but is easier to implement for now.

Each digit of BigInt is a base-256 number represented by a single byte. In the future, this could potentially be extrapolated to use base-4294967296 (32-bits per digit) or base-18446744073709551616 (64-bits per digit). For now it is easier to compare multiple bytes against native 32/64-bit numbers for testing.


# What is complete and what is there to do

(✅ Completed, ❌ To do)

## BigInt

- ✅ `+` Addition
- ✅ `-` Subtraction
- ✅ `*` Multiplication
- ❌ `/` Division
- ✅ `>>`/`<<` Bit shifting
- ❌ `==`/`<`/`>` Comparison
- ❌ Exponents
- ❌ Roots
- ❌ Logorithms

## BigFloat (TBD)

- ❌ `+` Addition
- ❌ `-` Subtraction
- ❌ `*` Multiplication
- ❌ `/` Division
- ❌ `>>`/`<<` Bit shifting
- ❌ `==`/`<`/`>` Comparison
- ❌ Exponents
- ❌ Roots
- ❌ Logorithms
