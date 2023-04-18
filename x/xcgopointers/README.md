# Pointers

> Go is a garbage collected language, and the garbage collector needs to know the location of every pointer to Go memory. Because of this, there are restrictions on passing pointers between Go and C. 

Pointers are part of slices, maps, etc.

> Note that values of some Go types, other than the type's zero value, always include Go pointers. This is true of string, slice, interface, channel, map, and function types. A pointer type may hold a Go pointer or a C pointer. Array and struct types may or may not include Go pointers, depending on the element types.