#ifdef __cplusplus
extern "C" {
#endif

typedef void* Rect;

void* rect_new(int w, int h);

int rect_area(void *handle);

#ifdef __cplusplus
}
#endif
