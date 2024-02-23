<script>
  import Task from "../components/Task.svelte";
  import Accordion from "../components/shared/Accordion.svelte";
  import Controls from "../components/shared/Controls.svelte";
  import { D_TASK_STATUS as STATUS, D_TASK_STATUS_LABELS } from "../const";
  import { groupTasksByStatus } from "../libs/helpers/tasks";

  const defaultAccordion = {
    [STATUS.DONE]: false,
    [STATUS.IN_PROGRESS]: true,
    [STATUS.TODO]: true,
    [STATUS.BACKLOG]: false,
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
      links: [],
      ...override,
    };
  }

  const tasks = [
    m({ status: STATUS.TODO, links: [{ id: "yo" }] }),
    m({ status: STATUS.TODO }),
    m({ status: STATUS.TODO }),
    m({ status: STATUS.DONE }),
    m({ status: STATUS.TODO }),
    m({ status: STATUS.DONE }),
    m({ status: STATUS.DONE }),
    m({ status: STATUS.DONE }),
    m({ status: STATUS.IN_PROGRESS, flag: "Something is wrong!" }),
    m({ status: STATUS.IN_PROGRESS }),
    m({ status: STATUS.IN_PROGRESS }),
    m({ status: STATUS.BACKLOG, flag: "Stuff", description: "Some Stuff" }),
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
  {#each Object.values(STATUS) as status}
    <div class="groupWrapper">
      <!-- WIP LIMIT RULE -->
      <div class="group">
        <Accordion open={defaultAccordion[status]}>
          <h1 slot="header" class="taskGroupHead">
            {D_TASK_STATUS_LABELS[status]}
          </h1>
          <div slot="content">
            {#each groupedTasks[status] as task}
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
    text-align: left;
    padding-left: 3rem;
  }

  .group > div {
    min-width: 60%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
  }

  .board {
    padding-bottom: 2rem;
  }
</style>
