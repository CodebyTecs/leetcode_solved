func isIsomorphic(s string, t string) bool {
    mappingStoT := make(map[byte]byte)
    mappingTtoS := make(map[byte]byte)

    for i := 0; i < len(s); i++ {
        c1 := s[i]
        c2 := t[i]

        if mappingStoT[c1] == 0 && mappingTtoS[c2] == 0 {
            mappingStoT[c1] = c2
            mappingTtoS[c2] = c1
        } else if mappingStoT[c1] != c2 || mappingTtoS[c2] != c1 {
            return false
        }
    }

    return true
}