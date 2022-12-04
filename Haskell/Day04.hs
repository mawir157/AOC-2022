import AdventHelper

import Data.List.Split (splitOn)

check :: String -> (Bool, Bool)
check s = (((l1 <= r1) && (r2 <= l2)) || ((r1 <= l1) && (l2 <= r2)), not ((l2 < r1) || (r2 < l1)))
  where [lhs,rhs] = splitOn "," s
        [l1,l2] = map read (splitOn "-" lhs) :: [Integer]
        [r1,r2] = map read (splitOn "-" rhs) :: [Integer]

main = do
  f <- readFile "../input/input04.txt"

  printSoln 4 (length $ filter (fst . check) $ lines f) (length $ filter (snd . check) $ lines f)
