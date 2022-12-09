import AdventHelper

import Data.List.Split (splitOn)
import Data.List (nub, scanl')

parseInput :: String -> (String, Int)
parseInput s = (d, n')
  where [d, n] = splitOn " " s
        n' = read n :: Int

updateT :: (Integer, Integer) -> (Integer, Integer) -> (Integer, Integer)
updateT (x1, y1) (x0, y0)  = if' ((abs dx > 1) || (abs dy > 1)) (x0 + signum dx, y0 + signum dy) (x0, y0)
  where (dx, dy) = (x1 - x0, y1 - y0)

pullChain :: [(Integer, Integer)] -> [(Integer, Integer)]
pullChain [x] = [x]
pullChain (x:y:xs) = [x] ++ pullChain ((updateT x y):xs)

moveHead :: (Integer, Integer) -> String -> (Integer, Integer)
moveHead (x, y) d
  | d == "U" = (x + 1, y)
  | d == "D" = (x - 1, y)
  | d == "L" = (x, y - 1)
  | d == "R" = (x, y + 1)

applyMove :: [(Integer, Integer)] -> String -> [(Integer, Integer)]
applyMove (c:cs) d = pullChain ((moveHead c d):cs)

expandMoves :: [(String, Int)] -> [String]
expandMoves [] = []
expandMoves ((d, n):xs) = (replicate n d) ++ expandMoves xs

main = do
  f <- readFile "../input/input09.txt"
  let ms = expandMoves $ map parseInput $ lines f

  let c1 = scanl' applyMove (replicate 2 (0,0)) ms
  let c2 = scanl' applyMove (replicate 10 (0,0)) ms

  printSoln 9 (length $ nub $ map last c1) (length $ nub $ map last c2) 
