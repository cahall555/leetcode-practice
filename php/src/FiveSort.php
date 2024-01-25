<?php declare(strict_types=1);

class FiveSort {
	public function fivesort(array $nums): array {
		$i = 0;
		$j = count($nums) - 1;

		while ($i <= $j) {
			if ($nums[$j] === 5) {
				$j--;
			} else if ($nums[$i] === 5) {
				$nums[$i] = $nums[$j];
				$nums[$j] = 5;
				$i++;
			} else {
				$i++;
			}
		}
		return $nums;
	}
}

//$fs = new FiveSort();
//print_r($fs->fivesort([1, 5, 3, 5, 5, 2, 5, 1, 5, 5, 5, 5, 4, 5, 5, 5]));
