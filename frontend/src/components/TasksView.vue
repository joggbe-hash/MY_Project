<template>
  <div class="tasks-view">
    <!-- View Header (Figma Screen 05: Clean "我的任務") -->
    <div class="view-header">
      <div class="view-header-title">我的任務</div>
    </div>

    <!-- Section 1: 今日私人任務 -->
    <div class="task-category-section">
      <div class="section-title-row">
        <h3 class="section-title">今日私人任務</h3>
        <button class="btn-outline btn-xs" @click="showAddPrivate = !showAddPrivate">
          {{ showAddPrivate ? '取消' : '＋ 新增私人任務' }}
        </button>
      </div>

      <!-- Add Private Task Inline Form -->
      <div v-if="showAddPrivate" class="inline-add-box">
        <input 
          v-model="newPrivateTitle" 
          placeholder="輸入私人任務名稱..." 
          class="input-control"
          @keyup.enter="submitPrivateTask"
        />
        <button class="btn-dark btn-sm" @click="submitPrivateTask">確認新增</button>
      </div>

      <!-- Private Tasks Cards (Matching Figma 05) -->
      <div class="task-cards-list">
        <div 
          v-for="task in privateTasks" 
          :key="task.id" 
          class="figma-card task-card"
        >
          <div class="task-text-content">
            <h4 class="task-card-title">{{ task.title }}</h4>
            <p class="task-card-sub">{{ task.description || '私人任務，永久留在自己的清單' }}</p>
            <div v-if="task.progress_note" class="task-note-pill">
              進度紀錄：{{ task.progress_note }}
            </div>
          </div>

          <div class="task-actions-row">
            <button class="btn-outline btn-sm" @click="openProgressModal(task)">
              進度紀錄
            </button>
            <button class="btn-dark btn-sm" @click="handleComplete(task.id)">
              完成任務
            </button>
          </div>
        </div>

        <div v-if="privateTasks.length === 0" class="empty-state-box">
          目前沒有未完成的私人任務
        </div>
      </div>
    </div>

    <!-- Section 2: 公共任務 -->
    <div class="task-category-section">
      <div class="section-title-row">
        <h3 class="section-title">公共任務</h3>
      </div>

      <!-- Public Tasks Cards (Matching Figma 05) -->
      <div class="task-cards-list">
        <div 
          v-for="task in publicTasks" 
          :key="task.id" 
          class="figma-card task-card"
        >
          <div class="task-text-content">
            <h4 class="task-card-title">{{ task.title }}</h4>
            <p class="task-card-sub">{{ task.description || '來自任務貼文，完成後需分享成果' }}</p>
          </div>

          <div class="task-actions-row">
            <button class="btn-dark btn-sm" @click="handleComplete(task.id)">
              完成任務
            </button>
          </div>
        </div>

        <div v-if="publicTasks.length === 0" class="empty-state-box">
          尚未加入任何公共任務，前往貼文區點擊「我也要」即可加入！
        </div>
      </div>
    </div>

    <!-- Section 3: 完成紀錄 (Matching Figma 05) -->
    <div class="task-category-section">
      <div class="figma-card completed-container-card">
        <div class="completed-header-row" @click="isCompletedExpanded = !isCompletedExpanded">
          <h3 class="section-title">完成紀錄</h3>
          <span class="count-badge">{{ completedTasks.length }} 個已完成 {{ isCompletedExpanded ? '▲' : '▼' }}</span>
        </div>

        <div v-if="isCompletedExpanded && completedTasks.length > 0" class="completed-tasks-list">
          <div 
            v-for="task in completedTasks" 
            :key="task.id" 
            class="completed-item-row"
          >
            <span class="completed-item-title">{{ task.title }}</span>
            <span class="completed-check">✓</span>
          </div>
        </div>

        <div v-else-if="isCompletedExpanded && completedTasks.length === 0" class="empty-completed-text">
          尚無完成紀錄，完成今日任務後會在此保留成果！
        </div>
      </div>
    </div>

    <!-- Progress Note Modal -->
    <div v-if="showProgressModal" class="modal-backdrop" @click.self="showProgressModal = false">
      <div class="modal-card">
        <h3 class="card-title">更新任務進度紀錄</h3>
        <p class="card-desc">{{ activeTask?.title }}</p>
        <input 
          v-model="progressInput" 
          class="input-control" 
          placeholder="例如：已完成第 1~3 章節草稿..." 
          @keyup.enter="saveProgressNote"
        />
        <div class="modal-footer">
          <button class="btn-outline" @click="showProgressModal = false">取消</button>
          <button class="btn-dark" @click="saveProgressNote">儲存進度</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
  privateTasks: {
    type: Array,
    default: () => []
  },
  publicTasks: {
    type: Array,
    default: () => []
  },
  completedTasks: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['create-private-task', 'complete-task', 'update-progress'])

const showAddPrivate = ref(false)
const newPrivateTitle = ref('')
const isCompletedExpanded = ref(false)

const showProgressModal = ref(false)
const activeTask = ref(null)
const progressInput = ref('')

function submitPrivateTask() {
  if (!newPrivateTitle.value.trim()) return
  emit('create-private-task', newPrivateTitle.value.trim())
  newPrivateTitle.value = ''
  showAddPrivate.value = false
}

function handleComplete(id) {
  emit('complete-task', id)
}

function openProgressModal(task) {
  activeTask.value = task
  progressInput.value = task.progress_note || ''
  showProgressModal.value = true
}

function saveProgressNote() {
  if (!activeTask.value) return
  emit('update-progress', { id: activeTask.value.id, note: progressInput.value })
  activeTask.value.progress_note = progressInput.value
  showProgressModal.value = false
}
</script>

<style scoped>
.tasks-view {
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.task-category-section {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.section-title-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 4px;
}

.section-title {
  font-size: 16px;
  font-weight: 700;
  color: var(--text-main);
}

.count-badge {
  font-size: 12px;
  color: var(--text-secondary);
  background: var(--bg-subtle);
  padding: 3px 10px;
  border-radius: var(--radius-pill);
}

.inline-add-box {
  display: flex;
  gap: 8px;
  padding: 10px;
  background: var(--bg-card);
  border-radius: var(--radius-md);
  border: 1px solid var(--border-main);
}

.task-cards-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.task-card {
  padding: 18px 20px;
  gap: 14px;
}

.task-text-content {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.task-card-title {
  font-size: 15px;
  font-weight: 700;
  color: var(--text-main);
  line-height: 1.4;
}

.task-card-sub {
  font-size: 13px;
  color: var(--text-secondary);
}

.task-note-pill {
  font-size: 11px;
  color: #1A5424;
  background: #EBFBEE;
  padding: 3px 8px;
  border-radius: 4px;
  align-self: flex-start;
  margin-top: 6px;
}

.task-actions-row {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: 8px;
}

.completed-container-card {
  padding: 16px 20px;
  cursor: pointer;
}

.completed-header-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.completed-tasks-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin-top: 12px;
  padding-top: 12px;
  border-top: 1px solid var(--border-light);
}

.completed-item-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 12px;
  background: var(--bg-subtle);
  border-radius: var(--radius-md);
  font-size: 13px;
}

.completed-item-title {
  text-decoration: line-through;
  color: var(--text-secondary);
}

.completed-check {
  color: #2B8A3E;
  font-weight: 700;
}

.empty-completed-text {
  font-size: 12px;
  color: var(--text-secondary);
  text-align: center;
  padding: 12px 0 4px;
}

.empty-state-box {
  padding: 24px 16px;
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
}

.btn-xs {
  padding: 4px 10px;
  font-size: 11px;
}
</style>
