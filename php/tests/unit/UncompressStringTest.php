<?php declare(strict_types=1);
require_once __DIR__ . '/../../src/UncompressString.php';
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

    public function testUncompressString2(): void
    {
	    $uncompress = new UncompressString();
	    $this->assertEquals(
	        'abc',
	        $uncompress->uncompress('1a1b1c')
	    );
    }

    public function testUncompressString3(): void
    {
	    $uncompress = new UncompressString();
	    $this->assertEquals(
	        'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
	        $uncompress->uncompress('130a')
	    );
    }
}
