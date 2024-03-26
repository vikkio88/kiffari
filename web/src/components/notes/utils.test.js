import { expect, test } from 'vitest';
import { getNoteItemConfig } from './utils';


test('should correctly set showcased info for done button', () => {
    const tests = [
        { due_date: null, body: "Lorem ipsum dolor sit amet.", expectedDoneBtn: false },
        { due_date: new Date(), body: "Quis nostrud exercitation ullamco laboris.", expectedDoneBtn: true }
    ];

    tests.forEach(({ due_date, body, expectedDoneBtn }) => {
        const config = getNoteItemConfig({ body, due_date });
        expect(config.doneBtn).toEqual(expectedDoneBtn);
    });
});