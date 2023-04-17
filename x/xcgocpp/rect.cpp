#include "rect.h"

Rectangle::Rectangle(int w, int h)
{
    width = w;
    height = h;
};

Rectangle::~Rectangle() { }

void Rectangle::set_values(int w, int h)
{
    width = w;
    height = h;
};
int Rectangle::area(void)
{
    return width * height;
}
;
