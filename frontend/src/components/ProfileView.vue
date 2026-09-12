<template>
  <div class="profile-view">
    <!-- View Header (Figma Screen 07) -->
    <div class="profile-top-bar">
      <div class="top-title">個人</div>
      <button class="btn-outline btn-xs edit-btn" @click="showEditModal = true">
        編輯資料
      </button>
    </div>

    <!-- User Profile Details -->
    <div class="user-intro-section">
      <div class="avatar-circle"></div>
      <div class="user-text-info">
        <h2 class="user-username">{{ profile?.user?.username || 'lavender_nation' }}</h2>
        <p class="user-motto">{{ profile?.user?.bio || '把靈感變成行動，把行動留下成果' }}</p>
        <p class="user-subline">{{ profile?.user?.weekly_info || '★ 週進度：4 個任務，已達成連續 3 天記錄' }}</p>
      </div>
    </div>

    <!-- Stats Row (12 成果, 28 任務, 46 靈感, 4 發起) -->
    <div class="stats-counter-bar">
      <div class="stat-col" @click="activeSubTab = '成果'">
        <div class="stat-num">{{ profile?.stats?.outcomes ?? 12 }}</div>
        <div class="stat-label">成果</div>
      </div>
      <div class="stat-col" @click="activeSubTab = '任務'">
        <div class="stat-num">{{ profile?.stats?.tasks ?? 28 }}</div>
        <div class="stat-label">任務</div>
      </div>
      <div class="stat-col" @click="activeSubTab = '靈感'">
        <div class="stat-num">{{ profile?.stats?.inspirations ?? 46 }}</div>
        <div class="stat-label">靈感</div>
      </div>
      <div class="stat-col" @click="activeSubTab = '發起'">
        <div class="stat-num">{{ profile?.stats?.launches ?? 4 }}</div>
        <div class="stat-label">發起</div>
      </div>
    </div>

    <!-- Badges Row (完成 5, 連續 3 天, 收藏 8, 發起 4) -->
    <div class="badges-row">
      <div class="badge-pill active">完成 5</div>
      <div class="badge-pill">連續 3 天</div>
      <div class="badge-pill">收藏 8</div>
      <div class="badge-pill">發起 4</div>
    </div>

    <!-- Tabs (成果, 任務, 靈感, 發起) -->
    <div class="profile-tabs-nav">
      <div 
        v-for="tab in ['成果', '任務', '靈感', '發起']" 
        :key="tab"
        class="profile-tab-item"
        :class="{ active: activeSubTab === tab }"
        @click="activeSubTab = tab"
      >
        <span>{{ tab }}</span>
      </div>
    </div>

    <!-- Showcase Info Card (Figma Screen 07: 成果展示) -->
    <div class="figma-card showcase-card">
      <div class="showcase-card-title">成果展示</div>
      <p class="showcase-card-desc">
        預設顯示完成任務後分享的照片與文字。也保留從靈感延伸出的創作。
      </p>
    </div>

    <!-- 3x3 Grid (Matching Figma 07) -->
    <div class="portfolio-grid">
      <div 
        v-for="item in displayedItems" 
        :key="item.id" 
        class="grid-item-box"
        :class="{ 'has-thumbnail': item.has_image }"
        @click="openItemDetail(item)"
      >
        <div v-if="item.has_image" class="grid-thumb"></div>
        <div class="grid-text-wrap">
          <div class="grid-item-title">{{ item.title }}</div>
          <div class="grid-item-tag">{{ item.tag }}</div>
        </div>
      </div>
    </div>

    <!-- Edit Profile Modal -->
    <div v-if="showEditModal" class="modal-backdrop" @click.self="showEditModal = false">
      <div class="modal-card">
        <h3 class="card-title">編輯個人檔案</h3>
        <label class="form-label">個人標語 / 座右銘</label>
        <input v-model="editBio" class="input-control" />
        <label class="form-label">週進度說明</label>
        <input v-model="editWeekly" class="input-control" />
        <div class="modal-footer">
          <button class="btn-outline" @click="showEditModal = false">取消</button>
          <button class="btn-dark" @click="saveProfile">儲存更新</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'

const props = defineProps({
  profile: {
    type: Object,
    default: () => ({})
  }
})

const emit = defineEmits(['update-profile', 'logout'])

const activeSubTab = ref('成果')
const showEditModal = ref(false)
const editBio = ref(props.profile?.user?.bio || '把靈感變成行動，把行動留下成果')
const editWeekly = ref(props.profile?.user?.weekly_info || '★ 週進度：4 個任務，已達成連續 3 天記錄')

watch(() => props.profile, (newProfile) => {
  if (newProfile?.user) {
    if (newProfile.user.bio) editBio.value = newProfile.user.bio
    if (newProfile.user.weekly_info) editWeekly.value = newProfile.user.weekly_info
  }
}, { deep: true, immediate: true })

