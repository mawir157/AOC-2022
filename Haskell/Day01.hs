import AdventHelper

import Data.List (sort)
import Data.List.Split (splitOn)
 
calorieCount :: String -> String -> Integer
calorieCount sep s = sum (map read $ splitOn sep s :: [Integer])

main = do
  f <- readFile "../input/input01.txt"
  let elfs = parseLineGroups "|" (lines f)
  let calories = reverse $ sort $ map (calorieCount "|") elfs

  printSoln 1 (head calories) (sum $ take 3 calories)
