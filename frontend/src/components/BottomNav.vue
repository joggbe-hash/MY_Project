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

    <!-- 2. 任務清單 -->
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
      <span>任務清單</span>
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

    <!-- 4. 靈感紀錄 -->
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
      <span>靈感紀錄</span>
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
.mobile-bottom-nav {
  display: flex;
  position: fixed;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 100%;
  max-width: 390px;
  min-height: 66px;
  padding: 6px 8px calc(6px + env(safe-area-inset-bottom, 0px));
  background: #FFFFFF;
  border-top: 1px solid #DBDBD9;
  justify-content: space-around;
  align-items: center;
  z-index: 1000;
  box-sizing: border-box;
  box-shadow: 0 -2px 10px rgba(0, 0, 0, 0.04);
}

@media (max-width: 480px) {
  .mobile-bottom-nav {
    left: 0;
    right: 0;
    transform: none;
    max-width: 100%;
  }
}

.nav-item {
  width: 60px;
  height: 48px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 6px;
  cursor: pointer;
  color: #75757A;
  transition: all 0.15s ease;
  user-select: none;
}

.nav-item.active {
  color: #141417;
}

.nav-item span {
  font-size: 10px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 400;
  text-align: center;
}

.nav-item.active span {
  font-weight: 700;
}

.nav-icon {
  width: 18px;
  height: 18px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.avatar-icon {
  width: 20px;
  height: 20px;
  border-radius: 9999px;
  background: #D1D6DB;
  border: 1px solid #DBDBD9;
}

.nav-item-center-btn {
  width: 78px;
  height: 60px;
  padding: 6px 8px;
  background: #141417;
  border-radius: 30px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 3px;
  cursor: pointer;
  color: #FFFFFF;
  transition: transform 0.15s ease, background-color 0.15s ease;
  user-select: none;
}

.nav-item-center-btn:hover {
  background: #25252b;
}

.nav-item-center-btn:active {
  transform: scale(0.96);
}

.nav-center-icon {
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #B8B8B8;
}

.nav-center-text {
  font-size: 10px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 400;
  line-height: 1.1;
  text-align: center;
  color: #FFFFFF;
}
</style>
