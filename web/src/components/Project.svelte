<script>
    import Task from "./Task.svelte";
    import Accordion from "./shared/Accordion.svelte";
    import Controls from "./shared/Controls.svelte";
    import { D_TASK_STATUS as STATUS, D_TASK_STATUS_LABELS } from "../const";
    import { groupTasksByStatus } from "../libs/helpers/tasks";
    export let project = {};

    const defaultAccordion = {
        [STATUS.DONE]: false,
        [STATUS.IN_PROGRESS]: true,
        [STATUS.TODO]: true,
        [STATUS.BACKLOG]: false,
    };
    let groupedTasks = groupTasksByStatus(
        Boolean(project) ? project.tasks : [],
    );
    function onTaskUpdate({ detail }) {
        const { prevStatus, task } = detail;
        //TODO: Update task status here
        groupedTasks[task.status].unshift(task);
        groupedTasks[prevStatus] = groupedTasks[prevStatus].filter(
            (t) => t.id != task.id,
        );

        groupedTasks = groupedTasks;
    }
</script>

<div class="board">
    <h1>{project.name}</h1>
    <p>{project.description}</p>

    {#if Array.isArray(project.links)}
        <div class="links">
            <strong>🔗 Links</strong>
            {#each project.links as link}
                <div class="link">
                    <a href={link.href} target="_blank">{link.label}</a>
                </div>
            {/each}
        </div>
    {/if}

    <!-- Progress -->
    {#each Object.values(STATUS) as status}
        <div class="groupWrapper">
            <!-- WIP LIMIT RULE -->
            <div class="group">
                <Accordion open={defaultAccordion[status]}>
                    <h1 slot="header" class="taskGroupHead">
                        {`(${groupedTasks[status].length}) ${D_TASK_STATUS_LABELS[status]}`}
                    </h1>
                    <div slot="content">
                        {#each groupedTasks[status] as task}
                            <Task {task} on:updatedTask={onTaskUpdate} />
                        {:else}
                            <h4>No tasks here...</h4>
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
        padding: 1rem 0;
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
