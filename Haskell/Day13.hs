import AdventHelper

import Data.List (sort, elemIndex)
import Data.Maybe (fromJust)

parseLinePairs :: [String] -> [(String, String)]
parseLinePairs [] = []
parseLinePairs ss = [(l, r)] ++ parseLinePairs ss' 
  where [l,r] = takeWhile (not . null) ss
        ss' = drop 1 $ dropWhile (not . null) ss

data Tree = Leaf Integer | Node [Tree] deriving (Show, Eq)

instance Ord Tree where
  compare (Leaf i)    (Leaf i')   = compare i i'
  compare (Node is)  (Node is') = cmp is is'
  compare l@(Leaf _)  n@(Node _) = compare (Node [l]) n
  compare n@(Node _) l@(Leaf _)  = compare n (Node [l])

cmp :: [Tree] -> [Tree] -> Ordering
cmp [] is = if is == [] then EQ else LT
cmp _  [] = GT
cmp (i:is) (i':is') = if' (c == EQ) (cmp is is') c
  where c = compare i i'

-- TIDY UP THIS JUNK --
parseHelper :: String -> String -> [String]
parseHelper _ [] = []
parseHelper curr (x:xs)
  | x == ','    = if' (curr == "") p (curr:p)
  | elem x "[]" = if' (curr == "") ([x]:p) (curr:[x]:p)
  | otherwise   = parseHelper (curr ++ [x]) xs
  where p = parseHelper "" xs

getList :: Int -> [String] -> ([String], [String])
getList c (x:xs)
  | x == "]" && c == 0 = ([], xs)
  | otherwise = let (l, res) = getList nc xs in (x:l, res)
  where nc = c + if' (x == "]") (-1) (if' (x == "[") 1 0)  

buildHelper :: [String] -> [Tree]
buildHelper [] = []
buildHelper (x:xs)
  | x == "["  = (Node (buildHelper l)):(buildHelper res)
  | otherwise = (Leaf (read x)):(buildHelper xs)
  where (l, res) = getList 0 xs

buildTree :: String -> Tree
buildTree s = head $ buildHelper $ parseHelper "" s
--
check :: (Tree, Tree) -> Bool
check (l, r) = (compare l r /= GT)

flattenPairs :: [(Tree, Tree)] -> [Tree]
flattenPairs [] = []
flattenPairs (x:xs) = [fst x, snd x] ++ flattenPairs xs

main = do
  f <- readFile "../input/input13.txt" 
  let pairs = parseLinePairs $ lines f
  let treePairs = map (\(l, r) -> (buildTree l, buildTree r)) pairs

  let good = map snd $ filter (\x -> fst x) $ zip (map (check) treePairs) [1,2..]
  let p21 = buildTree "[[2]]"
  let p22 = buildTree "[[6]]"
  let allTrees = (flattenPairs treePairs) ++ [p21] ++ [p22]
  let sorted = sort allTrees

  let a1 = fromJust (elemIndex p21 sorted)
  let a2 = fromJust (elemIndex p22 sorted)

  printSoln 13 (sum good) ((a1 + 1)*(a2 + 1))
