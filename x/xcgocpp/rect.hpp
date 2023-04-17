#ifndef RECT_HPP
#define RECT_HPP

class Rectangle {
public:
    Rectangle(int, int);
    ~Rectangle();
    void set_values(int, int);
    int area(void);

private:
    int width, height;
};

#endif // RECT_HPP
