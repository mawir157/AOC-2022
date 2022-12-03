import AdventHelper

import Data.Char (ord)
import Data.List.Split (chunksOf) 

scoreChar :: Char -> Integer
scoreChar c = c' - (if' (c' <= 90) 38 96)
  where c' = toInteger $ ord c

score :: String -> Integer
score s = scoreChar $ head i
  where lhs = take (div (length s) 2) s
        rhs = drop (div (length s) 2) s
        i = filter (\x -> elem x lhs) rhs

assignBadge :: [String] -> Integer
assignBadge [e1,e2,e3] = scoreChar $ head i'
  where i  = filter (\x -> elem x e1) e2
        i' = filter (\x -> elem x i) e3

main = do
  f <- readFile "../input/input03.txt"
  let l = lines f

  printSoln 3 (sum $ map score l) (sum . map (assignBadge) $ chunksOf 3 l)
