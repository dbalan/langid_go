#include "langid_go.h"
// #include "liblangid.h"
#include <string.h>

LanguageIdentifier *idfr;

void init() {
  idfr = get_default_identifier();
}

void destroy() {
  destroy_identifier(idfr);
}

const char *detect_language(char *text) {
  int text_len;
  const char *lang;

  text_len = strlen(text);

  lang = identify(idfr, text, text_len);

  return lang;
}
