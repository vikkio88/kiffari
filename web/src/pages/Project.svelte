<script>
  import Task from "../components/Task.svelte";
  import Accordion from "../components/shared/Accordion.svelte";
  import Controls from "../components/shared/Controls.svelte";
  import { D_TASK_STATUS as STATUS } from "../const";
  import { groupTasksByStatus } from "../libs/helpers/tasks";

  const taskGroupLabels = {
    [STATUS.DONE]: "Done ✅",
    [STATUS.IN_PROGRESS]: "In Progress 🔨",
    [STATUS.TODO]: "To Do 📋",
    [STATUS.BACKLOG]: "Backlog 🗃️",
  };

  const project = {
    name: "Kiffari",
    description: "A Kanban-ish approach to managing side projects.",
    links: [
      { label: "Source", href: "https://github.com" },
      { label: "Website", href: "https://github.com" },
    ],
  };

  function m(override = {}) {
    return {
      title: "Do Stuff",
      description: null,
      status: STATUS.BACKLOG,
      flag: null,
      ...override,
    };
  }

  const tasks = [
    m({ status: STATUS.TODO }),
    m({ status: STATUS.TODO }),
    m({ status: STATUS.TODO }),
    m({ status: STATUS.DONE }),
    m({ status: STATUS.TODO }),
    m({ status: STATUS.DONE }),
    m({ status: STATUS.DONE }),
    m({ status: STATUS.DONE }),
    m({ status: STATUS.IN_PROGRESS }),
    m({ status: STATUS.IN_PROGRESS }),
    m({ status: STATUS.IN_PROGRESS }),
    m({ status: STATUS.BACKLOG }),
    m({ status: STATUS.BACKLOG }),
    m({ status: STATUS.BACKLOG }),
    m({ status: STATUS.BACKLOG }),
    m({ status: STATUS.BACKLOG }),
  ];

  const groupedTasks = groupTasksByStatus(tasks);
</script>

<div class="board">
  <h1>{project.name}</h1>
  <p>{project.description}</p>
  <div class="links">
    <strong>🔗 Links</strong>
    {#each project.links as link}
      <div class="link">
        <a href={link.href} target="_blank">{link.label}</a>
      </div>
    {:else}
      <h3>No Links</h3>
    {/each}
  </div>
  <!-- Progress -->
  {#each Object.values(STATUS) as taskGroup}
    <div class="groupWrapper">
      <!-- WIP LIMIT RULE -->
      <div class="group">
        <Accordion>
          <h1 slot="header" class="taskGroupHead">{taskGroupLabels[taskGroup]}</h1>
          <div slot="content">
            {#each groupedTasks[taskGroup] as task}
              <Task {task} />
            {/each}
          </div>
        </Accordion>
      </div>
    </div>
  {/each}
</div>

<Controls>
  <button class="big-control">Add Task 🎫</button>
</Controls>

<style>
  h1 {
    margin: 0;
  }

  p {
    margin: 0;
  }

  div.links {
    padding: 1rem 0;
  }

  .taskGroupHead {
    font-size: 1.8rem;
  }


  .group > div {
    min-width: 60%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
  }
</style>
