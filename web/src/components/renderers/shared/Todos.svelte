<script>
    import { createEventDispatcher } from "svelte";
    import { extractTodos } from "../../../libs/renderers/extractors";
    import Adder from "../../shared/Adder.svelte";
    import Md from "svelte-markdown";
    const d = createEventDispatcher();

    export let source = "";
    let todos = extractTodos(source);

    function updateEvent() {
        d("todosUpdated", todos);
    }

    function onChange(index) {
        const wasDone = todos[index].done;
        const updated = todos.splice(index, 1);
        wasDone
            ? (todos = [...todos, ...updated])
            : (todos = [...updated, ...todos]);

        updateEvent();
    }

    function onPush({ detail: label }) {
        todos = [...todos, { label, done: false }];
        updateEvent();
    }

    function onStack({ detail: label }) {
        todos = [{ label, done: false }, ...todos];
        updateEvent();
    }

    function onDelete(index) {
        todos.splice(index, 1);
        todos = todos;
        updateEvent();
    }

    function bringTop(index) {
        todos.unshift(todos.splice(index, 1)[0]);
        todos = todos;
        updateEvent();
    }
</script>

{#if todos.length > 3}
    <Adder on:added={onStack} placeholder="The thing todo..." />
{/if}

<div class="wrapper">
    {#each todos as todo, i}
        <div class="todo">
            <div class="position">
                <button
                    title="Bring to the top"
                    disabled={i == 0 || todo.done}
                    on:click={() => bringTop(i)}
                >
                    ⬆️
                </button>
            </div>
            <input
                type="checkbox"
                bind:checked={todo.done}
                on:change={() => onChange(i)}
            />
            <div class="label" class:done={todo.done}>
                <Md source={todo.label} />
            </div>
            {#if todo.done}
                <button class="del" on:click={() => onDelete(i)}>❌</button>
            {/if}
        </div>
    {:else}
        <div class="frc">
            <h3>No todos... 🤷</h3>
        </div>
    {/each}

    <Adder on:added={onPush} placeholder="The thing todo..." />
</div>

<style>
    .wrapper {
        min-height: 30vh;
    }

    input[type="checkbox"] {
        transform: scale(2);
    }
    .todo {
        font-size: 1.3rem;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: flex-start;
        gap: 1rem;
    }

    .label {
        overflow: auto;
    }

    .position {
        font-size: 0.6rem;
        visibility: hidden;
    }

    .todo:hover .position {
        visibility: visible;
    }

    .done {
        text-decoration: line-through;
    }

    .del {
        margin-left: 1.5rem;
        visibility: hidden;
    }

    .todo:hover .del {
        visibility: visible;
    }
</style>
