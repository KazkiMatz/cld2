#ifdef __cplusplus
extern "C" {
#endif

#include <stdbool.h>

typedef struct {
    const char* code;
    int percent;
    bool is_reliable;
} LangInfo;

const LangInfo* DetectLang(char *data, int length);

#ifdef __cplusplus
}
#endif
