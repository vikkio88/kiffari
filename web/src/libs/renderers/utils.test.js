import { expect, test } from 'vitest';
import { exportTodos, extractTodos, extractYTVideos, parseLinks } from './extractors';


test('should extract links correctly', () => {
  const input = `
<!--
Plugin: Link
-->
mario
http://google.com

giacomo
https://google.it

lillo https://mamma.com

some stuff http://google.eu/path/to/resource?query=param&foo=bar
fabrizio https://example.com:8080/path/to/resource?query=param&foo=bar
`;

  const expectedOutput = [
    { href: 'http://google.com', label: 'mario' },
    { href: 'https://google.it', label: 'giacomo' },
    { href: 'https://mamma.com', label: 'lillo' },
    { href: 'http://google.eu/path/to/resource?query=param&foo=bar', label: 'some stuff' },
    { href: 'https://example.com:8080/path/to/resource?query=param&foo=bar', label: 'fabrizio' }
  ];

  expect(parseLinks(input)).toEqual(expectedOutput);
});


test("should extract yt videos correctly", () => {
  const input = `
<!--
Plugin: Youtube
-->
blabla
https://www.youtube.com/watch?v=tgbNymZ7vqY
https://youtu.be/tgbNymZ7vqY?si=o2CKaC88Hle4psd
https://youtu.be/tgbNymZ7vqY
http://google.eu/path/to/resource?query=param&foo=bar
xxxx
https://example.com:8080/path/to/resource?query=param&foo=bar
  `;

  const expectedOutput = [
    { original: "https://www.youtube.com/watch?v=tgbNymZ7vqY", embed: "https://www.youtube.com/embed/tgbNymZ7vqY" },
    { original: "https://youtu.be/tgbNymZ7vqY?si=o2CKaC88Hle4psd", embed: "https://www.youtube.com/embed/tgbNymZ7vqY" },
    { original: "https://youtu.be/tgbNymZ7vqY", embed: "https://www.youtube.com/embed/tgbNymZ7vqY" },
  ];

  expect(extractYTVideos(input)).toEqual(expectedOutput);
});

test("should extract todos correctly", () => {
  const input = `<!--
Plugin: todo
-->
- []Maurizio
  -   [ x]    Stuff
  - [x] Things 
-[]other things 
-[ x ] other stuff things 
`;
  const expectedOutput = [
    { label: "Maurizio", done: false },
    { label: "Stuff", done: true },
    { label: "Things", done: true },
    { label: "other things", done: false },
    { label: "other stuff things", done: true },
  ];

  expect(extractTodos(input)).toEqual(expectedOutput);
});

test("should export todos to md", () => {
  const todos = [
    { label: "Maurizio", done: false },
    { label: "Stuff", done: true },
    { label: "Things", done: true },
    { label: "other things", done: false },
    { label: "other stuff things", done: true },
  ];

  const expectedOutput = `<!--
plugin: todo
-->
- [] Maurizio
- [x] Stuff
- [x] Things
- [] other things
- [x] other stuff things
`;

  expect(exportTodos(todos)).toEqual(expectedOutput);
});