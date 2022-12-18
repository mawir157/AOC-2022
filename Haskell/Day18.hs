import AdventHelper

import Data.List (sortBy)
import Data.List.Split (splitOn)
import Data.Function (on)

type Brick = (Integer, Integer, Integer)
fst' :: (a,b,c) -> a
fst' (x,_,_) = x
snd' :: (a,b,c) -> b
snd' (_,x,_) = x
thd' :: (a,b,c) -> c
thd' (_,_,x) = x

parseInput :: String -> Brick
parseInput s = tuplify3 bs
  where bs = map read (splitOn "," s) :: [Integer]

countFaces :: Int -> [Brick] -> Integer
countFaces _ [x] = 0
countFaces n ((x1,y1,z1):(x2,y2,z2):bs)
  | n == 0 = (if' ((x1 + 1 == x2) && (y1 == y2) && (z1 == z2)) 1 0) + countFaces n ((x2,y2,z2):bs)
  | n == 1 = (if' ((x1 == x2) && (y1 + 1 == y2) && (z1 == z2)) 1 0) + countFaces n ((x2,y2,z2):bs)
  | n == 2 = (if' ((x1 == x2) && (y1 == y2) && (z1 + 1 == z2)) 1 0) + countFaces n ((x2,y2,z2):bs)

countExterior :: [Brick] -> Integer
countExterior bs = 6 * toInteger (length bs) - 2 * (x + y + z)
  where x = countFaces 0 $ sortBy ((compare `on` thd') <> (compare `on` snd') <> (compare `on` fst')) bs
        y = countFaces 1 $ sortBy ((compare `on` fst') <> (compare `on` thd') <> (compare `on` snd')) bs
        z = countFaces 2 $ sortBy ((compare `on` snd') <> (compare `on` fst') <> (compare `on` thd')) bs

generateBlocks :: (Integer, Integer, Integer) -> [Brick]
generateBlocks (nx,ny,nz) = [ (x,y,z) | x <- [0,1..nx+1], y <- [0,1..ny+1], z <- [0,1..nz+1] ]

maxDims :: [Brick] -> (Integer, Integer, Integer)
maxDims bs = (maximum $ map fst' bs, maximum $ map snd' bs, maximum $ map thd' bs)

isConnected :: ([Brick],[Brick]) -> ([Brick],[Brick])
isConnected (cc, []) = (cc, [])
isConnected (cc, q)
  | (length c) > 1 = (c, q')
  | otherwise      = (cc, q)
  where c = blockMatch cc q
        q' = filter (\x -> not $ elem x c) q

blockMatch :: [Brick] -> [Brick] -> [Brick]
blockMatch [] _ = []
blockMatch x [] = x
blockMatch ((x,y,z):xs) q = t ++ blockMatch xs q'
  where nbrs = [(x+1,y,z),(x-1,y,z),(x,y+1,z),(x,y-1,z),(x,y,z+1),(x,y,z-1)]
        t = filter (\b -> elem b nbrs) q
        q' = filter (\b -> not $ elem b t) q

getConnectedComponent :: ([Brick],[Brick]) -> ([Brick],[Brick])
getConnectedComponent p
  | p' == p   = p
  | otherwise = getConnectedComponent p'
  where p' = isConnected p

main = do
  f <- readFile "../input/input18.txt"
  let bs = map parseInput $ lines f

  let q = filter (\x -> not $ elem x bs) $ generateBlocks $ maxDims bs

  let t = getConnectedComponent ([(0,0,0)], q)
  let interior = init $ snd t -- for some reason we need to drop the final brick

  let part1 = countExterior bs
  let part2 = part1 - (countExterior interior)

  printSoln 18 part1 part2
