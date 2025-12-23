<script>
  import { SendHTTPRequest } from '../wailsjs/go/main/App'
  
  let method = 'GET'
  let url = 'https://jsonplaceholder.typicode.com/todos/1'
  let body = ''
  let headers = ''
  let response = ''
  let statusCode = 0
  let responseTime = ''
  let loading = false
  let error = ''

  const methods = ['GET', 'POST', 'PUT', 'DELETE', 'PATCH']

  async function sendRequest() {
    loading = true
    error = ''
    response = ''
    statusCode = 0
    responseTime = ''

    try {
      // Parse headers from text
      const headerMap = {}
      if (headers.trim()) {
        headers.split('\n').forEach(line => {
          const [key, ...valueParts] = line.split(':')
          if (key && valueParts.length) {
            headerMap[key.trim()] = valueParts.join(':').trim()
          }
        })
      }

      const req = {
        id: crypto.randomUUID(),
        name: '',
        method,
        url,
        headers: headerMap,
        body,
        type: 'http',
        query_vars: {}
      }

      const result = await SendHTTPRequest(req)
      
      if (result.error) {
        error = result.error
      } else {
        statusCode = result.status_code
        response = result.body
        responseTime = `${(result.time / 1000000).toFixed(2)}ms`
      }
    } catch (e) {
      error = e instanceof Error ? e.message : String(e)
    } finally {
      loading = false
    }
  }

  function formatJSON() {
    try {
      const parsed = JSON.parse(response)
      response = JSON.stringify(parsed, null, 2)
    } catch (e) {
      // Not valid JSON, leave as is
    }
  }

  $: if (response && !error) {
    formatJSON()
  }

  function handleKeydown(e) {
    // Cmd+Enter (Mac) or Ctrl+Enter (Windows/Linux) to send request
    if ((e.metaKey || e.ctrlKey) && e.key === 'Enter') {
      e.preventDefault()
      if (url && !loading) {
        sendRequest()
      }
    }
  }
</script>

<svelte:window on:keydown={handleKeydown} />

