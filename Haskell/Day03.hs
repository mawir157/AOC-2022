import AdventHelper

import Data.Char (ord)

scoreChar :: Char -> Integer
scoreChar c = if' ((ord c) <= 90) upper lower
  where upper = toInteger ((ord c) - 38)
        lower = toInteger ((ord c) - 96)

score :: String -> Integer
score s = scoreChar $ head i
  where lhs = take (div (length s) 2) s
        rhs = drop (div (length s) 2) s
        i = filter (\x -> elem x lhs) rhs

assignGroupBadge :: [String] -> Integer
assignGroupBadge [e1,e2,e3] = scoreChar $ head i'
  where i  = filter (\x -> elem x e1) e2
        i' = filter (\x -> elem x i) e3

assignBadges :: [String] -> [Integer]
assignBadges [] = []
assignBadges s = [assignGroupBadge (take 3 s)] ++ assignBadges (drop 3 s)

main = do
  f <- readFile "../input/input03.txt"

  printSoln 3 (sum $ map score $ lines f) (sum $ assignBadges $ lines f)
