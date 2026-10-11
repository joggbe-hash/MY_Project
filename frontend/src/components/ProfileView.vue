<template>
  <div class="profile-view">
    <!-- View Top Header (Screen 07: 編輯帳號設定 / Screen 11: 他人檔案) -->
    <div class="profile-header-bar">
      <div class="header-left-box">
        <button v-if="isOtherUser" class="profile-back-btn" @click="$emit('back-to-feed')">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <polyline points="15 18 9 12 15 6"/>
          </svg>
          <span>返回</span>
        </button>
        <span class="header-left-title">{{ isOtherUser ? '個人檔案' : '個人' }}</span>
      </div>

      <div class="header-right-actions">
        <button 
          v-if="!isOtherUser" 
          class="header-action-btn" 
          @click="$emit('open-settings')"
        >
          編輯帳號設定
        </button>
      </div>
    </div>

    <!-- User Info Card (Screen 07 & 11) -->
    <div class="profile-main-card">
      <div class="user-intro-row">
        <div 
          class="profile-avatar-circle" 
          :style="{ backgroundColor: (!isOtherUser && profile?.user?.avatar) ? 'transparent' : (!isOtherUser && profile?.user?.avatar_color ? profile?.user?.avatar_color : '#141417') }"
        >
          <img v-if="!isOtherUser && profile?.user?.avatar" :src="profile.user.avatar" class="profile-avatar-img" alt="avatar" />
          <span v-else class="profile-avatar-char">
            {{ (isOtherUser ? (viewingUsername || 'others_account') : (profile?.user?.nickname || profile?.user?.username || 'mine_account')).charAt(0).toUpperCase() }}
          </span>
        </div>
        <div class="user-meta-column">
          <div class="user-handle">{{ isOtherUser ? (viewingUsername || 'others_account') : (profile?.user?.username || 'mine_account') }}</div>
          <div class="user-motto-text">{{ profile?.user?.bio || '把靈感變成行動，把行動留下成果' }}</div>
          <div class="user-stats-summary">{{ profile?.user?.weekly_info || '本週完成 5 個任務，收藏 8 則靈感' }}</div>
        </div>
      </div>

      <!-- Stats Bar (28 位粉絲 | 12 成果 | 4 專輯) -->
      <div class="stats-counters-bar">
        <div class="counter-col">
          <div class="counter-num">28</div>
          <div class="counter-label">位粉絲</div>
        </div>
        <div class="counter-col">
          <div class="counter-num">{{ profile?.stats?.outcomes ?? 12 }}</div>
          <div class="counter-label">成果</div>
        </div>
        <div class="counter-col">
          <div class="counter-num">4</div>
          <div class="counter-label">專輯</div>
        </div>
      </div>

      <!-- Action Buttons Row -->
      <!-- Case A: Mine Account (Screen 07) -->
      <div v-if="!isOtherUser" class="action-pills-row">
        <button class="pill-btn-black">
          今日完成 5
        </button>
        <button class="pill-btn-white">
          收藏 8
        </button>
        <button class="pill-btn-white" @click="showEditModal = true">
          編輯個人檔案
        </button>
      </div>

      <!-- Case B: Other User Account (Screen 11) -->
      <div v-else class="action-pills-row split-half">
        <button 
          class="pill-btn-black half-btn"
          @click="isFollowing = !isFollowing"
        >
          {{ isFollowing ? '追蹤中' : '＋ 追蹤' }}
        </button>
        <button class="pill-btn-white half-btn" @click="handleSendMessage">
          發訊息
        </button>
      </div>
    </div>

    <!-- Subtabs Navigation -->
    <!-- Screen 07 has: 成果, 專輯, 發起, 私人 / Screen 11 has: 成果, 專輯, 發起 -->
    <div class="subtabs-nav-bar">
      <div 
        v-for="tab in currentSubtabs" 
        :key="tab"
        class="subtab-item"
        :class="{ active: activeSubtab === tab }"
        @click="activeSubtab = tab"
      >
        <div class="subtab-label">{{ tab }}</div>
        <div class="subtab-indicator" :class="{ 'indicator-active': activeSubtab === tab }"></div>
      </div>
    </div>

    <div class="subtab-divider"></div>

    <!-- 3-Column Portfolio Grid (Screen 07 & 11) -->
    <div class="portfolio-grid-container">
      <div 
        v-for="(item, index) in portfolioItems" 
        :key="index"
        class="grid-card-item"
        :class="{ 'card-gray-bg': item.isGray, 'card-white-bg': !item.isGray }"
        @click="openItemDetail(item)"
      >
        <div class="grid-card-thumbnail"></div>
        <div class="grid-card-title">{{ item.title }}</div>
        <div class="grid-card-tag">{{ item.tag }}</div>
      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'