<main>
  <div class="container">
    <header>
      <h1>🎯 Parley</h1>
      <p>HTTP Client <span class="hint">• Cmd+Enter to send</span></p>
    </header>

    <div class="request-panel">
      <div class="method-url">
        <select bind:value={method} class="method-select">
          {#each methods as m}
            <option value={m}>{m}</option>
          {/each}
        </select>
        <input 
          type="text" 
          bind:value={url} 
          placeholder="Enter request URL"
          class="url-input"
        />
        <button 
          on:click={sendRequest} 
          disabled={loading || !url}
          class="send-button"
        >
          {loading ? '⏳' : '▶'} Send
        </button>
      </div>

      <div class="tabs">
        <div class="tab-headers">
          <button class="tab-header active">Headers</button>
          <button class="tab-header">Body</button>
        </div>
        
        <div class="tab-content">
          <textarea
            bind:value={headers}
            placeholder="Header-Name: value&#10;Another-Header: value"
            rows="4"
          ></textarea>
        </div>
      </div>
    </div>

    <div class="response-panel">
      <div class="response-header">
        <h2>Response</h2>
        {#if statusCode}
          <div class="response-meta">
            <span class="status" class:success={statusCode >= 200 && statusCode < 300} class:error={statusCode >= 400}>
              {statusCode}
            </span>
            <span class="time">{responseTime}</span>
          </div>
        {/if}
      </div>

      {#if error}
        <div class="error-message">
          ❌ {error}
        </div>
      {:else if response}
        <pre class="response-body">{response}</pre>
      {:else if loading}
        <div class="placeholder">Sending request...</div>
      {:else}
        <div class="placeholder">Response will appear here</div>
      {/if}
    </div>
  </div>
</main>

<style>
  :global(body) {
    margin: 0;
    padding: 0;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
    background: #1a1a1a;
    color: #e0e0e0;
  }

  main {
    height: 100vh;
    overflow: hidden;
  }

  .container {
    display: flex;
    flex-direction: column;
    height: 100%;
    max-width: 1400px;
    margin: 0 auto;
    padding: 20px;
  }

  header {
    margin-bottom: 20px;
  }

  h1 {
    margin: 0;
    font-size: 32px;
    font-weight: 600;
  }

  header p {
    margin: 5px 0 0 0;
    color: #888;
    font-size: 14px;
  }

  .hint {
    color: #666;
    font-size: 12px;
  }

  .request-panel {
    background: #242424;
    border-radius: 8px;
    padding: 20px;
    margin-bottom: 20px;
  }

  .method-url {
    display: flex;
    gap: 10px;
    margin-bottom: 20px;
  }

  .method-select {
    padding: 10px 15px;
    background: #333;
    border: 1px solid #444;
    border-radius: 6px;
    color: #e0e0e0;
    font-size: 14px;
    font-weight: 600;
    cursor: pointer;
    min-width: 100px;
  }

  .url-input {
    flex: 1;
    padding: 10px 15px;
    background: #333;
    border: 1px solid #444;
    border-radius: 6px;
    color: #e0e0e0;
    font-size: 14px;
    font-family: 'SF Mono', Monaco, monospace;
  }

  .url-input:focus {
    outline: none;
    border-color: #4a9eff;
  }

  .send-button {
    padding: 10px 30px;
    background: #4a9eff;
    border: none;
    border-radius: 6px;
    color: white;
    font-size: 14px;
    font-weight: 600;
    cursor: pointer;
    transition: background 0.2s;
  }

  .send-button:hover:not(:disabled) {
    background: #357abd;
  }

  .send-button:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .tabs {
    border-top: 1px solid #333;
    padding-top: 15px;
  }

  .tab-headers {
    display: flex;
    gap: 10px;
    margin-bottom: 15px;
  }

  .tab-header {
    padding: 8px 16px;
    background: transparent;
    border: none;
    color: #888;
    font-size: 14px;
    cursor: pointer;
    border-bottom: 2px solid transparent;
  }

  .tab-header.active {
    color: #4a9eff;
    border-bottom-color: #4a9eff;
  }

  textarea {
    width: 100%;
    padding: 12px;
    background: #1a1a1a;
    border: 1px solid #333;
    border-radius: 6px;
    color: #e0e0e0;
    font-size: 13px;
    font-family: 'SF Mono', Monaco, monospace;
    resize: vertical;
  }

  textarea:focus {
    outline: none;
    border-color: #4a9eff;
  }

  .response-panel {
    flex: 1;
    background: #242424;
    border-radius: 8px;
    padding: 20px;
    display: flex;
    flex-direction: column;
    overflow: hidden;
  }

  .response-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 15px;
  }

  h2 {
    margin: 0;
    font-size: 18px;
    font-weight: 600;
  }

  .response-meta {
    display: flex;
    gap: 15px;
    align-items: center;
  }

  .status {
    padding: 4px 12px;
    border-radius: 4px;
    font-size: 13px;
    font-weight: 600;
  }

  .status.success {
    background: #1a4d2e;
    color: #4ade80;
  }

  .status.error {
    background: #4d1a1a;
    color: #ef4444;
  }

  .time {
    color: #888;
    font-size: 13px;
    font-family: 'SF Mono', Monaco, monospace;
  }

  .response-body {
    flex: 1;
    margin: 0;
    padding: 15px;
    background: #1a1a1a;
    border: 1px solid #333;
    border-radius: 6px;
    overflow: auto;
    font-size: 13px;
    font-family: 'SF Mono', Monaco, monospace;
    line-height: 1.6;
    color: #e0e0e0;
  }

  .placeholder {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #666;
    font-size: 14px;
  }

  .error-message {
    padding: 15px;
    background: #4d1a1a;
    border: 1px solid #ef4444;
    border-radius: 6px;
    color: #ef4444;
    font-size: 14px;
  }
</style>
