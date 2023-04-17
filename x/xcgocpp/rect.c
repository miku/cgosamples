#include "rect.hpp"
#include "rect.h"

Rect rect_new(int w, int h)
{
    Rectangle* r = new Rectangle(w, h);
    return (void*)r;
}

int rect_area(void* handle) {
    return reinterpret_cast<Rectangle*>(handle)->area();
}
