<script>
  import { updateNote } from "../../libs/api";
  import { exportTodos } from "../../libs/renderers/extractors";
  import Todos from "./shared/Todos.svelte";

  export let note = {};
  let { body } = note;

  function persistTodos({ detail: todos }) {
    updateNote(note.id, {
      id: note.id,
      title: note.title,
      tags: note.tags ?? [],
      pinned: note.pinned,
      body: exportTodos(todos),
    });
    //TODO: check result/reconciliate
  }
</script>

<Todos source={body} on:todosUpdated={persistTodos} />
