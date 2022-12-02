import AdventHelper

import Data.Char
import Data.List.Split (splitOn)

score :: String -> (Integer, Integer)
score s = (1 + rhs + final1, 1 + play2 + final2)
  where parts = splitOn " " s
        lhs = toInteger ((ord . head $ head parts) - 65)
        rhs = toInteger ((ord . head $ last parts) - 88)
        result1 = mod (rhs - lhs) 3
        final1 = if' (result1 == 0) 3 (if' (result1 == 1) 6 0)
        final2 = if' (rhs == 1) 3 (if' (rhs == 2) 6 0)
        play2 = mod (if' (rhs == 0) (lhs - 1) (if' (rhs == 1) lhs (lhs + 1))) 3

main = do
  f <- readFile "../input/input02.txt"
  let scores = map score (lines f)

  printSoln 2 (sum $ map fst scores) (sum $ map snd scores)
