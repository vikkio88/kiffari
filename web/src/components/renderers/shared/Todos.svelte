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

    function onAdd({ detail: label }) {
        todos = [...todos, { label, done: false }];
        updateEvent();
    }

    function onDelete(index) {
        todos.splice(index, 1);
        todos = todos;
        updateEvent();
    }
</script>

<div class="wrapper">
    {#each todos as todo, i}
        <div class="todo">
            -
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

    <Adder on:added={onAdd} placeholder="The thing todo..." />
</div>

<style>
    .wrapper {
        min-height: 30vh;
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

    .done {
        text-decoration: line-through;
    }
</style>
