<template>
  <div class="inspirations-view">
    <!-- View Header (Figma Screen 06) -->
    <div class="view-header">
      <div class="view-header-title">靈感清單</div>
      <div class="view-header-subtitle">蒐集靈感、隨手紀錄</div>
    </div>

    <!-- Search & Date Filter Bar -->
    <div class="search-date-filter-row">
      <div class="search-box-pill">
        <input 
          v-model="searchQuery" 
          placeholder="搜尋關鍵字" 
          class="search-input"
          @input="$emit('search', searchQuery)"
        />
      </div>

      <button class="date-filter-pill" @click="toggleDateFilter">
        <span>{{ selectedDateFilter }}</span>
      </button>
    </div>

    <!-- Quick Note Prompt Box -->
    <div class="quick-input-card">
      <input 
        v-model="quickNoteText" 
        class="quick-note-input"
        placeholder="今天有甚麼想法嗎？馬上記錄下來吧！"
        @keyup.enter="handleQuickSubmit"
      />
      <button 
        v-if="quickNoteText.trim()" 
        class="btn-save-note" 
        @click="handleQuickSubmit"
      >
        記下
      </button>
    </div>

    <!-- Grouped Inspirations (Matching Figma 06) -->
    <div class="date-sections-wrap">
      <div v-for="(items, date) in groupedInspirations" :key="date" class="date-group-block">
        <div class="date-header-label">{{ date }}</div>

        <div class="inspirations-list-group">
          <div 
            v-for="(item, idx) in items" 
            :key="item.id" 
            class="inspiration-card-row"
          >
            <div class="inspiration-num">{{ idx + 1 }}.</div>

            <div class="inspiration-content-column">
              <div class="inspiration-title-text">{{ item.title }}</div>

              <!-- Preview container for photo / link (Item 2 in Screen 06) -->
              <div v-if="item.preview_content || item.has_preview || idx === 1" class="preview-banner-box">
                <span class="preview-banner-text">貼文連結 / 照片預覽</span>
              </div>
            </div>

            <!-- Action Icons (Figma: Copy, Edit, Delete) -->
            <div class="actions-icon-group">
              <button class="action-icon-btn" title="複製靈感" @click="handleCopy(item)">
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="#4B4B50" stroke-width="2">
                  <rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
                  <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
                </svg>
              </button>

              <button class="action-icon-btn" title="編輯" @click="startEdit(item)">
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="#4B4B50" stroke-width="2">
                  <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                  <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                </svg>
              </button>

              <button class="action-icon-btn" title="刪除" @click="$emit('delete-inspiration', item.id)">
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="#4B4B50" stroke-width="2">
                  <polyline points="3 6 5 6 21 6"/>
                  <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>

      <div v-if="inspirations.length === 0" class="empty-hint">
        今天還沒有任何靈感，隨手寫下一則吧！
      </div>
    </div>

    <!-- Edit Inspiration Modal -->
    <div v-if="showEditModal" class="modal-backdrop" @click.self="showEditModal = false">
      <div class="modal-box">
        <h3 class="modal-title">編輯靈感內容</h3>
        <textarea 
          v-model="editItemText" 
          class="modal-textarea"
          rows="4"
        ></textarea>
        <div class="modal-actions">
          <button class="btn-cancel" @click="showEditModal = false">取消</button>
          <button class="btn-save" @click="saveEditItem">儲存更新</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  inspirations: {
    type: Array,
    default: () => [
      { id: 1, title: '去租書店把柯南漫畫裡的壞人都圈出來。', date_group: '2026/06/24' },
      { id: 2, title: '蛇的排遺出口只有一個，尿液有可能結晶堵塞。', preview_content: '貼文連結 / 照片預覽', date_group: '2026/06/24' },
      { id: 3, title: '舉辦社區吃麻糬大賽解決人口老化問題。', date_group: '2026/06/24' }
    ]
  }
})

const emit = defineEmits(['add-inspiration', 'delete-inspiration', 'search'])

const searchQuery = ref('')
const selectedDateFilter = ref('所有日期')
const quickNoteText = ref('')

const showEditModal = ref(false)
const editingItem = ref(null)
const editItemText = ref('')

const groupedInspirations = computed(() => {
  const groups = {}
  const list = props.inspirations || []
  for (const item of list) {
    const date = item.date_group || '2026/06/24'
    if (!groups[date]) groups[date] = []
    groups[date].push(item)
  }
  return groups
})

function toggleDateFilter() {
  selectedDateFilter.value = selectedDateFilter.value === '所有日期' ? '2026/06/24' : '所有日期'
}

function handleQuickSubmit() {
  if (!quickNoteText.value.trim()) return
  emit('add-inspiration', {
    title: quickNoteText.value.trim(),
    preview_content: '',
    image_url: ''
  })
  quickNoteText.value = ''
}

