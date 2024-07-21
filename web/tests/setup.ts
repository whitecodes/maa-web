import { expect } from 'vitest'
import * as matchers from '@testing-library/jest-dom/matchers'
import { config } from '@vue/test-utils'
import { createRouter, createWebHistory } from 'vue-router';
import { setActivePinia, createPinia } from 'pinia';

// Extend expect with jest-dom matchers
expect.extend(matchers);

// 设置 Pinia
const pinia = createPinia();
setActivePinia(pinia);

// 配置路由
const routes = [
    { path: '/', component: { template: 'one-click-grass' } },
    { path: '/one-click-grass', component: { template: 'one-click-grass' } },
    { path: '/auto-battle', component: { template: 'auto-battle' } },
    { path: '/tools', component: { template: 'tools' } },
    { path: '/settings', component: { template: 'settings' } }
];

export const router = createRouter({
    history: createWebHistory(),
    routes
});

// 全局设置 Vue Test Utils 的插件
config.global.plugins = [router, pinia];