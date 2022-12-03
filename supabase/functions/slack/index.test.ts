import {
    assertEquals,
} from "https://deno.land/std@0.65.0/testing/asserts.ts";
import { extractFirstUrlFromUrlsConcatByPipe } from "./index.ts"

// extractFirstUrlFromUrlsConcatByPipe

Deno.test("URLに|が含まれていない場合は引数をそのまま返す", () => {
  const actual = extractFirstUrlFromUrlsConcatByPipe('http://example.com')
  
  assertEquals(actual, 'http://example.com');
});

Deno.test("URLに|が含まれている場合は、|で区切った最初の文字列を返す", () => {
    const actual = extractFirstUrlFromUrlsConcatByPipe('http://example.com|http://example.com')
    
    assertEquals(actual, 'http://example.com');
});

Deno.test("URLが空文字の場合は空文字を返す", () => {
    const actual = extractFirstUrlFromUrlsConcatByPipe('')
    
    assertEquals(actual, '');
});
