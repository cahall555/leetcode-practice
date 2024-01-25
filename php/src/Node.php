<?php declare(strict_types=1);

class Node
{
    public $data;
    public $next;

    public function __construct($data)
    {
	$this->data = $data;
	$this->next = null;
    }
}
