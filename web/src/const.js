export const KATO_API_URL = import.meta.env.VITE_API_URL;
export const LOGIN_TOKEN_KEY = 'token';

// TODO: this could also come from config
export const D_TASK_STATUS = {
    DONE: "done",
    IN_PROGRESS: "in_progress",
    TODO: "todo",
    BACKLOG: "backlog",
};

export const D_TASK_STATUS_LABELS = {
    [D_TASK_STATUS.DONE]: "Done ✅",
    [D_TASK_STATUS.IN_PROGRESS]: "In Progress 🔨",
    [D_TASK_STATUS.TODO]: "To Do 📋",
    [D_TASK_STATUS.BACKLOG]: "Backlog 🗃️",
};

export const D_TASK_STATUS_EMOJIS = {
    [D_TASK_STATUS.DONE]: "✅",
    [D_TASK_STATUS.IN_PROGRESS]: "🔨",
    [D_TASK_STATUS.TODO]: "📋",
    [D_TASK_STATUS.BACKLOG]: "🗃️",
};

export const D_TASK_WORKFLOW = {
    [D_TASK_STATUS.DONE]: { from: D_TASK_STATUS.IN_PROGRESS, to: null },
    [D_TASK_STATUS.IN_PROGRESS]: { from: D_TASK_STATUS.TODO, to: D_TASK_STATUS.DONE },
    [D_TASK_STATUS.TODO]: { from: D_TASK_STATUS.BACKLOG, to: D_TASK_STATUS.IN_PROGRESS },
    [D_TASK_STATUS.BACKLOG]: { from: null, to: D_TASK_STATUS.TODO },
};

export const D_TASK_CATEGORIES = {
    FEATURE: "feature",
    BUG: "bug",
    DOC: "doc",
    SPIKE: "spike",
    CLEANUP: "cleanup",
};

export const D_TASK_CATEGORY_LABELS = {
    [D_TASK_CATEGORIES.FEATURE]: "💡",
    [D_TASK_CATEGORIES.BUG]: "🐞",
    [D_TASK_CATEGORIES.DOC]: "📄",
    [D_TASK_CATEGORIES.SPIKE]: "⚡",
    [D_TASK_CATEGORIES.CLEANUP]: "🧹",
};


export const MARKDOWN_PLUGINS = {
    CODE: "code",
    LINK: "link",
    YOUTUBE: "youtube",
    TODO: "todo",

    NONE: "none"
};