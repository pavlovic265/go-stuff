<script>
  import { onMount } from 'svelte';
  let time = ''

  onMount(() => {
    const evtSrc = new EventSource("http://localhost:8090/event");
    evtSrc.onmessage = function(event) {
      time = event.data; 
    }
    evtSrc.onerror = function(event) {
      console.log('Error :>> event :>> ', event);
    }
  })

  async function getTime() {
    const res = await fetch("http://localhost:8090/time")
    if (res.status !== 200) {
      console.log('Could not connect to server :>> ');
    }
  }
</script>
<main>
  <h1>Server sent events</h1>
  <button on:click={getTime}>Get Time</button>
  <p>Time: { time }</p>
</main>
