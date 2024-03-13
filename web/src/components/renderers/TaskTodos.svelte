<script>
  import { updateTask } from "../../libs/api";
  import { exportTodos } from "../../libs/renderers/extractors";
  import Todos from "./shared/Todos.svelte";

  export let task = {};
  let { description } = task;

  function persistTodos({ detail: todos }) {
    updateTask(task.project_id, {
      id: task.id,
      title: task.title,
      tags: task.tags ?? [],
      description: exportTodos(todos),
      status: task.status,
      category: task.category,
    });
    //TODO: check result/reconciliate
  }
</script>

<Todos source={description} on:todosUpdated={persistTodos} />
