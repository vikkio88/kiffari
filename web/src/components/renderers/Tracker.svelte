<script>
  import { updateNote } from "../../libs/api";
  import {
    exportTrackers,
    extractTrackers,
  } from "../../libs/renderers/extractors";

  export let note = {};
  const { body } = note;
  let trackers = extractTrackers(body);

  function onTrackingChange(i, event) {
    const { value } = event.target;
    trackers[i].range.value = value;

    trackers = trackers;
    updateNote(note.id, {
      id: note.id,
      title: note.title,
      tags: note.tags ?? [],
      pinned: note.pinned,
      body: exportTrackers(trackers),
    });
  }
</script>

<div class="trackerBody">
    <h3 title="Tracker">⏲️</h3>
  {#each trackers as t, i}
    <h3>
      {t.label ||
        `Tracker ${i + 1}`}{`${t.range.max == t.range.value ? " ✅" : ""}`}
    </h3>
    <h2>
      {t.range.value} / {t.range.max}
    </h2>
    <input
      type="range"
      min={t.range.min}
      max={t.range.max}
      step={t.range.step}
      value={t.range.value}
      on:change={(e) => onTrackingChange(i, e)}
    />
  {/each}
</div>

<style>
  .trackerBody {
    min-height: 30vh;
    text-align: center;
  }

  input[type="range"] {
    min-width: 60vw;
  }
</style>
