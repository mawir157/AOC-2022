import AdventHelper

import Data.Char
import Data.List.Split (splitOn)

score :: String -> (Integer, Integer)
score s = (1 + rhs + 3 * result1, 1 + play2 + 3 * rhs)
  where parts = splitOn " " s
        lhs = toInteger ((ord . head $ head parts) - 65)
        rhs = toInteger ((ord . head $ last parts) - 88)
        result1 = mod (rhs - lhs + 4) 3
        play2 = mod (lhs + rhs + 2) 3

main = do
  f <- readFile "../input/input02.txt"
  let scores = map score (lines f)

  printSoln 2 (sum $ map fst scores) (sum $ map snd scores)
