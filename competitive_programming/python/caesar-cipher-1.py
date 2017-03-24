import sys
import string
import pdb

N = int(input())
S = input()
K = int(input())

LOWER_ALPHABET_MAP = list(string.ascii_lowercase)
UPPER_ALPHABET_MAP = list(string.ascii_uppercase)
# pdb.set_trace()

def find_index(str):
  for i in range(26):
    if LOWER_ALPHABET_MAP[i] == str or UPPER_ALPHABET_MAP[i] == str:
      return i
  return False

for i in range(N):
  # pdb.set_trace()
  if S[i].isupper():
    if find_index(S[i]) == False and S[i] != 'A':
      sys.stdout.write(S[i])
    else:
      index = find_index(S[i]) + K
      if index >= 26: index %= 26
      sys.stdout.write(UPPER_ALPHABET_MAP[index])
  else:
    if find_index(S[i]) == False and S[i] != 'a':
      sys.stdout.write(S[i])
    else:
      index = find_index(S[i]) + K
      if index >= 26: index %= 26
      sys.stdout.write(LOWER_ALPHABET_MAP[index])
