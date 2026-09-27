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

    <!-- Stats Row (28 好友, 12 成果, 4 專輯) -->
    <div class="stats-counter-bar">
      <div class="stat-col">
        <div class="stat-num">{{ profile?.stats?.friends ?? 28 }}</div>
        <div class="stat-label">好友</div>
      </div>
      <div class="stat-col" @click="activeSubTab = '成果'">
        <div class="stat-num">{{ profile?.stats?.outcomes ?? 12 }}</div>
        <div class="stat-label">成果</div>
      </div>
      <div class="stat-col" @click="activeSubTab = '專輯'">
        <div class="stat-num">{{ profile?.stats?.albums ?? 4 }}</div>
        <div class="stat-label">專輯</div>
      </div>
    </div>

    <!-- Badges Row (今日完成 5, 收藏 8, 連續登入 3 天) -->
    <div class="badges-row">
      <div class="badge-pill active">今日完成 5</div>
      <div class="badge-pill">收藏 8</div>
      <div class="badge-pill">連續登入 3 天</div>
    </div>

    <!-- Tabs (成果, 專輯, 發起, 私人) -->
    <div class="profile-tabs-nav">
      <div 
        v-for="tab in ['成果', '專輯', '發起', '私人']" 
        :key="tab"
        class="profile-tab-item"
        :class="{ active: activeSubTab === tab }"
        @click="activeSubTab = tab"
      >
        <span>{{ tab }}</span>
      </div>
    </div>

    <!-- RUICHIEN 設計註釋: 預設顯示完成任務後分享的照片與文字，也保留從靈感延伸出的創作。可自訂是否展示。 -->
    <div v-if="activeSubTab === '成果'" class="outcome-header-toolbar">
      <div class="outcome-subtitle">
        預設顯示完成任務後分享的照片與文字，也保留從靈感延伸出的創作
      </div>
      <button 
        class="btn-xs customize-toggle-btn"
        :class="isCustomizingDisplay ? 'btn-dark' : 'btn-outline'"
        @click="isCustomizingDisplay = !isCustomizingDisplay"
      >
        {{ isCustomizingDisplay ? '完成自訂' : '自訂展示' }}
      </button>
    </div>

    <!-- Divider -->
    <div class="profile-divider"></div>

    <!-- 3x4 Grid (Matching Figma Screen 07: 12 cards) - 點擊後可以進入個別頁面，如ig可滑動 -->
    <div class="portfolio-grid">
      <div 
        v-for="(item, idx) in displayedItems" 
        :key="item.id" 
        class="grid-item-box"
        :class="{ 'item-hidden': !item.is_displayed, 'in-customizing': isCustomizingDisplay }"
        :style="{ background: item.bg || 'white' }"
        @click="handleCardClick(item, idx)"
      >
        <div class="grid-thumb">
          <!-- 自訂展示時顯示切換勾選或狀態圖示 -->
          <div 
            v-if="isCustomizingDisplay" 
            class="display-toggle-badge"
            :class="{ active: item.is_displayed }"
            @click.stop="toggleItemDisplay(item)"
          >
            {{ item.is_displayed ? '👁 展示中' : '🚫 已隱藏' }}
          </div>
          <div v-else-if="!item.is_displayed" class="hidden-pill">
            隱藏
          </div>
        </div>
        <div class="grid-text-wrap">
          <div class="grid-item-title">{{ item.title }}</div>
          <div class="grid-item-tag">{{ item.tag }}</div>
        </div>
      </div>
    </div>

    <!-- Instagram-like Slideable Outcome Detail Modal (RUICHIEN 設計註釋: 點擊後可以進入個別頁面，如ig可滑動) -->
    <div v-if="showDetailModal && currentDetailItem" class="modal-backdrop ig-modal-backdrop" @click.self="showDetailModal = false">
      <div class="ig-detail-container">
        <!-- Close Button -->
        <button class="ig-close-btn" @click="showDetailModal = false" title="關閉">✕</button>

        <!-- Slide Left Button -->
        <button 
          class="ig-nav-btn ig-prev-btn" 
          :disabled="currentDetailIndex === 0"
          @click="prevItem"
          title="上一則"
        >
          ‹
        </button>

        <!-- Slide Right Button -->
        <button 
          class="ig-nav-btn ig-next-btn" 
          :disabled="currentDetailIndex === portfolioList.length - 1"
          @click="nextItem"
          title="下一則"
        >
          ›
        </button>

        <!-- IG Card Content -->
        <div class="ig-card">
          <!-- Top author & info bar -->
          <div class="ig-card-header">
            <div class="ig-author-col">
              <div class="avatar avatar-sm"></div>
              <div class="ig-author-meta">
                <span class="ig-author-name">{{ profile?.user?.username || 'lavender_nation' }}</span>
                <span class="ig-post-time">{{ currentDetailItem.time || '今天' }}</span>
              </div>
            </div>
            <div class="ig-header-right">
              <span class="ig-tag-badge">{{ currentDetailItem.tag }}成果</span>
              <button 
                class="ig-visibility-btn" 
                :class="{ 'is-hidden': !currentDetailItem.is_displayed }"
                @click="toggleItemDisplay(currentDetailItem)"
                title="可自訂是否展示"
              >
                {{ currentDetailItem.is_displayed ? '👁 公開展示中' : '🚫 已設為隱藏' }}
              </button>
            </div>
          </div>

          <!-- Main Visual Showcase (Photo or Creation Canvas) -->
          <div class="ig-media-wrapper" :style="{ background: currentDetailItem.mediaBg || '#E5E8EB' }">
            <div class="ig-media-content">
              <div class="ig-media-icon">
                {{ currentDetailItem.tag === '任務' ? '📷' : '🎨' }}
              </div>
              <div class="ig-media-label">{{ currentDetailItem.title }} 成果紀錄</div>
            </div>
            <!-- Slider Index Indicator -->
            <div class="ig-index-pill">{{ currentDetailIndex + 1 }} / {{ portfolioList.length }}</div>
          </div>

          <!-- Bottom Story Description & Interaction -->
          <div class="ig-card-body">
            <div class="ig-body-actions">
              <div class="ig-action-item">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path>
                </svg>
                <span>{{ currentDetailItem.likes || 42 }} 喜歡</span>
              </div>
              <div class="ig-action-item">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path>
                </svg>
                <span>{{ currentDetailItem.comments || 6 }} 留言</span>
              </div>
            </div>

            <h4 class="ig-story-title">{{ currentDetailItem.title }}</h4>
            <p class="ig-story-desc">{{ currentDetailItem.description }}</p>
          </div>
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
import { ref, computed, watch, onMounted, onUnmounted } from 'vue'

