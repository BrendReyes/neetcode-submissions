func isValidSudoku(board [][]byte) bool {
    digits := make(map[byte]bool)

    // Loop through each row
    for _, row := range board {
        for _, val := range row {
            _, exist := digits[val]
            if exist && unicode.IsDigit(rune(val)) {
                return false
            }
            digits[val] = true
        }
        clear(digits)
    }

    // Loop through each column
    for col := 0; col < len(board[0]); col++ {
        for row := 0; row < len(board); row++ {
            _, exist := digits[board[row][col]]
            if exist && unicode.IsDigit(rune(board[row][col])) {
                return false
            }
            digits[board[row][col]] = true
        }
        clear(digits)
    }
    // Loop through each 3x3 box
    for boxRow := 0; boxRow < 3; boxRow++ {
        for boxCol := 0; boxCol < 3; boxCol++ {
            for r := 0; r < 3; r++ {
                for c := 0; c < 3; c++ {
                    val := board[boxRow*3+r][boxCol*3+c]
                    _, exist := digits[val]
                    if exist && unicode.IsDigit(rune(val)) {
                        return false
                    }
                    digits[val] = true
                }
            }
            clear(digits)
        }
    }
    return true
}