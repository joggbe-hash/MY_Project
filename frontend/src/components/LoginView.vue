<template>
  <div class="desktop-auth-page">
    <div class="auth-layout-grid">
      <!-- Left Hero Showcase Column (Desktop Only) -->
      <div class="auth-hero-col">
        <div class="hero-content">
          <!-- Brand Badge -->
          <div class="hero-brand">
            <div class="hero-badge">我</div>
            <div class="hero-brand-text">
              <span class="hero-logo-title">我也要</span>
              <span class="hero-logo-sub">Me Too Community</span>
            </div>
          </div>

          <h1 class="hero-headline">
            把靈感變成行動，<br />
            把行動留下成果。
          </h1>
          <p class="hero-desc">
            一個專為創作者、學生與實踐者打造的生活社群。看見好點子立即點擊「我也要」，今天就一起完成一件小事。
          </p>

          <!-- Interactive Feature Preview Cards on Desktop -->
          <div class="hero-showcase-cards">
            <!-- Card 1: Task Preview -->
            <div class="showcase-mini-card">
              <div class="mini-card-header">
                <span class="mini-tag tag-task">任務貼文</span>
                <span class="mini-time">剩 18 小時</span>
              </div>
              <div class="mini-card-title">拍一張今天的天空</div>
              <div class="mini-card-desc">用一張照片記錄今天的狀態，完成後分享一句話。</div>
              <div class="mini-card-footer">
                <div class="mini-pill-btn">我也要</div>
                <span class="mini-stat">成果 3 · 720 人參與</span>
              </div>
            </div>

            <!-- Card 2: Inspiration Preview -->
            <div class="showcase-mini-card">
              <div class="mini-card-header">
                <span class="mini-tag tag-insp">靈感貼文</span>
                <span class="mini-time">收藏保留 30 天</span>
              </div>
              <div class="mini-card-title">把通勤路上的顏色做成色票</div>
              <div class="mini-card-desc">看到有趣的配色先收藏，之後可以轉成自己的創作。</div>
              <div class="mini-palette-preview">
                <span style="background: #D8B4E2"></span>
                <span style="background: #A3C9E2"></span>
                <span style="background: #E8D3A7"></span>
                <span style="background: #C4E2C7"></span>
              </div>
            </div>
          </div>

          <!-- Feature Highlights -->
          <div class="hero-features-row">
            <div class="feature-bullet">
              <span class="bullet-check">✓</span>
              <span>今日任務連續設定</span>
            </div>
            <div class="feature-bullet">
              <span class="bullet-check">✓</span>
              <span>靈感隨手筆記</span>
            </div>
            <div class="feature-bullet">
              <span class="bullet-check">✓</span>
              <span>成果九宮格展示</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Auth Form Column (Spacious Desktop Form) -->
      <div class="auth-form-col">
        <div class="auth-box">
          <!-- Mobile Brand Header (Visible only on mobile) -->
          <div class="mobile-only-brand">
            <div class="hero-badge">我</div>
            <h2>我也要</h2>
            <p>把靈感變成行動，把行動留下成果</p>
          </div>

          <!-- Tabs Switcher -->
          <div class="auth-tabs-wrapper">
            <button 
              type="button"
              class="auth-tab" 
              :class="{ active: mode === 'login' }"
              @click="switchMode('login')"
            >
              帳號登入
            </button>
            <button 
              type="button"
              class="auth-tab" 
              :class="{ active: mode === 'register' }"
              @click="switchMode('register')"
            >
              新用戶註冊
            </button>
          </div>

          <!-- Alert Notification -->
          <div v-if="errorMsg" class="auth-alert error-alert">
            <span>{{ errorMsg }}</span>
          </div>
          <div v-if="successMsg" class="auth-alert success-alert">
            <span>{{ successMsg }}</span>
          </div>

          <!-- Auth Form -->
          <form class="auth-form-fields" @submit.prevent="handleSubmit">
            <div class="form-field">
              <label class="field-label">帳號 / 使用者名稱</label>
              <input 
                v-model="username" 
                type="text" 
                placeholder="請輸入帳號" 
                class="desktop-input" 
                required
              />
            </div>

            <div v-if="mode === 'register'" class="form-field">
              <label class="field-label">暱稱 (選填)</label>
              <input 
                v-model="nickname" 
                type="text" 
                placeholder="例如：薰衣草國度" 
                class="desktop-input" 
              />
            </div>

            <div class="form-field">
              <div class="field-label-row">
                <label class="field-label">密碼</label>
                <span class="field-hint" v-if="mode === 'login'">預設測試密碼：123456</span>
              </div>
              <div class="password-input-box">
                <input 
                  v-model="password" 
                  :type="showPassword ? 'text' : 'password'" 
                  placeholder="請輸入密碼" 
                  class="desktop-input" 
                  required
                />
                <button 
                  type="button" 
                  class="pwd-toggle-text" 
                  @click="showPassword = !showPassword"
                >
                  {{ showPassword ? '隱藏' : '顯示' }}
                </button>
              </div>
            </div>

            <button 
              type="submit" 
              class="desktop-submit-btn" 
              :disabled="loading"
            >
              {{ loading ? '處理中...' : (mode === 'login' ? '登入帳號' : '註冊並登入') }}
            </button>
          </form>

          <!-- Divider -->
          <div class="auth-or-divider">
            <span>快捷測試與體驗</span>
          </div>

          <!-- Quick Actions -->
          <div class="auth-quick-actions">
            <button 
              type="button" 
              class="quick-demo-card" 
              @click="quickDemoLogin"
              :disabled="loading"
            >
              <div class="quick-demo-left">
                <div class="demo-badge-icon">🚀</div>
                <div class="demo-info">
                  <div class="demo-name">一鍵以測試帳號登入</div>
                  <div class="demo-desc">帳號：lavender_nation (已注入完整 Figma 貼文與成果)</div>
                </div>
              </div>
              <div class="demo-arrow">→</div>
            </button>

            <!-- Direct Guest Preview Button -->
            <button 
              type="button" 
              class="guest-preview-btn" 
              @click="enterAsGuest"
            >
              免登入直接進入主頁瀏覽 →
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import * as api from '../api/client.js'

