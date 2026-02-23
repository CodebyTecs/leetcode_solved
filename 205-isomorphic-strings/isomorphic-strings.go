func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    sMap := make(map[byte]byte, len(s))
    tMap := make(map[byte]byte, len(t))

    for i := 0; i < len(s); i++ {
        sCh, tCh := s[i], t[i]
        if v, ok := sMap[sCh]; ok {
            if v != tCh {
                return false
            }
        } else {
            sMap[sCh] = tCh
        }
        if v, ok := tMap[tCh]; ok {
            if v != sCh {
                return false
            }
        } else {
            tMap[tCh] = sCh
        }
    }

    return true
}