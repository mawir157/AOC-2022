import AdventHelper

import Data.List (nub, sortBy)
import Data.List.Split (splitOn)
import Data.Function (on)

type BeaconPair = ((Integer, Integer),(Integer, Integer))

parseLine :: String -> BeaconPair
parseLine s = ((sx, sy), (bx, by))
  where [lhs, rhs] = splitOn ": " s
        [xx,yy] = splitOn ", " $ drop 10 lhs
        sx = read (drop 2 xx) :: Integer
        sy = read (drop 2 yy) :: Integer
        [uu,vv] = splitOn ", " $ drop 21 rhs
        bx = read (drop 2 uu) :: Integer
        by = read (drop 2 vv) :: Integer

deadCellsOnRow :: Integer -> BeaconPair -> (Integer, Integer)
deadCellsOnRow n ((sx, sy), (bx, by)) = (w + n', e - n')
  where d = abs (sx - bx) + abs (sy - by)
        w = sx - d
        e = sx + d
        n' = abs (n - sy)

beaconsOnRow :: Integer -> [BeaconPair] -> [Integer]
beaconsOnRow _ [] = []
beaconsOnRow n ((_,(bx,by)):xs) = (if' (by == n) [bx] []) ++ beaconsOnRow n xs

combineRanges :: [(Integer, Integer)] -> [(Integer, Integer)]
combineRanges [x] = [x]
combineRanges ((xlo, xhi):(ylo, yhi):xs)
  | ((ylo - 1 <= xhi) && (xlo - 1 <= yhi)) = combineRanges (new:xs)
  | otherwise                              = [(xlo, xhi)] ++ combineRanges ((ylo, yhi):xs)
  where new = (min xlo ylo, max xhi yhi)

combine' :: Int -> [(Integer, Integer)] -> [(Integer, Integer)]
combine' 0 x = x
combine' n x
  | length x > 1 = combine' (n-1) $ sortBy (compare `on` fst) $ combineRanges x
  | otherwise    = x

processRow :: [BeaconPair] -> Integer -> Integer
processRow ps n = hi - lo + 1
  where deadRanges = map (deadCellsOnRow n) ps
        goodDeadRanges = filter (\(x1,x2) -> x1 <= x2) deadRanges
        (lo, hi) = head $ combine' 100 goodDeadRanges

isContiguous :: (Integer, Integer) -> [BeaconPair] -> Integer -> [(Integer, Integer)]
isContiguous (s, t) ps n
  | (length reduced == 1) = []
  | otherwise             = reduced
  where deadRanges = map (clamp (s, t)) $ map (deadCellsOnRow n) ps
        goodDeadRanges = filter (\(x1, x2) -> x1 <= x2) deadRanges
        reduced = combine' 100 goodDeadRanges

clamp :: (Integer, Integer) -> (Integer, Integer) -> (Integer, Integer)
clamp (limS, limE) (x, y) = (max limS x, min limE y)

main = do
  f <- readFile "../input/input15.txt"
  let row = 2000000
  let search = (0,4000000)
  let pairs = map (parseLine) $ lines f

  let part1 = (processRow pairs row) - (toInteger $ length $ nub $ beaconsOnRow row pairs)
  let [(index, intervals)] = take 1 $ filter (\(_,v) -> 1 < length v) $ zipWithFn (isContiguous search pairs) [(fst search)..(snd search)]
  printSoln 15 part1 (index + 4000000 * ((snd $ head intervals) + 1))
