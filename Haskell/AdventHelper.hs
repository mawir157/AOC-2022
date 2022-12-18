module AdventHelper where
import Data.List
import Data.List.Split
import Data.Maybe
import qualified Data.Map as Map

splitOnAnyOf :: Eq a => [[a]] -> [a] -> [[a]]
splitOnAnyOf ds xs = foldl' (\ys d -> ys >>= splitOn d) [xs] ds

if' :: Bool -> a -> a -> a
if' True  x _ = x
if' False _ y = y

zipWithFn :: (a -> b) -> [a] -> [(a,b)]
zipWithFn fn as  = zip as (map fn as)

printSoln :: (Show a, Show b) => Integer -> a -> b -> IO()
printSoln d p1 p2 = putStrLn $ unlines [l0, l1, l2]
  where l0 = "Day " ++ (show d)
        l1 = "  Part 1: "  ++ (show p1)
        l2 = "  Part 2: "  ++ (show p2)

tuplify2 :: [a] -> (a,a)
tuplify2 [x,y] = (x,y)
tuplify2 _ = error "Can't tuplify this array"

tuplify3 :: [a] -> (a,a,a)
tuplify3 [x,y,z] = (x,y,z)
tuplify3 _ = error "Can't tuplify this array"

concat' :: [[a]] -> [a] -> [a]
concat' [] _ = []
concat' [s] _ = s
concat' (s:ss) c  = s ++ c ++ concat' ss c

parseLineGroups :: String -> [String] -> [String]
parseLineGroups _ [] = []
parseLineGroups c ss = concat' b c : parseLineGroups c ss' 
  where b = takeWhile (not . null) ss
        ss' = drop 1 $ dropWhile (not . null) ss

readInt :: String -> Integer
readInt ss = read (takeWhile (`elem` "1234567890") ss) :: Integer

replaceFirst :: String -> String -> String -> String
replaceFirst _ _ [] = []
replaceFirst old new ss
  | take l ss == old = new ++ drop l ss
  | otherwise        = head ss : replaceFirst old new (tail ss)
  where l = length old

diff :: [Integer] -> [Integer]
diff [] = error "Diff of single element is nonsensical"
diff [_] = []
diff (x:y:xs) = (x - y) : diff (y : xs)

minPair :: (Ord a) => [(a,b)] -> (a,b)
minPair [] = error "Empty List"
minPair [x] = x
minPair (x:y:xs) = if' (fst x < fst y) (minPair (x:xs)) (minPair (y:xs))

chiRemThm :: (Integer, Integer) -> (Integer, Integer) -> (Integer, Integer)
chiRemThm (a1, p1) (a2, p2) = (a3, p1 * p2)
  where a3 = head $ filter (\x -> x `mod` p2 == a2 `mod` p2) cands
        cands = [ a1 + n * p1 | n <- [1..p2]]

range :: (Int, Int) -> [Int]
range (x,y) = [x, (x + signum (y - x)) .. y]

freq :: (Eq a, Ord a) => [a] -> [(a, Int)]
freq xs = map (\x -> (head x, length x)) . group . sort $ xs

incr :: (Ord a) => Int -> Map.Map a Int -> a -> Map.Map a Int
incr i m k
  | Map.member k m = Map.adjust (\v -> v + i) k m
  | otherwise      = Map.insert k i m

decr :: (Ord k) => k -> Int -> Map.Map k Int -> Map.Map k Int
decr k d m
  | fromJust (Map.lookup k m) > d = Map.adjust (\v -> v - d) k m
  | otherwise                     = Map.delete k m

takeWhileInclusive :: (a -> Bool) -> [a] -> [a]
takeWhileInclusive _ [] = []
takeWhileInclusive p (x:xs) = x : if p x then takeWhileInclusive p xs
                                         else []

takeWhileInc :: Ord a => [a] -> [a]
takeWhileInc (x:x':xs)
  | x < x'    = x : takeWhileInc (x':xs)
  | otherwise = [x]
takeWhileInc xs = xs
