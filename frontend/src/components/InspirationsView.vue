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

    <!-- Quick Note Prompt Box (Figma Screen 06 - 依設計註釋：直接打字後，儲存/取消) -->
    <div class="figma-card quick-note-card">
      <input 
        v-model="quickNoteText" 
        class="quick-note-input"
        placeholder="今天有甚麼想法嗎？馬上記錄下來吧！"
        @keyup.enter="handleQuickSubmit"
      />
      <div v-if="quickNoteText.trim()" class="quick-note-actions">
        <button class="btn-dark btn-xs" @click="handleQuickSubmit">儲存</button>
        <button class="btn-outline btn-xs" @click="quickNoteText = ''">取消</button>
      </div>
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
                
                <!-- 點擊直接編輯 (依設計註釋) -->
                <div v-if="editingItemId === item.id" class="inline-edit-wrapper">
                  <input 
                    v-model="editingText" 
                    class="inline-edit-input"
                    autofocus
                    @keyup.enter="saveInlineEdit(item)"
                    @keyup.esc="cancelInlineEdit"
                    @blur="saveInlineEdit(item)"
                  />
                  <button class="inline-save-btn" title="儲存" @mousedown.prevent="saveInlineEdit(item)">
                    ✓
                  </button>
                </div>
                <span 
                  v-else 
                  class="item-text clickable-edit-text" 
                  title="點擊直接編輯"
                  @click="startInlineEdit(item)"
                >
                  {{ item.title }}
                </span>
              </div>

              <!-- Preview box for item with link or image (Item 2 in Figma) -->
              <div v-if="item.preview_content || item.image_url" class="preview-box inspiration-preview-box">
                <span>貼文連結 / 照片預覽</span>
              </div>
            </div>

            <!-- Action Icons (Figma: 推至私人任務區 ↑, 點擊直接編輯 ✎, 刪除) -->
            <div class="inspiration-actions">
              <button 
                class="icon-btn push-task-btn" 
                title="推至私人任務區" 
                @click="pushToPrivate(item)"
              >
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <line x1="12" y1="19" x2="12" y2="5"/>
                  <polyline points="5 12 12 5 19 12"/>
                </svg>
              </button>

              <button 
                class="icon-btn edit-btn" 
                title="點擊直接編輯" 
                @click="startInlineEdit(item)"
              >
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                  <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                </svg>
              </button>

              <button 
                class="icon-btn delete-btn" 
                title="刪除" 
                @click="promptDelete(item)"
              >
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

    <!-- 二次確認是否刪除彈窗 (RUICHIEN 設計註釋: 二次確認是否刪除) -->
    <div v-if="itemToDelete" class="modal-backdrop" @click.self="itemToDelete = null">
      <div class="modal-card confirm-modal">
        <h3 class="card-title">確定刪除這則靈感？</h3>
        <p class="confirm-message">「{{ itemToDelete.title }}」</p>
        <p class="confirm-submessage">刪除後將無法復原，是否確認刪除？</p>
        <div class="modal-footer">
          <button class="btn-outline" @click="itemToDelete = null">取消</button>
          <button class="btn-danger" @click="confirmDelete">確認刪除</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, nextTick } from 'vue'

const props = defineProps({
  inspirations: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits([
  'add-inspiration', 
  'delete-inspiration', 
  'update-inspiration',
  'push-to-private-task',
  'search'
])

const searchQuery = ref('')
const quickNoteText = ref('')

// 直接編輯狀態 (RUICHIEN 設計註釋: 點擊直接編輯)
const editingItemId = ref(null)
const editingText = ref('')

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

// RUICHIEN 設計註釋: 推至私人任務區
function pushToPrivate(item) {
  emit('push-to-private-task', item)
}

// RUICHIEN 設計註釋: 點擊直接編輯
function startInlineEdit(item) {
  editingItemId.value = item.id
  editingText.value = item.title
}

function saveInlineEdit(item) {
  if (editingItemId.value !== item.id) return
  if (editingText.value.trim()) {
    item.title = editingText.value.trim()
    emit('update-inspiration', { id: item.id, title: item.title })
  }
  editingItemId.value = null
  editingText.value = ''
}

// RUICHIEN 設計註釋: 二次確認是否刪除
const itemToDelete = ref(null)

function promptDelete(item) {
  itemToDelete.value = item
}

function confirmDelete() {
  if (itemToDelete.value) {
    emit('delete-inspiration', itemToDelete.value.id)
    itemToDelete.value = null
  }
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

.quick-note-actions {
  display: flex;
  align-items: center;
  gap: 6px;
}

.clickable-edit-text {
  cursor: pointer;
  border-radius: 4px;
  padding: 1px 4px;
  transition: var(--transition);
}

.clickable-edit-text:hover {
  background: var(--bg-subtle);
  color: var(--text-main);
  text-decoration: underline;
  text-decoration-style: dashed;
}

.inline-edit-wrapper {
  display: flex;
  align-items: center;
  gap: 6px;
  flex: 1;
}

.inline-edit-input {
  flex: 1;
  padding: 4px 8px;
  border-radius: var(--radius-sm);
  border: 1px solid var(--border-card);
  background: var(--bg-card);
  font-size: 13px;
  color: var(--text-main);
  font-family: inherit;
  outline: none;
}

.inline-edit-input:focus {
  border-color: var(--text-main);
}

.inline-save-btn {
  background: var(--bg-dark);
  color: white;
  border: none;
  border-radius: 4px;
  padding: 4px 8px;
  font-size: 12px;
  font-weight: 700;
  cursor: pointer;
}

.push-task-btn:hover {
  color: #2F9E44;
}

.edit-btn:hover {
  color: #1971C2;
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

.confirm-modal {
  max-width: 320px;
  text-align: center;
}

.confirm-message {
  font-size: 14px;
  font-weight: 700;
  color: var(--text-main);
  margin: 12px 0 6px;
  line-height: 1.4;
}

.confirm-submessage {
  font-size: 12px;
  color: var(--text-secondary);
  margin-bottom: 16px;
}

.modal-footer {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
}

.btn-danger {
  background: #FA5252;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: var(--radius-pill);
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
  transition: var(--transition);
}

.btn-danger:hover {
  background: #E03131;
}
</style>
