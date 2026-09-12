<template>
  <div class="inspirations-view">
    <!-- View Header (Figma Screen 06) -->
    <div class="view-header">
      <div class="view-header-title">靈感清單</div>
      <div class="view-header-subtitle">蒐集靈感，隨手紀錄</div>
    </div>

    <!-- Search & Date Filter Bar -->
    <div class="search-filter-bar">
      <div class="search-input-wrapper">
        <input 
          v-model="searchQuery" 
          placeholder="搜尋關鍵字" 
          class="search-input"
          @input="$emit('search', searchQuery)"
        />
      </div>

      <div class="date-filter-pill">
        <span>所有日期</span>
      </div>
    </div>

    <!-- Quick Note Prompt Box (Figma Screen 06) -->
    <div class="figma-card quick-note-card">
      <input 
        v-model="quickNoteText" 
        class="quick-note-input"
        placeholder="今天有甚麼想法嗎？馬上記錄下來吧！"
        @keyup.enter="handleQuickSubmit"
      />
      <button 
        v-if="quickNoteText.trim()" 
        class="btn-dark btn-xs" 
        @click="handleQuickSubmit"
      >
        記下
      </button>
    </div>

    <!-- Grouped Inspirations (Matching Figma 06) -->
    <div class="inspirations-content">
      <div v-for="(items, date) in groupedInspirations" :key="date" class="date-group-section">
        <div class="date-group-header">{{ date }}</div>

        <div class="group-items-list">
          <div 
            v-for="(item, idx) in items" 
            :key="item.id" 
            class="figma-card inspiration-card"
          >
            <div class="inspiration-main-col">
              <div class="inspiration-title-row">
                <span class="item-num">{{ idx + 1 }}.</span>
                <span class="item-text">{{ item.title }}</span>
              </div>

              <!-- Preview box for item with link or image (Item 2 in Figma) -->
              <div v-if="item.preview_content || item.image_url" class="preview-box inspiration-preview-box">
                <span>貼文連結 / 照片預覽</span>
              </div>
            </div>

            <!-- Action Icons (Figma: Arrow Up, Pencil Edit, Trash) -->
            <div class="inspiration-actions">
              <button class="icon-btn" title="分享 / 複製" @click="handleShare(item)">
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <line x1="12" y1="19" x2="12" y2="5"/>
                  <polyline points="5 12 12 5 19 12"/>
                </svg>
              </button>

              <button class="icon-btn" title="編輯" @click="startEdit(item)">
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                  <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                </svg>
              </button>

              <button class="icon-btn delete-btn" title="刪除" @click="$emit('delete-inspiration', item.id)">
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
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
      <div class="modal-card">
        <h3 class="card-title">編輯靈感內容</h3>
        <textarea 
          v-model="editItemText" 
          class="textarea-control"
          placeholder="修改靈感..."
        ></textarea>
        <div class="modal-footer">
          <button class="btn-outline" @click="showEditModal = false">取消</button>
          <button class="btn-dark" @click="saveEditItem">儲存更新</button>
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
    default: () => []
  }
})

const emit = defineEmits(['add-inspiration', 'delete-inspiration', 'search'])

const searchQuery = ref('')
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

function handleQuickSubmit() {
  if (!quickNoteText.value.trim()) return
  emit('add-inspiration', {
    title: quickNoteText.value.trim(),
    preview_content: '',
    image_url: ''
  })
  quickNoteText.value = ''
}

function handleShare(item) {
  if (navigator.clipboard) {
    navigator.clipboard.writeText(item.title)
    alert('已複製靈感內容：' + item.title)
  }
}

function startEdit(item) {
  editingItem.value = item
  editItemText.value = item.title
  showEditModal.value = true
}

function saveEditItem() {
  if (editingItem.value && editItemText.value.trim()) {
    editingItem.value.title = editItemText.value.trim()
  }
  showEditModal.value = false
}
</script>

<style scoped>
.inspirations-view {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.search-filter-bar {
  display: flex;
  align-items: center;
  gap: 10px;
}

.search-input-wrapper {
  flex: 1;
}

.search-input {
  width: 100%;
  padding: 10px 16px;
  border-radius: var(--radius-pill);
  border: 1px solid var(--border-main);
  background: var(--bg-card);
  font-size: 13px;
  font-family: inherit;
  outline: none;
  transition: var(--transition);
}

.search-input:focus {
  border-color: var(--bg-dark);
}

.date-filter-pill {
  padding: 10px 18px;
  border-radius: var(--radius-pill);
  border: 1px solid var(--border-main);
  background: var(--bg-card);
  font-size: 13px;
  color: var(--text-main);
  cursor: pointer;
  flex-shrink: 0;
}

.quick-note-card {
  padding: 14px 18px;
  display: flex;
  align-items: center;
  gap: 10px;
}

.quick-note-input {
  flex: 1;
  border: none;
  background: transparent;
  font-size: 13px;
  font-family: inherit;
  color: var(--text-main);
  outline: none;
}

.quick-note-input::placeholder {
  color: var(--text-secondary);
}

.inspirations-content {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.date-group-section {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.date-group-header {
  font-size: 14px;
  font-weight: 700;
  color: var(--text-main);
  padding: 2px 4px;
}

.group-items-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.inspiration-card {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: flex-start;
  padding: 16px 18px;
  gap: 12px;
}

.inspiration-main-col {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.inspiration-title-row {
  display: flex;
  align-items: flex-start;
  gap: 6px;
}

.item-num {
  font-size: 14px;
  font-weight: 700;
  color: var(--text-main);
  flex-shrink: 0;
  line-height: 1.5;
}

.item-text {
  font-size: 13px;
  color: var(--text-main);
  line-height: 1.5;
}

.inspiration-preview-box {
  width: 100%;
  height: 84px;
  background: var(--bg-subtle);
  border-radius: var(--radius-md);
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text-secondary);
  font-size: 12px;
  border: 1px solid var(--border-card);
}

.inspiration-actions {
  display: flex;
  align-items: center;
  gap: 6px;
  flex-shrink: 0;
  margin-top: 2px;
}

.icon-btn {
  background: none;
  border: none;
  color: var(--icon-color);
  cursor: pointer;
  padding: 4px;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: var(--transition);
}

.icon-btn:hover {
  color: var(--text-main);
  background: var(--bg-subtle);
}

.icon-btn.delete-btn:hover {
  color: #FA5252;
}

.empty-hint {
  padding: 30px;
  text-align: center;
  color: var(--text-secondary);
  font-size: 13px;
  background: var(--bg-card);
  border-radius: var(--radius-md);
  border: 1px dashed var(--border-main);
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  margin-top: 10px;
}
</style>