const props = defineProps({
  profile: {
    type: Object,
    default: () => ({
      user: {
        username: 'mine_account',
        nickname: '我的帳號',
        bio: '把靈感變成行動，把行動留下成果',
        weekly_info: '本週完成 5 個任務，收藏 8 則靈感',
        avatar: '',
        avatar_color: '#141417',
        website: ''
      },
      stats: {
        outcomes: 12,
        followers: 28,
        albums: 4
      }
    })
  },
  initialIsOtherUser: {
    type: Boolean,
    default: false
  },
  viewingUsername: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update-profile', 'logout', 'navigate-onboarding', 'back-to-feed', 'open-settings'])

const isOtherUser = ref(props.initialIsOtherUser)
watch(() => props.initialIsOtherUser, (val) => {
  isOtherUser.value = val
})
const isFollowing = ref(true)
const activeSubtab = ref('成果')

const currentSubtabs = computed(() => {
  if (isOtherUser.value) {
    return ['成果', '專輯', '發起']
  }
  return ['成果', '專輯', '發起', '私人']
})

// 12 Items matching Figma Screen 07 & 11
const portfolioItems = ref([
  { title: '天空紀錄', tag: '任務', isGray: true },
  { title: '色票靈感', tag: '靈感', isGray: false },
  { title: '書桌整理', tag: '任務', isGray: true },
  { title: '租書店任務', tag: '任務', isGray: true },
  { title: '晚霞照片', tag: '任務', isGray: true },
  { title: '書桌整理', tag: '任務', isGray: true },
  { title: '手寫觀察', tag: '靈感', isGray: false },
  { title: '晚霞照片', tag: '任務', isGray: true },
  { title: '素材備忘', tag: '靈感', isGray: false },
  { title: '手寫觀察', tag: '靈感', isGray: false },
  { title: '晚霞照片', tag: '任務', isGray: true },
  { title: '素材備忘', tag: '靈感', isGray: false }
])

function handleSendMessage() {
  alert('發送訊息至 others_account')
}

function openItemDetail(item) {
  alert(`【${item.title}】\n類型：${item.tag}\n已記錄存檔於個人成果庫。`)
}
</script>

<style scoped>
.profile-view {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.profile-header-bar {
  background: #FFFFFF;
  padding: 10px 18px;
  border-radius: 8px;
  border: 1px solid #DBDBD9;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-left-title {
  color: #141414;
  font-size: 22px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
  line-height: 30px;
}

.header-left-box {
  display: flex;
  align-items: center;
  gap: 10px;
}

.profile-back-btn {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  padding: 4px 10px;
  border-radius: 14px;
  border: 1px solid #DBDBD9;
  background: #FFFFFF;
  font-size: 12px;
  font-family: inherit;
  font-weight: 500;
  color: #141417;
  cursor: pointer;
  transition: all 0.15s ease;
}

.profile-back-btn:hover {
  background: #F0F0ED;
}

.header-action-btn {
  height: 28px;
  padding: 0 12px;
  background: #FFFFFF;
  border-radius: 14px;
  outline: 1px solid #E6E6E6;
  outline-offset: -1px;
  border: none;
  color: #141414;
  font-size: 12px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 400;
  cursor: pointer;
  transition: all 0.15s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.header-action-btn:hover {
  outline-color: #B8B8B8;
  background: #F8F8F6;
}

.profile-main-card {
  background: #FFFFFF;
  border-radius: 8px;
  border: 1px solid #DBDBD9;
  padding: 12px 18px 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  box-sizing: border-box;
}

.user-intro-row {
  display: flex;
  align-items: center;
  gap: 14px;
}

.profile-avatar-circle {
  width: 64px;
  height: 64px;
  border-radius: 9999px;
  border: 1px solid #E6E6E6;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.profile-avatar-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.profile-avatar-char {
  color: #FFFFFF;
  font-size: 24px;
  font-weight: 700;
}

.user-meta-column {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.user-handle-row {
  display: flex;
  align-items: center;
  gap: 8px;
}

.user-handle {
  color: #141414;
  font-size: 18px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
  line-height: 24px;
}

.user-nickname-badge {
  display: inline-block;
  padding: 1px 8px;
  border-radius: 12px;
  background: #F4F4F1;
  border: 1px solid #E2E2DF;
  font-size: 11px;
  font-weight: 600;
  color: #555558;
}

.user-motto-text {
  color: #7A7A7A;
  font-size: 13px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 400;
  line-height: 18px;
}

.user-stats-summary {
  color: #7A7A7A;
  font-size: 12px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 400;
  line-height: 16px;
}

.stats-counters-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 4px 10px;
}

.counter-col {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2px;
}

.counter-num {
  color: #141414;
  font-size: 18px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
  line-height: 24px;
}

.counter-label {
  color: #7A7A7A;
  font-size: 12px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 400;
  line-height: 16px;
}

.action-pills-row {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.action-pills-row.split-half {
  gap: 12px;
}

.half-btn {
  flex: 1;
}

.pill-btn-black {
  height: 28px;
  padding: 6px 14px;
  background: #1A1A1A;
  border-radius: 14px;
  outline: 1px solid #1A1A1A;
  outline-offset: -1px;
  border: none;
  color: #FFFFFF;
  font-size: 12px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.15s ease;
}

.pill-btn-black:hover {
  background: #333338;
}

.pill-btn-white {
  height: 28px;
  padding: 6px 14px;
  background: #FFFFFF;
  border-radius: 14px;
  outline: 1px solid #E6E6E6;
  outline-offset: -1px;
  border: none;
  color: #141414;
  font-size: 12px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 400;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.15s ease;
}

.pill-btn-white:hover {
  background: #F8F8F6;
  outline-color: #B8B8B8;
}

.subtabs-nav-bar {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 13px;
  padding: 0 10px;
}

.subtab-item {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
  cursor: pointer;
  user-select: none;
}

.subtab-label {
  color: #7A7A7A;
  font-size: 14px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 400;
  line-height: 19px;
  transition: color 0.15s ease;
}

.subtab-item.active .subtab-label {
  color: #141414;
  font-weight: 700;
}

.subtab-indicator {
  width: 32px;
  height: 2px;
  background: transparent;
  border-radius: 1px;
  transition: all 0.15s ease;
}

.subtab-indicator.indicator-active {
  background: #141414;
  width: 100%;
}

.subtab-divider {
  width: 100%;
  height: 1px;
  background: #E6E6E6;
}

.portfolio-grid-container {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 8px;
}

.grid-card-item {
  height: 134px;
  padding: 8px;
  border-radius: 6px;
  outline: 1px solid #E6E6E6;
  outline-offset: -1px;
  display: flex;
  flex-direction: column;
  gap: 5px;
  box-sizing: border-box;
  cursor: pointer;
  transition: transform 0.15s ease, outline-color 0.15s ease;
}

.grid-card-item:hover {
  transform: translateY(-2px);
  outline-color: #B8B8B8;
}

.card-gray-bg {
  background: #E5E8EB;
}

.card-white-bg {
  background: #FFFFFF;
}

.grid-card-thumbnail {
  width: 100%;
  height: 58px;
  background: #CCCFD1;
  border-radius: 4px;
}

.grid-card-title {
  color: #141414;
  font-size: 12px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 500;
  line-height: 16px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.grid-card-tag {
  color: #7A7A7A;
  font-size: 10px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 400;
  line-height: 14px;
}

.modal-backdrop {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.4);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-window {
  width: 90%;
  max-width: 360px;
  background: #FFFFFF;
  border-radius: 12px;
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.modal-title {
  margin: 0;
  font-size: 16px;
  color: #141417;
}

.input-label {
  font-size: 12px;
  color: #75757A;
  margin-bottom: -6px;
}

.modal-input {
  width: 100%;
  padding: 8px 12px;
  border: 1px solid #DBDBD9;
  border-radius: 6px;
  font-size: 13px;
  box-sizing: border-box;
  outline: none;
}

.modal-btn-row {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  margin-top: 6px;
}

.modal-cancel-btn {
  padding: 6px 14px;
  border: 1px solid #DBDBD9;
  background: white;
  border-radius: 14px;
  font-size: 12px;
  cursor: pointer;
}

.modal-quick-links {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin: 6px 0;
}

.quick-link-divider {
  height: 1px;
  background: #EFEFEA;
  margin: 4px 0;
}

.quick-link-row-btn {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 12px;
  background: #F7F7F5;
  border: 1px solid #E5E5E3;
  border-radius: 8px;
  font-size: 13px;
  color: #141417;
  cursor: pointer;
  transition: background 0.15s ease;
}

.quick-link-row-btn:hover {
  background: #EFEFEA;
}

.link-arrow {
  color: #75757A;
  font-weight: 700;
}

.modal-save-btn {
  padding: 6px 14px;
  background: #141417;
  color: white;
  border: none;
  border-radius: 14px;
  font-size: 12px;
  cursor: pointer;
}

.modal-logout-row {
  display: flex;
  justify-content: center;
  margin-top: 14px;
  padding-top: 10px;
  border-top: 1px solid #EDEDEC;
}

.modal-logout-text-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  background: none;
  border: none;
  color: #E53935;
  font-size: 13px;
  font-family: inherit;
  font-weight: 500;
  cursor: pointer;
  padding: 6px 12px;
  border-radius: 8px;
  transition: background 0.15s ease;
}

.modal-logout-text-btn:hover {
  background: #FEECEB;
}
</style>
