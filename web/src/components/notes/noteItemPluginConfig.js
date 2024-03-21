import { MARKDOWN_PLUGINS as PLUGINS } from "../../const";
import { mdConfigParser } from "../../libs/renderers/mdConfigParser";

const defaultNoteItemConfig = {
    info: {
        preview: true,
        dates: true,
        viewBtn: true,
        editBtn: true,
    },
    icon: null,
};

const pluginConfigMapper = {
    [PLUGINS.CODE]: defaultNoteItemConfig,
    [PLUGINS.LINK]: {
        info: {
            preview: false,
            dates: true,
            viewBtn: true,
            editBtn: false,
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
        },
        icon: "✅",
        title: "Todos",
    },

    [PLUGINS.NONE]: defaultNoteItemConfig,
};

export function getConfigFromPlugin(body) {
    const mdConfig = mdConfigParser(body);

    return pluginConfigMapper[mdConfig.plugin] ?? pluginConfigMapper[PLUGINS.NONE];
}