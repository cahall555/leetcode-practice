<?php declare(strict_types=1);
require_once __DIR__ . '/../src/UncompressString.php';
use PHPUnit\Framework\TestCase;

final class UncompressStringTest extends TestCase
{
    public function testUncompressString(): void
    {
	$uncompress = new UncompressString();
	$this->assertEquals(
	    'aabbbbbc',
	    $uncompress->uncompress('2a5b1c')
	);
    }
}
