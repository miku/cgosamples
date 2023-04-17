#include "rect.hpp"
#include "rect.h"

Rect rect_new(int w, int h)
{
    Rectangle* r = new Rectangle(w, h);
    return (void*)r;
}

int rect_area(void* handle) {
    // https://en.cppreference.com/w/cpp/language/reinterpret_cast
    // It is purely a compile-time directive which instructs the compiler to
    // treat expression as if it had the type new-type.
    //
    // A practical use case of "opaque" data types is when you want to expose
    // an API to C but write the implementation in C++.
    //
    // https://stackoverflow.com/questions/573294/when-to-use-reinterpret-cast#comment73033255_573574
    return reinterpret_cast<Rectangle*>(handle)->area();
}
