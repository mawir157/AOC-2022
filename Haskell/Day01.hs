import AdventHelper

import Data.List
import Data.List.Split

calorieCount :: String -> String -> Integer
calorieCount sep s = sum cs
  where cs = map read $ splitOn sep s :: [Integer]

main = do
  putStrLn "Day 1"
  f <- readFile "../input/input01.txt"
  let elfs = parseLineGroups "|" (lines f)
  let calories = reverse $ sort $ map (calorieCount "|") elfs

  printSoln 1 $ head calories
  printSoln 2 $ sum $ take 3 calories
