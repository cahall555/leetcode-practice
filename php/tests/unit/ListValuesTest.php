<?php declare(strict_types=1);

require_once __DIR__ . '/../../src/ListValues.php';
use PHPUnit\Framework\TestCase;

final class ListValuesTest extends TestCase
{
    public function testListValues(): void
    {
	$a = new Node('a');
	$b = new Node('b');
	$c = new Node('c');
	$d = new Node('d');

	$a->next = $b;
	$b->next = $c;
	$c->next = $d;


	$listValues = new ListValues();

	$this->assertEquals(
	    ['a', 'b', 'c', 'd'],
	    $listValues->listvalues($a)
	);
    }

    public function testListValues2(): void {
	    $a = new Node('a');

	    $listValues = new ListValues();
	    $this->assertEquals(
		    ['a'],
		    $listValues->listvalues($a)
	    );
    }

    public function testListValues3(): void {
	    
	    $listValues = new ListValues();
	    $this->assertEquals(
		    [],
		    $listValues->listvalues(null)
	    );
    }
}


