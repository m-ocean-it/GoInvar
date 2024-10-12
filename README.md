# GoInvar: encoding invariants with types

Examples of using the library can be found [here](https://github.com/m-ocean-it/GoInvarCollection).

The library allows you to define *invariant holders*, i.e. types of elements that upheld specified conditions, or invariants. The invariants are checked upon initialization and subsequent unwrapping of holders.

Checking the invariants upon unwrapping is relevant only when the wrapped type is externally modifiable via a pointer (for example, a map or a slice), since directly mutating the internal value of a invariant holder is impossible. But, for now, at least, the checks during unwrapping are performed on any type of internal value.