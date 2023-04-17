#ifndef RECT_FILE_H
#define RECT_FILE_H

class Rectangle {
public:
    Rectangle(int, int);
    ~Rectangle();
    void set_values(int, int);
    int area(void);

private:
    int width, height;
};

#endif // RECT_FILE_H
