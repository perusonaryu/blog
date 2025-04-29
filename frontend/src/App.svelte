<script lang="ts">
  import { onMount } from 'svelte';
  import axios from 'axios';

  let message = '';
  let dbStatus = '';

  onMount(async () => {
    try {
      const response = await axios.get('/api/');
      message = response.data;
    } catch (error) {
      console.error('APIエラー:', error);
      message = 'API接続エラー';
    }
  });

  async function checkDB() {
    try {
      const response = await axios.get('/api/db-check');
      console.log(response.data);
      dbStatus = response.data;
    } catch (error) {
      console.error('DB接続エラー:', error);
      dbStatus = 'DB接続エラー';
    }
  }
</script>

<main>
  <h1>Svelte + Go アプリケーション</h1>
  <div class="card">
    <p>バックエンドからのメッセージ: {message}</p>
    <button on:click={checkDB}>DB接続確認</button>
    {#if dbStatus}
      <p>DBステータス: {dbStatus}</p>
    {/if}
  </div>

  <div class="bg-blue-500 text-white p-4 rounded-lg">Hello, Tailwind CSS!</div>
</main>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
  }

  .card {
    padding: 2em;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    background-color: #f9f9f9;
  }

  button {
    background-color: #ff3e00;
    color: white;
    border: none;
    padding: 8px 16px;
    border-radius: 4px;
    cursor: pointer;
    margin: 1em 0;
  }

  button:hover {
    background-color: #ff5e20;
  }
</style>
