import { expect, test } from 'vitest';
import { parseLinks } from '../extractors';


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

      http://google.eu/path/to/resource?query=param&foo=bar
      fabrizio https://example.com:8080/path/to/resource?query=param&foo=bar
    `;

    const expectedOutput = [
        { href: 'http://google.com', label: 'mario' },
        { href: 'https://google.it', label: 'giacomo' },
        { href: 'https://mamma.com', label: 'lillo' },
        { href: 'http://google.eu/path/to/resource?query=param&foo=bar', label: 'http://google.eu/path/to/resource?query=param&foo=bar' },
        { href: 'https://example.com:8080/path/to/resource?query=param&foo=bar', label: 'fabrizio' }
    ];

    expect(parseLinks(input)).toEqual(expectedOutput);
});
