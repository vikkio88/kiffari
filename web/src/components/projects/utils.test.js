import { expect, test } from 'vitest';
import { toProgress } from './utils';
import { D_TASK_STATUS as S } from '../../const';

test("grouped tasks to progress", () => {
    expect(toProgress()).toEqual({ total: 0, value: 0 });
    expect(toProgress({ [S.DONE]: [...Array(10).keys()] })).toEqual({ total: 10, value: 10 });
    expect(toProgress({
        [S.DONE]: [...Array(1).keys()], [S.TODO]: [...Array(8).keys()], [S.IN_PROGRESS]: [...Array(1).keys()]
    })).toEqual({ total: 10, value: 1 });
});