# show a next word in lexicographically by using given(input) letters
# think as base-26 symbol
# bad slow program. it calclates a number 26^100 on 100 letters word.

ALPHABET= {
  a: 1, b: 2, c: 3, d: 4, e: 5, f: 6, g: 7, h: 8, i: 9, j: 10,
  k: 11, l: 12, m: 13, n: 14, o: 15, p: 16, q: 17, r: 18, s: 19,
  t: 20, u: 21, v: 22, w: 23, x: 24, y: 25, z: 26
}

#def numbarize_alphabet(letter)
#  ALPHABET[letter.to_sym]
#end
#
#def is_same_letter?(word)
#  word.chomp.split("").uniq.length == 1
#end
#
#def twenty_six_baslize(word)
#  base_26_value = 0
#  word.chomp.split("").each_with_index do |letter, i|
#    base_26_value += numbarize_alphabet(letter) * 26**i
#  end
#  base_26_value
#end
#
#base_words = []
#gets.to_i.times do
#  base_words.push gets
#end
#
## create a hash the key is the word and the value is base-26 value
#base_words.each do |word|
#  replaced_words = {}
#  word = word.chomp
#  if is_same_letter?(word)
#    puts "no answer"
#  else
#    word.split("").permutation(word.length).to_a.each do |w|
#      replaced_words.store(w.join, twenty_six_baslize(w.join))
#    end
#
#    lexicographically_ordered_words = replaced_words.sort
#    puts lexicographically_ordered_words[lexicographically_ordered_words.index([word, twenty_six_baslize(word)]) + 1][0]
#  end
#end

# finally, I took an another technique below this
# https://www.nayuki.io/page/next-lexicographical-permutation-algorithm

class Array
  def swap!(a, b)
    raise ArgumentError unless a.between?(0, self.count-1) && b.between?(0, self.count-1)

    self[a], self[b] = self[b], self[a]

    self
  end

  def swap(a, b)
    self.dup.swap!(a, b)
  end
end

# first value of return array is swap value
def find_longest_no_increasing_word(word)
  reversed_word_letters = word.split("").reverse
  longest_word_letters = []
  for i in 0..word.length-2
    if reversed_word_letters[i] <= reversed_word_letters[i+1]
      longest_word_letters.push(reversed_word_letters[i])
    else
      longest_word_letters.push(reversed_word_letters[i])
      swap_letter = reversed_word_letters[i+1]
      break
    end
  end

  if swap_letter
    longest_word_letters.push(swap_letter).reverse
    longest_word_letters.push(swap_letter).reverse
  else
    #puts "no answer"
    false
  end
end

def find_pivot(letters)
  swap_letter = letters.shift
  least_gap = 10**5
  index = -1
  letters.each_with_index do |letter, i|
    gap = ALPHABET[letter.to_sym] - ALPHABET[swap_letter.to_sym]
    #gap = letter.to_i - swap_letter.to_i
    if swap_letter < letter && gap != 0 && gap <= least_gap
      least_gap = gap
      index = i
    end
  end
  index
end

base_words.each do |word|
  longest_no_increasing_word_letters = find_longest_no_increasing_word(word.chomp)
  if longest_no_increasing_word_letters
    pivot_index = find_pivot(longest_no_increasing_word_letters)
    longest_no_increasing_word_letters.swap!(0, pivot_index)
    swap = longest_no_increasing_word_letters.shift
    ordered_letters = longest_no_increasing_word_letters.reverse.unshift(swap)
    puts word.chomp.split("").take(word.chomp.length - ordered_letters.length).concat(ordered_letters).join
  else
    puts "no answer"
  end
end
