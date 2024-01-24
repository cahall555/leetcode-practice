<?php declare(strict_types=1);

final class UncompressString {

public function uncompress (string $str): string {
 $NUMBERS = "0123456789";
 $result = [];
 $i = 0;
 $j = 0;

 while ($j < strlen($str)) {
 	if (strpos($NUMBERS, $str[$j]) !== false) {
		$j++;
 	} else {
 		$number = (int)substr($str, $i, $j - $i);
		$letter = $str[$j];
 		$result[] = str_repeat($letter, $number);
 		$i = $j + 1;
 		$j = $i;
	}
 }
 return implode("", $result);
 
}
}

$uncompress = new UncompressString();
print_r($uncompress->uncompress("2a5b1c"));
