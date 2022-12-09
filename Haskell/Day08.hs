import AdventHelper

import Data.Char (digitToInt)
import Data.List (nub, reverse, transpose)

parseLine :: String -> [Integer]
parseLine l = map (toInteger . digitToInt) l :: [Integer]

parseInput :: [String] -> [[Integer]]
parseInput ss = map parseLine ss

rotateForestL :: [[a]] -> [[a]]
rotateForestL m = reverse $ transpose m

rotateForestR :: [[a]] -> [[a]]
rotateForestR m = transpose $ reverse m

rotateForest2 :: [[a]] -> [[a]]
rotateForest2 m = rotateForestL $ rotateForestL m

edgeSpotterL :: Integer -> [Integer] -> [Bool]
edgeSpotterL _ [] = []
edgeSpotterL n (l:s)
  | l > n     = [True] ++ edgeSpotterL l s
  | otherwise = [False] ++ edgeSpotterL n s

edgeSpotter :: [[Integer]] -> [[Bool]]
edgeSpotter m = map (edgeSpotterL (-1)) m

rowOr :: [Bool] -> [Bool] -> [Bool]
rowOr xs ys = map (\(l,r) -> l || r) $ zip xs ys

orMatrix :: [[Bool]] -> [[Bool]] -> [[Bool]]
orMatrix [] _ = []
orMatrix _ [] = []
orMatrix (x:xs) (y:ys) = [(rowOr x y)] ++ orMatrix xs ys

forestSpotter :: [[Integer]] -> [[Bool]]
forestSpotter m = orMatrix f3 $ orMatrix f2 $ orMatrix f1 f0
  where f0 = edgeSpotter m
        f1 = rotateForestR $ edgeSpotter $ rotateForestL m
        f2 = rotateForest2 $ edgeSpotter $ rotateForest2 m
        f3 = rotateForestL $ edgeSpotter $ rotateForestR m

part1 :: [[Bool]] -> Integer
part1 [] = 0
part1 (x:xs) = s + part1 xs
  where s = sum $ map (\b -> if' b 1 0) x

edgeCounterR :: [Integer] -> [Integer]
edgeCounterR [x] = [0]
edgeCounterR (x:xs) = [l] ++ edgeCounterR xs
  where l = toInteger $ length $ takeWhileInclusive (< x) xs

edgeCounter :: [[Integer]] -> [[Integer]]
edgeCounter m = map edgeCounterR m

rowMul :: [Integer] -> [Integer] -> [Integer]
rowMul xs ys = map (\(l,r) -> l * r) $ zip xs ys

mulMatrix :: [[Integer]] -> [[Integer]] -> [[Integer]]
mulMatrix [] _ = []
mulMatrix _ [] = []
mulMatrix (x:xs) (y:ys) = [(rowMul x y)] ++ mulMatrix xs ys

forestCounter :: [[Integer]] -> [[Integer]]
forestCounter m = mulMatrix f3 $ mulMatrix f2 $ mulMatrix f1 f0
  where f0 = edgeCounter m
        f1 = rotateForestR $ edgeCounter $ rotateForestL m
        f2 = rotateForest2 $ edgeCounter $ rotateForest2 m
        f3 = rotateForestL $ edgeCounter $ rotateForestR m

part2 :: [[Integer]] -> Integer
part2 m = maximum $ map maximum m

main = do
  f <- readFile "../input/input08.txt"
  let forest = map (parseLine) $ lines f

  printSoln 8 (part1 $ forestSpotter forest) (part2 $ forestCounter forest)
