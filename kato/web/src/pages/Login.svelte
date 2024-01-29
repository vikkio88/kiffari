<script>
  import { navigate } from "svelte-routing";
  import Spinner from "../components/shared/Spinner.svelte";

  import { LOGIN_TOKEN_KEY } from "../const";
  import { userToken } from "../store";
  import { onMount } from "svelte";

  let loginPromise = null;

  function login() {
    window.localStorage.setItem(LOGIN_TOKEN_KEY, "user");
    $userToken = "user";
    navigate("/", { replace: true });
  }

  onMount(() => {
    if (Boolean(window.localStorage.getItem(LOGIN_TOKEN_KEY))) {
      navigate("/", { replace: true });
    }
  });
</script>

{#if !loginPromise}
  <button on:click={login}>Login</button>
{:else}
  <Spinner />
{/if}
