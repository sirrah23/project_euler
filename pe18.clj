(ns pe18)
(use 'clojure.java.io)
(require '[clojure.string :as str])

(defn vec-builder 
	"Parses the file and creates a vector of vectors"
	[]
	(with-open [rdr (reader "/home/harris/pyramid.txt")]
		(loop [ret [] line (first (line-seq rdr))]
			(if (= line nil) 
				ret 
				(recur (conj ret (into [] (map read-string (str/split line #" ")))) 
				(first (line-seq rdr)))))))

(defn max-reduce
	"Takes vector and performs pairwise max reduction" 
	[v] 
	(map (partial reduce max) (map vector v (rest v))))


(defn maximal-path
	"Computes maximal path in pyramid" 
	[vec-built]
	(if (= () (rest vec-built))
		(first vec-built)
		(map + (first vec-built) (max-reduce (maximal-path (rest vec-built))))))

(defn -main
	"run main"
	[]
	(maximal-path (vec-builder)))

(-main)