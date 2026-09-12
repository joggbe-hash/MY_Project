<template>
  <nav class="mobile-bottom-nav">
    <!-- 1. 貼文區 -->
    <div 
      class="nav-item" 
      :class="{ active: currentTab === 'feed' }"
      @click="$emit('change-tab', 'feed')"
    >
      <div class="nav-icon">
        <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <rect x="3" y="3" width="18" height="18" rx="2"/>
          <line x1="3" y1="9" x2="21" y2="9"/>
          <line x1="9" y1="21" x2="9" y2="9"/>
        </svg>
      </div>
      <span>貼文區</span>
    </div>

    <!-- 2. 靈感清單 -->
    <div 
      class="nav-item" 
      :class="{ active: currentTab === 'inspirations' }"
      @click="$emit('change-tab', 'inspirations')"
    >
      <div class="nav-icon">
        <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
          <polyline points="14 2 14 8 20 8"/>
          <line x1="16" y1="13" x2="8" y2="13"/>
        </svg>
      </div>
      <span>靈感清單</span>
    </div>

    <!-- 3. 我也要發文 (Center Raised Pill) -->
    <div 
      class="nav-item-center-btn" 
      :class="{ active: currentTab === 'create' }"
      @click="$emit('change-tab', 'create')"
    >
      <div class="nav-center-icon">
        <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
          <line x1="12" y1="5" x2="12" y2="19"/>
          <line x1="5" y1="12" x2="19" y2="12"/>
        </svg>
      </div>
      <div class="nav-center-text">我也要<br/>發文</div>
    </div>

    <!-- 4. 我的任務 -->
    <div 
      class="nav-item" 
      :class="{ active: currentTab === 'tasks' }"
      @click="$emit('change-tab', 'tasks')"
    >
      <div class="nav-icon">
        <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M9 11l3 3L22 4"/>
          <path d="M21 12v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11"/>
        </svg>
      </div>
      <span>我的任務</span>
    </div>

    <!-- 5. 個人 -->
    <div 
      class="nav-item" 
      :class="{ active: currentTab === 'profile' }"
      @click="$emit('change-tab', 'profile')"
    >
      <div class="avatar-icon"></div>
      <span>個人</span>
    </div>
  </nav>
</template>

<script setup>
defineProps({
  currentTab: {
    type: String,
    default: 'feed'
  }
})
defineEmits(['change-tab', 'open-create-post'])
</script>

<style scoped>
.nav-icon {
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.avatar-icon {
  width: 18px;
  height: 18px;
  border-radius: 9999px;
  background: var(--avatar-default);
  border: 1px solid var(--border-main);
}

.nav-center-icon {
  display: flex;
  align-items: center;
  justify-content: center;
}

.nav-center-text {
  line-height: 1.1;
  font-size: 10px;
  text-align: center;
}
</style>
