<template>
  <aside class="desktop-sidebar">
    <!-- Brand Logo -->
    <div class="sidebar-brand" @click="$emit('change-tab', 'feed')">
      <div class="brand-badge">我</div>
      <div class="brand-text">
        <div class="brand-title">我也要</div>
        <div class="brand-sub">靈感與行動社群</div>
      </div>
    </div>

    <!-- Navigation Menu -->
    <nav class="sidebar-nav">
      <button 
        class="nav-btn" 
        :class="{ active: currentTab === 'feed' }"
        @click="$emit('change-tab', 'feed')"
      >
        <div class="icon-wrap">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <rect x="3" y="3" width="18" height="18" rx="2"/>
            <line x1="9" y1="3" x2="9" y2="21"/>
          </svg>
        </div>
        <span>貼文區</span>
      </button>

      <button 
        class="nav-btn" 
        :class="{ active: currentTab === 'inspirations' }"
        @click="$emit('change-tab', 'inspirations')"
      >
        <div class="icon-wrap">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
            <polyline points="14 2 14 8 20 8"/>
            <line x1="16" y1="13" x2="8" y2="13"/>
            <line x1="16" y1="17" x2="8" y2="17"/>
          </svg>
        </div>
        <span>靈感清單</span>
      </button>

      <button 
        class="nav-btn" 
        :class="{ active: currentTab === 'tasks' }"
        @click="$emit('change-tab', 'tasks')"
      >
        <div class="icon-wrap">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <polyline points="9 11 12 14 22 4"/>
            <path d="M21 12v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11"/>
          </svg>
        </div>
        <span>我的任務</span>
      </button>

      <button 
        class="nav-btn" 
        :class="{ active: currentTab === 'profile' }"
        @click="$emit('change-tab', 'profile')"
      >
        <div class="icon-wrap">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
            <circle cx="12" cy="7" r="4"/>
          </svg>
        </div>
        <span>個人主頁</span>
      </button>
    </nav>

    <!-- Prominent Action Buttons -->
    <div class="sidebar-actions">
      <button class="create-post-btn" @click="$emit('open-create-post')">
        <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
          <line x1="12" y1="5" x2="12" y2="19"/>
          <line x1="5" y1="12" x2="19" y2="12"/>
        </svg>
        <span>我也要發文</span>
      </button>

      <button class="daily-task-btn" @click="$emit('open-daily-task')">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="12" cy="12" r="10"/>
          <polyline points="12 6 12 12 16 14"/>
        </svg>
        <span>今日任務規劃</span>
      </button>
    </div>

    <!-- User Profile Micro Card & Logout -->
    <div class="sidebar-user-section">
      <div class="sidebar-user" @click="$emit('change-tab', 'profile')">
        <div class="avatar avatar-sm">
          <span>{{ (currentUser?.username || 'U').charAt(0).toUpperCase() }}</span>
        </div>
        <div class="user-info">
          <div class="user-name">{{ currentUser?.username || 'lavender_nation' }}</div>
          <div class="user-handle">{{ currentUser?.weekly_info || '本週完成 5 個任務' }}</div>
        </div>
      </div>
      <button class="logout-btn" title="登出帳號" @click="$emit('logout')">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/>
          <polyline points="16 17 21 12 16 7"/>
          <line x1="21" y1="12" x2="9" y2="12"/>
        </svg>
        <span>登出</span>
      </button>
    </div>
  </aside>
</template>

<script setup>
defineProps({
  currentTab: {
    type: String,
    default: 'feed'
  },
  currentUser: {
    type: Object,
    default: () => ({})
  }
})
defineEmits(['change-tab', 'open-create-post', 'open-daily-task', 'logout'])
</script>

<style scoped>
.sidebar-brand {
  display: flex;
  align-items: center;
  gap: 12px;
  cursor: pointer;
  padding-bottom: 20px;
  margin-bottom: 16px;
  border-bottom: 1px solid var(--border-light);
}

.brand-badge {
  width: 36px;
  height: 36px;
  background: var(--bg-dark);
  color: #fff;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 16px;
}

.brand-title {
  font-size: 18px;
  font-weight: 700;
  color: var(--text-main);
  line-height: 1.2;
}

.brand-sub {
  font-size: 11px;
  color: var(--text-secondary);
}

.sidebar-nav {
  display: flex;
  flex-direction: column;
  gap: 6px;
  flex: 1;
}

.nav-btn {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 14px;
  border-radius: 10px;
  border: none;
  background: transparent;
  color: var(--text-secondary);
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: var(--transition);
  text-align: left;
  width: 100%;
}

.nav-btn:hover {
  background: var(--bg-secondary);
  color: var(--text-main);
}

.nav-btn.active {
  background: var(--bg-dark);
  color: #fff;
  font-weight: 700;
}

.icon-wrap {
  display: flex;
  align-items: center;
  justify-content: center;
}

.sidebar-actions {
  display: flex;
  flex-direction: column;
  gap: 10px;
  padding: 16px 0;
}

.create-post-btn {
  background: var(--bg-dark);
  color: #fff;
  border: 1px solid var(--bg-dark);
  border-radius: 12px;
  padding: 12px 16px;
  font-size: 14px;
  font-weight: 700;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  transition: var(--transition);
  box-shadow: 0 4px 12px rgba(20, 20, 23, 0.15);
}

.create-post-btn:hover {
  background: var(--bg-dark-hover);
  transform: translateY(-1px);
}

.daily-task-btn {
  background: transparent;
  color: var(--text-main);
  border: 1px solid var(--border-main);
  border-radius: 12px;
  padding: 10px 16px;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  transition: var(--transition);
}

.daily-task-btn:hover {
  background: var(--bg-secondary);
  border-color: var(--border-strong);
}

.sidebar-user-section {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-top: auto;
}

.sidebar-user {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 12px;
  border-radius: 12px;
  border: 1px solid var(--border-light);
  background: var(--bg-app);
  cursor: pointer;
  transition: var(--transition);
}

.sidebar-user:hover {
  background: #ECECE9;
}

.logout-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  padding: 6px 12px;
  border-radius: 8px;
  border: 1px solid var(--border-light);
  background: transparent;
  color: var(--text-secondary);
  font-size: 12px;
  cursor: pointer;
  transition: var(--transition);
}

.logout-btn:hover {
  background: #FFF5F5;
  color: #E03131;
  border-color: #FFC9C9;
}

.user-name {
  font-size: 13px;
  font-weight: 700;
  color: var(--text-main);
}

.user-handle {
  font-size: 11px;
  color: var(--text-secondary);
}
</style>
