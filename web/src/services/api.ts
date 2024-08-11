const baseURL = `${window.location.protocol}//${window.location.hostname}:${window.location.port}`

export default {
  async getVersion() {
    const response = await fetch(`${baseURL}/version`, {
      headers: {
        'Content-Type': 'application/json'
      }
    })
    if (!response.ok) {
      throw new Error('Network response was not ok')
    }
    return response.json()
  },

  async getMaaConnected() {
    const response = await fetch(`${baseURL}/maa/connected`, {
      headers: {
        'Content-Type': 'application/json'
      }
    })
    if (!response.ok) {
      throw new Error('Network response was not ok')
    }
    return response.json()
  }
}