const props = defineProps({
  profile: {
    type: Object,
    default: () => ({})
  }
})

const emit = defineEmits(['update-profile', 'logout'])

const activeSubTab = ref('成果')
const showEditModal = ref(false)
const isCustomizingDisplay = ref(false)
const editBio = ref(props.profile?.user?.bio || '把靈感變成行動，把行動留下成果')
const editWeekly = ref(props.profile?.user?.weekly_info || '★ 週進度：4 個任務，已達成連續 3 天記錄')

watch(() => props.profile, (newProfile) => {
  if (newProfile?.user) {
    if (newProfile.user.bio) editBio.value = newProfile.user.bio
    if (newProfile.user.weekly_info) editWeekly.value = newProfile.user.weekly_info
  }
}, { deep: true, immediate: true })

// 12 items exactly matching Figma Screen 07 + Rich Story Content (RUICHIEN 設計註釋: 完成任務分享的照片與文字 + 靈感延伸創作)
const portfolioList = ref([
  { 
    id: 1, 
    title: '天空紀錄', 
    tag: '任務', 
    bg: '#E5E8EB',
    mediaBg: 'linear-gradient(135deg, #a1c4fd 0%, #c2e9fb 100%)',
    description: '完成任務後分享：今天的天空很特別，陽光穿過層疊的雲朵，泛著藍灰與金色的光澤，記錄下此時此刻的平靜。',
    time: '今天 14:20',
    likes: 72,
    comments: 14,
    is_displayed: true
  },
  { 
    id: 2, 
    title: '色票靈感', 
    tag: '靈感', 
    bg: 'white',
    mediaBg: 'linear-gradient(135deg, #ff9a9e 0%, #fecfef 99%, #fecfef 100%)',
    description: '從靈感延伸出的創作：在通勤捷運上採集到的三個對比色票，未來可以用在筆記排版上！',
    time: '昨天 19:40',
    likes: 86,
    comments: 4,
    is_displayed: true
  },
  { 
    id: 3, 
    title: '書桌整理', 
    tag: '任務', 
    bg: '#E5E8EB',
    mediaBg: 'linear-gradient(135deg, #fdfbfb 0%, #ebedee 100%)',
    description: '完成任務後分享：把桌上的筆記本與便利貼全部歸位，換上乾淨的新桌墊，工作效率直接翻倍。',
    time: '2天前',
    likes: 53,
    comments: 8,
    is_displayed: true
  },
  { 
    id: 4, 
    title: '租書店任務', 
    tag: '任務', 
    bg: '#E5E8EB',
    mediaBg: 'linear-gradient(135deg, #e0c3fc 0%, #8ec5fc 100%)',
    description: '完成了社區租書店的舊書巡禮，重溫了小時候看過的柯南全套！',
    time: '3天前',
    likes: 65,
    comments: 11,
    is_displayed: true
  },
  { 
    id: 5, 
    title: '晚霞照片', 
    tag: '任務', 
    bg: '#E5E8EB',
    mediaBg: 'linear-gradient(135deg, #f6d365 0%, #fda085 100%)',
    description: '下班路上的夕陽橘紅色晚霞，隨手拍下分享給社群夥伴們。',
    time: '4天前',
    likes: 98,
    comments: 19,
    is_displayed: true
  },
  { 
    id: 6, 
    title: '書桌整理', 
    tag: '任務', 
    bg: '#E5E8EB',
    mediaBg: 'linear-gradient(135deg, #cfd9df 0%, #e2ebf0 100%)',
    description: '定期每週維持乾淨整潔的桌面環境。',
    time: '5天前',
    likes: 31,
    comments: 2,
    is_displayed: true
  },
  { 
    id: 7, 
    title: '手寫觀察', 
    tag: '靈感', 
    bg: 'white',
    mediaBg: 'linear-gradient(135deg, #fbc2eb 0%, #a6c1ee 100%)',
    description: '紀錄老招牌上的傳統手寫字體筆觸與字距比例。',
    time: '6天前',
    likes: 47,
    comments: 5,
    is_displayed: true
  },
  { 
    id: 8, 
    title: '晚霞照片', 
    tag: '任務', 
    bg: '#E5E8EB',
    mediaBg: 'linear-gradient(135deg, #fdcbf1 0%, #e6dee9 100%)',
    description: '夏日傍晚的粉紫色魔幻時刻，留下一張成果。',
    time: '1週前',
    likes: 64,
    comments: 7,
    is_displayed: true
  },
  { 
    id: 9, 
    title: '素材備忘', 
    tag: '靈感', 
    bg: 'white',
    mediaBg: 'linear-gradient(135deg, #a8edea 0%, #fed6e3 100%)',
    description: '公園木頭長椅與落葉的天然材質紋理紀錄。',
    time: '1週前',
    likes: 39,
    comments: 3,
    is_displayed: true
  },
  { 
    id: 10, 
    title: '手寫觀察', 
    tag: '靈感', 
    bg: 'white',
    mediaBg: 'linear-gradient(135deg, #e0c3fc 0%, #8ec5fc 100%)',
    description: '古早味小吃攤的菜單手繪插圖，充滿人情味。',
    time: '2週前',
    likes: 58,
    comments: 9,
    is_displayed: true
  },
  { 
    id: 11, 
    title: '晚霞照片', 
    tag: '任務', 
    bg: '#E5E8EB',
    mediaBg: 'linear-gradient(135deg, #f093fb 0%, #f5576c 100%)',
    description: '雲海透光的奇幻晚霞。',
    time: '2週前',
    likes: 81,
    comments: 12,
    is_displayed: true
  },
  { 
    id: 12, 
    title: '素材備忘', 
    tag: '靈感', 
    bg: 'white',
    mediaBg: 'linear-gradient(135deg, #5ee7df 0%, #b490ca 100%)',
    description: '日常聲音與環境色彩的隨筆。',
    time: '3週前',
    likes: 29,
    comments: 1,
    is_displayed: true
  }
])

