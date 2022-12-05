import AdventHelper

import Data.List.Split (chunksOf, splitOn)
import Data.List (foldl')

parseChunks :: String -> [Char]
parseChunks (s:c:ss) = if' (s == ' ') [] [c]

parseInput :: [[Char]] -> String -> [[Char]]
parseInput xs s = map (\(lst, str) -> lst ++ parseChunks str) z
  where ss = chunksOf 4 s
        z = zip xs ss

parseRules :: String -> (Int, Int, Int)
parseRules s = t'
  where t = splitOn " " s
        t' = (read (t!!1) :: Int, read (t!!3) :: Int, read (t!!5) :: Int)

applyMove :: Bool -> [[Char]] -> (Int, Int, Int) -> [[Char]]
applyMove b ss (n, f, t) =  map fst ts'
  where ss' = zip ss [1,2..]
        ts  = map(\(s, i) -> (if' (f == i) (drop n s) s, i)) ss'
        tomove = take n $ ss!!(f-1)
        tomove' = if' b (reverse tomove) tomove
        ts' = map(\(s, i) -> (if' (t == i) (tomove' ++ s) s, i)) ts

main = do
  f <- readFile "../input/input05.txt"
  let temp = ["", "", "", "", "", "", "", "", ""]
  let blocks = foldl' (parseInput) temp $ take 8 $ lines f
  let moves = map (parseRules) $ drop 10 $ lines f
  let part1 = foldl' (applyMove True) blocks moves
  let part2 = foldl' (applyMove False) blocks moves

  printSoln 5 (map head part1) (map head part2)
