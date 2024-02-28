import { D_TASK_STATUS as STATUS } from "../../const";

export function groupTasksByStatus(tasks) {
    const grouped = {
        [STATUS.DONE]: [],
        [STATUS.IN_PROGRESS]: [],
        [STATUS.TODO]: [],
        [STATUS.BACKLOG]: [],
    };
    for (const t of tasks) {
        grouped[t.status] = [...grouped[t.status] ?? [], t];
    }
    return grouped;
}