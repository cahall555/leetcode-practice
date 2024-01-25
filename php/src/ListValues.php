<?php declare(strict_types=1);
require_once __DIR__ . '/Node.php';

class ListValues {
	
	public function fillValues(?Node $head, array &$values): void {
		if ($head == null) {
			return;
		}
		$values[] = $head->data;
		$this->fillValues($head->next, $values);
	}
	
	public function listvalues(?Node $head): array {
		$values = [];
		$this->fillValues($head, $values);
		return $values;
	}
}


//$a = new Node('a');
//$b = new Node('b');
//$c = new Node('c');
//$d = new Node('d');

//$a->next = $b;
//$b->next = $c;
//$c->next = $d;

//$listValues = new ListValues();
//$val = $listValues->listvalues($a);
//print_r($val);

