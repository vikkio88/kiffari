<script>
  import TagsList from "../components/TagsList.svelte";
  import Controls from "../components/shared/Controls.svelte";
  import { navigate } from "svelte-routing";
  import ConfirmButton from "../components/shared/ConfirmButton.svelte";
  import BodyRenderer from "../components/BodyRenderer.svelte";
  import { protectedRoute } from "../libs/routes";
  import {
    getNoteDetails,
    del,
    catchLogout,
    archive,
    unarchive,
  } from "../libs/api";
  import DateSeverity from "../components/shared/DateSeverity.svelte";
  import Header from "../components/shared/Header.svelte";
  import { getDate } from "../libs/dates";
  import DashedHead from "../components/shared/DashedHead.svelte";
  protectedRoute();

  export let id = "";
  let notePromise = getNoteDetails(id);

  async function onDelete() {
    const resp = await del("notes", id);
    if (resp.status === 401) {
      catchLogout();
    }

    if (resp.ok) {
      navigate("/kato", { replace: true });
    }
  }

  async function onArchiveToggle(isArchived) {
    const resp = await (isArchived
      ? unarchive("notes", id)
      : archive("notes", id));
    if (resp.ok) {
      window.location.reload();
    }

    //TODO: add catchlogout
  }

  function onDownload(title, body, tags) {
    let tagsString = "";
    for (const tag of tags) {
      tagsString += ` #${tag.label}`;
    }
    const element = document.createElement("a");
    element.setAttribute(
      "href",
      "data:text/markdown;charset=utf-8," +
        encodeURIComponent(`# ${title}\n${body}\n\n${tagsString.trim()}`),
    );
    element.download = `${title}.md`;
    element.style.display = "none";
    document.body.appendChild(element);
    element.click();
    document.body.removeChild(element);
  }
</script>

{#await notePromise then note}
  <div class="note">
    {#if note.archived}
      <DashedHead>Archived</DashedHead>
    {/if}
    <h2>{note.title}</h2>
    <div class="dates">
      {getDate(note)}
      <DateSeverity date={note.due_date} />
    </div>
    <div class="body">
      <BodyRenderer body={note.body} />
    </div>
    {#if Array.isArray(note.tags) && note.tags.length > 0}
      <div class="tags">
        <Header text="Tags" />
        <div class="tagList">
          <TagsList tags={note.tags} />
        </div>
      </div>
    {/if}
  </div>
  <Controls background>
    <button
      title="Export as MD"
      on:click={() => onDownload(note.title, note.body, note.tags)}
    >
      🔽
    </button>
    {#if !note.archived}
      <button title="Edit" on:click={() => navigate(`/edit-note/${id}`)}>
        📝
      </button>
    {/if}
    <ConfirmButton
      title={note.archived ? "Un-Archive" : "Archive"}
      confirmLabel={note.archived ? "Un-Archive?" : "Archive?"}
      onConfirmed={() => onArchiveToggle(note.archived)}
    >
      {#if !note.archived}
        🗄️
      {:else}
        🔄
      {/if}
    </ConfirmButton>
    <ConfirmButton title="Delete" confirmLabel="Delete?" onConfirmed={onDelete}>
      🗑️
    </ConfirmButton>
  </Controls>
{/await}

<style>
  .note {
    border-radius: 10px;
    padding: 0 2rem;
  }

  .note h2 {
    margin: 1rem 0 0 0;
  }

  .note div.dates {
    font-size: smaller;
    color: #a3a3a3;
  }

  .body {
    border-radius: 10px;
    text-align: left;
    padding: 1.5rem;
    margin: 0.5rem 0;
  }

  .tagList {
    display: flex;
    flex-direction: row;
    padding-bottom: 1.5rem;
    flex-wrap: wrap;
    gap: 0.2rem;
  }
</style>
