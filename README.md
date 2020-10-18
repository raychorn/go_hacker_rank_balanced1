# Hacker Rank - Most Balanced Partition

## Table of Contents

- [Preface](#preface)
- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Contributing](../CONTRIBUTING.md)

## Preface <a name = "preface"></a>

I am not sure why this problem exists other than the folks at Hacker Rank who probably spend all their time dreaming this stuff up might have thought this was
interesting.

The only part of this problem that seems the least bit interesting to me is the obvious programming effort required to make logical sense of the really crappy way
the data was architected.  Two arrays to describe a set of trees?  Really?!?  Who would dream this up?  I cannot say anything charitable about any of this but this
was the problem as best I could recall.  I am sure Hacker Rank has more of these problems than anyone might be able to document and I am sure none of it makes
any real sense.

And what's the deal about putting an artificial time limit on coding the problem?  When does this occur in real life?  Never?  I have been working in the industry for
longer than I like to recall and NEVER have I ever been faced with getting only 30 minutes to get code written to resolve any issue.  

Ok.  Let's assume you can buy into the problem.  What follows here was my take on it.  Your mileage may vary.

## About <a name = "about"></a>

Consider two arrays as parent {-1, 0, 1, 2, -1} and sizes {1, 4, 3, 4, 1} where each node of the tree has a pointer to another node with the base of the tree as -1 and
each node has a size. Make the trees, however many as may exist, be as balanced as possible and return the difference between the minimum and maximum as a positive integer.

You may cut one edge to balance the trees. (I am a bit fuzzy as to this part of the problem statement.)  You can cut any edge but once but the edge is gone?  (Again I am fuzzy and dubious
about this too.)

## Getting Started <a name = "getting_started"></a>

Checkout the code and have fun.

### Prerequisites

Pay me if you need help. Seriously, I need the cash. Otherwise you're on your own.

```
Give examples
```

### Installing

Go figure.  Really. It's GO.  Or it's Go time. So Go.

```
Give the example
```

And repeat

```
until finished
```

Make it work.

## Usage <a name = "usage"></a>

Use it any way you wish.

Yeah, there a License and a Copyright notice.  Deal with it.
(c). Copyright Ray C Horn, All Rights Reserved.