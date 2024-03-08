<script>
  import { formatYHM } from "../../libs/dates";

  export let date = null;
  export let compareDate = new Date();
  function severityClass(dueDate) {
    const date = new Date(dueDate);

    const diff = (date.getTime() - compareDate.getTime()) / 1000;

    if (diff <= 24 * 3600) {
      return "due-now";
    }

    if (diff <= 7 * 24 * 3600) {
      return "due-soon";
    }

    return "due-later";
  }
</script>

{#if date}
  <div
    class={`due ${severityClass(date)}`}
    title={`due on: ${new Date(date).toLocaleString()}`}
  >
    {`${formatYHM(date)} ⏰`}
  </div>
{/if}

<style>
  .due-now {
    color: var(--danger-color);
  }

  .due-soon {
    color: var(--warning-color);
  }

  .due-later {
    color: var(--success-color);
  }
</style>
