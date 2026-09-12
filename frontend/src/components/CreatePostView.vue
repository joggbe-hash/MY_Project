<template>
  <div class="create-post-view">
    <!-- Header (Figma Screen 04) -->
    <div class="view-header">
      <div class="view-header-title">我也要發文</div>
      <div class="view-header-subtitle">發布任務或靈感，讓別人也能回應</div>
    </div>

    <!-- Main Card Form -->
    <div class="figma-card create-card">
      <!-- Type Selector -->
      <div class="post-type-group">
        <label class="section-label">選擇貼文類型</label>
        <div class="type-pills-row">
          <button 
            type="button"
            class="type-pill" 
            :class="{ active: postType === 'task' }"
            @click="postType = 'task'"
          >
            任務貼文
          </button>
          <button 
            type="button"
            class="type-pill" 
            :class="{ active: postType === 'inspiration' }"
            @click="postType = 'inspiration'"
          >
            靈感貼文
          </button>
        </div>
      </div>

      <!-- Title Input -->
      <div class="input-group">
        <label class="input-label">貼文標題：</label>
        <input 
          v-model="postTitle" 
          placeholder="請輸入標題..." 
          class="input-control"
        />
      </div>

      <!-- Content Textarea -->
      <div class="input-group">
        <textarea 
          v-model="postContent" 
          class="textarea-control"
          placeholder="寫下你想邀請別人一起做的事，或記錄一則靈感。"
          rows="4"
        ></textarea>
      </div>

      <!-- Photo Upload Mock Box (Figma Screen 04 matching) -->
      <div class="photo-upload-box" @click="toggleMockImage">
        <div class="upload-icon-circle">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
            <line x1="12" y1="19" x2="12" y2="5"/>
            <polyline points="5 12 12 5 19 12"/>
          </svg>
        </div>
        <div class="upload-label">
          {{ hasImage ? '已加入照片 (點擊移除)' : '可加入照片' }}
        </div>
      </div>

      <!-- Submit Button -->
      <button 
        class="btn-dark publish-btn" 
        :disabled="!postTitle.trim()"
        @click="handlePublish"
      >
        發布貼文
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const emit = defineEmits(['publish'])

const postType = ref('task')
const postTitle = ref('')
const postContent = ref('')
const hasImage = ref(false)

function toggleMockImage() {
  hasImage.value = !hasImage.value
}

function handlePublish() {
  if (!postTitle.value.trim()) return
  emit('publish', {
    post_type: postType.value,
    title: postTitle.value.trim(),
    content: postContent.value.trim() || (postType.value === 'task' ? '一起來挑戰看看！' : '看到的好點子，分享給大家。'),
    image_url: hasImage.value ? 'uploaded-photo-preview' : ''
  })
  postTitle.value = ''
  postContent.value = ''
  hasImage.value = false
}
</script>

<style scoped>
.create-post-view {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.create-card {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.post-type-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.section-label {
  font-size: 14px;
  font-weight: 700;
  color: var(--text-main);
}

.input-label {
  font-size: 14px;
  font-weight: 600;
  color: var(--text-main);
  margin-bottom: 6px;
  display: block;
}

.type-pills-row {
  display: flex;
  gap: 8px;
}

.type-pill {
  flex: 1;
  padding: 8px 16px;
  border-radius: var(--radius-pill);
  font-size: 13px;
  font-weight: 600;
  border: 1px solid var(--border-main);
  background: transparent;
  color: var(--text-main);
  cursor: pointer;
  transition: var(--transition);
}

.type-pill.active {
  background: var(--bg-dark);
  color: var(--text-white);
  border-color: var(--bg-dark);
}

.photo-upload-box {
  border: 1px dashed var(--border-strong);
  border-radius: var(--radius-md);
  padding: 30px 16px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 8px;
  cursor: pointer;
  background: var(--bg-subtle);
  transition: var(--transition);
}

.photo-upload-box:hover {
  background: #ECECE9;
  border-color: var(--text-main);
}

.upload-icon-circle {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background: var(--bg-card);
  border: 1px solid var(--border-main);
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text-main);
}

.upload-label {
  font-size: 13px;
  color: var(--text-secondary);
  font-weight: 500;
}

.publish-btn {
  width: 100%;
  padding: 12px;
  font-size: 15px;
  font-weight: 700;
  margin-top: 6px;
}
</style>