const emit = defineEmits(['login-success'])

const mode = ref('login')
const username = ref('lavender_nation')
const password = ref('123456')
const nickname = ref('')
const showPassword = ref(false)
const loading = ref(false)
const errorMsg = ref('')
const successMsg = ref('')

function switchMode(newMode) {
  mode.value = newMode
  errorMsg.value = ''
  successMsg.value = ''
  if (newMode === 'register') {
    username.value = ''
    password.value = ''
    nickname.value = ''
  } else {
    username.value = 'lavender_nation'
    password.value = '123456'
  }
}

async function handleSubmit() {
  if (!username.value.trim() || !password.value.trim()) {
    errorMsg.value = '請填寫帳號與密碼'
    return
  }

  errorMsg.value = ''
  successMsg.value = ''
  loading.value = true

  try {
    if (mode.value === 'login') {
      const res = await api.login(username.value.trim(), password.value.trim())
      emit('login-success', res.user)
    } else {
      await api.register(username.value.trim(), password.value.trim(), nickname.value.trim())
      successMsg.value = '註冊成功！請直接點擊登入。'
      mode.value = 'login'
    }
  } catch (err) {
    errorMsg.value = err.message || '發生錯誤，請稍後再試'
  } finally {
    loading.value = false
  }
}

async function quickDemoLogin() {
  username.value = 'lavender_nation'
  password.value = '123456'
  errorMsg.value = ''
  loading.value = true
  try {
    const res = await api.login('lavender_nation', '123456')
    emit('login-success', res.user)
  } catch (err) {
    emit('login-success', {
      id: 1,
      username: 'lavender_nation',
      nickname: '薰衣草國度',
      bio: '把靈感變成行動，把行動留下成果',
      weekly_info: '本週完成 5 個任務，收藏 8 則靈感'
    })
  } finally {
    loading.value = false
  }
}