// 依據自訂展示狀態過濾或全顯
const displayedItems = computed(() => {
  if (isCustomizingDisplay.value) {
    return portfolioList.value
  }
  return portfolioList.value.filter(item => item.is_displayed)
})

// RUICHIEN 設計註釋: 點擊後可以進入個別頁面，如ig可滑動
const showDetailModal = ref(false)
const currentDetailIndex = ref(0)
const currentDetailItem = computed(() => portfolioList.value[currentDetailIndex.value] || null)

function handleCardClick(item, idx) {
  if (isCustomizingDisplay.value) {
    toggleItemDisplay(item)
    return
  }
  const realIndex = portfolioList.value.findIndex(p => p.id === item.id)
  currentDetailIndex.value = realIndex !== -1 ? realIndex : idx
  showDetailModal.value = true
}

function prevItem() {
  if (currentDetailIndex.value > 0) {
    currentDetailIndex.value--
  }
}

function nextItem() {
  if (currentDetailIndex.value < portfolioList.value.length - 1) {
    currentDetailIndex.value++
  }
}

// 鍵盤左右鍵滑動切換
function handleKeydown(e) {
  if (!showDetailModal.value) return
  if (e.key === 'ArrowLeft') prevItem()
  if (e.key === 'ArrowRight') nextItem()
  if (e.key === 'Escape') showDetailModal.value = false
}

