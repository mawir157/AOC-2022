import AdventHelper

import Data.List (nub)

getMarker :: Int -> Int -> String -> Int
getMarker l n s
  | length hd == l = n + (l - 1)
  | otherwise      = getMarker l (n + 1) (tail s)
  where hd = nub $ take l s

main = do
  f <- readFile "../input/input06.txt"

  printSoln 6 (getMarker 4 1 $ head $ lines f) (getMarker 14 1 $ head $ lines f)
