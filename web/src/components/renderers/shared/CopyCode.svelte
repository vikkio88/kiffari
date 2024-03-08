<script>
  import { hasCodeSnippets } from "../../../libs/renderers/helpers";
  export let body = "";
  let hasCode = hasCodeSnippets(body);

  function copyText(e) {
    navigator.clipboard.writeText(e.target.innerText);
    const span = document.createElement("span");
    span.classList.add("tip");
    span.innerText = "📋 Copied!";
    e.target.parentElement.appendChild(span);
    setTimeout(() => {
      e.target.parentElement.removeChild(span);
    }, 800);
  }

  function handleCopyPaste(e) {
    const { checked } = e.currentTarget;
    const snippets = document.querySelectorAll("code");
    if (checked) {
      snippets.forEach((s) => {
        s.parentElement.classList.add("p-r");
        s.parentElement.classList.add("crs-pointer");
        s.addEventListener("click", copyText);
      });
    } else {
      snippets.forEach((s) => {
        s.parentElement.classList.remove("p-r");
        s.parentElement.classList.remove("crs-pointer");
        s.removeEventListener("click", copyText);
      });
    }
  }
</script>

<div class="hidden" class:hasCode>
  <label for="copyPaste">Copy <i>Code</i> 📋</label>
  <input name="copyPaste" type="checkbox" on:change={handleCopyPaste} />
</div>

<style>
  .hasCode {
    visibility: visible;
  }
</style>
