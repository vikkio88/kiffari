<script>
  import { navigate } from "svelte-routing";
  import Spinner from "../components/shared/Spinner.svelte";

  import { KATO_API_URL, LOGIN_TOKEN_KEY } from "../const";
  import { userToken } from "../store";
  import { onMount } from "svelte";

  let loginPromise = null;
  let value;

  function login() {
    fetch(`${KATO_API_URL}/login`, {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ passkey: value }),
    }).then(resp => resp.json()).then(console.log);
    // $userToken = "user";
    // navigate("/", { replace: true });
  }

  onMount(() => {
    if (Boolean(window.localStorage.getItem(LOGIN_TOKEN_KEY))) {
      navigate("/", { replace: true });
    }
  });
</script>

{#if !loginPromise}
  <input type="text" placeholder="Passkey" bind:value />
  <button on:click={login}>Login</button>
{:else}
  <Spinner />
{/if}
