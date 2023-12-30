// Big o notation: Time: O(n + m) Space: O(n + m)

// problem: given two strings, write a function to determine if the second string is an anagram of the first. An anagram is a word, phrase, or name formed by rearranging the letters of another, such as cinema, formed from iceman.

// stratagy: create a count object adding a count for each character of s1 and subtracting a count for each character of s2. If the count is not 0 for each character in the count object, return false. If the count is 0 for each character in the count object, return true.

const anagrams = (s1, s2) => {

};


module.exports = anagrams;








// Practice from memory before checking solution below:
//const anagrams = (s1, s2) => {
//	const count = {};
//
//	for (let char of s1) {
//		if (!(char in count)) {
//			count[char] = 0;
//		}
//		count[char] += 1;
//	}
//	console.log("s1: ",count);
//	for (let char of s2) {
//		if (count[char] === undefined) {
//			return false;
//		} else {
//			count[char] -= 1;
//		}
//	}
//	console.log("s2: ",count);
//	for (let char in count) {
//		if (count[char] !== 0) {
//			return false;
//		}
//	}
//	console.log("final: ",count);
//	return true;
//};

//anagrams("hello", "olleh");