function enterAsGuest() {
  emit('login-success', {
    id: 1,
    username: '訪客體驗者',
    nickname: '訪客',
    bio: '把靈感變成行動，把行動留下成果',
    weekly_info: '本週完成 5 個任務，收藏 8 則靈感'
  })
}
</script>

<style scoped>
.desktop-auth-page {
  width: 100%;
  min-height: 100vh;
  background-color: var(--bg-app);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 40px 24px;
}

.auth-layout-grid {
  width: 100%;
  max-width: 1140px;
  min-height: 640px;
  display: flex;
  gap: 48px;
  align-items: center;
  justify-content: space-between;
}

/* Left Hero Showcase Column */
.auth-hero-col {
  flex: 1.1;
  display: flex;
  flex-direction: column;
}

.hero-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
  max-width: 520px;
}

.hero-brand {
  display: flex;
  align-items: center;
  gap: 14px;
}

.hero-badge {
  width: 44px;
  height: 44px;
  background: var(--bg-dark);
  color: #fff;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 20px;
  box-shadow: 0 4px 12px rgba(20, 20, 23, 0.15);
}

.hero-brand-text {
  display: flex;
  flex-direction: column;
}

.hero-logo-title {
  font-size: 22px;
  font-weight: 700;
  color: var(--text-main);
  line-height: 1.2;
}

.hero-logo-sub {
  font-size: 12px;
  color: var(--text-secondary);
}

.hero-headline {
  font-size: 36px;
  font-weight: 700;
  color: var(--text-main);
  line-height: 1.3;
  letter-spacing: -0.5px;
}

.hero-desc {
  font-size: 15px;
  color: var(--text-secondary);
  line-height: 1.6;
}

.hero-showcase-cards {
  display: flex;
  gap: 16px;
  margin-top: 8px;
}

.showcase-mini-card {
  flex: 1;
  background: var(--bg-card);
  border: 1px solid var(--border-main);
  border-radius: 12px;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 8px;
  box-shadow: var(--shadow-sm);
}

.mini-card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 11px;
}

.mini-tag {
  padding: 2px 6px;
  border-radius: 4px;
  font-weight: 700;
}

.tag-task {
  background: var(--bg-dark);
  color: #fff;
}

.tag-insp {
  background: var(--bg-secondary);
  color: var(--text-main);
}

.mini-time {
  color: var(--text-muted);
}

.mini-card-title {
  font-size: 14px;
  font-weight: 700;
  color: var(--text-main);
}

.mini-card-desc {
  font-size: 12px;
  color: var(--text-secondary);
  line-height: 1.4;
}

.mini-card-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 4px;
}

.mini-pill-btn {
  padding: 4px 10px;
  background: var(--bg-dark);
  color: #fff;
  border-radius: 14px;
  font-size: 11px;
  font-weight: 700;
}

.mini-stat {
  font-size: 11px;
  color: var(--text-secondary);
}

.mini-palette-preview {
  display: flex;
  height: 20px;
  border-radius: 4px;
  overflow: hidden;
  margin-top: 4px;
}

.mini-palette-preview span {
  flex: 1;
}

.hero-features-row {
  display: flex;
  gap: 18px;
  margin-top: 6px;
}

.feature-bullet {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: var(--text-main);
  font-weight: 500;
}

.bullet-check {
  width: 18px;
  height: 18px;
  border-radius: 50%;
  background: #EBFBEE;
  color: #2B8A3E;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 11px;
  font-weight: 700;
}

/* Right Auth Form Column */
.auth-form-col {
  flex: 0.9;
  display: flex;
  justify-content: center;
}

.auth-box {
  width: 100%;
  max-width: 440px;
  background: var(--bg-card);
  border: 1px solid var(--border-main);
  border-radius: 20px;
  padding: 36px 32px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.06);
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.mobile-only-brand {
  display: none;
  flex-direction: column;
  align-items: center;
  text-align: center;
  gap: 6px;
  margin-bottom: 8px;
}

.mobile-only-brand h2 {
  font-size: 24px;
  font-weight: 700;
}

.mobile-only-brand p {
  font-size: 13px;
  color: var(--text-secondary);
}