onMounted(() => {
  window.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  window.removeEventListener('keydown', handleKeydown)
})

// RUICHIEN 設計註釋: 可自訂是否展示
function toggleItemDisplay(item) {
  item.is_displayed = !item.is_displayed
}

function saveProfile() {
  emit('update-profile', { bio: editBio.value, weekly_info: editWeekly.value })
  showEditModal.value = false
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

.profile-divider {
  width: 100%;
  height: 1px;
  background: #E6E6E6;
  margin: 2px 0 10px;
}

/* 3x4 Grid matching Figma */
.portfolio-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 8px;
}

.grid-item-box {
  padding: 8px;
  border: 1px solid #E6E6E6;
  border-radius: 6px;
  display: flex;
  flex-direction: column;
  height: 134px;
  cursor: pointer;
  overflow: hidden;
  transition: var(--transition);
}

.grid-item-box:hover {
  border-color: #B8B8B8;
  transform: translateY(-1px);
}

.grid-thumb {
  width: 100%;
  height: 58px;
  border-radius: 4px;
  background: #CCCFD1;
  flex-shrink: 0;
}

.grid-text-wrap {
  padding-top: 5px;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.grid-item-title {
  font-size: 12px;
  font-weight: 500;
  color: #141414;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.grid-item-tag {
  font-size: 10px;
  color: #7A7A7A;
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

/* RUICHIEN 設計註釋: 自訂展示工具列 */
.outcome-header-toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 6px 4px;
  gap: 12px;
  background: var(--bg-subtle);
  border-radius: var(--radius-sm);
  margin-top: -6px;
}

.outcome-subtitle {
  font-size: 11px;
  color: var(--text-secondary);
  line-height: 1.4;
  flex: 1;
}

.customize-toggle-btn {
  flex-shrink: 0;
  padding: 4px 10px;
  font-size: 11px;
}

.grid-item-box.item-hidden {
  opacity: 0.45;
  filter: grayscale(0.5);
}

.grid-item-box.in-customizing {
  outline: 2px dashed #4dabf7;
  outline-offset: -2px;
}

.display-toggle-badge {
  font-size: 10px;
  padding: 2px 4px;
  background: rgba(0, 0, 0, 0.7);
  color: white;
  border-radius: 4px;
  margin: 4px;
  display: inline-block;
  font-weight: 600;
}

.display-toggle-badge.active {
  background: #2b8a3e;
}

.hidden-pill {
  font-size: 10px;
  padding: 2px 6px;
  background: rgba(0, 0, 0, 0.6);
  color: white;
  border-radius: 4px;
  margin: 4px;
  display: inline-block;
}

/* RUICHIEN 設計註釋: Instagram-like Slideable Detail Modal */
.ig-modal-backdrop {
  z-index: 1000;
  background: rgba(0, 0, 0, 0.75);
}

.ig-detail-container {
  position: relative;
  width: 100%;
  max-width: 420px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.ig-close-btn {
  position: absolute;
  top: -42px;
  right: 0;
  background: rgba(255, 255, 255, 0.2);
  border: none;
  color: white;
  font-size: 18px;
  width: 32px;
  height: 32px;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: var(--transition);
}

.ig-close-btn:hover {
  background: rgba(255, 255, 255, 0.4);
}

.ig-nav-btn {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  background: rgba(255, 255, 255, 0.9);
  border: none;
  font-size: 24px;
  width: 36px;
  height: 36px;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
  z-index: 10;
  transition: var(--transition);
}

.ig-prev-btn {
  left: -18px;
}

.ig-next-btn {
  right: -18px;
}

.ig-nav-btn:disabled {
  opacity: 0.3;
  cursor: not-allowed;
}

.ig-card {
  width: 100%;
  background: white;
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.3);
  display: flex;
  flex-direction: column;
}

.ig-card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 14px 16px;
  border-bottom: 1px solid var(--border-card);
}