function handleCopy(item) {
  if (navigator.clipboard) {
    navigator.clipboard.writeText(item.title)
    alert('已複製靈感內容：\n' + item.title)
  }
}

function startEdit(item) {
  editingItem.value = item
  editItemText.value = item.title
  showEditModal.value = true
}

function saveEditItem() {
  if (editingItem.value) {
    editingItem.value.title = editItemText.value.trim()
  }
  showEditModal.value = false
}
</script>

<style scoped>
.inspirations-view {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.view-header {
  background: #FFFFFF;
  padding: 12px 18px;
  border-radius: 8px;
  border: 1px solid #DBDBD9;
  display: flex;
  flex-direction: column;
  gap: 3px;
}

.view-header-title {
  color: #141417;
  font-size: 24px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
  line-height: 1.2;
}

.view-header-subtitle {
  color: #75757A;
  font-size: 12px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 400;
}

.search-date-filter-row {
  display: flex;
  gap: 8px;
}

.search-box-pill {
  flex: 1;
  padding: 7px 12px;
  background: #FFFFFF;
  border-radius: 18px;
  outline: 1px solid #DBDBD9;
  outline-offset: -1px;
  display: flex;
  align-items: center;
}

.search-input {
  width: 100%;
  border: none;
  outline: none;
  background: transparent;
  font-size: 13px;
  font-family: inherit;
  color: #141417;
}

.search-input::placeholder {
  color: #141417;
}

.date-filter-pill {
  width: 118px;
  padding: 7px 12px;
  background: #FFFFFF;
  border-radius: 18px;
  outline: 1px solid #DBDBD9;
  outline-offset: -1px;
  border: none;
  color: #141417;
  font-size: 13px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.15s ease;
}

.date-filter-pill:hover {
  outline-color: #B8B8B8;
  background: #F8F8F6;
}

.quick-input-card {
  padding: 12px 14px;
  background: #FFFFFF;
  border-radius: 8px;
  outline: 1px solid #DBDBD9;
  outline-offset: -1px;
  display: flex;
  align-items: center;
  gap: 8px;
}

.quick-note-input {
  flex: 1;
  border: none;
  outline: none;
  background: transparent;
  font-size: 13px;
  font-family: inherit;
  color: #141417;
}

.quick-note-input::placeholder {
  color: #75757A;
}

.btn-save-note {
  padding: 4px 12px;
  background: #141417;
  color: #FFFFFF;
  border: none;
  border-radius: 14px;
  font-size: 12px;
  cursor: pointer;
}

.date-sections-wrap {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.date-group-block {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.date-header-label {
  color: #141417;
  font-size: 14px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
  padding: 2px 2px;
}

.inspirations-list-group {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.inspiration-card-row {
  padding: 14px;
  background: #FFFFFF;
  border-radius: 8px;
  outline: 1px solid #DBDBD9;
  outline-offset: -1px;
  display: flex;
  align-items: flex-start;
  gap: 10px;
  box-sizing: border-box;
}

.inspiration-num {
  color: #141417;
  font-size: 14px;
  font-weight: 700;
  padding-top: 1px;
}

.inspiration-content-column {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.inspiration-title-text {
  color: #141417;
  font-size: 13px;
  font-weight: 400;
  line-height: 1.4;
}

.preview-banner-box {
  width: 100%;
  height: 110px;
  background: #F0F0ED;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.preview-banner-text {
  color: #75757A;
  font-size: 13px;
}

.actions-icon-group {
  display: flex;
  align-items: center;
  gap: 6px;
}

.action-icon-btn {
  background: none;
  border: none;
  padding: 4px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 4px;
  transition: background 0.15s ease;
}

.action-icon-btn:hover {
  background: #F0F0ED;
}

.empty-hint {
  padding: 30px;
  text-align: center;
  color: #75757A;
  background: #FFFFFF;
  border-radius: 8px;
  border: 1px solid #DBDBD9;
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

.modal-box {
  width: 90%;
  max-width: 360px;
  background: #FFFFFF;
  border-radius: 12px;
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.modal-title {
  margin: 0;
  font-size: 16px;
  color: #141417;
}

.modal-textarea {
  width: 100%;
  padding: 10px;
  border: 1px solid #DBDBD9;
  border-radius: 8px;
  font-family: inherit;
  font-size: 13px;
  outline: none;
  resize: vertical;
  box-sizing: border-box;
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

.btn-cancel {
  padding: 6px 14px;
  border: 1px solid #DBDBD9;
  background: #FFFFFF;
  border-radius: 14px;
  font-size: 12px;
  cursor: pointer;
}

.btn-save {
  padding: 6px 14px;
  background: #141417;
  color: #FFFFFF;
  border: none;
  border-radius: 14px;
  font-size: 12px;
  cursor: pointer;
}
</style>
