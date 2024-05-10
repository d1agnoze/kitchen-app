<script setup lang="ts">
  import {ref, onMounted} from 'vue'
  const message =ref('goos')
  async function getData(){
    try {
      const res = await fetch('http://localhost:8080/ping', { method: 'GET' });
      if (!res.ok) { 
        console.log(res)
        throw new Error(`HTTP error stat: ${res.status}`); }

      const data = await res.json();
      message.value = data.message; // Make sure the data structure matches what you expect
    } catch (error) {
      console.error('Error fetching data:', error);
      message.value = 'Error'; // Make sure the data structure matches what you expect
    }
  }
</script>
<template>
  <div>
    <h1>Message from beyond: {{ message }}</h1>
    <button @click="getData">Get</button>
  </div>
</template>
