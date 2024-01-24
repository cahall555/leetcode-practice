const anagrams = require('./anagrams');

test('anagrams function exists', () => {
  expect(typeof anagrams).toEqual('function');
});

test('"hello" is an anagram of "llohe"', () => {
  expect(anagrams('hello', 'llohe')).toBeTruthy();
});

test('"Hey there!" is not an anagram of "How are you!"', () => {
  expect(anagrams('Hey there', 'How are you!')).not.toBeTruthy();
});
