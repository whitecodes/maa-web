import { createPinia, defineStore } from 'pinia'

export const useMainStore = defineStore('main', {
  state: () => ({
    selectedTab: 'one-click-grass' as string
  }),
  actions: {
    setSelectedTab(tab: string) {
      this.selectedTab = tab;
    }
  }
})

const pinia = createPinia()
export default pinia
