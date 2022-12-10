import AdventHelper

import Data.List.Split (chunksOf)
import Data.List (scanl')

parseInput :: String -> (String, Integer)
parseInput s = if' (d == "addx") (d, n) (d, 0)
  where d = take 4 s
        n = read (drop 5 s) :: Integer

expandInstructions :: [(String, Integer)] -> [(String, Integer)]
expandInstructions [] = []
expandInstructions (x:xs) = x' ++ expandInstructions xs
  where x' = if' (fst x == "addx") [("noop", 0), x] [x]

sumAtIndices :: [Int] -> [Integer] -> Integer
sumAtIndices [] _ = 0
sumAtIndices (i:is) l = ((toInteger i) * (l!!(i-1))) + sumAtIndices is l

applyInstruction :: (Integer, Integer) -> (String, Integer) -> (Integer, Integer)
applyInstruction (reg, cycles) (ins, val) = (reg', cycles + 1)
  where reg' = if' (ins == "addx") (reg + val) reg

renderCRT :: Integer -> [Integer] -> String
renderCRT _ [] = []
renderCRT pixel (s:ss) = s' ++ renderCRT pixel' ss
  where pixel' = mod (pixel + 1) 40
        s' = if' (elem pixel [s - 1, s, s + 1]) "#" " "

main = do
  f <- readFile "../input/input10.txt"
  let ins = expandInstructions $ map parseInput $ lines f
  let vs = map fst $ scanl' applyInstruction (1,0) ins
  let crt = chunksOf 40 $ renderCRT 0 vs

  printSoln 10 (sumAtIndices [20, 60 ,100, 140, 180, 220] vs) "See Below:"
  mapM_ print crt