.ig-author-col {
  display: flex;
  align-items: center;
  gap: 10px;
}

.ig-author-meta {
  display: flex;
  flex-direction: column;
}

.ig-author-name {
  font-size: 13px;
  font-weight: 700;
  color: var(--text-main);
}

.ig-post-time {
  font-size: 11px;
  color: var(--text-secondary);
}

.ig-header-right {
  display: flex;
  align-items: center;
  gap: 8px;
}

.ig-tag-badge {
  font-size: 11px;
  background: var(--bg-subtle);
  color: var(--text-main);
  padding: 3px 8px;
  border-radius: var(--radius-pill);
  font-weight: 600;
}

.ig-visibility-btn {
  font-size: 11px;
  padding: 3px 8px;
  border-radius: var(--radius-pill);
  border: 1px solid var(--border-main);
  background: var(--bg-card);
  cursor: pointer;
  font-weight: 500;
}

.ig-visibility-btn.is-hidden {
  border-color: #FA5252;
  color: #FA5252;
}

.ig-media-wrapper {
  width: 100%;
  height: 240px;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
}

.ig-media-content {
  text-align: center;
  color: #333;
}

.ig-media-icon {
  font-size: 48px;
  margin-bottom: 6px;
}

.ig-media-label {
  font-size: 14px;
  font-weight: 600;
  color: #141414;
}

.ig-index-pill {
  position: absolute;
  top: 12px;
  right: 12px;
  background: rgba(0, 0, 0, 0.6);
  color: white;
  padding: 3px 8px;
  border-radius: var(--radius-pill);
  font-size: 11px;
  font-weight: 600;
}

.ig-card-body {
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.ig-body-actions {
  display: flex;
  align-items: center;
  gap: 16px;
  color: var(--text-secondary);
  font-size: 12px;
  margin-bottom: 4px;
}

.ig-action-item {
  display: flex;
  align-items: center;
  gap: 4px;
}

.ig-story-title {
  font-size: 15px;
  font-weight: 700;
  color: var(--text-main);
}

.ig-story-desc {
  font-size: 13px;
  color: var(--text-main);
  line-height: 1.6;
}
</style>
