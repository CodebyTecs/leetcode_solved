func compress(chars []byte) int {
    writeIndex := 0
    readIndex := 0

    for readIndex < len(chars) {
        char := chars[readIndex]
        count := 0

        for readIndex < len(chars) && chars[readIndex] == char {
            readIndex++
            count++
        }

        chars[writeIndex] = char
        writeIndex++

        if count > 1 {
            for _, c := range fmt.Sprintf("%d", count) {
                chars[writeIndex] = byte(c)
                writeIndex++
            }
        }
    }

    return writeIndex
}