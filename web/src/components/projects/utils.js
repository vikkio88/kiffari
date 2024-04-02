import { D_TASK_STATUS as STATUS, } from "../../const";

const defaultAccordionStatus = {
    [STATUS.DONE]: false,
    [STATUS.IN_PROGRESS]: true,
    [STATUS.TODO]: false,
    [STATUS.BACKLOG]: false,
};

export function getAccordionStatus(groupedTasks = {}) {
    const status = {
        ...defaultAccordionStatus
    };
    if (groupedTasks[STATUS.IN_PROGRESS].length > 0) {
        return status;
    }

    status[STATUS.IN_PROGRESS] = false;
    status[STATUS.BACKLOG] = true;
    return status;
}

export function toProgress(groupedTasks = {}) {
    const total = (groupedTasks[STATUS.TODO]?.length ?? 0) + (groupedTasks[STATUS.IN_PROGRESS]?.length ?? 0) + (groupedTasks[STATUS.DONE]?.length ?? 0);
    return { total, value: (groupedTasks[STATUS.DONE]?.length ?? 0) };
}