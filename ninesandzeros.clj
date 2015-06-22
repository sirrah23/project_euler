;;;;Solves the ninesandzeros problem as posted on programming praxis
(ns ninesandzeros
	(:require [clojure.set :as set]))
(defn ninesandzeros 
	"Computes the smallest number made of only nines and zeroes
	that the number n will divide"
	[n]
	(defn n_and_z_helper
		"Takes the current 9*10^n and using all the previously
		computed values sees if any of their sums is divisible
		by n"
		[curr_num,history_set]
		(loop [new_set #{} curr_iter curr_num history 
			(sort (if (> curr_num 9) (set/union history_set #{9}) history_set))]
			(cond (= 0 (mod curr_iter n)) curr_iter
				  (= (set history) #{}) (n_and_z_helper (* 10 curr_num) (set/union history_set new_set))
				  :else (recur (conj new_set curr_iter) (+ curr_num (first history)) (rest history)))))
	(n_and_z_helper 9 #{}))

(ninesandzeros 23)
