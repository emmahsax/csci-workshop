;;These are some examples of some problems in 4clojure

;Replicate a Sequence
(defn rep [coll n]
  (loop [c coll times n result []]
    (if (empty? c)
      result
      (if (= times 0)
      	(recur (rest c) n result)
      	(recur c (dec times) (conj result (first c)))))))

;Interpose a Sequence
(defn my-interpose [arbitrary collection]
  (reverse
    (loop [arb arbitrary coll collection result []]
       (if (empty? (rest coll))
           (conj result (first coll))
           (recur arb (rest coll) (flatten (conj result
                                                 (list arb
                                                       (first coll)))))))))

;Pack a Sequence
(defn pack [thing]
  (partition-by identity thing))

;Drop Every Nth Item
(defn dropn [collection n]
    (loop [coll collection index 1 result []]
      (if (empty? coll)
        result
        (if (= 0 (rem index n))
          (recur (rest coll) (inc index) result)
          (recur (rest coll) (inc index) (conj result (first coll)))))))

;Split a Sequence
(defn split [number collection]
  (list (take number collection)
        (drop number collection)))

;A Half-Truth
(defn half-truth [& bool]
  (if (every? true? bool)
    false
    (if (every? false? bool)
      false
      true)))

;Map Construction
(defn map-construct [vec1 vec2]
  (apply hash-map (interleave vec1 vec2)))

;Greatest Common Divisor
(defn greatest-common-divisor [n1 n2]
  (loop [n1 n1 n2 n2 current 1 gcd 0]
    (if (or (= current n1)
            (= current n2))
        (cond
         (= 0 (rem n2 n1)) n1
         (= 0 (rem n1 n2)) n2
         :else gcd)
        (if (and (= 0 (rem n1 current))
                 (= 0 (rem n2 current)))
            (recur n1 n2 (inc current) current)
            (recur n1 n2 (inc current) gcd)))))

;Set Intersection
(defn sect [s1 s2]
  (set (filter #(contains? s1 %) s2)))

;Re-implement Iterate
(defn my-iterate [func value]
  (lazy-seq
    (cons value (my-iterate func (func value)))))
;lazy-seq helps to not create a stackoverflow error or infinite recursion

;Product Digits
(defn product [n1 n2]
  (map #(Integer/parseInt %) (re-seq #"\d" (str (* n1 n2)))))

;Cartesian Products
(defn cartesian [s1 s2]
  (set (for [x s1 y s2] [x y])))

;Group a Sequence
(defn group [f coll]
  (into {}
        (map #(vector (f (first %)) (vec %))
             (partition-by f (sort coll)))))

;Read a Binary Number
(defn binary-to-decimal [string]
  (loop [binary (map #(- (int %) (int \0)) (seq string)) result 0 counter 1]
    (if (empty? binary)
      result
      (recur (butlast binary) (+ result (* (last binary) counter)) (* counter 2)))))

;Symmetric Difference
(defn set-difference [s1 s2]
  (set
   (concat
    (clojure.set/difference s1 s2)
    (clojure.set/difference s2 s1))))

;Dot Product
(defn dot-product [v1 v2]
  (reduce + (map #(* %1 %2) v1 v2)))

;Infix Calculator
(defn infix-calculator [& args]
  (reduce (fn [a [op b]] (op a b))
          (first args)
          (partition 2 (rest args))))

;Pascal's Triangle
(defn pascal [number]
  (condp = number
    1 [1]
    (let [row-before (pascal (dec number))
          row-before-half1 (concat [0] row-before)
          row-before-half2 (concat row-before [0])]
      (vec (map + row-before-half1 row-before-half2)))))

;Re-implement Map
(defn my-map [f coll]
  (if (false? (empty? coll))
        (lazy-seq
          (cons (f (first coll)) (my-map f (rest coll))))))

;Indexing Sequences
(defn index [coll]
  (map #(vector % (.indexOf coll %)) coll))

;Sum of Square of Digits
(defn sum-square-digits [c-of-ints]
  (loop [counter 0 c-of-ints c-of-ints]
    (if (empty? c-of-ints)
      counter
      (if (< (first c-of-ints)
             (reduce +
                (map #(* % %)
                     (map #(- (int %) (int \0))
                           (seq (str (first c-of-ints)))))))
          (recur (inc counter) (rest c-of-ints))
          (recur counter (rest c-of-ints))))))

;To Tree or Not To Tree
(defn binary-tree? [coll]
  (or (= nil coll)
      (and (= 3 (count coll))
           (not (sequential? (first coll)))
           (not (some true? coll))
           (not (some false? coll))
           (binary-tree? (second coll))
           (binary-tree? (last coll)))))

;Recognize Playing Cards
(defn card [string]
  (let [x (first (re-seq #"\w" string))
        y (first (re-seq #"\d" string))
        z (last (re-seq #"\w" string))]
    (hash-map :suit
              (cond
                (= "S" x) :spade
                (= "D" x) :diamond
                (= "H" x) :heart
                (= "C" x) :club)
              :rank
              (cond
                (= "2" y) 0
                (= "3" y) 1
                (= "4" y) 2
                (= "5" y) 3
                (= "6" y) 4
                (= "7" y) 5
                (= "8" y) 6
                (= "9" y) 7
                (= "T" z) 8
                (= "J" z) 9
                (= "Q" z) 10
                (= "K" z) 11
                (= "A" z) 12))))

;Least Common Multiple
(defn least-multiple [& nums]
  (letfn [(gcd [x y]
            (let [maximum (max x y)
                  minimum (min x y)
                  modd (mod maximum minimum)]
              (if (zero? modd)
                minimum
                (recur minimum modd))))
          (lcm [maximum minimum]
            (/ (* maximum minimum) (gcd maximum minimum)))]
    (reduce lcm nums)))

;Pascal's Trapezoid
(defn pascal-trap [v]
  (let [p (fn [new-v]
            (let [v1 (concat [0] new-v)
                  v2 (concat new-v [0])]
              (map +' v1 v2)))]
     (lazy-seq
       (concat [v] (pascal-trap (p v))))))

;Beauty is Symmetry
(defn symmetric-b-tree? [coll]
  (letfn [(mirror? [a b]
             (cond
               (not= (sequential? a) (sequential? b)) false
               (and (sequential? a) (sequential? b))
                  (let [[root-a Left-a Right-a] a
                        [root-b Left-b Right-b] b]
                     (and (= root-a root-b)
                          (mirror? Left-a Right-b)
                          (mirror? Left-b Right-a)))
               :else (= a b)))]
     (mirror? (second coll) (last coll))))
