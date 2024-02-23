export function groupTasksByStatus(tasks) {
    const grouped = {};
    for (const t of tasks) {
        grouped[t.status] = [...grouped[t.status] ?? [], t];
    }
    return grouped;
}