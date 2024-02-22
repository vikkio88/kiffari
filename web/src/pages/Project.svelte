<script>
  import Controls from "../components/shared/Controls.svelte";

  const STATUS = {
    DONE: "done",
    IN_PROGRESS: "in_progress",
    TODO: "todo",
    BACKLOG: "backlog",
  };

  const project = {
    name: "Kiffari",
    links: [
      { label: "Source", href: "https://github.com" },
      { label: "Website", href: "https://github.com" },
    ],
  };

  function m(override = {}) {
    return {
      title: "Do Stuff",
      ...override,
    };
  }

  const tasks = {
    [STATUS.DONE]: [m(), m()],
    [STATUS.IN_PROGRESS]: [m(), m()],
    [STATUS.TODO]: [m(),m(),m(),m(),],
    [STATUS.BACKLOG]: [m(), m(), m(), m(), m()],
  };
</script>

<div class="board">
  <h1>{project.name}</h1>
  <!-- LINKS -->
  <!-- Progress -->
  {#each Object.keys(tasks) as taskGroup}
    <div class="groupWrapper">
      <h1>{taskGroup}</h1>
      <!-- WIP LIMIT RULE -->
      <div class="group">
        {#each tasks[taskGroup] as task}
          <h2>{task.title}</h2>
        {/each}
      </div>
    </div>
  {/each}
</div>

<Controls>
  <button>Add Task</button>
</Controls>

<style>
</style>
