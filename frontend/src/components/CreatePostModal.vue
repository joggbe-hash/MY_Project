<template>
  <div class="modal-backdrop" @click.self="$emit('close')">
    <div class="modal-card create-post-modal">
      <!-- Modal Header -->
      <div class="modal-header-row">
        <div>
          <h2 class="modal-title">我也要發文</h2>
          <p class="modal-subtitle">發布任務或靈感，讓別人也能回應</p>
        </div>
        <button class="close-btn" @click="$emit('close')">✕</button>
      </div>

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
        <input 
          v-model="postTitle" 
          placeholder="貼文標題：" 
          class="input-control title-input"
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

      <!-- Photo Upload Mock Box (Figma matching) -->
      <div class="photo-upload-box" @click="toggleMockImage">
        <div class="upload-icon-circle">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z"/>
            <circle cx="12" cy="13" r="4"/>
          </svg>
        </div>
        <div class="upload-label">
          {{ hasImage ? '已附帶照片預覽 (點擊移除)' : '可加入照片' }}
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

const emit = defineEmits(['close', 'publish'])

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
}
</script>

<style scoped>
.create-post-modal {
  max-width: 440px;
  gap: 16px;
}

.modal-header-row {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}

.modal-title {
  font-size: 20px;
  font-weight: 700;
  color: var(--text-main);
}

.modal-subtitle {
  font-size: 12px;
  color: var(--text-secondary);
  margin-top: 2px;
}

.close-btn {
  background: none;
  border: none;
  font-size: 18px;
  color: var(--text-secondary);
  cursor: pointer;
  padding: 4px;
}

.close-btn:hover {
  color: var(--text-main);
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

.type-pills-row {
  display: flex;
  gap: 10px;
}

.type-pill {
  flex: 1;
  padding: 8px 16px;
  border-radius: var(--radius-pill);
  border: 1px solid var(--border-main);
  background: var(--bg-card);
  color: var(--text-main);
  font-size: 13px;
  cursor: pointer;
  transition: var(--transition);
  display: flex;
  align-items: center;
  justify-content: center;
}

.type-pill.active {
  background: var(--bg-dark);
  color: #fff;
  border-color: var(--bg-dark);
  font-weight: 700;
}

.title-input {
  font-weight: 500;
}

.photo-upload-box {
  width: 100%;
  height: 110px;
  background: var(--bg-secondary);
  border-radius: var(--radius-md);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 6px;
  cursor: pointer;
  transition: var(--transition);
  border: 1px dashed var(--border-strong);
}

.photo-upload-box:hover {
  background: #E8E8E4;
}

.upload-icon-circle {
  color: var(--icon-color);
}

.upload-label {
  color: var(--text-secondary);
  font-size: 13px;
}

.publish-btn {
  width: 100%;
  padding: 12px;
  font-size: 14px;
  margin-top: 4px;
}
</style>
