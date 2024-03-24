import { MARKDOWN_PLUGINS as PLUGINS } from "../../const";
import { mdConfigParser } from "../../libs/renderers/mdConfigParser";

const defaultNoteItemConfig = {
    info: {
        preview: true,
        dates: true,
        viewBtn: true,
        editBtn: true,
        doneBtn: false,
    },
    icon: null,
    title: null,
};

const pluginConfigMapper = {
    [PLUGINS.CODE]: defaultNoteItemConfig,
    [PLUGINS.LINK]: {
        info: {
            preview: false,
            dates: true,
            viewBtn: true,
            editBtn: false,
            doneBtn: false,

        },
        icon: "🔗",
        title: "Links",
    },
    [PLUGINS.YOUTUBE]: {
        info: {
            preview: false,
            dates: true,
            viewBtn: true,
            editBtn: false,
            doneBtn: false,
        },
        icon: "📺",
        title: "Youtube Video",
    },
    [PLUGINS.TODO]: {
        info: {
            preview: false,
            dates: true,
            viewBtn: true,
            editBtn: false,
            doneBtn: false,
        },
        icon: "✅",
        title: "Todos",
    },

    [PLUGINS.NONE]: defaultNoteItemConfig,
};

export function getNoteItemConfig({ body = "", due_date = null } = {}) {
    const mdConfig = mdConfigParser(body);

    const base = { ...(pluginConfigMapper[mdConfig.plugin] ?? pluginConfigMapper[PLUGINS.NONE]) };

    if (Boolean(due_date)) {
        base.info.editBtn = false;
        base.info.doneBtn = true;
    }

    return base;
}