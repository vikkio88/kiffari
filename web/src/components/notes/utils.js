import { MARKDOWN_PLUGINS as PLUGINS } from "../../const";
import { mdConfigParser } from "../../libs/renderers/mdConfigParser";

const defaultNoteItemConfig = {
    preview: true,
    dates: true,
    viewBtn: true,
    editBtn: true,
    doneBtn: false,
    icon: null,
    title: null,
};

const pluginConfigMapper = {
    [PLUGINS.CODE]: { ...defaultNoteItemConfig },
    [PLUGINS.LINK]: {
        preview: false,
        dates: true,
        viewBtn: true,
        editBtn: false,
        doneBtn: false,
        icon: "🔗",
        title: "Links",
    },
    [PLUGINS.YOUTUBE]: {
        preview: false,
        dates: true,
        viewBtn: true,
        editBtn: false,
        doneBtn: false,
        icon: "📺",
        title: "Youtube Video",
    },
    [PLUGINS.TODO]: {
        preview: false,
        dates: true,
        viewBtn: true,
        editBtn: false,
        doneBtn: false,
        icon: "✅",
        title: "Todos",
    },
    [PLUGINS.NONE]: { ...defaultNoteItemConfig },
};

export function getNoteItemConfig({ body = "", due_date = null } = {}) {
    const mdConfig = mdConfigParser(body);
    const base = { ...pluginConfigMapper[mdConfig.plugin] ?? pluginConfigMapper[PLUGINS.NONE] };

    if (Boolean(due_date)) {
        base.editBtn = false;
        base.doneBtn = true;
    }

    return base;
}
