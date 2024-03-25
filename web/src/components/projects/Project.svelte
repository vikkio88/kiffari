<script>
  import Task from "../task/Task.svelte";
  import Accordion from "../shared/Accordion.svelte";
  import Controls from "../shared/Controls.svelte";
  import {
    D_TASK_STATUS as STATUS,
    D_TASK_STATUS_LABELS,
    D_TASK_CATEGORIES as CATEGORIES,
  } from "../../const";
  import { groupTasksByStatus } from "../../libs/helpers/tasks";
  import Adder from "../shared/Adder.svelte";
  import {
    addTask,
    catchLogout,
    del,
    moveTask,
    updateProject,
  } from "../../libs/api";
  import { navigate } from "svelte-routing";
  import ProjectEditor from "./ProjectEditor.svelte";
  import ConfirmButton from "../shared/ConfirmButton.svelte";
  import Progress from "../shared/Progress.svelte";
  import { getAccordionStatus } from "./utils";
  export let project = {};

  let groupedTasks = groupTasksByStatus(Boolean(project) ? project.tasks : []);

  let accordionsState = getAccordionStatus(groupedTasks);
  let editingProject = false;

  function onTaskUpdate({ detail }) {
    const { prevStatus, task } = detail;
    groupedTasks[task.status].unshift(task);
    groupedTasks[prevStatus] = groupedTasks[prevStatus].filter(
      (t) => t.id != task.id
    );

    accordionsState[prevStatus] = false;
    accordionsState[task.status] = true;
    groupedTasks = groupedTasks;

    //TODO: add fallback if it fails
    moveTask(project.id, task.id, task.status);
  }

  function onAdd(task, status) {
    task = { ...task, status, category: CATEGORIES.FEATURE, tags: [] };
    groupedTasks[task.status].unshift(task);
    groupedTasks = groupedTasks;
    addTask(project.id, task)
      .then((resp) => {
        if (resp.ok) {
          return resp.json();
        }

        if (resp.status === 401) {
          catchLogout();
          return;
        }

        if (resp.status === 400) {
          //TODO: notify error
          groupedTasks[status] = groupedTasks[status].filter((t) =>
            Boolean(t.id)
          );
          groupedTasks = groupedTasks;
          return null;
        }

        return null;
      })
      .then((newTask) => {
        if (!Boolean(newTask)) return;

        groupedTasks[status] = groupedTasks[status].filter((t) =>
          Boolean(t.id)
        );
        groupedTasks[status].unshift(newTask);

        groupedTasks = groupedTasks;
      });
  }

  async function onProjectSave(newProject) {
    editingProject = false;
    const updatedProject = await updateProject(newProject.id, newProject);
    project = { ...project, ...updatedProject };

    //TODO: handle errors in case?
  }

  async function onDelete() {
    const res = await del("projects", project.id);
    if (res.ok) {
      navigate("/kiffari");
      return;
    }

    //TODO: handle error
  }
</script>

<div class="board">
  <div class="details">
    {#if !editingProject}
      <h2>{project.name}</h2>
      <p>{project.description}</p>

      {#if Array.isArray(project.links) && project.links.length > 0}
        <div class="links">
          <strong>🔗 Links</strong>
          {#each project.links as link}
            <div class="link">
              <a href={link.href} target="_blank">{link.label}</a>
            </div>
          {/each}
        </div>
      {/if}
    {:else}
      <ProjectEditor
        id={project.id}
        name={project.name}
        description={project.description}
        links={project.links}
        onSave={onProjectSave}
        onClose={() => (editingProject = false)}
      />
    {/if}

    <div class="projectControls">
      {#if !editingProject}
        <div>
          <ConfirmButton
            title="Delete"
            confirmLabel="Delete?"
            onConfirmed={onDelete}
          >
            🗑️
          </ConfirmButton>
          <button
            title="Archived Tasks"
            on:click={() => navigate(`/projects/${project.id}/archived-tasks`)}
          >
            Tasks 🗄️
          </button>
        </div>
        <button title="Edit" on:click={() => (editingProject = true)}>
          📝
        </button>
      {/if}
    </div>
  </div>
  {#each Object.values(STATUS) as status}
    <div class="groupWrapper">
      <!-- WIP LIMIT RULE -->
      <div class="group">
        <Accordion open={accordionsState[status]}>
          <h3 class="header" slot="header">
            {`(${groupedTasks[status].length}) ${D_TASK_STATUS_LABELS[status]}`}
          </h3>
          <div class="taskList" slot="content">
            {#each groupedTasks[status] as task}
              <Task {task} on:updatedTask={onTaskUpdate} />
            {:else}
              <h4>No tasks here...</h4>
            {/each}
            <Adder
              placeholder="New Task title..."
              on:added={({ detail: title }) => onAdd({ title }, status)}
              centered
            />
          </div>
        </Accordion>
      </div>
    </div>
  {/each}
</div>

<Controls>
  <button
    class="big-control"
    on:click={() => navigate(`/projects/${project.id}/create-task`)}
  >
    Add Task 🎫
  </button>
</Controls>

<style>
  .details {
    border-radius: var(--default-border-radius);
    padding: 2rem 0 0 0;
  }
  .projectControls {
    visibility: hidden;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }
  .details:hover {
    background-color: var(--div-item-hover-color);
  }

  .details:hover .projectControls {
    visibility: visible;
  }
  h2 {
    margin: 0;
  }

  p {
    margin: 0;
    padding: 1rem 0;
  }

  div.links {
    padding: 1rem 0;
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

  .header {
    margin: 0;
  }

  .taskList {
    margin-bottom: 1rem;
  }
</style>
