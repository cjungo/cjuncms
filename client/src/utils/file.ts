const GET_SUFFIX_PATTERN = /(?:.+?)\.(.+)$/;
const GET_SUFFIX_PATTERN_WITH_DOT = /(?:.+?)(\..+)$/;

export const getSuffix = (path: string, withDot: boolean = false): string => {
  const m = path.match(
    withDot ? GET_SUFFIX_PATTERN_WITH_DOT : GET_SUFFIX_PATTERN
  );
  return m ? m[1] : "";
};
