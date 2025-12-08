func backspaceCompare(s string, t string) bool {
    pleft := len(s) - 1
    pright := len(t) - 1

    skipS, skipT := 0, 0

    for pleft >= 0 || pright >= 0 {
        for pleft >= 0 {
            if s[pleft] == '#' {
                skipS++
                pleft--
            } else if skipS > 0 {
                skipS--
                pleft--
            } else {
                break
            }
        }

        for pright >= 0 {
            if t[pright] == '#' {
                skipT++
                pright--
            } else if skipT > 0 {
                skipT--
                pright--
            } else {
                break
            }
        }

        if pleft >= 0 && pright >= 0 && s[pleft] != t[pright] {
            return false
        }

        if (pleft >= 0) != (pright >= 0) {
            return false
        }

        pleft--
        pright--
    }

    return true
}