<div class="spinner" />

<style>
  .spinner {
    border: 3px solid #333; /* Light grey */
    border-top: 3px solid white; /* Blue */
    border-radius: 50%;
    width: 20px;
    height: 20px;
    animation: spin .5s linear infinite;
  }
  @keyframes spin {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }
</style>
