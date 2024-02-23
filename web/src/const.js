export const KATO_API_URL = import.meta.env.VITE_API_URL;
export const LOGIN_TOKEN_KEY = 'token';

// TODO: this could also come from config
export const D_TASK_STATUS = {
    DONE: "done",
    IN_PROGRESS: "in_progress",
    TODO: "todo",
    BACKLOG: "backlog",
};

export const D_TASK_WORKFLOW = {
    [D_TASK_STATUS.DONE]: { from: D_TASK_STATUS.IN_PROGRESS, to: null },
    [D_TASK_STATUS.IN_PROGRESS]: { from: D_TASK_STATUS.TODO, to: D_TASK_STATUS.DONE },
    [D_TASK_STATUS.TODO]: { from: D_TASK_STATUS.BACKLOG, to: D_TASK_STATUS.IN_PROGRESS },
    [D_TASK_STATUS.BACKLOG]: { from: null, to: D_TASK_STATUS.TODO },
};