// 9 items exactly matching Figma Screen 07
const defaultItems = [
  { id: 1, title: '天空紀錄', tag: '任務成果', has_image: true },
  { id: 2, title: '書桌整理', tag: '任務成果', has_image: false },
  { id: 3, title: '色票靈感', tag: '靈感轉化', has_image: true },
  { id: 4, title: '租書店任務', tag: '整理任務', has_image: false },
  { id: 5, title: '專案草圖', tag: '靈感轉化', has_image: false },
  { id: 6, title: '晚餐照片', tag: '任務成果', has_image: true },
  { id: 7, title: '手寫觀察', tag: '私人筆記', has_image: false },
  { id: 8, title: '桌面改造', tag: '任務成果', has_image: true },
  { id: 9, title: '素材筆記', tag: '靈感整理', has_image: false }
]

const displayedItems = computed(() => {
  return defaultItems
})

function saveProfile() {
  emit('update-profile', { bio: editBio.value, weekly_info: editWeekly.value })
  showEditModal.value = false
}

function openItemDetail(item) {
  alert(`【${item.title}】\n類型：${item.tag}\n完成狀態：已打卡存檔！`)
}
</script>

<style scoped>
.profile-view {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.profile-top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 4px 2px;
}

.top-title {
  font-size: 24px;
  font-weight: 700;
  color: var(--text-main);
}

.edit-btn {
  padding: 5px 14px;
  font-size: 12px;
}

.user-intro-section {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 6px 2px;
}

.avatar-circle {
  width: 66px;
  height: 66px;
  border-radius: 50%;
  background: var(--avatar-default);
  flex-shrink: 0;
}

.user-text-info {
  display: flex;
  flex-direction: column;
  gap: 3px;
}

.user-username {
  font-size: 18px;
  font-weight: 700;
  color: var(--text-main);
}

.user-motto {
  font-size: 13px;
  color: var(--text-secondary);
}

.user-subline {
  font-size: 11px;
  color: var(--text-muted);
}

.stats-counter-bar {
  display: flex;
  justify-content: space-around;
  padding: 10px 0;
}

.stat-col {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2px;
  cursor: pointer;
}

.stat-num {
  font-size: 18px;
  font-weight: 700;
  color: var(--text-main);
}

.stat-label {
  font-size: 12px;
  color: var(--text-secondary);
}

.badges-row {
  display: flex;
  gap: 8px;
  overflow-x: auto;
  scrollbar-width: none;
}
.badges-row::-webkit-scrollbar {
  display: none;
}

.badge-pill {
  padding: 5px 14px;
  border-radius: var(--radius-pill);
  font-size: 12px;
  border: 1px solid var(--border-light);
  background: var(--bg-card);
  color: var(--text-main);
  white-space: nowrap;
}

.badge-pill.active {
  background: var(--bg-dark);
  color: var(--text-white);
  border-color: var(--bg-dark);
  font-weight: 700;
}

.profile-tabs-nav {
  display: flex;
  justify-content: space-around;
  border-bottom: 1px solid var(--border-light);
  padding: 6px 0 0;
}

.profile-tab-item {
  padding: 8px 12px 10px;
  font-size: 14px;
  color: var(--text-secondary);
  cursor: pointer;
  position: relative;
}

.profile-tab-item.active {
  color: var(--text-main);
  font-weight: 700;
}

.profile-tab-item.active::after {
  content: '';
  position: absolute;
  bottom: -1px;
  left: 0;
  right: 0;
  height: 2px;
  background: var(--text-main);
}

.showcase-card {
  padding: 14px 18px;
  gap: 4px;
}

.showcase-card-title {
  font-size: 14px;
  font-weight: 700;
  color: var(--text-main);
}

.showcase-card-desc {
  font-size: 12px;
  color: var(--text-secondary);
  line-height: 1.5;
}

/* 3x3 Grid */
.portfolio-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 12px;
}

.grid-item-box {
  background: var(--bg-card);
  border: 1px solid var(--border-main);
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  min-height: 110px;
  cursor: pointer;
  overflow: hidden;
  transition: var(--transition);
}

.grid-item-box:hover {
  border-color: var(--border-strong);
}

.grid-thumb {
  width: 100%;
  height: 64px;
  background: #D9D9D9;
}

.grid-text-wrap {
  padding: 8px 10px;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.grid-item-title {
  font-size: 12px;
  font-weight: 700;
  color: var(--text-main);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.grid-item-tag {
  font-size: 10px;
  color: var(--text-secondary);
}

.form-label {
  font-size: 13px;
  font-weight: 700;
  color: var(--text-main);
  margin-top: 4px;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  margin-top: 10px;
}
</style>
