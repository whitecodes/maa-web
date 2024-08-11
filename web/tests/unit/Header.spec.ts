import { mount } from '@vue/test-utils'
import VersionPane from '@/components/VersionPanel.vue'
import { useMainStore } from '@/stores'
import { describe, it, expect, vi } from 'vitest'
import { router } from '../setup'

// Mock the API call
vi.mock('@/services/api', () => ({
    default: {
        getVersion: vi.fn(() => Promise.resolve({ version: 'v5.4.0-beta.2' })),
        getMaaConnected: vi.fn(() => Promise.resolve({ connected: false }))
    }
}))

describe('Version.vue', () => {
    it('renders version from API', async () => {
        const wrapper = mount(VersionPane)

        // Wait for the API call to resolve
        await new Promise(resolve => setTimeout(resolve, 0))
        await wrapper.vm.$nextTick()

        expect(wrapper.text()).toContain('v5.4.0-beta.2')
    })

    it('renders connected from API', async () => {
        const wrapper = mount(VersionPane)

        // Wait for the API call to resolve
        await new Promise(resolve => setTimeout(resolve, 0))
        await wrapper.vm.$nextTick()

        expect(wrapper.text()).toContain('disconnected')
    })


    it('sets selected tab when clicked', async () => {
        const wrapper = mount(VersionPane)

        // 确保路由已经加载
        await router.isReady();
    
        const store = useMainStore();
    
        // 触发点击事件
        await wrapper.find('a[href="/auto-battle"]').trigger('click');
    
        // 验证 store 的 selectedTab 是否更新
        expect(store.selectedTab).toBe('auto-battle');
    })
})

