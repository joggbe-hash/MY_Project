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
            <div class="status-pill-tag executing-tag">
              執行任務
            </div>
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
            <div class="status-pill-tag executing-tag">
              執行任務
            </div>
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

    <!-- Task Execution Modal (Matching Figma Screen 10) -->
    <div v-if="showProgressModal" class="modal-backdrop" @click.self="showProgressModal = false">
      <div class="modal-card execution-modal-card">
        <div class="execution-header">
          <h3 class="card-title">執行記錄</h3>
        </div>

        <div class="figma-card task-info-card">
          <div class="status-pill">進行中</div>
          <h4 class="execution-task-title">{{ activeTask?.title || '拍一張今天的天空' }}</h4>
          <p class="execution-task-desc">用照片或文字逐步記錄執行過程，完成後發布成果。</p>
        </div>

        <div class="execution-section-header">
          <div class="section-heading">過程記錄</div>
          <div class="section-subtext">依序加入每個步驟，可使用文字或照片</div>
        </div>

        <div class="execution-steps-list">
          <!-- Step 1 -->
          <div class="step-card">
            <div class="step-badge">1</div>
            <div class="step-media-box">
              <span>{{ progressInput || '一張相片' }}</span>
            </div>
            <div class="step-time">記錄時間 今天 14:20</div>
          </div>

          <!-- Step 2 -->
          <div class="step-card">
            <div class="step-badge">2</div>
            <div class="step-actions-row">
              <div class="step-add-box" @click="addTextStep">
                <span class="plus-icon">＋</span>
                <span>新增文字</span>
              </div>
              <div class="step-add-box" @click="addPhotoStep">
                <span class="plus-icon">＋</span>
                <span>拍照或從相簿選擇</span>
              </div>
            </div>
          </div>
        </div>

        <div class="execution-modal-footer">
          <button class="btn-outline flex-1" @click="saveProgressNote">儲存目前進度</button>
          <button class="btn-dark flex-1" @click="finishTask">完成任務</button>
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
  if (activeTask.value) activeTask.value.progress_note = progressInput.value
  showProgressModal.value = false
}

function finishTask() {
  if (activeTask.value) {
    handleComplete(activeTask.value.id)
    showProgressModal.value = false
  }
}

function addTextStep() {
  const note = prompt('請輸入新增的步驟文字：', '完成初步紀錄')
  if (note) {
    progressInput.value = note
  }
}

function addPhotoStep() {
  progressInput.value = '已選擇相片照片'
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

.status-pill-tag {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 6px 14px;
  background: #FFFFFF;
  border: 1px solid #E5E5E5;
  border-radius: 18px;
  color: #141417;
  font-size: 13px;
  font-weight: 500;
  cursor: default;
  user-select: none;
}

/* Figma Screen 10 Execution Modal Styles */
.execution-modal-card {
  max-width: 380px;
  width: 90%;
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.status-pill {
  display: inline-block;
  padding: 4px 10px;
  background: #1F1F21;
  color: #FFFFFF;
  font-size: 11px;
  font-weight: 500;
  border-radius: 14px;
  align-self: flex-start;
}

.task-info-card {
  padding: 14px;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.execution-task-title {
  font-size: 17px;
  font-weight: 500;
  color: #1A1A1C;
}

.execution-task-desc {
  font-size: 12px;
  color: #616369;
  line-height: 1.4;
}

.execution-section-header {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.section-heading {
  font-size: 16px;
  font-weight: 500;
  color: #1A1A1C;
}

.section-subtext {
  font-size: 11px;
  color: #616369;
}

.execution-steps-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.step-card {
  padding: 14px;
  background: #FFFFFF;
  border: 1px solid #D6D9DB;
  border-radius: 6px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.step-badge {
  width: 26px;
  height: 26px;
  background: #1F1F21;
  color: #FFFFFF;
  border-radius: 13px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 500;
}

.step-media-box {
  height: 112px;
  background: #F5F5F7;
  border-radius: 5px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #616369;
  font-size: 12px;
}

.step-time {
  font-size: 11px;
  color: #616369;
}

.step-actions-row {
  display: flex;
  gap: 10px;
}

.step-add-box {
  flex: 1;
  height: 112px;
  background: #F5F5F7;
  border-radius: 5px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 4px;
  color: #616369;
  font-size: 12px;
  cursor: pointer;
}

.plus-icon {
  font-size: 20px;
  line-height: 1;
}

.execution-modal-footer {
  display: flex;
  gap: 10px;
  padding-top: 6px;
}

.flex-1 {
  flex: 1;
}
</style>
