<?php declare(strict_types=1);
require_once __DIR__ . "/../../src/FiveSort.php";
use PHPUnit\Framework\TestCase;

final class FiveSortTest extends TestCase {
	public function testFiveSort(): void {
		$fs = new FiveSort();	
		$this->assertEquals([1, 4, 3, 2, 5], $fs->fiveSort([5, 4, 3, 2, 1]));
	}

	public function testFiveSort2(): void {
		$fs = new FiveSort();	
		$this->assertEquals([4, 1, 2, 1, 2, 3, 5, 5, 5, 5, 5, 5, 5, 5], $fs->fiveSort([5, 1, 2, 5, 5, 3, 2, 5, 1, 5, 5, 5, 4, 5]));
	}

}
