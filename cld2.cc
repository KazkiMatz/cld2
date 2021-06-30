#include <cstddef>
#include <string.h>
#include <stdio.h>
#include <string>
#include <cstdlib>

#include "compact_lang_det.h"
#include "cld2.h"

const LangInfo* DetectLang(char *data, int length) {

    bool is_plain_text = true;
    CLD2::CLDHints cldhints = {NULL, NULL, 0, CLD2::UNKNOWN_LANGUAGE};
    bool allow_extended_lang = true;
    int flags = 0;
    CLD2::Language language3[3];
    int percent3[3];
    double normalized_score3[3];
    CLD2::ResultChunkVector resultchunkvector;
    int text_bytes;
    bool is_reliable;

    if (length <= 0) {
        length = strlen(data);
    }

    CLD2::Language summary_lang = CLD2::UNKNOWN_LANGUAGE;

    summary_lang = CLD2::ExtDetectLanguageSummary(data, 
            length,
            is_plain_text,
            &cldhints,
            flags,
            language3,
            percent3,
            normalized_score3,
            &resultchunkvector,
            &text_bytes,
            &is_reliable);

    LangInfo res[3] = {
      { CLD2::LanguageCode(language3[0]), percent3[0], is_reliable },
      { CLD2::LanguageCode(language3[1]), percent3[1], is_reliable },
      { CLD2::LanguageCode(language3[2]), percent3[2], is_reliable }
    };

    LangInfo* info_set = (LangInfo*)std::malloc(3*sizeof(LangInfo));
    memcpy(info_set, &res, sizeof(res));

    return info_set;
}