.auth-tabs-wrapper {
  display: flex;
  background: var(--bg-app);
  padding: 4px;
  border-radius: 12px;
  border: 1px solid var(--border-light);
}

.auth-tab {
  flex: 1;
  padding: 9px 16px;
  border: none;
  background: transparent;
  color: var(--text-secondary);
  font-size: 14px;
  font-weight: 500;
  border-radius: 8px;
  cursor: pointer;
  transition: var(--transition);
}

.auth-tab.active {
  background: var(--bg-card);
  color: var(--text-main);
  font-weight: 700;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.05);
}

.auth-form-fields {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.form-field {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.field-label-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.field-label {
  font-size: 13px;
  font-weight: 700;
  color: var(--text-main);
}

.field-hint {
  font-size: 11px;
  color: var(--text-secondary);
}

.desktop-input {
  width: 100%;
  padding: 12px 14px;
  border-radius: 10px;
  border: 1px solid var(--border-main);
  background: var(--bg-app);
  font-size: 14px;
  font-family: inherit;
  color: var(--text-main);
  outline: none;
  transition: var(--transition);
}

.desktop-input:focus {
  border-color: var(--bg-dark);
  background: var(--bg-card);
}

.password-input-box {
  position: relative;
  display: flex;
  align-items: center;
}

.pwd-toggle-text {
  position: absolute;
  right: 12px;
  background: none;
  border: none;
  font-size: 12px;
  color: var(--text-secondary);
  cursor: pointer;
  padding: 4px;
}

.pwd-toggle-text:hover {
  color: var(--text-main);
}

.desktop-submit-btn {
  width: 100%;
  padding: 13px;
  border-radius: 10px;
  background: var(--bg-dark);
  color: #fff;
  border: none;
  font-size: 15px;
  font-weight: 700;
  cursor: pointer;
  transition: var(--transition);
  margin-top: 4px;
}

.desktop-submit-btn:hover {
  background: var(--bg-dark-hover);
  transform: translateY(-1px);
}

.auth-or-divider {
  display: flex;
  align-items: center;
  text-align: center;
  color: var(--text-muted);
  font-size: 12px;
}

.auth-or-divider::before,
.auth-or-divider::after {
  content: '';
  flex: 1;
  border-bottom: 1px solid var(--border-light);
}

.auth-or-divider span {
  padding: 0 10px;
}

.auth-quick-actions {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.quick-demo-card {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 16px;
  background: var(--bg-app);
  border: 1px solid var(--border-main);
  border-radius: 12px;
  cursor: pointer;
  transition: var(--transition);
  text-align: left;
}

.quick-demo-card:hover {
  background: #ECECE9;
  border-color: var(--border-strong);
  transform: translateY(-1px);
}

.quick-demo-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.demo-badge-icon {
  font-size: 22px;
}

.demo-name {
  font-size: 13px;
  font-weight: 700;
  color: var(--text-main);
}

.demo-desc {
  font-size: 11px;
  color: var(--text-secondary);
}

.demo-arrow {
  font-size: 16px;
  color: var(--text-secondary);
  font-weight: 700;
}

.guest-preview-btn {
  background: none;
  border: none;
  color: var(--text-secondary);
  font-size: 13px;
  cursor: pointer;
  text-align: center;
  padding: 6px;
  border-radius: 6px;
  transition: var(--transition);
}

.guest-preview-btn:hover {
  color: var(--text-main);
  text-decoration: underline;
}

.auth-alert {
  padding: 10px 14px;
  border-radius: 8px;
  font-size: 13px;
}

.error-alert {
  background: #FFF5F5;
  color: #E03131;
  border: 1px solid #FFC9C9;
}

.success-alert {
  background: #EBFBEE;
  color: #2B8A3E;
  border: 1px solid #B2F2BB;
}

/* =========================================================
   Responsive Breakpoints for Auth Page
========================================================= */

@media (max-width: 960px) {
  .auth-hero-col {
    display: none;
  }

  .auth-layout-grid {
    justify-content: center;
    max-width: 480px;
  }

  .mobile-only-brand {
    display: flex;
  }

  .desktop-auth-page {
    padding: 24px 16px;
  }
}
</style>
