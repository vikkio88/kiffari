<script>
    import { updateNote } from "../../libs/api";
    import { exportTodos, extractTodos } from "../../libs/renderers/extractors";

    export let note = {};
    let { body } = note;
    let todos = extractTodos(body);

    function onChange() {
        updateNote(note.id, {
            id: note.id,
            title: note.title,
            tags: note.tags ?? [],
            body: exportTodos(todos),
        });
        //TODO: check result/reconciliate
    }

    //TODO: add manual TODO
    //TODO: add button to create todo on note creation
</script>

{#each todos as todo}
    <div class="todo">
        -
        <input type="checkbox" bind:checked={todo.done} on:change={onChange} />
        {todo.label}
    </div>
{:else}
    <div class="fcr">No Todos</div>
{/each}

<style>
    .todo {
        font-size: 1.3rem;
        padding: 0.5rem;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: flex-start;
        gap: 1rem;
    }
</style>
